package model
type ParsedEmail struct {
	Subject     string
	Body        string
	Attachments []Attachment
}

type Attachment struct {
	ContentDisposition string
	ContentID          string
	ContentType        string
	Data               []byte
}