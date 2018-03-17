package util

import (
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
)

func ListDirByFix(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() {
			// 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			//匹配文件
			files = append(files, dirPth + PthSep + fi.Name())
		}
	}
	return files, nil
}

func ListDir(dirPth string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		files = append(files, dirPth + PthSep + fi.Name())
		if fi.IsDir() {
			// 忽略目录
			item_files, err := ListDir(dirPth + PthSep + fi.Name())
			if err != nil {
				return nil, err
			}
			for _, f := range item_files {
				files = append(files, f)
			}
		}

	}
	return files, nil
}

//获取第二级目录
func GetSecondPathName(path string) string {
	path = filepath.Clean(strings.TrimLeft(path, "/"))
	path_strs := strings.Split(path, "/")
	if len(path_strs) >= 2 {
		return path_strs[1]
	} else {
		return ""
	}
}


