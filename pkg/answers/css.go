package answers

import (
	"github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
)

type CssQuestion struct {
	Answer       string
	CheckMessage string
	AnswerCss    *css.Stylesheet
}

func NewCssQuestion(answer, checkMessage string) (Answer, error) {
	answerCss, err := parser.Parse(answer)
	if err != nil {
		return nil, err
	}
	return &CssQuestion{
		Answer:       answer,
		CheckMessage: checkMessage,
		AnswerCss:    answerCss,
	}, nil
}

func (q *CssQuestion) Check(content any) error {
	return nil
}

func (q *CssQuestion) GetAnswer() any {
	return q.Answer
}
