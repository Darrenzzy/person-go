package _00_init_code

import (
	"fmt"
	"testing"
)

type test struct {
	IntEs []int
	Lists []int
	K     int
}

var tests = []test{}

//

func Test_upToDayUp(t *testing.T) {
	for k1 := range tests {
		fmt.Println("初始化")

		for k := range tests[k1].IntEs {
			fmt.Println(tests[k1].IntEs[k])

		}
		pre := 1
		fmt.Println("结果：", pre)

	}
}