package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var FLICKR_SITE = "https://www.flickr.com"

type FlickrDownloader struct {
	DebugMode   bool
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	FatalLogger *log.Logger
}

// Find all pages url by user or set.
func findAllPages(url string) (urls []string, err error) {
	maxpage := 131
	returls := []string{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	doc.Find(".Paginator .pages .rapidnofollow").Each(func(i int, s *goquery.Selection) {
		num, err := strconv.Atoi(s.Text())
		if err == nil {
			if num > maxpage {
				maxpage = num
			}
		}
	})
	if maxpage == 0 {
		returls = append(returls, url+"/page1")
	}
	for i := 1; i <= maxpage; i++ {
		oneurl := url + "/page" + strconv.Itoa(i)
		returls = append(returls, oneurl)
	}
	return returls, nil
}

// Find all photo url from page.
func findPhotoUrls(url string) (uris []string, err error) {
	urls := []string{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".photo-display-item .hover-target .thumb .photo_container .rapidnofollow").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Attr("href")
		urls = append(urls, FLICKR_SITE+t)
	})
	return urls, nil

}

// Find photo .jpg link by photo url.
// It will depance on size.
func findPhotoTrueLink(url, size string) (uri string, err error) {
	photoId := parsePhotoId(url)
	doc, err := goquery.NewDocument(FLICKR_SITE + "/" + photoId + "/sizes/" + size)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	trueurl := ""
	doc.Find("#allsizes-photo img").Each(func(i int, s *goquery.Selection) {
		tr, isFind := s.Attr("src")
		if isFind {
			trueurl = tr
		}
	})
	return trueurl, nil
}

// Parse photo id by photo link.
// ex; https://www.flickr.com/photos/marksein/9448406987/in/set-72157634949960809
// will return marksein/9448406987
// ex: https://www.flickr.com/photos/marksein/9448406987
// will return marksein/9448406987 too.
func parsePhotoId(urls string) string {
	fileURL, err := url.Parse(urls)
	if err != nil {
		panic(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	id := segments[2] + "/" + segments[3]
	return id
}

// Get Filename by url.
// ex: http://example.com/ex.jpg
// It will return ex.jpg
func parseFileName(urls string) string {
	fileURL, err := url.Parse(urls)
	if err != nil {
		panic(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	return fileName
}

// Create FlickerDownloader. Use debug para to setup logger in debug mode.
func InitDownloader(debug bool) *FlickrDownloader {
	downloader := new(FlickrDownloader)
	downloader.DebugMode = debug
	downloader.InitLogger(os.Stdout, os.Stdout, os.Stderr)
	return downloader
}

// Init logger. This method will init INFO,DEBUG,ERROR three logger to
// FlickerDownloader.
func (downloader *FlickrDownloader) InitLogger(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	downloader.InfoLogger = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)
	downloader.DebugLogger = log.New(warningHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	downloader.FatalLogger = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// Save all photo by url. Here's url can contain manay page,it often be a set or
// an user's all photo.
// path is where you want to storage downloaded photo.
// A url often has many page, sometimes you dont want download all page at one times,
// you can use maxPage para.
// imageSize:
//      o means origin.
//      l means large.
//      m means Medium.
func (downloader *FlickrDownloader) SaveAllPhoto(url, path string, maxPage int, imageSize string) {
	pageUrls, err := downloader.getPagesUrls(url)
	if err != nil {
		downloader.errors(err)
		return
	}
	photoPageUrls := []string{}
	for pageIndex, element := range pageUrls {
		if (pageIndex + 1) > maxPage {
			break
		}
		us, err := downloader.getPhotoUrls(element)
		if err != nil {
			downloader.errors(err)
		}
		photoPageUrls = append(photoPageUrls, us...)
	}
	downloader.info("Finded " + strconv.Itoa(len(photoPageUrls)) + " photos.In " + url)
	var wg sync.WaitGroup
	for _, photoUrl := range photoPageUrls {
		wg.Add(1)
		go downloader.savePhoto(photoUrl, path, imageSize, &wg)
		// wait for a little time.
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func (downloader *FlickrDownloader) getPagesUrls(url string) (uris []string, err error) {
	downloader.info("Find Page Urls " + url)
	return findAllPages(url)
}

func (downloader *FlickrDownloader) getPhotoUrls(url string) (uris []string, err error) {
	downloader.info("Find Photo Urls " + url)
	return findPhotoUrls(url)
}

func (downloader *FlickrDownloader) savePhoto(url, path, imageSize string, wg *sync.WaitGroup) {
	trueLink, _ := findPhotoTrueLink(url, imageSize)
	downloader.debug("Download " + trueLink)
	resp, err := http.Get(trueLink)
	defer resp.Body.Close()
	if err != nil {
		downloader.errors(err)
		return
	}
	filename := parseFileName(trueLink)
	downloader.debug("Save " + filename)
	out, err := os.Create(path + "/" + filename)
	defer out.Close()
	if err != nil {
		downloader.errors(err)
		return
	}
	_, ferr := io.Copy(out, resp.Body)
	if ferr != nil {
		downloader.errors(ferr)
		return
	}
	downloader.info("File :" + filename + " Saved.")
	wg.Done()
}

func (downloader *FlickrDownloader) info(v ...interface{}) {
	go downloader.InfoLogger.Println(v)
}
func (downloader *FlickrDownloader) debug(v ...interface{}) {
	if downloader.DebugMode {
		go downloader.InfoLogger.Println(v)
	}
}
func (downloader *FlickrDownloader) errors(v ...interface{}) {
	go downloader.InfoLogger.Println(v)
}

func main() {
	VERSION := "v0.2"
	urlPtr := flag.String("u", "", "Photo Set Url")
	pathPtr := flag.String("p", "/tmp", "Download path.")
	maxPagePtr := flag.Int("m", 99999, "Max Pages to Download")
	debugPtr := flag.Bool("d", false, "Enable debug mode.")
	sizePtr := flag.String("s", "o", "What image size you want to download.'o' is means origin.'l' means large.")
	flag.Parse()
	if len(*urlPtr) < 1 {
		fmt.Println("FlickrDownloader " + VERSION)
		fmt.Println("Url can not be empty.Use -h to get some help.")
		return
	}
	fmt.Println("Start download photo from:" + *urlPtr)
	fmt.Println("The max number of download page is " + strconv.Itoa(*maxPagePtr))
	fmt.Println("Download image size is " + *sizePtr)

	downloader := InitDownloader(*debugPtr)
	fmt.Println("end")
	downloader.SaveAllPhoto(*urlPtr, *pathPtr, *maxPagePtr, *sizePtr)
}