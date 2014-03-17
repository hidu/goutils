package goutils
import (
	"fmt"
//	"github.com/bmizerany/assert"
	"testing"
//	"reflect"
)
func TestHtml_input_text(t *testing.T){
   test:=Html_input_text("a","b","style='color:red'","class='c'")
   fmt.Println(test)
   html_option:=new(Html_Options)
   html_option.AddOption("a","b",false)
   sele:=Html_select("a",html_option,"id='a'")
   fmt.Println(sele)
//    fmt.Println("type:",c.Kind())
   link:=Html_link("http://www.baidu.com","baidu-百度")
   fmt.Println(link)
   
   checkBox:=Html_checkBox("name","good","haohaohao",false,"style='width:100px'")
   fmt.Println(checkBox)
}