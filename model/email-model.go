package model
import "github.com/meta-node-blockchain/meta-node/types"
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
type ResultData struct {
    Receipt types.Receipt
    Err     error
}