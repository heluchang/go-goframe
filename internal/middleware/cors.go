package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddlewareCORS 完整的跨域中间件配置
func MiddlewareCORS(r *ghttp.Request) {
	// GF内置默认跨域配置（包含所有必要的头）
	r.Response.CORSDefault()
	// 额外：允许自定义响应头（比如Content-Type）
	r.Response.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	// 继续执行后续中间件/控制器
	r.Middleware.Next()
}
