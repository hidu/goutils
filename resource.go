package goutils
import (
 "gopkg.in/cookieo9/resources-go.v2"
 "io/ioutil"
 "strings"
 "log"
 "mime"
 "time"
 "path/filepath"
 "net/http"
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

func (re *Resource)HandleStatic(w http.ResponseWriter,r *http.Request,path string){
    res,err:=re.Get(path)
    if(err!=nil){
        http.NotFound(w,r)
        return;
     }
    finfo,_:=res.Stat()
    modtime:=finfo.ModTime()
    if t, err := time.Parse(http.TimeFormat, r.Header.Get("If-Modified-Since")); err == nil && modtime.Before(t.Add(1*time.Second)) {
           h := w.Header()
           delete(h, "Content-Type")
           delete(h, "Content-Length")
           w.WriteHeader(http.StatusNotModified)
           return
           }
   mimeType:= mime.TypeByExtension(filepath.Ext(path))
   if(mimeType!=""){
       w.Header().Set("Content-Type",mimeType)
     }
    w.Header().Set("Last-Modified",modtime.UTC().Format(http.TimeFormat))
    w.Write(re.Load(path))
}
