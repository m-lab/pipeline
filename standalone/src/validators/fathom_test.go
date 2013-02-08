package validators

import (
  "testing"
)

func TestFathom_Invalid(t *testing.T) {
  err := Fathom("", []byte(""))
  if err == nil {
    t.Error("Empty string should be invalid")
  }

  err = Fathom("", []byte("1"))
  if err == nil {
    t.Error("Simple number should be invalid")
  }

  err = Fathom("", []byte("hello world"))
  if err == nil {
    t.Error("Simple string should be invalid")
  }

  err = Fathom("", []byte("[1, 2]"))
  if err == nil {
    t.Error("JSON array should be invalid")
  }
}

func TestFathom_Malformed(t *testing.T) {
  err := Fathom("", []byte("{\"a\": 1, b: 2}"))
  if err == nil {
    t.Error("Bad key should be invalid")
  }

  err = Fathom("", []byte("{\"a\": 1, \"b\": 2"))
  if err == nil {
    t.Error("Incomplete map should be invalid")
  }
}

func TestFathom_Valid(t *testing.T) {
  err := Fathom("", []byte("{\"a\": 1, \"b\": 2}"))
  if err != nil {
    t.Error("Valid JSON should be valid: ", err.Error())
  }
}
