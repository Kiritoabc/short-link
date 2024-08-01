package qrHandler

// NewQrHandler 创建一个QrHandler
func NewQrHandler() QrIntf {
	return &qrHandler{}
}

// QrIntf QrHandler接口
type QrIntf interface {
	GenerateQrCode(url string) ([]byte, error)
}
