package router

import (
	"supplyChainPrivacyComputeSystem/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", service.Hello)
	// 上传文件
	r.POST("/supplychain/file/upload", service.Upload)



	// r.GET("/supplychain/rawdata", service.GetRawData)

	return r
}
