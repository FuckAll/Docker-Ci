/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod  on 2016/05/13 16:57
 */

package filewalk

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type fl struct {
	list []string
}

func WalkDir(dirPath, suffix string) *fl {
	f := &fl{}
	f.fileWalk(dirPath, suffix)
	return f
}

func (f *fl) FileList() []string {
	return f.list
}

func (f *fl) fileWalk(dirPath, suffix string) {
	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	//先递归找文件夹
	for _, v := range fis {
		if v.IsDir() {
			f.fileWalk(filepath.Join(dirPath, v.Name()), suffix)
		}
	}

	//在递归的文件夹中依次罗列sql文件
	fis, err = ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	suffix = strings.ToUpper(suffix)
	for _, v := range fis {
		if !v.IsDir() && strings.HasSuffix(strings.ToUpper(v.Name()), suffix) {
			f.list = append(f.list, filepath.Join(dirPath, v.Name()))
		}
	}
}
