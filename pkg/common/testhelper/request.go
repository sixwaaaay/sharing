package testhelper

import (
	"io"
	"net/url"
)

type TestCase struct {
	Name     string
	Method   string
	Path     string
	Body     io.Reader
	Form     url.Values
	Expected string
}
