package goutils
import (
	"fmt"
//	"github.com/bmizerany/assert"
	"testing"
)
func TestHtml_input_text(t *testing.T){
   test:=Html_input_text("a","b","style='color:red'","class='c'")
   fmt.Println(test)
   sele:=Html_select("a","v",map[string]string{"a":"aasss","v":"sss","s":"tttt\"ss>s"})
   fmt.Println(sele)
   
   link:=Html_link("http://www.baidu.com","baidu-百度")
   fmt.Println(link)
   
   checkBox:=Html_checkBox("name","good","haohaohao",false,"style='width:100px'")
   fmt.Println(checkBox)
}