package test

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
	"path"
	"runtime/debug"
)

var (
	imagesRoot  = "C:/Users/pc/Desktop/gifs/gif"
	gifOutRoot  = "C:/Users/pc/Desktop/"
	imgNmeFiles []string
)

func test() {

	dirInfo, _ := os.ReadDir(imagesRoot)
	for i := 0; i < len(dirInfo); i++ {
		if !dirInfo[i].IsDir() {
			imgNmeFiles = append(imgNmeFiles, dirInfo[i].Name())
		}
	}
	inputGifBoss := &gif.GIF{}
	for i, currentGifName := range imgNmeFiles {
		g, _ := os.Open(path.Join(imagesRoot, currentGifName))
		currentGifImage, err := gif.Decode(g)
		if err != nil {
			println(err.Error())
			debug.PrintStack()
			return
		}
		_ = g.Close()
		fmt.Printf("%v\n", currentGifImage.Bounds())
		inputGifBoss.Image = append(inputGifBoss.Image, currentGifImage.(*image.Paletted))
		inputGifBoss.Delay = append(inputGifBoss.Delay, (i+1)*100)
	}
	inputGifBoss.Config = image.Config{ColorModel: inputGifBoss.Config.ColorModel, Width: 500, Height: 500}
	inputGifBoss.LoopCount = -1
	finalGif, _ := os.OpenFile(path.Join(gifOutRoot, UuidGenerator()+"out.gif"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	err := gif.EncodeAll(finalGif, inputGifBoss)
	if err != nil {
		println(err.Error())
		debug.PrintStack()
		return
	}
	err = finalGif.Close()
	if err != nil {
		println(err.Error())
		debug.PrintStack()
		return
	}
}

func UuidGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
