//************************************************************************//
// API "Omnibus": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/omnibus
// --design=omnibus/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// PostHref returns the resource href.
func PostHref(postID interface{}) string {
	return fmt.Sprintf("/posts/%v", postID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	return fmt.Sprintf("/users/%v", userID)
}
