package answers

import (
	"YangCodeCamp/model"
	"errors"
)

type Answer interface {
	Check(content any) (string, error)
	GetAnswer() any
}

func GetAnswerChecker(typ model.Type, answer, checkMessage string) (Answer, error) {
	if typ == model.HTML {
		return NewHtmlQuestion(answer, checkMessage)
	} else if typ == model.CSS {
		return NewCssQuestion(answer, checkMessage)
	} else if typ == model.JavaScript {
		return NewJSQuestion(answer, checkMessage)
	} else if typ == model.PythonChoice {
		return NewPythonChoiceQuestion(answer, checkMessage)
	}
	return nil, errors.New("not support type")
}
