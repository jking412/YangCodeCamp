package answers

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strings"
)

var (
	ErrAnswerNotMatch = errors.New("answer not match")
)

type HtmlQuestion struct {
	Answer       string
	CheckMessage string
	AnswerDom    *goquery.Document
}

func NewHtmlQuestion(answer, checkMessage string) (*HtmlQuestion, error) {

	answerDom, err := goquery.NewDocumentFromReader(strings.NewReader(answer))
	if err != nil {
		return nil, err
	}

	return &HtmlQuestion{
		Answer:       answer,
		CheckMessage: checkMessage,
		AnswerDom:    answerDom,
	}, nil

}

func (q *HtmlQuestion) Check(content any) (string, error) {

	c := content.(string)

	contentDom, err := goquery.NewDocumentFromReader(strings.NewReader(content.(string)))
	if err != nil {
		return c, ErrAnswerNotMatch
	}

	contentRoot := contentDom.Nodes[0]
	answerRoot := q.AnswerDom.Nodes[0]

	answerNode := answerRoot.FirstChild
	contentNode := contentRoot.FirstChild

	err = check(answerNode, contentNode)

	return c, err
}

func (q *HtmlQuestion) GetAnswer() any {
	return q.Answer
}

var space = []string{
	" ",
	"\n",
	"\t",
	"\r",
}

func check(node1, node2 *html.Node) (err error) {
	if node1.Type != node2.Type {
		return ErrAnswerNotMatch
	}

	if node1.Attr != nil || node2.Attr != nil {
		if node1.Attr == nil || node2.Attr == nil {
			return ErrAnswerNotMatch
		}
		if len(node1.Attr) != len(node2.Attr) {
			return ErrAnswerNotMatch
		}
		for _, a := range node1.Attr {
			var exist bool
			for _, b := range node2.Attr {
				if a.Key == b.Key && a.Val == b.Val {
					exist = true
					break
				}
			}
			if !exist {
				return ErrAnswerNotMatch
			}
		}
	}

	// 删除/n/t/r空格前缀和后缀后比较
	data1 := node1.Data
	data2 := node2.Data

	for _, c := range space {
		for {
			if strings.HasPrefix(data1, c) {
				data1 = strings.TrimPrefix(data1, c)
			} else {
				break
			}
		}
		for {
			if strings.HasSuffix(data1, c) {
				data1 = strings.TrimSuffix(data1, c)
			} else {
				break
			}
		}

		for {
			if strings.HasPrefix(data2, c) {
				data2 = strings.TrimPrefix(data2, c)
			} else {
				break
			}
		}

		for {
			if strings.HasSuffix(data2, c) {
				data2 = strings.TrimSuffix(data2, c)
			} else {
				break
			}
		}
	}

	if data1 != data2 {
		return ErrAnswerNotMatch
	}
	fmt.Println(data1, data2)
	if node1.NextSibling == nil || node2.NextSibling == nil {
		if node1.NextSibling != nil || node2.NextSibling != nil {
			return ErrAnswerNotMatch
		}
	} else {
		err = check(node1.NextSibling, node2.NextSibling)
	}

	if err != nil {
		return err
	}

	if node1.FirstChild == nil || node2.FirstChild == nil {
		if node1.FirstChild != nil || node2.FirstChild != nil {
			return ErrAnswerNotMatch
		}
	} else {
		err = check(node1.FirstChild, node2.FirstChild)
	}

	return err
}
