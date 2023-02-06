package downloader

import (
	"errors"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/gogf/gf/v2/os/gfile"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type FilePart struct {
	Index int
	Start uint64
	End   uint64
	Data  []byte
}

type Downloader struct {
	// 下载地址
	url string
	// 输出文件名
	fileName string
	// 输出目录
	outputPath string
	// 要下载文件的大小
	fileSize uint64
	// 线程数
	taskNum     int
	dividedFile []FilePart
}

func NewDownloader(url string, fileName string, path string, taskNum int) *Downloader {
	if path == "" {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		log.Println("输出目录：", currentPath)
		path = currentPath
	}

	if !gfile.Exists(path) {
		gfile.Mkdir(path)
		gfile.Chmod(path, 644)
	}

	return &Downloader{
		url:         url,
		fileName:    fileName,
		outputPath:  path,
		fileSize:    0,
		taskNum:     taskNum,
		dividedFile: make([]FilePart, taskNum),
	}
}

func (d *Downloader) Download() error {
	fileSize, err := d.GetFileSize()
	if err != nil {
		return err
	}
	d.fileSize = fileSize
	tasks := make([]FilePart, d.taskNum)
	taskLength := fileSize / uint64(d.taskNum)
	for i := 0; i < d.taskNum; i++ {
		tasks[i].Index = i
		tasks[i].Start = uint64(i) * taskLength
		if i == d.taskNum-1 {
			tasks[i].End = d.fileSize - 1
		} else {
			tasks[i].End = tasks[i].Start + taskLength - 1
		}
	}
	wg := sync.WaitGroup{}
	for _, t := range tasks {
		wg.Add(1)
		go func(task FilePart) {
			defer wg.Done()
			if err = d.downloadPart(task); err != nil {
				log.Printf("文件下载失败 %v %v", err, task)
			}
		}(t)
	}
	wg.Wait()

	return d.merge()
}

func (d *Downloader) GetFileSize() (uint64, error) {
	var client = http.Client{}
	resp, err := client.Head(d.url)
	if err != nil {
		log.Println("获取长度失败, ", err.Error())
		return 0, err
	}
	if resp.StatusCode > 299 {
		return 0, errors.New(fmt.Sprintf("Can't process, response is %v", resp.StatusCode))
	}

	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("服务器不支持文件断点续传")
	}

	fileSize := resp.ContentLength
	log.Printf("要下载的文件大小为 %v\n", funs.ByteCountIEC(fileSize))
	return uint64(fileSize), err
}

func (d *Downloader) downloadPart(p FilePart) error {
	client := http.Client{}
	log.Printf("开始[%d]下载from:%v to:%v\n", p.Index, funs.ByteCountIEC(p.Start), funs.ByteCountIEC(p.End))
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%v-%v", p.Start, p.End))
	resp, err := client.Do(req)

	if err != nil {
		log.Println("请求响应失败", err)
		return err
	}

	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("服务器状态码: %v", resp.StatusCode))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("关闭resp失败, %v", err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if uint64(len(data)) != (p.End - p.Start + 1) {
		return errors.New("分片下载长度不正确")
	}
	p.Data = data
	d.dividedFile[p.Index] = p
	return nil
}

func (d *Downloader) merge() error {
	path := filepath.Join(d.outputPath, d.fileName)
	log.Println("下载完毕，开始合并文件", path)
	file, err := os.Create(path)
	if err != nil {
		log.Println("文件创建失败")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("文件关闭失败", err)
		}
	}(file)
	var totalSize uint64 = 0
	for i := 0; i < d.taskNum; i++ {
		writeLen, err := file.Write(d.dividedFile[i].Data)
		if err != nil {
			log.Printf("合并文件时失败, %v\n", err)
			return err
		}
		totalSize += uint64(writeLen)
	}
	if totalSize != d.fileSize {
		return errors.New("文件不完整")
	}
	return nil
}
