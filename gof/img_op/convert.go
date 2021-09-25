package img_op

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"strings"
)

// ConvertToGif 将所有图片格式编码为gif图片格式并存储到指定位置中
// decodeImagePaths是要进行编码成gif格式的原图片绝对路径列表,他们的文件名不应该重复。outImageRoot是转换后的存储的绝对路径
func ConvertToGif(decodeImagePaths []string, outImageRoot string) {
	var c rune
	for _, v := range decodeImagePaths {
		f, err := os.Open(v)
		decodeImg, _, err := image.Decode(f)
		if err != nil {
			println("error in ConvertToGif:", err.Error())
			_, _ = fmt.Scanf("%c", &c)
			os.Exit(1)
		}
		err = f.Close()
		if err != nil {
			println("error in ConvertToGif:", err.Error())
			_, _ = fmt.Scanf("%c", &c)
			os.Exit(1)
		}
		endNameIndex := strings.LastIndex(v, string(os.PathSeparator)) + 1
		fileName := v[endNameIndex:]
		formatStrIndex := strings.LastIndex(fileName, ".") + 1
		outName := strings.Replace(fileName, fileName[formatStrIndex:], "gif", -1)
		f, err = os.OpenFile(path.Join(outImageRoot, outName), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			println("error in ConvertToGif:", err.Error())
			_, _ = fmt.Scanf("%c", &c)
			os.Exit(1)
		}
		var p = &gif.Options{NumColors: 256}
		err = gif.Encode(f, decodeImg, p)
		if err != nil {
			println("error in ConvertToGif Encode:", err.Error())
			_, _ = fmt.Scanf("%c", &c)
			os.Exit(1)
		}
		err = f.Close()
		if err != nil {
			println("error in ConvertToGif final close write *f, operate directly return:", err.Error())
			return
		}
	}
}
