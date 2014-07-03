package utils

import (
    "fmt"
    "regexp"
     "strings"
     "crypto/md5"
)

/**
* parse str
    style='width:1' class="hello" checked=on
  as
  map[style:width:1 class:hello checked:on]
*/
func StringToMap(str string) (data map[string]string) {
    re := regexp.MustCompile(`\s*([\w-]+)\s*=\s*(['"]?)(.*)`)
    data = make(map[string]string)
    matches := re.FindAllStringSubmatch(str, -1)
    if len(matches) > 0 {
        first := matches[0]
        var reg2_txt string
        if first[2] == "'" || first[2] == `"` {
            reg2_txt = fmt.Sprintf(`([^%s]*)%s(\s+.*)?`, first[2], first[2])
        } else if first[2] == "" {
            reg2_txt = `(\S+)\s*(.*)`
        }
        re2 := regexp.MustCompile(reg2_txt)
        subResult := re2.FindAllStringSubmatch(first[3], -1)

        if len(subResult) > 0 && len(subResult[0]) > 1 {
            data[first[1]] = subResult[0][1]
            if len(subResult[0][2]) > 0 {
                _subResult := StringToMap(subResult[0][2])
                for k, v := range _subResult {
                    data[k] = v
                }
            }
        }
    }
    return
}

func isChar(ru rune) bool {
    return (ru >= 0 && ru <= 9) || (ru >= 'a' && ru <= 'z') || (ru >= 'A' && ru <= 'Z') || ru == '_' || ru == '-'
}

func LoadText2Slice(text string) (result [][]string){
  lines:=strings.Split(text,"\n")
  r:=regexp.MustCompile(`\s+`)
  for _,v:=range lines{
     index:=strings.IndexByte(v,'#')
     if(index==0){
       continue
     }
     if(index>0){
        v=v[:index]
     }
     v=strings.TrimSpace(v)
     if (v==""){
        continue
     }
     result=append(result,r.Split(v,-1))
  }
 return 
}
 
func StrMd5(mystr string)string{
  h := md5.New()
  h.Write([]byte(mystr))
  return fmt.Sprintf("%x",h.Sum(nil))
}
