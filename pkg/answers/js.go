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

func (q *JSQuestion) Check(content any) (string, error) {

	c := content.(string) + q.CheckMessage

	output, err := docker.CheckJS(c, q.Answer)
	if err != nil {
		return output, ErrAnswerNotMatch
	}

	return output, nil
}

func (q *JSQuestion) GetAnswer() any {
	return q.Answer
}
