package goutils

import (
      "fmt"
    "encoding/json"
    "github.com/bmizerany/assert"
    "testing"
)

func TestGetVal(t *testing.T) {
   if (1>2){
	    str := `{"a":{"c":1,"f":1.1},"b":[1,2],"d":{"1":{"a":"ccc"},"2":3}}`
	    var m map[string]interface{}
	    err := json.Unmarshal([]byte(str), &m)
	    assert.Equal(t, err, nil)
	    //    fmt.Println(m,err)
	    w := NewInterfaceWalker(m)
	    cases := make(map[string]interface{})
	    cases["a/c"] = "1"
	    cases["d//1/a"] = "ccc"
	    for k, v := range cases {
	        _v := w.GetString(k)
	        fmt.Println(_v)
	        assert.Equal(t, v, _v)
	    }
    }
    d:=make(map[int]int)
    d[1]=2
    d[2]=9
    q:=NewInterfaceWalker(d)
         fmt.Println("q_test:",q.GetString("1"))
//         fmt.Println("a/f:",w.GetInt("a/f"))
}
