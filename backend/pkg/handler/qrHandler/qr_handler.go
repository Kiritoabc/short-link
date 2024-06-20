package qrHandler

import "github.com/Kiritoabc/short-link/backend/pkg/utils"

type qrHandler struct {
}

func (q *qrHandler) GenerateQrCode(url string) ([]byte, error) {
	// todo: 其他操作
	return utils.GenerateQrNoImage(url)
}
