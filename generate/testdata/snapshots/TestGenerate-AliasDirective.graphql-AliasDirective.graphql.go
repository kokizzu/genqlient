// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// AliasDirectiveResponse is returned by AliasDirective on success.
type AliasDirectiveResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	MyUser AliasDirectiveUser `json:"user"`
}

// GetMyUser returns AliasDirectiveResponse.MyUser, and is useful for accessing the field via an interface.
func (v *AliasDirectiveResponse) GetMyUser() AliasDirectiveUser { return v.MyUser }

// AliasDirectiveUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type AliasDirectiveUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	UserID   testutil.ID `json:"id"`
	UserName string      `json:"name"`
}

// GetUserID returns AliasDirectiveUser.UserID, and is useful for accessing the field via an interface.
func (v *AliasDirectiveUser) GetUserID() testutil.ID { return v.UserID }

// GetUserName returns AliasDirectiveUser.UserName, and is useful for accessing the field via an interface.
func (v *AliasDirectiveUser) GetUserName() string { return v.UserName }

// The query executed by AliasDirective.
const AliasDirective_Operation = `
query AliasDirective {
	user {
		id
		name
	}
}
`

func AliasDirective(
	client_ graphql.Client,
) (data_ *AliasDirectiveResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "AliasDirective",
		Query:  AliasDirective_Operation,
	}

	data_ = &AliasDirectiveResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

