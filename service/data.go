package service

import (
	"supplyChainPrivacyComputeSystem/model"

	"github.com/gin-gonic/gin"
)

func GetRawData(c *gin.Context) {
	data, _ := model.GetRawData()
	// TODO:分页
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取原始数据成功",
		"data":    data,
	})
}
