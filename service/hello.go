package service

import (
	"supplyChainPrivacyComputeSystem/utils"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	w := c.Writer
	utils.RespOK(w, "OK", "OK")
}