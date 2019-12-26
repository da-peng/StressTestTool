package utils

import (
	"fmt"
	"runtime/debug"
	"time"
)

//Errors msg stdout
func Errors(msg string) {
	printOut("Error", msg)
	debug.PrintStack()
}

//Wrong msg stdout
func Wrong(msg string) {
	printOut("Wrong", msg)
}

//Info msg stdout
func Info(msg string) {
	printOut("Info", msg)
}

//Debug msg stdout
func Debug(msg string) {
	printOut("Debug", msg)
}

func printOut(flag string, msg string) {
	fmt.Printf("%s [%s]: %s", time.Now().Format("20060102150405"), flag, msg)
}
