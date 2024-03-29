package person

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	parentsRepoMock := NewMockParentsRepository(ctrl)
	personsRepoMock := NewMockPersonsRepository(ctrl)

	service := NewService(personsRepoMock, parentsRepoMock)

	t.Run("Should not create person when create error in db", func(t *testing.T) {
		createIntent := CreatePersonIntent{
			Name: "Ana",
		}

		expectedErr := errors.New("db error")
		personsRepoMock.EXPECT().Create(ctx, createIntent).Return(&Person{}, expectedErr)

		person, err := service.Create(ctx, createIntent)

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, 0, person.ID)
	})

	//TO DO: create tests
}
