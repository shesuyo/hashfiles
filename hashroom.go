package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
)

func NewHashRoom(dir, output, ignoreDir, ignoreFile string, goroutine int) *HashRoom {
	var err error
	hr := new(HashRoom)
	_, err = os.Stat(dir)
	if err != nil {
		dir, err := filepath.Abs(dir)
		if err != nil {
			panic(err)
		}
		if _, err = os.Stat(dir); err != nil {
			panic(err)
		}
	}
	hr.DirPath = dir
	hr.OutPath = output
	if ignoreDir != "" {
		hr.regDir, err = regexp.CompilePOSIX(ignoreDir)
		if err != nil {
			panic(err)
		}
	}
	if ignoreFile != "" {
		hr.regFile, err = regexp.CompilePOSIX(ignoreFile)
		if err != nil {
			panic(err)
		}
	}
	hr.maxGoroutine = goroutine

	return hr

}

type HashRoom struct {
	DirPath string
	OutPath string

	regFile *regexp.Regexp
	regDir  *regexp.Regexp

	maxGoroutine int

	fileinfos []FileInfo
}

func (hr *HashRoom) Run() {
	hr.analysis()
	hr.calculate()
	hr.WriteToFile()
}

//讲需要计算SHA1的文件挑选出来
func (hr *HashRoom) analysis() {
	err := filepath.Walk(hr.DirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if hr.regDir != nil && hr.regDir.MatchString(info.Name()) {
				return filepath.SkipDir
			}
		} else {
			if hr.regFile != nil && hr.regFile.MatchString(info.Name()) {
				return nil
			}
			hr.fileinfos = append(hr.fileinfos, FileInfo{Path: path, Size: info.Size()})
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

//进行计算SHA1
func (hr *HashRoom) calculate() {
	orders := make(chan struct{}, hr.maxGoroutine)
	length := len(hr.fileinfos)
	wg := sync.WaitGroup{}

	for i := 0; i < length; i++ {
		orders <- struct{}{}
		wg.Add(1)
		go func(i int) {
			bs, err := ioutil.ReadFile(hr.fileinfos[i].Path)
			if err != nil {
				panic(err)
			}
			hr.fileinfos[i].Hash = fmt.Sprintf("%X", sha1.Sum(bs))
			<-orders
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//将数据写到指定文件里面
func (hr *HashRoom) WriteToFile() {
	bf := new(bytes.Buffer)
	length := len(hr.fileinfos)
	for i := 0; i < length; i++ {
		bf.WriteString(hr.fileinfos[i].Path)
		bf.WriteString(",")
		bf.WriteString(hr.fileinfos[i].Hash)
		bf.WriteString(",")
		bf.WriteString(strconv.FormatInt(hr.fileinfos[i].Size, 10))
		bf.WriteString("\r\n")
	}
	ioutil.WriteFile(hr.OutPath, bf.Bytes(), os.ModePerm)
}
