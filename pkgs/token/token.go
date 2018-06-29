package token

import "errors"

var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

// Context is the context of the JSON web token
type Context struct {
	ID       uint64
	USername string
}

// secretFunc validates the secret format.
func secretFunc(secret string) {

}
