package goutils

import (
    "fmt"
    "strconv"
    "strings"
)

type myMap struct {
    configs map[string]interface{}
    paths   []string
}

func NewMapWalker(m map[string]interface{}) *myMap {
    return &myMap{configs: m, paths: []string{}}
}

func (myMap *myMap) Gointo(name string) {
    path_num := len(myMap.paths)
    if len(name) > 1 {
        name = strings.TrimRight(name, "/")
    }
    if name == "." || name == "" {
        return
    }
    if name == "/" {
        myMap.paths = []string{}
    } else if name == ".." {
        if path_num > 0 {
            myMap.paths = myMap.paths[0 : path_num-1]
        }
    } else if strings.Contains(name, "/") {
        if name[0] == '/' {
            myMap.paths = []string{}
        }
        subNames := strings.Split(name, "/")
        for _, v := range subNames {
            myMap.Gointo(v)
        }
    } else {
        myMap.paths = append(myMap.paths, name)
    }
}

/*
*读取指定项的值
 */
func (myMap *myMap) GetString(name string, def ...string) string {
    val := myMap.Get(name)
    switch val.(type) {
    case string, float64, bool:
        return fmt.Sprint(val)
    }
    if len(def) > 0 {
        return def[0]
    } else {
        return ""
    }
}

func (myMap *myMap) GetStringArray(name string, def ...[]string) []string {
    val := myMap.Get(name)
    var defVal []string
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = []string{}
    }
    if val == nil {
        return defVal
    }

    switch val.(type) {
    case []interface{}:
        val_arr := val.([]interface{})
        result := make([]string, 0, len(val_arr))
        for i, v := range val_arr {
            result[i] = fmt.Sprint(v)
        }
        return result
    }
    return defVal
}

func (myMap *myMap) GetInt(name string, def ...int) int {
    var defVal int
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = 0
    }
    val := myMap.GetFloat(name, float64(defVal))
    return int(val)
}

func (myMap *myMap) GetIntArray(name string, def ...[]int) []int {
    var defVal []int
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = []int{}
    }
    float_arr := myMap.GetFloatArray(name, intArr2Float(defVal))
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

func (myMap *myMap) GetFloat(name string, def ...float64) float64 {
    var defVal float64
    if len(def) > 0 {
        defVal = def[0]
    } else {
        defVal = 0
    }
    ret, err := strconv.ParseFloat(myMap.GetString(name, fmt.Sprint(defVal)), 64)
    if err != nil && len(def) > 0 {
        return def[0]
    }
    return ret
}

func (myMap *myMap) GetFloatArray(name string, def ...[]float64) []float64 {
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

    str_arr := myMap.GetStringArray(name, strArr)
    result := make([]float64, 0, len(str_arr))
    for i, v := range str_arr {
        ret, _ := strconv.ParseFloat(v, 64)
        result[i] = ret
    }
    return result
}

func (myMap *myMap) GetBool(name string) bool {
    val := myMap.GetString(name)
    if val == "" {
        return false
    }
    bv, err := strconv.ParseBool(val)
    if err == nil {
        return bv
    }
    return false
}

func (myMap *myMap) Get(name string) (val interface{}) {
    if name != "" && name[0] != '/' {
        name = strings.Join(myMap.paths, "/") + "/" + name
    }
    _, val = MapWalk(myMap.configs, name)
    return
}

func MapWalk(m map[string]interface{}, name string) (has bool, val interface{}) {
    name = strings.TrimSpace(name)
    val = m
    if name == "/" || name == "" {
        return true, val
    }
    has = false
    paths := strings.Split(strings.Trim(name, "/"), "/")
    //   fmt.Println("name:",len(paths),paths)
    for _, v := range paths {
        if v == "" {
            continue
        }
        switch val.(type) {
        case map[string]interface{}:
            cur_val := val.(map[string]interface{})
            val, has = cur_val[v]
            if !has {
                return
            }
        default:
            return
        }
    }
    return
}
