package haserr

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// 判断有无错误,并返回true或者false
func HasErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// 判断有无错误,并返回true或者false,有错误时打印错误
func Println(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		file = file[strings.LastIndex(file, `/`)+1:]
		log.Printf("%v,第%v行,错误类型:%v", file, line, err)
		return true
	}
	return false
}

//判断有无错误,并返回true或者false,有错误时调用log.Fatal退出
func Fatal(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		file = file[strings.LastIndex(file, `/`)+1:]
		log.Fatalf("%v,第%v行,错误类型:%v", file, line, err)
		return true
	}
	return false
}

//判断有无错误,并返回true或者false,有错误时调用panic退出
func Panic(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		file = file[strings.LastIndex(file, `/`)+1:]
		panic(fmt.Sprintf("%v,第%v行,错误类型:%v", file, line, err))
		return true
	}
	return false
}
