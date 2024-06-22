package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func reverseProxyHandler(target *url.URL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ServeHTTP(w, r)
	}
}

func ProxyHandler(ctx *gin.Context) {
	targetURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func main() {
	engine := gin.Default()

	engine.Use(ProxyHandler)
	log.Println("Starting reverse proxy server on :8000")

	log.Fatalln(engine.Run(":8080"))

}
