package goutils
import (
  "html/template"
  "fmt"
)

func Html_input_tag(tagType string,name string,value string,other_params ... interface{}) (html string) {
	params:=make(map[string]string)
	if(len(name)>0){
		params["name"]=name	
	}
	params["type"]=tagType	
	params["value"]=value
	params=params_merge(params,other_params)
	html = "<input "+paramsAsString(params)+"/>"
	return
}

func paramsAsString(more_params ...interface{}) string{
  params:=make(map[string]string)
  params=params_merge(params,more_params)
  html:=""
  for k,v:=range params{
		html+=" "+k+`="`+template.HTMLEscapeString(v)+`"`
	}
	return html
}

func params_merge(params map[string]string,more_params []interface{}) map[string]string{
  for _,param:=range more_params{
      switch param.(type){
      case map[string]string:
	      for k,v:=range param.(map[string]string){
	     			 params[k]=v
	      	}
	    case string:
	    _params:=StringToMap(fmt.Sprint(param))
	     for k,v:=range _params{
	     			 params[k]=v
	      	} 	
      }
  }
  return params
}

func Html_input_text(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("text",name,value,other_params ...)
}

func Html_input_hidden(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("hidden",name,value,other_params ...)
}
func Html_input_password(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("password",name,value,other_params ...)
}
func Html_input_email(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("email",name,value,other_params ...)
}
func Html_input_url(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("url",name,value,other_params ...)
}
func Html_input_search(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("search",name,value,other_params ...)
}
func Html_input_file(name string,value string,other_params ... interface{}) string{
	return Html_input_tag("file",name,value,other_params ...)
}
func Html_input_submit(value string,other_params ... interface{}) string{
	return Html_input_tag("submit","",value,other_params ...)
}
func Html_input_reset(value string,other_params ... interface{}) string{
	return Html_input_tag("reset","",value,other_params ...)
}

func Html_select(name string,selected string,values map[string]string,other_params ...interface{}) string{
   params:=make(map[string]string)
   params["name"]=name
   html:="<select"+paramsAsString(params_merge(params,other_params))+">\n"
   for k,v:=range values{
		option_fmt:="<option name='%s'%s>%s</option>\n"
		select_str:=""
		if(selected==k){
			select_str=" selected='selected'"
		}
		option:=fmt.Sprintf(option_fmt,template.HTMLEscapeString(k),select_str,template.HTMLEscapeString(v))
		html+=option
   }
   html+="</select>";
	return html
}

func Html_link(url string,text string,more_params ...interface{}) string{
 params:=make(map[string]string)
 params["href"]=url
 params["title"]=text
 html:="<a "+paramsAsString(params_merge(params,more_params))+">"+template.HTMLEscapeString(text)+"</a>"
 return html
}

func Html_checkBox(name string,value string,label string,isChecked bool,other_params ...interface{}) string{
	params:=make(map[string]string)
	if(isChecked){
		params["checked"]="checked"
	}
	params=params_merge(params,other_params)
	html:="<label>"+Html_input_tag("checkbox",name,value,params)+template.HTMLEscapeString(label)+"</label>"
	return html
}

func Html_datalist(id string,values []string) string{
  html:="<datalist id='"+id+"'>"
  for _,v:=range values{
 	 html+="<option value='"+template.HTMLEscapeString(v)+"'>"
  }
  html+="</datalist>"
  return html
}
