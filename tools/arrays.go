package tools

import (
	"fmt"
)

/**
* implode  函数 
*/
func ArrayToStr(split string,arr []string) string {
	var res string = ""

	if len(arr) < 1 {
		return res
	}

	for _,i := range arr {
		res = fmt.Sprintf("%s%s%s", res, split, i)
	}

	return res
}