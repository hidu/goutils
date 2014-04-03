package goutils

import (
    "fmt"
    "strconv"
    "strings"
    os_path "path" 
    "reflect"
    "log"
)

type Object struct {
    data interface{}
    paths []string
}
func NewInterfaceWalker(obj interface{}) *Object {
    return &Object{data:obj, paths: []string{}}
}


func (obj *Object) GetInterface(path interface{}) (has bool,val interface{}) {
    path_str:=os_path.Clean(fmt.Sprint(path))
    if path_str != "." && path_str[0] != '/' {
        path_str = strings.Join(obj.paths, "/") + "/" + path_str
    }
    has, val= InterfaceWalk(obj.data, path_str)
    return
}

func (obj *Object) GetObject(path interface{}) (has bool,val *Object) {
    has, tmp:= obj.GetInterface(path)
    if has{
       val=NewInterfaceWalker(tmp)
     }
    return
}

func (obj *Object) Gointo(path interface{}) {
    path_str:=strings.Join(obj.paths,"/")+"/"+fmt.Sprint(path)
    path_str_clean:=os_path.Clean(path_str)
    obj.paths=strings.Split(path_str_clean,"/")
}

/*
*读取指定项的值
 */
func (obj *Object) GetString(path interface{}, def ...string) string {
    has,val := obj.GetInterface(path)
    if has{
	   return fmt.Sprint(val)
    }
    if len(def) > 0 {
        return def[0]
    } else {
        return ""
    }
}

func (obj *Object) GetStringSlice(path interface{}, def ...[]string) []string {
   interface_array:=obj.GetInterfaceSlice(path)
   if(len(interface_array)>0){
	   result:=make([]string,len(interface_array))
	   for i,v:=range interface_array{
	      result[i]=fmt.Sprint(v)
	    }
	   return result
   }else{
      if (len(def)>0){
        return def[0]
      }else{
       return []string{}
      }
   }
}

func (obj *Object)GetInterfaceSlice(path interface{},def ...[]interface{}) []interface{}{
   has,val := obj.GetInterface(path)
    if has{
        _type:=reflect.TypeOf(val).String()
        if len(_type)<3 || _type[:2]!="[]"{
         log.Printf("GetStringArray failed,[%s] not slice",_type)
       }else{
          value_of_t:=reflect.ValueOf(val)
          val_len:=value_of_t.Len()
	       result := make([]interface{}, val_len)
	       for i:=0;i<val_len;i++{
	         result[i]=value_of_t.Index(i).Interface()
	         }
	         return result
           
         }
    }
    if len(def) > 0 {
        return def[0]
    } else {
        return []interface{}{}
    }
}

func (obj *Object) GetInt(path interface{}, def ...int) int {
    val := obj.GetFloat(path)
    if val!=float64(-1){
       return int(val)
     }else{
	    if len(def) > 0 {
	        return def[0]
	    } else {
	        return -1
	    }
     }
}

func (obj *Object) GetIntSlice(path interface{}, def ...[]int) []int {
    float_arr := obj.GetFloatSlice(path)
    if len(float_arr)>0{
	    result := make([]int, len(float_arr))
	    for i, v := range float_arr {
	        result[i] = int(v)
	    }
	    return result
    }else{
    if len(def) > 0 {
        return def[0]
    } else {
        return []int{}
    }
    }
}

func intArr2Float(intArr []int) []float64 {
    floatArr := make([]float64, len(intArr))
    for index, v := range intArr {
        floatArr[index] = float64(v)
    }
    return floatArr
}

func (obj *Object) GetFloat(path interface{}, def ...float64) float64 {
    str:=obj.GetString(path)
    if(str!=""){
	    ret, err := strconv.ParseFloat(str, 64)
	    if err != nil {
	        log.Printf("GetFloat faild [%s] value:[%v]",path,ret)
	        return -1
	    }
      return ret
    }else{
       if len(def) > 0 {
        return def[0]
	    } else {
	     return -1
	    }
    }
}

func (obj *Object) GetFloatSlice(path interface{}, def ...[]float64) []float64 {
    str_arr := obj.GetStringSlice(path)
    if len(str_arr)>0{
	    result := make([]float64, len(str_arr))
	    for i, v := range str_arr {
	        ret, _ := strconv.ParseFloat(v, 64)
	        result[i] = ret
	    }
	    return result
    }else{
	     if len(def) > 0 {
	        return def[0]
	    } else {
	       return []float64{}
	    }
    }
}

func (obj *Object) GetBool(path interface{}) bool {
    val := obj.GetString(path)
    if val == "" {
        return false
    }
    bv, err := strconv.ParseBool(val)
    if err == nil {
        return bv
    }
    return false
}

/**
*quick get the val from a interface
*/
func InterfaceWalk(obj interface{}, path interface{}) (has bool, val interface{}) {
    path_str:= strings.TrimSpace(os_path.Clean(fmt.Sprint(path)))
//    fmt.Println("path:",path_str)
    if path_str == "/" || path_str == "." {
        return true, obj
    }
    val_tmp:= obj
    paths := strings.Split(strings.Trim(path_str, "/"), "/")
    n:=0
    for _, sub_name := range paths {
        _type:=reflect.TypeOf(val_tmp).String()
         _value:=reflect.ValueOf(val_tmp)
        if(len(_type)>3 && _type[:3]=="map"){
          has_match:=false
          for _,_key:=range _value.MapKeys(){
            _key_str:=fmt.Sprint(_key.Interface())
            if(sub_name==_key_str){
	             val_tmp=_value.MapIndex(_key).Interface()
	             has_match=true
	             break
               }
            }
         if !has_match{
           break
            }
        }else if(len(_type)>3 && _type[:2]=="[]"){
          index,err:=strconv.Atoi(sub_name)
          if err!=nil{
            log.Printf("now here is slice,[%s] must int,input path is:[%s]",sub_name,path)
            break
             }
          total_len:=_value.Len()
          if (index>0 && index>total_len) || (index<0 && index*-1>total_len){
            log.Printf("slice index out of range,index:[%d],slice size:[%d]",index,total_len)
            break
             }
         if(index<0){
              index=total_len+index
            }
          val_tmp=_value.Index(index).Interface()
        }else{
           log.Printf("not support now:%s",_type)
            break
          }
       n++
    }
    if n==len(paths){
   	 return true,val_tmp
     }
    return false,nil
}
