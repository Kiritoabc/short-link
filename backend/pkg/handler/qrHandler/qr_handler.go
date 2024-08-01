package qrHandler

import "github.com/Kiritoabc/short-link/backend/pkg/utils"

// qrHandler
type qrHandler struct {
}

// GenerateQrCode 创建二维码
func (q *qrHandler) GenerateQrCode(url string) ([]byte, error) {
	// todo: 其他操作
	return utils.GenerateQrNoImage(url)
}
