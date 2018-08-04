package xlog

import (
	"fmt"
)

var colorFormat = "\x1b[%dm%s\x1b[0m"

func colorize(colorCode int, str string) string {
	return fmt.Sprintf(colorFormat, colorCode, str)
}

// ColorBlue code 蓝色
var ColorBlue = 34

// Blue color
func Blue(s string) string {
	return colorize(ColorBlue, s)
}

// ColorLightGreen code 天蓝色
var ColorLightGreen = 36

// LightGreen color
func LightGreen(s string) string {
	return colorize(ColorLightGreen, s)
}

// ColorPurple code 淡红
var ColorPurple = 35

// Purple color
func Purple(s string) string {
	return colorize(ColorPurple, s)
}

// ColorWhite code 白色
var ColorWhite = 0

// White coloe
func White(s string) string {
	return colorize(ColorWhite, s)
}

// ColorGray code 灰色
var ColorGray = 37

// Gray color
func Gray(s string) string {
	return colorize(ColorGray, s)
}

// ColorRed code 红色
var ColorRed = 31

// Red color
func Red(s string) string {
	return colorize(ColorRed, s)
}

// ColorGreen code 绿色
var ColorGreen = 32

// Green color
func Green(s string) string {
	return colorize(ColorGreen, s)
}

// ColorYellow code 黄色
var ColorYellow = 33

// Yellow color
func Yellow(s string) string {
	return colorize(ColorYellow, s)
}
