package utils

import "github.com/skip2/go-qrcode"

const size = 256

// GenerateQrNoImage /* 生成二维码，非图片格式，使用的是byte
func GenerateQrNoImage(url string) ([]byte, error) {
	return qrcode.Encode(url, qrcode.Medium, size)
}

// GenerateQr /* 在本地生成图片
func GenerateQr(url string) error {
	return qrcode.WriteFile(url, qrcode.Medium, size, "qr.png")
}
