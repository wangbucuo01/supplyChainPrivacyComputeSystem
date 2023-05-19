package model

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Data struct {
	ID         int    `json:"id"`
	PlainText  string `json:"plaintext"`
	CipherText string `json:"ciphertext"`
	UploadDate string `json:"upload_date"`
	State      int    `json:"state"`
	UID        int    `json:"uid"`
	Proof      string `json:"proof"`
	DataDesc   string `json:"data_desc"`
}

func (d Data) TableName() string {
	return "data"
}

type DataBasic []int

func GetRawData() (DataBasic, error) {
	data, err := ioutil.ReadFile("path")
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
