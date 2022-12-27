package main

import (
	"errors"
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

	//嵌套error
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)

	//errors.Unwrap 函数
	fmt.Println(errors.Unwrap(w))

	//errors.Is 函数的定义，可以解释为：
	//如果 err 和 target 是同一个，那么返回 true。
	//如果 err 是一个 wrapping error，target 也包含在这个嵌套 error 链中的话，也返回 true。
	fmt.Println(errors.Is(w, e))

	//errors.As 函数
	//有了嵌套之后，error的断言也失效了，因此使用errors.As
	var cm *commonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	//defer关键字，保证后面的代码一定被执行，用来修饰方法或函数，常用于成对的操作

	//panic函数，是go的内置函数，运行时的异常，可以接受 interface{} 类型的参数，也就是任何类型的值都可以传递给 panic 函数
	//panic("手动抛出panic")
	//panic 异常是一种非常严重的情况，会让程序中断运行，使程序崩溃，所以如果是不影响程序运行的错误，不要使用 panic，使用普通错误 error 即可。

	//recover函数：配合defer关键字，从panic异常中恢复
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("恢复了panic异常", p)
		}
	}()
	panic("手动抛出panic")
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
