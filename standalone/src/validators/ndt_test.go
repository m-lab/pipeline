package validators

import (
  "testing"
)

func TestNdt_Invalid(t *testing.T) {
  err := Ndt("", []byte("hello world"));
  if err == nil {
    t.Error("Should be unimplemented")
  }
}

func TestNdt_Valid(t *testing.T) {
  err := Ndt("", []byte("hello world"));
  if err == nil {
    t.Error("Should be unimplemented")
  }
}
