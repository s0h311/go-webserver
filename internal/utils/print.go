package utils

import "fmt"

var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintlnColorful(color string, text string) {
	var Reset = "\033[0m"
	fmt.Println(color + text + Reset)
}

func SprintfColorful(color string, text string, vars ...any) string {
	return fmt.Sprintf(color+text, vars...)
}

func PrintDivider() {
	PrintlnColorful(Cyan, "======================")
}

func GetDivider() string {
	return SprintfColorful(Cyan, "======================")
}
