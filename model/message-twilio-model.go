package model

type MessageTwilio struct {
	SmsSid      string `json:"SmsSid"`
	ProfileName string `json:"ProfileName"`
	WaId        string `json:"WaId"`
	SmsStatus   string `json:"SmsStatus"`
	Body        string `json:"Body"`
	To          string `json:"To"`
	From        string `json:"From"`
}
