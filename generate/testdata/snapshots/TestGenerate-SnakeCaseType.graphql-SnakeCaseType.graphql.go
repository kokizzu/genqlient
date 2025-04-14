// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// SnakeCaseTypeResponse is returned by SnakeCaseType on success.
type SnakeCaseTypeResponse struct {
	Snake_case_type SnakeCaseTypeSnake_case_type `json:"snake_case_type"`
}

// GetSnake_case_type returns SnakeCaseTypeResponse.Snake_case_type, and is useful for accessing the field via an interface.
func (v *SnakeCaseTypeResponse) GetSnake_case_type() SnakeCaseTypeSnake_case_type {
	return v.Snake_case_type
}

// SnakeCaseTypeSnake_case_type includes the requested fields of the GraphQL type snake_case_type.
type SnakeCaseTypeSnake_case_type struct {
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns SnakeCaseTypeSnake_case_type.Id, and is useful for accessing the field via an interface.
func (v *SnakeCaseTypeSnake_case_type) GetId() testutil.ID { return v.Id }

// GetName returns SnakeCaseTypeSnake_case_type.Name, and is useful for accessing the field via an interface.
func (v *SnakeCaseTypeSnake_case_type) GetName() string { return v.Name }

// The query executed by SnakeCaseType.
const SnakeCaseType_Operation = `
query SnakeCaseType {
	snake_case_type {
		id
		name
	}
}
`

func SnakeCaseType(
	client_ graphql.Client,
) (data_ *SnakeCaseTypeResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "SnakeCaseType",
		Query:  SnakeCaseType_Operation,
	}

	data_ = &SnakeCaseTypeResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

