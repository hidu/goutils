package goutils

import (
    "fmt"
    "strconv"
    "strings"
    "path"
     "reflect"
)

type object struct {
    data interface{}
    paths []string
}

func NewInterfaceWalker(obj interface{}) *object {
    return &object{data:obj, paths: []string{}}
}


func (obj *object) GetInterface(name string) (has bool,val interface{}) {
    name=path.Clean(name)
    if name != "." && name[0] != '/' {
        name = strings.Join(obj.paths, "/") + "/" + name
    }
    has, val= InterfaceWalk(obj.data, name)
    return
}

func (obj *object) Get(name string) (has bool,val *object) {
    has, tmp:= obj.GetInterface(name)
    if has{
       val=NewInterfaceWalker(tmp)
     }
    return
}

func (obj *object) Gointo(name string) {
    path_str:=strings.Join(obj.paths,"/")+"/"+name
    path_str_clean:=path.Clean(path_str)
    obj.paths=strings.Split(path_str_clean,"/")
}

/*
*读取指定项的值
 */
func (obj *object) GetString(name string, def ...string) string {
    has,val := obj.GetInterface(name)
    if has{
	        return fmt.Sprint(val)
    }
    if len(def) > 0 {
        return def[0]
    } else {
        return ""
    }
}

func (obj *object) GetStringArray(name string, def ...[]string) []string {
    has,val := obj.GetInterface(name)
    if has{
	    switch val.(type) {
	    case []interface{}:
	        val_arr := val.([]interface{})
	        result := make([]string, 0, len(val_arr))
	        for i, v := range val_arr {
	            result[i] = fmt.Sprint(v)
	        }
	        return result
	    }
    }
    if len(def) > 0 {
        return def[0]
    } else {
        return []string{}
    }
}

func (obj *object) GetInt(name string, def ...int) int {
    var defVal int
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = -1
    }
    val := obj.GetFloat(name, float64(defVal))
    return int(val)
}

func (obj *object) GetIntArray(name string, def ...[]int) []int {
    var defVal []int
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = []int{}
    }
    float_arr := obj.GetFloatArray(name, intArr2Float(defVal))
    result := make([]int, 0, len(float_arr))
    for i, v := range float_arr {
        result[i] = int(v)
    }
    return result
}

func intArr2Float(intArr []int) []float64 {
    floatArr := make([]float64, len(intArr))
    for index, v := range intArr {
        floatArr[index] = float64(v)
    }
    return floatArr
}

func (obj *object) GetFloat(name string, def ...float64) float64 {
    var defVal float64
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = -1
    }
    ret, err := strconv.ParseFloat(obj.GetString(name, fmt.Sprint(defVal)), 64)
    if err != nil && len(def) > 0 {
        return def[0]
    }
    return ret
}

func (obj *object) GetFloatArray(name string, def ...[]float64) []float64 {
    var defVal []float64
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = []float64{}
    }
    strArr := make([]string, len(defVal))
    for index, v := range defVal {
        strArr[index] = fmt.Sprint(v)
    }

    str_arr := obj.GetStringArray(name, strArr)
    result := make([]float64, 0, len(str_arr))
    for i, v := range str_arr {
        ret, _ := strconv.ParseFloat(v, 64)
        result[i] = ret
    }
    return result
}

func (obj *object) GetBool(name string) bool {
    val := obj.GetString(name)
    if val == "" {
        return false
    }
    bv, err := strconv.ParseBool(val)
    if err == nil {
        return bv
    }
    return false
}


func InterfaceWalk(m interface{}, name string) (has bool, val interface{}) {
    name = strings.TrimSpace(path.Clean(name))
    val_tmp:= m
    if name == "/" || name == "." {
        return true, val
    }
    paths := strings.Split(strings.Trim(name, "/"), "/")
    n:=0
    for _, sub_name := range paths {
        n++
        _type:=reflect.TypeOf(val_tmp).String()
        if(len(_type)>3 && _type[:3]=="map"){
          value_of_t:=reflect.ValueOf(val_tmp)
          for _,_key:=range value_of_t.MapKeys(){
            _key_str:=fmt.Sprint(_key.Interface())
            if(sub_name==_key_str){
	             val_tmp=value_of_t.MapIndex(_key).Interface()
	             break
               }
            }
          }
    }
    if n==len(paths){
   	 return true,val_tmp
     }
    return false,nil
}
