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

func (q *PythonChoiceQuestion) Check(content any) error {
	if content.(string) != q.Answer {
		return ErrAnswerNotMatch
	}
	return nil
}

func (q *PythonChoiceQuestion) GetAnswer() any {
	return q.Answer
}
