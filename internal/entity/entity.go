package entity

import (
	"database/sql/driver"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type ID ulid.ULID

func NewID() ID {
	return ID(ulid.Make())
}

func Parse(v string) (ID, error) {
	id, err := ulid.Parse(v)
	return ID(id), err
}

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, id.String())), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	v, err := Parse(str)
	if err != nil {
		return err
	}
	*id = v
	return nil
}

func (id ID) String() string {
	return ulid.ULID(id).String()
}

// Scan Implement gorm.Scanner interface
func (id *ID) Scan(value interface{}) error {
	if value == nil {
		*id = ID{}
		return nil
	}

	var err error
	switch v := value.(type) {
	case []byte:
		*id, err = Parse(string(v))
		if err != nil {
			return err
		}
	case string:
		*id, err = Parse(v)
		if err != nil {
			return err
		}
	default:
		return errors.New("failed to scan ID: invalid type")
	}

	return nil
}

// Value Implement gorm.Valuer interface
func (id ID) Value() (driver.Value, error) {
	return id.String(), nil
}
