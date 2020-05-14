package utils

import (
	"fmt"
	"runtime/debug"
	"time"
)

//Errors 信息输出
func Errors(msg string) {
	printOut("Error", msg)
	debug.PrintStack()
}

//Wrong 信息输出
func Wrong(msg string) {
	printOut("Wrong", msg)
}

//Info 信息输出
func Info(msg string) {
	printOut("Info", msg)
}

//Debug 信息输出
func Debug(msg string) {
	printOut("Debug", msg)
}

func printOut(flag string, msg string) {
	fmt.Printf("%s [%s]: %s", time.Now().Format("20060102150405"), flag, msg)
}
