package demo

import "fmt"

var share string = ""

func accessInner(funcName string) {
	// TODO: 去掉次循环后，有大概率出现 race detector 检测失效情况
	for i := 1; i < 10000; i++ {
		share = ""
	}
	share = "access Share inner demo package, function: " + funcName
	fmt.Println(share)
}

func AccessOuter(pkgName string, funcName string) {
	for i := 1; i < 10000; i++ {
		share = ""
	}
	share = "access ShareOuter in " + pkgName + " package, function: " + funcName
	fmt.Println(share)
}