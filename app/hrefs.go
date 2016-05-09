//************************************************************************//
// API "Atlas": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/goa-atlas
// --design=goa-atlas/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	return fmt.Sprintf("/users/%v", userID)
}
