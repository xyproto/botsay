package main

import (
       "testing"
       "fmt"
)

func TestSizeLatin(t *testing.T) {
  s, _ := size("abc")
  expected := 3
  if s != expected {
    t.Error(fmt.Sprintf("Expected %d but instead got %d!", expected, s))
  }
}

func TestSizeUnicode(t *testing.T) {
  s, _ := size("абв")
  expected := 3
  if s != expected {
    t.Error(fmt.Sprintf("Expected %d but instead got %d!", expected, s))
  }
}
