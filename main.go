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
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		log.Panic(err)
		return
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

	_, err = file.Write([]byte(a.String()))
	if err != nil {
		log.Panic(err)
		return
	}
	defer file.Close()

}

func read(delim rune) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString(byte(delim))
}
