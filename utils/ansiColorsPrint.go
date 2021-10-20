package utils

import "fmt"

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Color("\033[1;30m")
	Red     = Color("\033[1;31m")
	Green   = Color("\033[1;32m")
	Yellow  = Color("\033[1;33m")
	Purple  = Color("\033[1;34m")
	Magenta = Color("\033[1;35m")
	Teal    = Color("\033[1;36m")
	White   = Color("\033[1;37m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return colorString + fmt.Sprint(args...) + "\033[0m"
	}
	return sprint
}
