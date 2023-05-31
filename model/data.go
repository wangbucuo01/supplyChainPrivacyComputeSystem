package model

import (
	"fmt"
	"strconv"
	"strings"
	"supplyChainPrivacyComputeSystem/algorithm"
	"supplyChainPrivacyComputeSystem/utils"

	"math/big"

	"gorm.io/gorm"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

type Data struct {
	ID         int    `json:"id"`
	PlainText  string `json:"plain_text"`
	CipherText string `json:"cipher_text"`
	UploadDate string `json:"upload_date"`
	State      int    `json:"state"`
	UID        int    `json:"uid"`
	Proof      string `json:"proof"`
	DataDesc   string `json:"data_desc"`
}

func (d Data) TableName() string {
	return "data"
}

// 只加密，不计算
func EncryptData(plaintext string) string {
	M, _ := strconv.Atoi(plaintext)
	// Generate a 128-bit private key.
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
	// encrypt this
	m := new(big.Int).SetInt64(int64(M))
	c, _ := algorithm.Encrypt(&privKey.PublicKey, m.Bytes())
	r := ""
	for j := 0; j < len(c); j++ {
		r += strconv.Itoa(int(c[j]))
		r += " "
	}
	return r
}

type CubicCircuit struct {
	X frontend.Variable `gnark:"x"`       // 输入值
	Y frontend.Variable `gnark:",public"` // 最小值
	Z frontend.Variable `gnark:",public"` // 最大值
}

// Define declares the circuit constraints
// x < 100  x > 0
func (circuit *CubicCircuit) Define(api frontend.API) error {
	api.AssertIsEqual(api.Cmp(circuit.X, circuit.Y), 1)
	api.AssertIsEqual(api.Cmp(circuit.Z, circuit.X), 1)
	return nil
}

// 只生成证明，不验证
func CreateProof(plaintext string) string {
	x, _ := strconv.Atoi(plaintext)
	// witness definition
	assignment := CubicCircuit{
		X: x,
		Y: 0,
		Z: 100,
	}
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	proofStr := fmt.Sprintf("%v", witness)
	return proofStr
}

func CreateData(data Data) *gorm.DB {
	return utils.DB.Create(&data)
}

func GetDataList(uid int, query_date string) []*Data {
	var data []*Data
	if uid != 0 && query_date != "" {
		utils.DB.Raw("SELECT * FROM data WHERE uid = ? and DATE(upload_date) = ?", uid, query_date).Scan(&data)
		// utils.DB.Where("query_date = ?", uid, query_date).Find(&data)
	} else if uid != 0 {
		utils.DB.Where("uid = ?", uid).Find(&data)
	} else if query_date != "" {
		// utils.DB.Where("query_date = ?", query_date).Find(&data)
		utils.DB.Raw("SELECT * FROM data WHERE DATE(upload_date) = ?", query_date).Scan(&data)
	} else {
		utils.DB.Find(&data)
	}
	return data
}

func ComputeData(data []*Data) string {
	fileContentEnc := [][]byte{}
	// 计算
	for _, dv := range data {
		c := dv.CipherText
		fileContent := []byte{}
		for _, b := range strings.Split(c, " ") {
			bi, _ := strconv.Atoi(b)
			fileContent = append(fileContent, byte(bi))
		}
		fileContentEnc = append(fileContentEnc, fileContent[:len(fileContent)-1])
	}
	// fmt.Println(fileContentEnc)
	// 解密
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
	sum := new(big.Int).SetInt64(int64(0))
	sumC, _ := algorithm.Encrypt(&privKey.PublicKey, sum.Bytes())
	for i := 0; i < len(fileContentEnc); i++ {
		sumC = algorithm.AddCipher(&privKey.PublicKey, fileContentEnc[i], sumC)
	}
	decryptedAddition, _ := algorithm.Decrypt(&privKey, sumC)
	res := new(big.Int).SetBytes(decryptedAddition).String()
	return res
}

func FindDataByID(data_id int) Data {
	data := Data{}
	utils.DB.Where("id = ?", data_id).First(&data)
	return data
}

func DataVerify(data Data) bool {
	// compiles our circuit into a R1CS
	var circuit CubicCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	// groth16 zkSNARK: Setup
	pk, vk, _ := groth16.Setup(ccs)

	// 数据解密
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

	c := data.CipherText
	fileContent := []byte{}
	for _, b := range strings.Split(c, " ") {
		bi, _ := strconv.Atoi(b)
		fileContent = append(fileContent, byte(bi))
	}

	decryptedData, _ := algorithm.Decrypt(&privKey, fileContent[:len(fileContent)-1])
	assignment := CubicCircuit{
		X: int(decryptedData[0]),
		Y: 0,
		Z: 100,
	}
	// res := new(big.Int).SetBytes(decryptedAddition).String()

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()

	// groth16: Prove & Verify
	verifyRes := true
	proof, _ := groth16.Prove(ccs, pk, witness)
	defer func() {
		if err := recover(); err != nil {
			verifyRes = false
		}
	}()
	groth16.Verify(proof, vk, publicWitness) 
	return verifyRes
}
