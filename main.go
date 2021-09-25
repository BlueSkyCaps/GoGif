package main

import (
	"GoGif/gof"
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	var w, h, dur, order int64 = 0, 0, 0, 0
	var inputRoot = ""
	if len(args) > 1 {
		w, _ = strconv.ParseInt(args[0], 10, 32)
		h, _ = strconv.ParseInt(args[1], 10, 32)
		dur, _ = strconv.ParseInt(args[2], 10, 32)
		order, _ = strconv.ParseInt(args[3], 10, 32)
		inputRoot = args[4]
	}
	gof.Run(int(w), int(h), float32(dur), int(order), inputRoot)
	var ce rune
	_, _ = fmt.Scanf("%c", &ce)
	return
}
