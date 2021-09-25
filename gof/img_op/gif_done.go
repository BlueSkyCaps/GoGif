package img_op

import (
	"GoGif/common"
	"fmt"
	"image"
	"image/gif"
	"os"
	"path"
	"runtime/debug"
)

var ce rune

// OpGifFileToGifDone 将带有素材图片的文件夹路径以及此目录下所有gif图片的文件名输出最终gifOutRoot文件路径的gif动图
// imgNameFiles必须在传递之前确保不为空且图片格式皆为gif，imagesInputRoot必须是imgNameFiles元素的有效上级路径
func OpGifFileToGifDone(imagesInputRoot string, imgNameFiles []string, gifOutRoot string) {

	inputGifBoss := &gif.GIF{}
	for _, currentGifName := range imgNameFiles {
		g, _ := os.Open(path.Join(imagesInputRoot, currentGifName))
		currentGifImage, err := gif.Decode(g)
		if err != nil {
			println(err.Error())
			fmt.Scanf("%c", &ce)
			debug.PrintStack()
			os.Exit(1)
		}
		_ = g.Close()
		fmt.Printf("%v\n", currentGifImage.Bounds())
		inputGifBoss.Image = append(inputGifBoss.Image, currentGifImage.(*image.Paletted))
		inputGifBoss.Delay = append(inputGifBoss.Delay, 100)
	}
	inputGifBoss.Config = image.Config{ColorModel: inputGifBoss.Config.ColorModel, Width: 200, Height: 200}
	inputGifBoss.LoopCount = 0
	finalGif, _ := os.OpenFile(path.Join(gifOutRoot, common.UuidGenerator()+"out.gif"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	err := gif.EncodeAll(finalGif, inputGifBoss)
	if err != nil {
		println(err.Error())
		fmt.Scanf("%c", &ce)
		debug.PrintStack()
		os.Exit(1)
	}
	err = finalGif.Close()
	if err != nil {
		println(err.Error())
		fmt.Scanf("%c", &ce)
		debug.PrintStack()
		os.Exit(1)
	}
}
