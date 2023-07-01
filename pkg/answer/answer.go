package answer

type Answer interface {
	Check(content any) error
	GetAnswer() any
}
