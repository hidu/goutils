package goutils

import (
//    "fmt"
    "github.com/bmizerany/assert"
    "testing"
)

func TestIsInArray(t *testing.T) {
  arr:=[]string{"a","b","c"}
  has:=IsInArray("a",arr,[]string)
  assert.Equal(t,true,has)
}