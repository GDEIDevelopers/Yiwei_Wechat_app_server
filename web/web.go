package web

import (
	"context"
	"log"
	"net/http"
	"time"

	config "github.com/GDEIDevelopers/Yiwei_Wechat_app_server/config"
	"github.com/gin-gonic/gin"
)

type Web struct {
	server    *http.Server
	waitUntil chan struct{}
	c         *config.Config
}

// 关闭HTTP服务器函数
func (w *Web) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := w.server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
		return
	}
	<-w.waitUntil
}

// 设置HTTP路由
func (w *Web) setUpRouter(r *gin.Engine) {
	r.POST("")
}

// 运行HTTP服务器
func (s *Web) Run() {
	defer close(s.waitUntil)
	if s.c.Cert != "" && s.c.Key != "" {
		if err := s.server.ListenAndServeTLS(s.c.Cert, s.c.Key); err != http.ErrServerClosed {
			log.Printf("%v", err)
			return
		}
	} else {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("%v", err)
			return
		}
	}
}

func New(c *config.Config) *Web {
	r := gin.Default()
	s := &Web{
		server: &http.Server{
			Handler: r,
			Addr:    c.HTTPAddr,
		},
		c: c,
	}
	s.setUpRouter(r)
	go s.Run()
	return s
}
