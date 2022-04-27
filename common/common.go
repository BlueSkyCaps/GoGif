package common

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"sort"
)

func MatchRegexString(p, v string) bool {
	var e error
	var m bool
	if m, e = regexp.MatchString(p, v); m && (e == nil) {
		return true
	}
	if e != nil {
		println("error in MatchRegexString", e.Error())
	}
	return false
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

// CreateFolder 在指定路径下创建一个新文件夹 存在则删除再创建
func CreateFolder(createPath string) {
	// 有没有存在此文件夹
	_, err := os.Stat(createPath)
	// 存在则先删除
	if err == nil {
		// 先删除此文件夹中的所有子数据，若有
		err = os.RemoveAll(createPath)
		// 删除子数据发生错误
		if err != nil {
			panic("removeAll tmp folder data failed.")
		}
		// 再删除此文件夹本身
		_ = os.Remove(createPath)
		// 再次判断此文件夹是否存在
		_, err = os.Stat(createPath)
		// 若nil,表示仍存在，预期不理想。中止
		if err == nil {
			panic("remove tmp folder failed.")
		}
	}
	err = os.MkdirAll(createPath, os.ModePerm)
	if err != nil {
		println("error in CreateFolder:", err.Error())
		var c int
		_, _ = fmt.Scanf("%c", &c)
		os.Exit(1)
	}
}

func SortStringSlice(sr []string, desc bool) {
	if desc {
		sort.Slice(sr, func(i, j int) bool {
			return sr[i] > sr[j]
		})
		return
	}
	sort.Slice(sr, func(i, j int) bool {
		return sr[i] < sr[j]
	})
}

// RemoveFolder 删除一个文件夹
func RemoveFolder(Path string) {
	// 有没有存在此文件夹
	_, err := os.Stat(Path)
	// 存在则先删除
	if err == nil {
		// 先删除此文件夹中的所有子数据，若有
		err = os.RemoveAll(Path)
		// 删除子数据发生错误
		if err != nil {
			panic("removeAll tmp folder data failed.")
		}
		// 再删除此文件夹本身
		_ = os.Remove(Path)
	}
}
