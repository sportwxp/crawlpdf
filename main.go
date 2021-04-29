package main

import (
	// "path"
	"hhxxcloud.com/crawlbooks/urlbooks"
)

var (
	// URL     string = "https://sjceodisha.in/wp-content/uploads/2019/09/"
	// URL string = "http://index-of.es/z0ro-Repository-2/Cyber/01%20-%20Computer%20Science/"
	URL string = "http://web.cecs.pdx.edu/~harry/cs201/slides/"
	// PATTERN string = `.+href=\"(.+\.pdf)\">`
	PATTERN string = `.+href=\"(.+\.pdf)\">`
	splite  string = "<tr>"
	// splite   string = "\n"
	savepath string = "./books"
)

func main() {
	urlmodel := urlbooks.NewBookList(URL, PATTERN, splite)
	urlmodel.GetParse()
	urlmodel.GetFile(savepath)
	// urlmodel.GetList("test")
	// for _, item := range urlmodel.BookList {
	// 	urlbooks.GetFile(item, "/Users/xpwang/Desktop/goCrawl")
	// 	fmt.Println(item)
	// }
	// for _, item := range urlmodel.BookList {
	// 	fmt.Println(item)
	// }
	// fmt.Println(urlmodel.BookList)
}
