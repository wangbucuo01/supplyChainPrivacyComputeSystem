package service

import (
	"testing"
)

func TestFileRead(t *testing.T) {
	res, err := FileRead("../data/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
