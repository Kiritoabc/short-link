package midware

import (
	"fmt"
	"github.com/Kiritoabc/short-link/gateway/cmd/pkg/config"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

var (
	Ip              = []string{}
	poolIndex int64 = 0
)

// 初始化我们需要的IP

// LoadBalancerMiddleware 负载均衡的中间件
func LoadBalancerMiddleware(ctx *gin.Context) {
	// 开启代理==> 不是开启代理，而是转发
	//poolLoadBalance(ctx)
	proxyModel := config.ProxyModel.Value

	proxyLoadBalance(ctx, ProxyModel(ctx, proxyModel))
	return
}

type backend struct {
	ctx *gin.Context
}

// 参数函数（用于实现 Option 接口的类型）
type optionFunc func(*backend)

func (f optionFunc) apply(o *backend) {
	f(o)
}

type ProxyOption interface {
	apply(*backend)
}

// 负载均衡
func proxyLoadBalance(ctx *gin.Context, opts ...ProxyOption) {
	// 获取model
	o := &backend{
		ctx: ctx,
	}
	for _, opt := range opts {
		opt.apply(o)
	}
}

// WithPoolBL 轮询
func WithPoolBL(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		n := len(Ip)
		if poolIndex >= int64(n) {
			atomic.StoreInt64(&poolIndex, 0)
		}
		// 拿到代理的Ip地址
		proxyIp := Ip[poolIndex]
		atomic.AddInt64(&poolIndex, 1)
		proxyHandler(proxyIp, ctx)
	})
}

// WithRandBL 随机
func WithRandBL(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		// 随机
		randIndex := rand.Intn(len(Ip))

		proxyIp := Ip[randIndex]
		proxyHandler(proxyIp, ctx)
	})
}

func proxyHandler(proxyIp string, ctx *gin.Context) {
	target, err := url.Parse(proxyIp)
	if err != nil {
		fmt.Printf("parse proxy url error: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func ProxyModel(ctx *gin.Context, model string) ProxyOption {
	switch model {
	case "pool":
		return WithPoolBL(ctx)
	case "rand":
		return WithRandBL(ctx)
	default:
		return WithRandBL(ctx)
	}
}
