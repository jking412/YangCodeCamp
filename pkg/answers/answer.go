package answers

import (
	"YangCodeCamp/model"
	"errors"
)

type Answer interface {
	Check(content any) error
	GetAnswer() any
}

func GetAnswerChecker(typ model.Type, answer, checkMessage string) (Answer, error) {
	if typ == model.HTML {
		return NewHtmlQuestion(answer, checkMessage)
	} else if typ == model.CSS {
		return NewCssQuestion(answer, checkMessage)
	}
	return nil, errors.New("not support type")
}
