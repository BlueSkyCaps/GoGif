package main

import (
	"GoGif/gof"
	"fmt"
	"os"
	"strconv"
)

const Author = "*** 此程序作者：BlueSkyCaps(芝士为了玩|比尔小贤)"
const Right = "*** 此程序由作者本人制作且免费向外开源，\n*** 若你是从第三方购买而获取此程序，则代表你可能受骗。"

func main() {
	fmt.Println("***")
	fmt.Println(Author)
	fmt.Println(Right)
	fmt.Println("***")
	fmt.Println()

	args := os.Args
	var w, h, dur, order int64 = 0, 0, 0, 0
	var inputRoot = ""
	if len(args) > 1 {
		// args[0]外部命令行调用时的此进程文件名
		w, _ = strconv.ParseInt(args[1], 10, 32)
		h, _ = strconv.ParseInt(args[2], 10, 32)
		dur, _ = strconv.ParseInt(args[3], 10, 32)
		order, _ = strconv.ParseInt(args[4], 10, 32)
		inputRoot = args[5]
	}
	for i := 0; i < len(args); i++ {
		fmt.Println("->->", args[i])
	}
	gof.Run(int(w), int(h), float32(dur), int(order), inputRoot)
	var ce rune
	_, _ = fmt.Scanf("%c", &ce)
	return
}
