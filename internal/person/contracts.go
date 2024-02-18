package person

type Service struct {
	personsRepository PersonsRepository
	parentsRepository ParentsRepository
}

func NewService(
	personsRepository PersonsRepository,
	parentsRepository ParentsRepository) *Service {
	return &Service{
		personsRepository: personsRepository,
		parentsRepository: parentsRepository,
	}
}

type Person struct {
	ID       string
	Name     string
	Parents  []Person
	Children []Person
}

type ParentsRepository interface{}

type PersonsRepository interface{}
