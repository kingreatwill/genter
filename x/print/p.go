package main

import (
	"github.com/fatih/color"
)

func main() {
	//fmt.Print("\x1b[4;30;46m") // 设置颜色样式
	//fmt.Print("Hello World")   // 打印文本内容
	//fmt.Println("\x1b[0m")     // 样式结束符, 清楚之前的显示属性

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")
}

// https://github.com/fatih/color
// \x1b[4;30;46m 由 3 部分组成
// \x1b[ ：控制序列导入器
// 4;30;46：由分号分隔的参数。4 表示下划线，30 表示设置前景色黑色，46 表示设置背景颜色青色
// m ：最后一个字符（总是一个字符）

// \x1b[0m 包含 0 用来表示清除显示属性。
