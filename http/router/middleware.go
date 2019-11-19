package router

import (
	"net/http"

	"github.com/infraboard/mcube/http/auth"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"
)

// Middleware 中间件的函数签名
type Middleware interface {
	Handler(http.Handler) http.Handler
}

// MiddlewareFunc is an adapter to allow the use of ordinary functions as Negroni handlers.
// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler object that calls f.
type MiddlewareFunc func(http.Handler) http.Handler

// Handler wrappe for function
func (h MiddlewareFunc) Handler(next http.Handler) http.Handler {
	return h(next)
}

// NewAutherMiddleware 初始化一个认证中间件
func NewAutherMiddleware(auther auth.Auther) Middleware {
	return &AutherMiddleware{
		auther: auther,
	}
}

// AutherMiddleware 认证中间件
type AutherMiddleware struct {
	auther auth.Auther
}

// Handler 实现中间件
func (m *AutherMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		// 使用auther进行认证
		authInfo, err := m.auther.Auth(r.Header)
		if err != nil {
			response.Failed(wr, err)
			return
		}

		rc := context.GetContext(r)
		rc.AuthInfo = authInfo
		context.WithContext(r, rc)

		// next handler
		next.ServeHTTP(wr, r)
	})
}
