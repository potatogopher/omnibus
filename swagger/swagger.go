//************************************************************************//
// Omnibus Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/omnibus
// --design=omnibus/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import "github.com/goadesign/goa"

// MountController mounts the swagger spec controller.
func MountController(service *goa.Service) {
	service.ServeFiles("/swagger.json", "swagger/swagger.json")
}
