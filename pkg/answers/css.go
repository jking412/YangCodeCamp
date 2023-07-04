package answers

type CssQuestion struct {
	Answer       string
	CheckMessage string
}

func NewCssQuestion(answer, checkMessage string) (Answer, error) {
	return &CssQuestion{
		Answer:       answer,
		CheckMessage: checkMessage,
	}, nil
}

func (q *CssQuestion) Check(content any) error {
	return nil
}

func (q *CssQuestion) GetAnswer() any {
	return q.Answer
}
