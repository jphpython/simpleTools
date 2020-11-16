package tools

import (
	"regexp"
)

/**
* 验证身份证号是否合法
*/
func VerifyIdenytifyCode(code string) bool {
	res := false

	switch len(code) {
	case 15://15位
		if m, _ := regexp.MatchString(`^(\d{15})$`, code); m {
			res = true
		}
	case 18://18位
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, code); m {
			res = true
		}
	}

	return res
}