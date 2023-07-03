package answers

import (
	"YangCodeCamp/model"
	"errors"
)

type Answer interface {
	Check(content any) error
	GetAnswer() any
}

func GetAnswerChecker(typ model.Type, answer string) (Answer, error) {
	if typ == 0 {
		return NewHtmlQuestion(answer)
	}
	return nil, errors.New("not support type")
}
