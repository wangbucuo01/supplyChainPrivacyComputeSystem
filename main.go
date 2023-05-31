package main

import (
	"supplyChainPrivacyComputeSystem/router"
	"supplyChainPrivacyComputeSystem/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8081")
}