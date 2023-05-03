package main

import "supplyChainPrivacyComputeSystem/router"

func main() {
	r := router.Router()
	r.Run(":8081")
}