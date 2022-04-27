package gof

import (
	"GoGIf/common"
	"GoGIf/gof/img_op"
	"bufio"
	"fmt"
	"image"
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
	interval        float32
	order           int
	// MatchImageFormat 匹配以常见图像文件格式为后缀的正则表达模式
	/*MatchImageFormat = fmt.Sprintf("%v.*(%v.png|%v.gif|%v.jpg|%v.jpeg|%v.PNG|%v.GIF|%v.JPG|%v.JPEG)$",
	separator, separator, separator, separator, separator, separator, separator, separator, separator)*/
	MatchImageFormat = fmt.Sprintf("\\w+\\.(png|gif|jpg|jpeg|PNG|GIF|JPG|JPEG)$")
)

var preinstallSize img_op.Size

func init() {
	if runtime.GOOS == "windows" {
		gifOutRoot = path.Join(gifOutRoot, "desktop")
	}
}

// 从标准控制流读取必须参数,Scanf需以\n字符接收最后的回车。路径采取bufio读取完整字符串,避免带有空格目录被截断
func readFormStd() bool {

	fmt.Println("请输入生成Gif动画的宽(应该大于等于你最大图片素材的宽且应该是整数)：")
	_, e := fmt.Scanf("%d\n", &preinstallSize.X)
	if e != nil {
		println("无效整数位,宽度已设为500", e.Error())
		preinstallSize.X = 500
	}
	fmt.Println("请输入生成最终Gif动图的高(应该大于等于你最大图片素材的高且应该是整数)：")
	_, e = fmt.Scanf("%d\n", &preinstallSize.Y)
	if e != nil {
		println("无效整数位,高度已设为500", e.Error())
		preinstallSize.Y = 500
	}
	fmt.Println("请输入图片间的停留时长(秒)：")
	_, e = fmt.Scanf("%f\n", &interval)
	if e != nil {
		println("无效间隔,停留已设为1", e.Error())
		interval = 1
	}
	if interval > 10.0 {
		println("间隔超过10s,太久了,这是制作gif动图啊，你停留在一帧那么久有啥用？短点吧！")
		return false
	}
	if interval < 0.1 {
		println("间隔小于0.1s,太快了,这是制作gif动图啊，你一帧闪那么快，别人都看不清楚有啥用！")
		return false
	}
	fmt.Println("是否倒序？(输入0忽略，输入1倒序生成)")
	_, e = fmt.Scanf("%d\n", &order)
	if e != nil {
		println(e.Error())
		return false
	}
	fmt.Println("请输入你想要用于制作动图Gif的图片所在的文件夹路径。(直接粘贴路径并回车即可,文件名默认按序号[数字大小、字母顺序]排序整理)")
	in := bufio.NewReader(os.Stdin)
	imagesInputRoot, e = in.ReadString('\n')
	// ReadString读\n结束并接收\n，此处去除最后的\n,windows是\r\n
	imagesInputRoot = strings.TrimSuffix(imagesInputRoot, "\n")
	imagesInputRoot = strings.TrimSuffix(imagesInputRoot, "\r")
	if e != nil {
		println(e.Error())
		return false
	}
	return true
}

func Run(wFromGui, hFromGui int, durFromGui float32, orderFromGui int, inputRootFromGui string) {
	// 若参数没有数据则从标准控制台输入素材地址 要生成的尺寸 每帧间隔
	if inputRootFromGui == "" && wFromGui <= 0 && hFromGui <= 0 && durFromGui <= 0 {
		success := readFormStd()
		if !success {
			return
		}
	} else {
		// 参数有数据 表示从特点平台(windows)gui控件传入
		preinstallSize = img_op.Size{X: wFromGui, Y: hFromGui}
		imagesInputRoot = inputRootFromGui
		interval = durFromGui
		order = orderFromGui
	}
	fi, e := os.Stat(imagesInputRoot)
	if e != nil || !fi.IsDir() {
		if e != nil {
			println(e.Error())
		}
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

	// 检查输入文件夹下每张图片素材的尺寸是否超出预设尺寸
	success := checkDimensions(inputImageAbsLists)
	if !success {
		return
	}
	// 输入文件夹拼接临时转换文件夹成为新路径 并生成这个新路径(临时转换文件夹)
	imagesInputRoot, _ = filepath.Abs(path.Join(imagesInputRoot, tmpConvertFolder))

	common.CreateFolder(imagesInputRoot)

	// 开始将要组合的连续图片先统一编码为gif格式，并且存放在临时转换文件夹中
	img_op.ConvertToGif(inputImageAbsLists, imagesInputRoot)

	// 把原先存储的输入图片素材文件名统一改为.gif后缀,因为有可能素材不是gif格式后缀,统一编码后文件名已经统一为类似{1|2..}.gif
	for i := 0; i < len(imgNameFiles); i++ {
		formatEnd := strings.LastIndex(imgNameFiles[i], ".") + 1
		formatStr := imgNameFiles[i][formatEnd:]
		imgNameFiles[i] = strings.Replace(imgNameFiles[i], formatStr, "gif", -1)
	}

	// 将imgNameFiles的素材文件名列表重新按升序排列(若设置了倒序，则倒序)
	common.SortStringSlice(imgNameFiles, order == 1)

	fmt.Println("开始生成..")
	// 使用此方法将imagesInputRoot临时文件夹路径下的素材图片(已经全为gif格式)生成为gifOutRoot路径下的gif动图
	img_op.OpGifFileToGifDone(imagesInputRoot, imgNameFiles, gifOutRoot, preinstallSize, interval)
	fmt.Println("制作成功，请在你的用户目录(或Windows桌面)查看。")
	// 生成完毕，删除临时文件夹
	common.RemoveFolder(imagesInputRoot)
}

// 检查图片文件路径列表中每一张的尺寸是否超出预设尺寸
func checkDimensions(imageAbsLists []string) bool {
	for i := 0; i < len(imageAbsLists); i++ {
		f, err := os.Open(imageAbsLists[i])
		if err != nil {
			println("error open in checkDimensions:", err.Error())
			return false
		}
		ic, _, err := image.DecodeConfig(f)
		if ic.Width == 0 || ic.Height == 0 {
			println("素材图片", imageAbsLists[i], "尺寸为0，无效或已损坏的图片。")
			return false
		}
		if ic.Width > preinstallSize.X {
			println("素材图片", imageAbsLists[i], "宽度大于你想要的宽，请缩小图片尺寸。"+
				"提示：要想生成良好效果的动图，所有图片素材应小于等于输入的尺寸且它们的尺寸应该一致。")
			return false
		}
		if ic.Height > preinstallSize.Y {
			println("素材图片", imageAbsLists[i], "高度大于你想要的高，请缩小图片尺寸。"+
				"提示：要想生成良好效果的动图，所有图片素材应小于等于输入的尺寸且它们的尺寸应该一致。")
			return false
		}
		err = f.Close()
		if err != nil {
			return false
		}
	}
	return true
}
