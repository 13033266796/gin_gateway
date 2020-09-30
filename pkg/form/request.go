package form

import (
	// "fmt"
	"github.com/gin-gonic/gin/binding"
)

type TestRequest struct {
	Name 	string	`form:"name" json:user binding:"required"`
	Age		uint	`from:"age" json:"age" binding:"required"`
}