package urlbooks

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// const (
// 	WorkDir string = "./books"
// )

type BookList struct {
	URL           string
	StringPattern *regexp.Regexp
	BookList      []string
	Split         string
}

func NewBookList(url, pattern, splite string) *BookList {
	return &BookList{
		URL:           url,
		StringPattern: regexp.MustCompile(pattern),
		Split:         splite,
	}
}

func (b *BookList) GetParse() error {
	res, err := http.Get(b.URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// fmt.Print(string(content))
	// os.WriteFile("book.html", content, 0666)
	// log.Fatal("stop")
	for _, line := range strings.Split(string(content), b.Split) {
		tokens := b.StringPattern.FindStringSubmatch(line)
		if len(tokens) > 1 {
			b.BookList = append(b.BookList, b.URL+tokens[1])
			// fmt.Println(b.BookList[len(b.BookList)-1])
		}
	}
	log.Print(b.BookList)
	return nil
}

func (b *BookList) GetFile(savepath string) error {
	wg := &sync.WaitGroup{}

	for _, href := range b.BookList {
		// response, err := http.Get(href)
		// if err != nil {
		// 	log.Fatal("http get file wrong")
		// }
		// fmt.Println("http get success")
		// defer response.Body.Close()

		// bytecontent, _ := io.ReadAll(response.Body)
		// fmt.Printf("read %d bytes in memroy\n", len(bytecontent))
		// savefilename := path.Join(savepath, strings.Split(href, "/")[len(strings.Split(href, "/"))-1])
		// err = os.WriteFile(savefilename, bytecontent, 0666)
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("get %s done\n", savefilename)
		wg.Add(1)

		go wgetfile(href, savepath, wg)
	}
	wg.Wait()

	return nil
}

func wgetfile(name, savepath string, wg *sync.WaitGroup) {

	fmt.Printf("Start ***** %s begin \n", name)
	cmd := exec.Command("wget", name)
	cmd.Dir = savepath
	// log.Print(savepath)
	err := cmd.Run()
	log.Printf("???Command finished with error: %v", err)
	fmt.Printf("Down %%%%%% %s done \n", name)
	wg.Done()
}
