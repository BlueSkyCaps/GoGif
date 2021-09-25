package gof

import (
	"GoGif/common"
	"GoGif/gof/img_op"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const separator = string(os.PathSeparator)

// 转换图片格式的临时文件夹
const tmpConvertFolder = "gof-e2c2ea85-276b-465b-9c75-7715f6cba9a8"

var (
	imagesInputRoot string // 要生成gif动图的输入文件夹
	gifOutRoot, _   = os.UserHomeDir()
	imgNameFiles    []string
	// MatchImageFormat 匹配以常见图像文件格式为后缀的正则表达模式
	MatchImageFormat = fmt.Sprintf("%v.*(%v.png|%v.gif|%v.jpg|%v.jpeg)$", separator, separator, separator, separator, separator)
)

func init() {
	if runtime.GOOS == "windows" {
		gifOutRoot = path.Join(gifOutRoot, "desktop")
	}
}
func Run() {
	fmt.Println("请输入你想要用于制作动图Gif的图片所在的文件夹路径。(直接粘贴路径并回车即可)")
	_, e := fmt.Scanf("%s", &imagesInputRoot)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(imagesInputRoot)
	fi, e := os.Stat(imagesInputRoot)
	if e != nil || !fi.IsDir() {
		println("不存在输入的这个文件夹哦(⊙o⊙)？检查：你输入的路径是文件夹吗，文件夹存在吗？")
		return
	}
	dirInfo, _ := os.ReadDir(imagesInputRoot)

	for i := 0; i < len(dirInfo); i++ {
		if !dirInfo[i].IsDir() {
			if common.MatchRegexString(MatchImageFormat, dirInfo[i].Name()) {
				imgNameFiles = append(imgNameFiles, dirInfo[i].Name())
			}
		}
	}

	if len(imgNameFiles) <= 0 {
		println("文件夹中没有一张图片？检查：文件夹有至少一张图片格式的文件吗？制作动态gif,你应该不止一张源图片。")
		return
	}
	// 把输入文件夹下要组合的图片素材的绝对路径存储在一个列表
	var inputImageAbsLists []string
	for i := 0; i < len(imgNameFiles); i++ {
		abs, _ := filepath.Abs(path.Join(imagesInputRoot, imgNameFiles[i]))
		inputImageAbsLists = append(inputImageAbsLists, abs)
	}

	// 输入文件夹拼接临时转换文件夹成为新路径 并生成这个新路径(临时转换文件夹)
	imagesInputRoot, _ = filepath.Abs(path.Join(imagesInputRoot, tmpConvertFolder))

	common.CreateFolder(imagesInputRoot)

	// 开始将要组合的连续图片先统一编码为gif格式，并且存放在临时转换文件夹中
	img_op.ConvertToGif(inputImageAbsLists, imagesInputRoot)

	// 把原先存储的输入图片素材文件名统一改为.gif后缀,因为有可能素材不是gif格式后缀,统一编码后文件名已经统一为{1|2..}.gif
	for i := 0; i < len(imgNameFiles); i++ {
		formatEnd := strings.LastIndex(imgNameFiles[i], ".") + 1
		formatStr := imgNameFiles[i][formatEnd:]
		imgNameFiles[i] = strings.Replace(imgNameFiles[i], formatStr, "gif", -1)
	}

	// 将imgNameFiles的素材文件名列表重新按升序排列(若设置了倒序，则倒序)
	common.SortStringSlice(imgNameFiles, false)

	// 使用此方法将imagesInputRoot临时文件夹路径下的素材图片(已经全为gif格式)生成为gifOutRoot路径下的gif动图
	img_op.OpGifFileToGifDone(imagesInputRoot, imgNameFiles, gifOutRoot)

	// 生成完毕，删除临时文件夹
	common.RemoveFolder(imagesInputRoot)
}
