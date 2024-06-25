package midware

import (
	"fmt"
	"github.com/Kiritoabc/short-link/gateway/cmd/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

var (
	Ip        []string
	WeightIps []WeightNode
)

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

// WithPoolLB 轮询
func WithPoolLB(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		proxyHandler(poolServer.getNode(), ctx)
	})
}

// WithRandLB 随机
func WithRandLB(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		proxyHandler(randServer.getNode(), ctx)
	})
}

// WithRandPoolLB 随机+轮询
func WithRandPoolLB(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		proxyHandler(randPoolServer.getNode(), ctx)
	})
}

// WithWightPoolBl 加权轮训
func WithWightPoolBl(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		proxyHandler(weightPoolServer.getNode(), ctx)
	})
}

// WithHashLb hash
func WithHashLb(ctx *gin.Context) ProxyOption {
	return optionFunc(func(o *backend) {
		proxyHandler(hashRing.getNode(ctx.RemoteIP()), ctx)
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
		return WithPoolLB(ctx)
	case "rand":
		return WithRandLB(ctx)
	case "randAndPool":
		return WithRandPoolLB(ctx)
	case "weightPool":
		return WithWightPoolBl(ctx)
	case "hash":
		return WithHashLb(ctx)
	default:
		fmt.Println("proxy model error")
		return WithRandLB(ctx)
	}
}
