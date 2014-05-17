package goutils
import (
 "gopkg.in/cookieo9/resources-go.v2"
 "io/ioutil"
 "strings"
 "log"
)
type Resource struct{}

var DefaultResource *Resource=&Resource{}

func (re *Resource)Load(path string) []byte{
     res,err:=re.Get(path)
     if(err!=nil){
        return []byte{}
      }
     r,_:=res.Open()
     bf,err:=ioutil.ReadAll(r)
     if(err!=nil){
        log.Println("read res[",path,"] failed",err.Error())
      }
     return bf
}

func (re *Resource)Get(path string)(resources.Resource,error){
    path=strings.TrimLeft(path,"/")
    res,err:=resources.Find(path)
    if(err!=nil){
      log.Println("load res[",path,"] failed",err.Error())
      return nil,err
     }
     return res,nil
}
