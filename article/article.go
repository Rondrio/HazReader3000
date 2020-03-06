package article

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

const (
	regex = `{"@context":"[\W,\w]*","@type":"[\W,\w]*","genre":"[\W,\w]*","datePublished":"[\W,\w]*","dateModified":"[\W,\w]*","keywords":"[\W,\w]*","thumbnailUrl":"[\W,\w]*","mainEntityOfPage":{"@type":"[\W,\w]*","@id":"[\W,\w]*"},"author":{"@type"[\W,\w]*","name":"[\W,\w]*"},"headline":"[\W,\w]*","description":"[\W,\w]*","articleBody":"[\W,\w]*","isAccessibleForFree":"[\W,\w]*","isPartOf":{"@type":[\W,\w]*,"name":"[\W,\w]*"},"publisher":{"@type":"[\W,\w]*","name":"[\W,\w]*","logo":{"@type":"[\W,\w]*","url":"[\W,\w]*"}},"image":{"@type":"[\W,\w]*","name":"[\W,\w]*","url":"[\W,\w]*","description":"[\W,\w]*","copyrightHolder":"[\W,\w]*","height":[\W,\w]*,"width":[\W,\w]*}}`
)

type Article struct {
	Headline    string `json:"headline"`
	Description string `json:"description"`
	ArticleBody string `json:"articleBody"`
}

func Parse(b []byte) (*Article, error) {
	var article Article


	reg,err := regexp.Compile(regex)
	if err != nil{
		log.Println(err)
	}

	cont := reg.FindString(string(b))
	err = json.Unmarshal([]byte(cont),&article)
	return &article,err
}
func (a Article) String() string {
	return fmt.Sprintf("####%s\n\n%s\n\n%s", format([]byte(a.Headline)), format([]byte(a.Description)), format([]byte(a.ArticleBody)))
}

func format(b []byte) []byte{
//	for i := range b {
//		if i%100 == 0 {//} && b[i] == byte(' ') {
//			b = insert(b,i,byte('\n'))
//		}
//	}

	return b
}

func insert(arr []byte, pos int, elem byte) []byte {
	if pos < 0 {
		pos = 0
	} else if pos >= len(arr) {
		pos = len(arr)
	}
	out := make([]byte, len(arr)+1)
	copy(out[:pos], arr[:pos])
	out[pos] = elem
	copy(out[pos+1:], arr[pos:])
	return out
}