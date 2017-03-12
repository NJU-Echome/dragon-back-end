package util

import (
	"fmt"
	"testing"
)

type fromStruct struct {
	List []innerFromStruct
}

type innerFromStruct struct {
	I int
}

type toStruct struct {
	List []innerToStruct
}

type innerToStruct struct {
	I int
}

func TestEmbedSlice(t *testing.T) {
	innerFromlist := []innerFromStruct{innerFromStruct{2}}
	fromList := []fromStruct{fromStruct{innerFromlist}}
	fmt.Printf("fromList: %v\n", fromList)

	toList := []toStruct{}
	fmt.Printf("toList: %v\n", toList)

	PoListToVoList(fromList, &toList)
	fmt.Printf("toList: %v\n", toList)
}
