package answers

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
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
		return c, err
	}

	contentRoot := contentDom.Nodes[0]
	answerRoot := q.AnswerDom.Nodes[0]

	answerNode := answerRoot.FirstChild
	contentNode := contentRoot.FirstChild

	for answerNode != nil && contentNode != nil {

		for {
			if answerNode.Type != contentNode.Type {
				return c, ErrAnswerNotMatch
			}

			if strings.Trim(answerNode.Data, " ") != strings.Trim(contentNode.Data, " ") {
				return c, ErrAnswerNotMatch
			}

			if answerNode.NextSibling == nil || contentNode.NextSibling == nil {
				break
			}

			answerNode = answerNode.NextSibling
			contentNode = contentNode.NextSibling
		}

		answerNode = answerNode.FirstChild
		contentNode = contentNode.FirstChild
	}

	if answerNode != nil || contentNode != nil {
		return c, ErrAnswerNotMatch
	}

	return c, nil
}

func (q *HtmlQuestion) GetAnswer() any {
	return q.Answer
}
