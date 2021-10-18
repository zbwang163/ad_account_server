package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zbwang163/ad_account_server/adapter/http_adapter"
)

func Register(r *gin.Engine) {
	g := r.Group("/ad_account_server/api")
	g.Use(http_adapter.UserInfoMiddleware, http_adapter.LogIdMiddleware, http_adapter.ResponseMiddleware)

	userRouter := g.Group("/user")
	userRouter.POST("/login", http_adapter.HandlerFunc(http_adapter.Login))
	userRouter.GET("/login/get_capture", http_adapter.HandlerFunc(http_adapter.GetCaptureImage))
	userRouter.POST("/register", http_adapter.HandlerFunc(http_adapter.Register))
	userRouter.POST("/register/email_capture", http_adapter.HandlerFunc(http_adapter.SendEmailCapture))
	userRouter.POST("/info", http_adapter.HandlerFunc(http_adapter.GetUserInfo))
	userRouter.POST("/update", http_adapter.UserPrivilegeManagement, http_adapter.HandlerFunc(http_adapter.UpdateUserInfo))

	policyRouter := g.Group("/policy")
	policyRouter.POST("/add", http_adapter.HandlerFunc(http_adapter.AddPolicy))
	policyRouter.POST("/remove", http_adapter.HandlerFunc(http_adapter.RemovePolicy))
}
