package answers

import "YangCodeCamp/pkg/docker"

type JSQuestion struct {
	Answer       string
	CheckMessage string
}

func NewJSQuestion(answer, checkMessage string) (*JSQuestion, error) {

	return &JSQuestion{
		Answer:       answer,
		CheckMessage: checkMessage,
	}, nil

}

func (q *JSQuestion) Check(content any) error {

	err := docker.CheckJS(content.(string), q.Answer)
	if err != nil {
		return ErrAnswerNotMatch
	}

	return nil
}

func (q *JSQuestion) GetAnswer() any {
	return q.Answer
}
