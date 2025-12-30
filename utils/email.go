package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net"
	"net/mail"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/emersion/go-msgauth/dmarc"
	"github.com/meta-node-blockchain/verify_server/model"
	"github.com/microcosm-cc/bluemonday"
	"github.com/miekg/dns"
	"github.com/toorop/go-dkim"
)
func SanitizeEmailHTML(html string) string {
	policy := bluemonday.UGCPolicy()
	policy.AllowElements("html", "head", "body", "label", "input", "font", "main", "nav", "header", "footer", "kbd", "legend", "map", "title", "div", "span")
	policy.AllowAttrs("style").Globally()
	policy.AllowAttrs("face", "size").OnElements("font")
	policy.AllowAttrs("name", "content", "http-equiv").OnElements("meta")
	policy.AllowAttrs("name", "data-id").OnElements("a")
	policy.AllowAttrs("for").OnElements("label")
	policy.AllowAttrs("type").OnElements("input")
	policy.AllowAttrs("rel", "href").OnElements("link")
	policy.AllowAttrs("topmargin", "leftmargin", "marginwidth", "marginheight", "yahoo").OnElements("body")
	policy.AllowAttrs("xmlns").OnElements("html")
	policy.AllowAttrs("style", "vspace", "hspace", "face", "bgcolor", "color", "border", "cellpadding", "cellspacing").Globally()
	policy.AllowAttrs("class", "id", "style").OnElements("div", "span")
	policy.AllowAttrs("bgcolor", "color", "align").OnElements("basefont", "font", "hr", "table", "td")
	policy.AllowAttrs("border").OnElements("img", "table", "basefont", "font", "hr", "td")
	policy.AllowAttrs("cellpadding", "cellspacing", "valign", "halign").OnElements("table")
	policy.AllowAttrs("src").OnElements("img")

	trustedImagePattern := regexp.MustCompile(`^(data:image/|https?://(?:m\.pro|payws\.com|payws\.net))`)
	policy.AllowAttrs("src").Matching(trustedImagePattern).OnElements("img")
	policy.AllowDataURIImages()
	policy.RequireNoFollowOnLinks(true)
	policy.RequireNoFollowOnFullyQualifiedLinks(true)
	policy.AddTargetBlankToFullyQualifiedLinks(true)

	return policy.Sanitize(html)
}
// ============================================
// EMAIL VALIDATION FUNCTIONS
// ============================================

func ExtractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func IsIPInCIDR(ip, cidr string) bool {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	return network.Contains(net.ParseIP(ip))
}

func CheckSPF(ip, domain string) (bool, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeTXT)

	client := new(dns.Client)
	resp, _, err := client.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return false, fmt.Errorf("DNS query failed: %v", err)
	}

	var spfRecord string
	for _, answer := range resp.Answer {
		if txt, ok := answer.(*dns.TXT); ok {
			for _, txtRecord := range txt.Txt {
				if strings.HasPrefix(txtRecord, "v=spf1") {
					spfRecord = txtRecord
					break
				}
			}
		}
	}

	if spfRecord == "" {
		return false, fmt.Errorf("no SPF record found for domain %s", domain)
	}

	spfParts := strings.Split(spfRecord, " ")
	for _, part := range spfParts {
		if strings.HasPrefix(part, "ip4:") {
			allowedIP := strings.TrimPrefix(part, "ip4:")
			if strings.Contains(allowedIP, "/") {
				if IsIPInCIDR(ip, allowedIP) {
					return true, nil
				}
			} else if ip == allowedIP {
				return true, nil
			}
		} else if strings.HasPrefix(part, "include:") {
			includedDomain := strings.TrimPrefix(part, "include:")
			result, err := CheckSPF(ip, includedDomain)
			if err == nil && result {
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("IP %s not authorized by SPF for domain %s", ip, domain)
}

func CheckDMARC(domain string) (bool, error) {
	policy, err := dmarc.Lookup(domain)
	if err != nil {
		return false, fmt.Errorf("DMARC check failed: %v", err)
	}
	if policy.Policy == dmarc.PolicyReject {
		return false, fmt.Errorf("DMARC reject policy applied for domain: %s", domain)
	}
	return true, nil
}

func IsUsingGmailMX(domain string) (bool, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeMX)

	client := new(dns.Client)
	resp, _, err := client.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return false, fmt.Errorf("DNS query failed: %v", err)
	}

	for _, answer := range resp.Answer {
		if mx, ok := answer.(*dns.MX); ok {
			if strings.HasSuffix(mx.Mx, ".google.com.") {
				return true, nil
			}
		}
	}

	return false, nil
}

func CheckDKIM(email []byte, senderDomain string) (bool, error) {
	result, err := dkim.Verify(&email)
	if err != nil {
		if strings.Contains(err.Error(), "signature has expired") {
			if senderDomain == "gmail.com" {
				return true, nil
			}

			usingMX, err := IsUsingGmailMX(senderDomain)
			if err == nil && usingMX {
				return true, nil
			}

			return false, fmt.Errorf("DKIM signature expired")
		}
		return false, err
	}
	if result != 1 {
		return false, fmt.Errorf("DKIM verification failed")
	}
	return true, nil
}

func IsValidRecipientName(name string) bool {
	return len(name) > 0 && len(name) <= 42 && !strings.ContainsAny(name, " !@#$%^&*()")
}
// ============================================
// EMAIL PROCESSING FUNCTIONS
// ============================================


func ParseEmail(emailData string) (*model.ParsedEmail, error) {
	reader := strings.NewReader(emailData)
	msg, err := mail.ReadMessage(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read MIME message: %w", err)
	}

	subject := msg.Header.Get("Subject")
	contentType := msg.Header.Get("Content-Type")
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Content-Type: %w", err)
	}

	parsedEmail := &model.ParsedEmail{Subject: subject}

	if strings.HasPrefix(mediaType, "multipart/") {
		multipartReader := multipart.NewReader(msg.Body, params["boundary"])
		return parseMultipartEmail(multipartReader, parsedEmail)
	}

	bodyContent, err := io.ReadAll(msg.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading email body: %w", err)
	}

	parsedEmail.Body = string(bodyContent)
	return parsedEmail, nil
}

func parseMultipartEmail(multipartReader *multipart.Reader, parsedEmail *model.ParsedEmail) (*model.ParsedEmail, error) {
	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading multipart part: %w", err)
		}

		contentType := part.Header.Get("Content-Type")
		contentDisposition := part.Header.Get("Content-Disposition")

		if strings.HasPrefix(contentType, "text/") {
			bodyContent, err := io.ReadAll(part)
			if err != nil {
				return nil, fmt.Errorf("error reading body content: %w", err)
			}

			if strings.HasPrefix(contentType, "text/html") || parsedEmail.Body == "" {
				parsedEmail.Body = string(bodyContent)
			}
		}

		if strings.HasPrefix(contentDisposition, "attachment") {
			attachmentData, err := io.ReadAll(part)
			if err != nil {
				return nil, fmt.Errorf("error reading attachment: %w", err)
			}

			parsedEmail.Attachments = append(parsedEmail.Attachments, model.Attachment{
				ContentDisposition: contentDisposition,
				ContentID:          part.Header.Get("Content-ID"),
				ContentType:        contentType,
				Data:               attachmentData,
			})
		}
	}

	return parsedEmail, nil
}



// ============================================
// ENCRYPTION FUNCTIONS
// ============================================

func GeneratePassword(email string) (string, error) {
	hash := sha256.Sum256([]byte(email))
	password := hex.EncodeToString(hash[:])
	return password, nil
}

func EncryptEmail(input, password string) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	if _, err := io.WriteString(gzipWriter, input); err != nil {
		return nil, fmt.Errorf("failed to gzip data: %v", err)
	}
	gzipWriter.Close()

	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	cipherText := gcm.Seal(nonce, nonce, buf.Bytes(), nil)

	return cipherText, nil
}

func DecryptEmail(cipherText []byte, password string) (string, error) {
	key := sha256.Sum256([]byte(password))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", fmt.Errorf("failed to create cipher block: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", fmt.Errorf("cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt data: %v", err)
	}

	var buf bytes.Buffer
	buf.Write(plainText)
	gzipReader, err := gzip.NewReader(&buf)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %v", err)
	}
	defer gzipReader.Close()

	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return "", fmt.Errorf("failed to read decompressed data: %v", err)
	}

	return string(decompressedData), nil
}

func SaveEmailLocally(encryptedEmail []byte) error {
	emailDir := "./email"

	if err := os.MkdirAll(emailDir, 0755); err != nil {
		return fmt.Errorf("failed to create email directory: %w", err)
	}

	filename := fmt.Sprintf("email_%d.txt.gz", time.Now().UnixNano())
	filePath := filepath.Join(emailDir, filename) // ✅ ĐÚNG

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(encryptedEmail)
	return err
}