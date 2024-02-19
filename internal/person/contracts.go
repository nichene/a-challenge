package person

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Service struct {
	personsRepository PersonsRepository
	parentsRepository ParentsRepository
}

func NewService(
	personsRepository PersonsRepository,
	parentsRepository ParentsRepository,
) *Service {
	return &Service{
		personsRepository: personsRepository,
		parentsRepository: parentsRepository,
	}
}

type Person struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name"`
	Parents   []Person  `json:"parents,omitempty" gorm:"-"`
	Children  []Person  `json:"children,omitempty" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Person) TableName() string {
	return "persons"
}

func (p *Person) GormDataType() string {
	return "json"
}

func (p *Person) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Person) Scan(src interface{}) error {
	*p = Person{}

	switch src := src.(type) {
	case nil:
		return nil

	case string:
		return json.Unmarshal([]byte(src), p)

	case []byte:
		return json.Unmarshal(src, p)

	default:
		return errors.New("scan: incompatible type for Person")
	}
}

type Parent struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	PersonID  int       `json:"person_id" gorm:"foreignKey"`
	ParentID  int       `json:"parent_id" gorm:"foreignKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Parent) GormDataType() string {
	return "json"
}

func (p *Parent) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Parent) Scan(src interface{}) error {
	*p = Parent{}

	switch src := src.(type) {
	case nil:
		return nil

	case string:
		return json.Unmarshal([]byte(src), p)

	case []byte:
		return json.Unmarshal(src, p)

	default:
		return errors.New("scan: incompatible type for Parent")
	}
}

func (p *Parent) TableName() string {
	return "parents"
}

type CreatePersonIntent struct {
	Name     string   `json:"name" validate:"required"`
	Parents  []string `json:"parents"`
	Children []string `json:"children"`
}

type ParentsRepository interface {
	Create(ctx context.Context, personID, parentID int) (Parent, error)
	Find(ctx context.Context, personID, parentID int) (Parent, error)
}

type PersonsRepository interface {
	FindByName(ctx context.Context, name string) (*Person, error)
	Create(ctx context.Context, intent CreatePersonIntent) (*Person, error)
}
