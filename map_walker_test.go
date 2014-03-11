package goutils

import (
	//  "fmt"
	"encoding/json"
	"github.com/bmizerany/assert"
	"testing"
)

func TestGetVal(t *testing.T) {
	str := `{"a":{"c":1,"f":1.1},"b":[1,2],"d":{"1":{"a":"ccc"},"2":3}}`
	var m map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	assert.Equal(t, err, nil)
	//    fmt.Println(m,err)
	w := NewMapWalker(m)
	cases := make(map[string]interface{})
	cases["a/c"] = "1"
	cases["d//1/a"] = "ccc"
	for k, v := range cases {
		_v := w.GetString(k)
		assert.Equal(t, v, _v)
	}
	//     fmt.Println("a/f",w.GetFloat("a/f"))
	//     fmt.Println("a/f",w.GetInt("a/f"))
}
