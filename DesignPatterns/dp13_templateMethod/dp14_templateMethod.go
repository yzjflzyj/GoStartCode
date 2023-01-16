package main

import "fmt"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

// Download 公共的模板方法
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

// 在此，子实现，可能有实现该方法，也可能没有实现。实现的时，调用自己的，没有实现时调用公共模板的
func (t *template) save() {
	fmt.Print("default save\n")
}

//实现类需要获得模板的身份，才能调用模板方法，因此实现类的结构体装的是模板类的指针类型

// HTTPDownloader 模板实现
type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

// FTPDownloader 模板实现
type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

func main() {
	ftpDownloader := NewFTPDownloader()
	ftpDownloader.Download("ftp.uri")

	httpDownloader := NewHTTPDownloader()
	httpDownloader.Download("http.uri")
}
