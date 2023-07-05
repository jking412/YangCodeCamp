package answers

import "os/exec"

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

	cmd := exec.Command("docker", "exec", "js", "node", "-e", content.(string))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	if string(output[0:len(output)-1]) != q.Answer {
		return ErrAnswerNotMatch
	}

	return nil
}

func (q *JSQuestion) GetAnswer() any {
	return q.Answer
}
