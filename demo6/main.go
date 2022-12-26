package main

import (
	"fmt"
)

func main() {
	sum, err := add(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}

	//error的断言
	sum1, err1 := add(-1, 2)
	if cm, ok := err1.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum1)
	}
}

//主动抛异常
func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		//return 0, errors.New("a或者b不能为负数")
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a或者b不能为负数"}
	} else {
		return a + b, nil
	}
}

//自定义异常的结构体
type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

// commonError实现Error接口
func (ce *commonError) Error() string {
	return ce.errorMsg
}
