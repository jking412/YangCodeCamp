package answers

type PythonChoiceQuestion struct {
	Answer       string
	CheckMessage string
}

func NewPythonChoiceQuestion(answer, checkMessage string) (*PythonChoiceQuestion, error) {

	return &PythonChoiceQuestion{
		Answer:       answer,
		CheckMessage: checkMessage,
	}, nil

}

func (q *PythonChoiceQuestion) Check(content any) (string, error) {
	c := content.(string)
	if c != q.Answer {
		return c, ErrAnswerNotMatch
	}
	return c, nil
}

func (q *PythonChoiceQuestion) GetAnswer() any {
	return q.Answer
}
