package tests

import (
	"github.com/goravel/framework/testing"

	"cariinkos/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
