package person

import "context"

func (s *Service) Find(ctx context.Context, name string) (Person, error) {
	return Person{}, nil
}
