package test

import (
	"fmt"
	"testing"
	"wwgo/utils"
)

func Test_CaseConversion(t *testing.T) {
	str := "aaabcD"
	str = utils.StringCase(str)
	fmt.Println("str=====", str)
}
