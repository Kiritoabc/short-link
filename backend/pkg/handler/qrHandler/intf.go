package qrHandler

func NewQrHandler() QrIntf {
	return &qrHandler{}
}

type QrIntf interface {
	GenerateQrCode(url string) ([]byte, error)
}
