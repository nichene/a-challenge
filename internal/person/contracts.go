package person

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type Person struct {
	Name string
}
