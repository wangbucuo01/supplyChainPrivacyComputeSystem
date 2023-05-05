package model

import (
	"io/ioutil"
	"strconv"
	"strings"
)

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
