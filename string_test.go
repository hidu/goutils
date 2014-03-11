package goutils

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestStringToMap(t *testing.T) {
	s := `style='width:1' class="hello" checked=on`
	ret := StringToMap(s)
	if (1>1){
		fmt.Println("input:", s, "\noutput:", ret)
		for k, v := range ret {
			fmt.Println(k, "==>", v)
		}
	}
	assert.Equal(t, len(ret), 3)
	assert.Equal(t, ret["style"], "width:1")
	assert.Equal(t, ret["class"], "hello")
	assert.Equal(t, ret["checked"], "on")
}
