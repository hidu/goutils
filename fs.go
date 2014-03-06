package goutils

import(
 "os"
 "io/ioutil"
)
const(
	FILE_APPEND=1
)

func File_get_contents(file_path string)(data []byte,err error){
    f,err:=os.Open(file_path)
    defer f.Close()
    if err!=nil{
       return nil,err
    }
    bf,err1:=ioutil.ReadAll(f)
    if err1!=nil{
       return nil,err1
    }
   return bf,nil
}


func File_put_contents(file_path string,data []byte,def ...int){
    flags:=os.O_RDWR|os.O_CREATE
    is_append:=false
    if len(def)>0 && def[0]==FILE_APPEND{
      is_append=true
      flags=flags|os.O_APPEND
     }
    f, err := os.OpenFile(file_path,flags , 0644) 
    defer f.Close();
    if err != nil { panic(err) }
    write_at:=int64(0)
    if is_append{
		 stat,_ := f.Stat()
		 write_at=stat.Size()
     }
	 f.WriteAt(data,write_at)
}

func File_exists(file_path string) bool{
	_,err := os.Stat( file_path )
	if err == nil {
	   return true
	}
	return os.IsExist( err )
}

func Md5_file(file_path string) string{
	return ""
}