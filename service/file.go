package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"supplyChainPrivacyComputeSystem/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	//上传到本地
	UploadLocal(c)
	// TODO 上传到阿里云OSS
	// UploadOSS(c)
}

func UploadLocal(c *gin.Context) {
	w := c.Writer
	req := c.Request
	srcFile, head, err := req.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
		return
	}
	suffix := ".txt"
	ofilName := head.Filename
	temp := strings.Split(ofilName, ".")
	if len(temp) > 1 {
		suffix = "." + temp[len(temp)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstFile, err := os.Create("./data/" + fileName)
	if err != nil {
		utils.RespFail(w, err.Error())
		return
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	url := "./data/" + fileName
	// read file
	fileContent, err := FileRead(url)
	if err != nil {
		// TODO:封装errors
		fmt.Println("read file error:", err)
		return 
	}

	fmt.Println("提交的明文为：", fileContent)
	utils.RespOK(w, fileContent, "上传文件成功!")
}

// 业务逻辑01 读取文件
func FileRead(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []int{}, err
	}

	res := strings.Split(string(data), "\n")
	resInts := []int{}
	for _, v := range res {
		v, _ := strconv.Atoi(v)
		resInts = append(resInts, v)
	}
	return resInts, nil
}
