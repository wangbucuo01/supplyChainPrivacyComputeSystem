package main

import (
	"fmt"
	"math/big"

	"supplyChainPrivacyComputeSystem/algorithm"
)

func compute(path string) string {
	// read file
	fileContent := []int{10, 20, 30, 40, 50, 60}

	fmt.Println("提交的明文为：", fileContent)

	// Generate a 128-bit private key.
	// privKey, _ := algorithm.GenerateKey(rand.Reader, 128)
	var privKey algorithm.PrivateKey
	s1 := "230248999022113540210087686017020063423"
	n1 := new(big.Int)
	fmt.Sscan(s1, n1)
	privKey.PublicKey.N = n1
	s2 := "230248999022113540210087686017020063424"
	n2 := new(big.Int)
	fmt.Sscan(s2, n2)
	privKey.PublicKey.G = n2
	s3 := "53014601550685241995926887475913498072750990676600102320512300914922942476929"
	n3 := new(big.Int)
	fmt.Sscan(s3, n3)
	privKey.PublicKey.NSquared = n3
	s4 := "16292079216082402919"
	n4 := new(big.Int)
	fmt.Sscan(s4, n4)
	privKey.P = n4
	s5 := "265431845183104204423962053249059720561"
	n5 := new(big.Int)
	fmt.Sscan(s5, n5)
	privKey.PP = n5
	s6 := "16292079216082402918"
	n6 := new(big.Int)
	fmt.Sscan(s6, n6)
	privKey.Pminusone = n6
	s7 := "14132573010989770217"
	n7 := new(big.Int)
	fmt.Sscan(s7, n7)
	privKey.Q = n7
	s8 := "199729619910956459810731731422460227089"
	n8 := new(big.Int)
	fmt.Sscan(s8, n8)
	privKey.QQ = n8
	s9 := "14132573010989770216"
	n9 := new(big.Int)
	fmt.Sscan(s9, n9)
	privKey.Qminusone = n9
	s10 := "6964544395643021245"
	n10 := new(big.Int)
	fmt.Sscan(s10, n10)
	privKey.Pinvq = n10
	s11 := "8028750950694167162"
	n11 := new(big.Int)
	fmt.Sscan(s11, n11)
	privKey.Hp = n11
	s12 := "7168028615346748972"
	n12 := new(big.Int)
	fmt.Sscan(s12, n12)
	privKey.Hq = n12
	s13 := "230248999022113540210087686017020063423"
	n13 := new(big.Int)
	fmt.Sscan(s13, n13)
	privKey.N = n13
	fmt.Println(privKey)

	// read every numbers and encrypt this
	fileContentEnc := [][]byte{}
	fileContentEnc = [][]byte{{10, 251, 24, 149, 80, 239, 186, 7, 59, 119, 103, 43, 82, 150, 36, 93, 113, 126, 196, 251, 185, 122, 47, 246, 20, 5, 234, 176, 118, 210, 124, 125},
		{110, 69, 14, 31, 40, 62, 160, 18, 245, 55, 144, 109, 143, 94, 66, 117, 76, 104, 30, 28, 239, 62, 55, 62, 69, 215, 33, 213, 167, 109, 152, 47},
		{41, 62, 76, 99, 62, 211, 164, 140, 195, 29, 141, 114, 200, 174, 159, 216, 92, 243, 228, 62, 218, 155, 198, 235, 1, 253, 94, 49, 26, 225, 118, 201}}
	sum := new(big.Int).SetInt64(int64(0))
	sumC, _ := algorithm.Encrypt(&privKey.PublicKey, sum.Bytes())
	for i := 0; i < len(fileContentEnc); i++ {
		// m := new(big.Int).SetInt64(int64(fileContent[i]))
		// c, _ := algorithm.Encrypt(&privKey.PublicKey, m.Bytes())
		// fileContentEnc = append(fileContentEnc, c)
		// // r := ""
		// for j := 0; j < len(c); j++ {
		// 	r += strconv.Itoa(int(c[j]))
		// 	r += " "
		// }
		// fmt.Println("c:", c)
		// fmt.Println("r:", r)
		d,_ := algorithm.Decrypt(&privKey, fileContentEnc[i])
		fmt.Println("***", d)
		sumC = algorithm.AddCipher(&privKey.PublicKey, fileContentEnc[i], sumC)
	}
	fmt.Println("加密后的密文数据为：", fileContentEnc)
	decryptedAddition, _ := algorithm.Decrypt(&privKey, sumC)
	fmt.Println("解密后的计算结果为: ", new(big.Int).SetBytes(decryptedAddition).String())
	res := new(big.Int).SetBytes(decryptedAddition).String()
	return res
}

func main() {
	compute("./test.txt")
}
