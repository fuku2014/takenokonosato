package main

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	if Foo() != "Foo" {
		t.Error(errors.New("foo is not foo"))
	}
}
