package main

import (
	"HAZReader/article"
	"bufio"
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	link,err := read('\n')
	if err != nil{
		log.Println(err)
		return
	}

	link = strings.ReplaceAll(link,"\r\n","")

	resp, err := http.Get(link)
	if err != nil {
		log.Panic(err)
		return
	}

	b := new(bytes.Buffer)
	n, err := b.ReadFrom(resp.Body)
	if err != nil {
		log.Panic(err)
		return
	}

	if n != resp.ContentLength {
		log.Printf("Read %v, Bodylength : %v", n, resp.ContentLength)
	}

	file, err := os.Create("artikel.md")
	if err != nil {
		log.Panic(err)
		return
	}

	a, err := article.Parse(b.Bytes())
	if err != nil {
		log.Panic(err)
		return
	}

	n1, err := file.Write([]byte(a.String()))
	if err != nil {
		log.Panic(err)
		return
	}
	defer file.Close()

	if int64(n1) != n {
		log.Printf("Wrote %v, Contentlength : %v", n1, n)
	}

}

func read(delim rune) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString(byte(delim))
}
