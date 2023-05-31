package service

import (
	"strconv"
	"supplyChainPrivacyComputeSystem/model"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateData(c *gin.Context) {
	// 获取明文数据
	plaintext := c.Request.FormValue("plaintext")
	uid, _ := strconv.Atoi(c.Request.FormValue("uid"))
	user := model.FindUserByID(uid)
	if user.UserName == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户不存在",
			"data":    user,
		})
		return
	}
	data_desc := c.Request.FormValue("data_desc")

	// 对数据进行处理
	// 对数据进行加密
	ciphertext := model.EncryptData(plaintext)
	// 生成证明
	// TODO: 不需要
	proof := model.CreateProof(plaintext)

	// 写入数据库
	data := model.Data{}
	data.PlainText = plaintext
	data.CipherText = ciphertext
	data.UploadDate = time.Now().Format("2006-01-02 15:04:05")
	data.State = 0
	data.UID = uid
	data.Proof = proof
	data.DataDesc = data_desc
	model.CreateData(data)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "上传数据成功!",
		"data":    data,
	})
}

func DataList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	query_date := c.Query("query_date")
	data := model.GetDataList(uid, query_date)
	//TODO 分页
	c.JSON(200, gin.H{
		"code":    0,
		"message": "查询数据列表成功!",
		"data":    data,
	})
}

func ComputeData(c *gin.Context) {
	// 聚合人id（非数据上传者id）
	uid, _ := strconv.Atoi(c.Query("uid"))
	// 查询日期 2023-05-31
	query_date := c.Query("query_date")
	res_desc := c.Query("res_desc")
	// 查询到的数据集（根据query_date)
	data := model.GetDataList(0, query_date)
	// 计算结果
	res := model.ComputeData(data)
	result := model.Result{}
	// 写入数据库
	result.UID = uid
	result.Res = res
	result.Date = time.Now().Format("2006-01-02 15:04:05") // 写入当前时间
	result.ResDesc = res_desc
	model.CreateResult(result)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "数据聚合成功!",
		"data":    result,
	})
}

func DataVerify(c *gin.Context) {
	data_id, _ := strconv.Atoi(c.Query("data_id"))
	data := model.FindDataByID(data_id)
	res := model.DataVerify(data)
	if !res {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "数据验证不通过!",
			"data":    data,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "数据验证成功!",
		"data":    data,
	})
}