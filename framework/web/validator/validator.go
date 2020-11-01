package validator

import (
	"fmt"
)

// CheckRule is used for constuct error messages slice
func CheckRule(errorArray *[]string, c bool, errMsg string, args ...interface{}) {
	if !c {
		if args != nil {
			*errorArray = append(*errorArray, fmt.Sprintf(errMsg, args...))
		} else {
			*errorArray = append(*errorArray, fmt.Sprintf(errMsg))
		}
	}
}
