package web

import (
	"context"
	"log"
	"net/http"
	"time"

	config "github.com/GDEIDevelopers/Yiwei_Wechat_app_server/config"
	"github.com/dgraph-io/badger/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Web struct {
	server    *http.Server
	waitUntil chan struct{}
	c         *config.Config
	db        *gorm.DB
	cache     *badger.DB
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
	// 登录
	r.POST("/login", w.Login)
	// 使用默认AccessToken鉴权的API接口
	adminGroup := r.Group("/admin").Use(w.AdminAuth)
	// Admin模块
	adminGroup.GET("/reserve/all", w.AdminAllReserve)
	adminGroup.GET("/reserve/today", w.AdminTodayReserve)
	adminGroup.GET("/reserve/:id/detail", w.AdminGetOneReserveDetail)
	adminGroup.PATCH("/reserve/:id/modify", w.AdminModifyReserveDetail)
	adminGroup.DELETE("/reserve/:id/delete", w.AdminDeleteReserve)
	adminGroup.PUT("/reserve/create", w.AdminCreateReserve)

	adminGroup.GET("/member/all/list", w.AdminListAllMember)
	adminGroup.GET("/member/:id/list", w.AdminListOneMember)
	adminGroup.PATCH("/member/:id/modfiy", w.AdminModifyMemebr)
	adminGroup.DELETE("/member/:id/delete", w.AdminRemoveMember)
	adminGroup.PUT("/member/create", w.AdminMemberCreate)

	// 预约模块
	userGroup := r.Group("/user").Use(w.UserAuth)
	userGroup.GET("/reserve/all", w.UserListAllReserve)
	userGroup.GET("/reserve/:id/detail", w.UserGetOneReserveDetail)
	userGroup.PATCH("/reserve/:id/modify", w.UserModifyReserveDetail)
	userGroup.DELETE("/reserve/:id/delete", w.UserDeleteReserve)
	userGroup.PUT("/reserve/create", w.UserCreateReserve)

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
