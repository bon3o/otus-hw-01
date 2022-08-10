package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	MinMax struct {
		Age int `validate:"min:18|max:50"`
	}

	In struct {
		Role   UserRole `validate:"in:admin,stuff"`
		Status string   `validate:"in:new,created"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidateErrors(t *testing.T) {
	var uRole UserRole = "www"

	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{in: "NotStructError", expectedErr: ErrNotStruct},
		{in: App{Version: ""}, expectedErr: ValidationErrors{{Field: "Version", Err: ErrLen}}},
		{in: MinMax{Age: 17}, expectedErr: ValidationErrors{{Field: "Age", Err: ErrMin}}},
		{in: MinMax{Age: 51}, expectedErr: ValidationErrors{{Field: "Age", Err: ErrMax}}},
		{in: In{Role: uRole, Status: "new"}, expectedErr: ValidationErrors{{Field: "Role", Err: ErrIn}}},
		{in: In{Role: uRole, Status: "new1"}, expectedErr: ValidationErrors{{Field: "Role", Err: ErrIn}, {Field: "Status", Err: ErrIn}}},
		{in: User{
			ID:     "asd",
			Name:   "Name",
			Age:    17,
			Email:  "asd.asd",
			Role:   uRole,
			Phones: []string{"8 8005553535 "},
			meta:   []byte("{}"),
		}, expectedErr: ValidationErrors{
			{Field: "ID", Err: ErrLen},
			{Field: "Age", Err: ErrMin},
			{Field: "Email", Err: ErrRegexp},
			{Field: "Role", Err: ErrIn},
			{Field: "Phones", Err: ErrLen},
		}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestValidateSuccess(t *testing.T) {
	var uRole UserRole = "admin"

	tests := []struct {
		in interface{}
	}{
		{in: App{Version: "0.0.1"}},
		{in: MinMax{Age: 19}},
		{in: MinMax{Age: 50}},
		{in: In{Role: uRole, Status: "created"}},
		{in: User{
			ID:     "5f04797b-e4ea-4ede-91c7-576a42d1f764",
			Name:   "Name",
			Age:    21,
			Email:  "test@test.ru",
			Role:   uRole,
			Phones: []string{"88005553535", "89995553535"},
			meta:   []byte("{}"),
		}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.NoError(t, err)
		})
	}
}
