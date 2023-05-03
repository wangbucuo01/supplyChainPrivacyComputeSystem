package encrypart

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"

	paillier "github.com/roasbeef/go-go-gadget-paillier"
)

func compute(path string) string {
	// read file
	fileContent, err := fileRead(path)
	if err != nil {
		// TODO:封装errors
		fmt.Println("read file error:", err)
		return ""
	}

	fmt.Println("提交的明文为：", fileContent)

	// Generate a 128-bit private key.
	privKey, _ := paillier.GenerateKey(rand.Reader, 128)

	// read every numbers and encrypt this
	fileContentEnc := [][]byte{}
	sum := new(big.Int).SetInt64(int64(0))
	sumC, _ := paillier.Encrypt(&privKey.PublicKey, sum.Bytes())
	for i := 0; i < len(fileContent); i++ {
		m := new(big.Int).SetInt64(int64(fileContent[i]))
		c, _ := paillier.Encrypt(&privKey.PublicKey, m.Bytes())
		fileContentEnc = append(fileContentEnc, c)

		sumC = paillier.AddCipher(&privKey.PublicKey, c, sumC)
	}

	fmt.Println("加密后的密文数据为：", fileContentEnc)
	decryptedAddition, _ := paillier.Decrypt(privKey, sumC)
	fmt.Println("解密后的计算结果为: ", new(big.Int).SetBytes(decryptedAddition).String())
	res := new(big.Int).SetBytes(decryptedAddition).String()
	return res
}

// 业务逻辑01 读取文件
func fileRead(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []int{}, err
	}

	res := strings.Split(string(data), "\r\n")
	resInts := []int{}
	for _, v := range res {
		v, _ := strconv.Atoi(v)
		resInts = append(resInts, v)
	}
	return resInts, nil
}
