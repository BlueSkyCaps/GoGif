package main

import (
	"fmt"
	"sort"
)

var (
	imgNmeFiles []string
)

func main() {
	imgNmeFiles = []string{"1.gif", "2.gif", "10.gif"}
	fmt.Printf("%v", imgNmeFiles)
	sort.Slice(imgNmeFiles, func(i, j int) bool {
		return imgNmeFiles[i] < imgNmeFiles[j]
	})
	fmt.Printf("%v", imgNmeFiles)

}
