package main

import (
	"github.com/gin-gonic/gin"
)

//func LoadBalancerMiddleware() {
//	parseUrl, err := url.Parse("http://localhost:8080/ping")
//	if err != nil {
//		panic(err)
//	}
//	rp := httputil.NewSingleHostReverseProxy(parseUrl)
//	http.HandleFunc("/gateway", rp.ServeHTTP)
//	log.Fatal(http.ListenAndServe(":8888", nil))
//}

func main() {
	// 启动 gateway
	//go LoadBalancerMiddleware()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "hello world 0",
		})
	})

	router.Run(":8000")
}
