package utils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code" example:"7"`
	Data interface{} `json:"data" `
	Msg  string      `json:"msg" example:"successed"`
}

const (
	ERROR   = 400
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func QrResult(qrCode []byte, c *gin.Context) {
	pngBase64Str := base64.StdEncoding.EncodeToString(qrCode)
	htmlResp := `
       <html>
           <head></head>
           <body>
               <img src="data:image/png;base64,` + pngBase64Str + `" alt="QR Code">
           </body>
       </html>
   `
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlResp))
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithQrCode(qrCode []byte, c *gin.Context) {
	QrResult(qrCode, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
