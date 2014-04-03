package goutils

import (
//      "fmt"
    "encoding/json"
    "github.com/bmizerany/assert"
    "testing"
)

func TestGetVal(t *testing.T) {
   if (3>2){
	    str := `{"a":{"c":1,"f":1.1},"b":[1,2],"d":{"1":{"a":"ccc"},"2":3}}`
	    var m map[string]interface{}
	    err := json.Unmarshal([]byte(str), &m)
	    assert.Equal(t, err, nil)
	    w := NewInterfaceWalker(m)
	    cases := make(map[string]interface{})
	    cases["a/c"] = "1"
	    cases["b/1"] = "2"
	    cases["b/0"] = "1"
	    cases["d//1/a"] = "ccc"
	    for k, v := range cases {
//	    fmt.Println(k,v, w.GetString(k))
	        assert.Equal(t, v, w.GetString(k))
	    }
	   assert.Equal(t, w.GetIntSlice("b"),[]int{1,2})
    }
    
    int_map:=make(map[int]int)
    int_map[1]=2
    int_map[2]=9
    slice_walker:=NewInterfaceWalker(int_map)
    for k,v := range int_map {
//       fmt.Println("int_map:",k,"->",v, "?=",slice_walker.GetInt(k))
       assert.Equal(t, v,slice_walker.GetInt(k))
     }
     arr:=[]int{1,3,5}
     walker_3:=NewInterfaceWalker(arr) 
     for i,v:=range arr{
        assert.Equal(t,v, walker_3.GetInt(i))
     }  
        assert.Equal(t,arr, walker_3.GetIntSlice(""))
}
