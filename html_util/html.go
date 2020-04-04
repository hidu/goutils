package html_util

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"sort"

	"github.com/hidu/goutils/str_util"
)

func InputTag(tagType string, name string, value string, otherParams ...interface{}) (html string) {
	params := make(map[string]string)
	if len(name) > 0 {
		params["name"] = name
	}
	params["type"] = tagType
	params["value"] = value
	params = paramsMerge(params, otherParams)
	var buf bytes.Buffer
	buf.WriteString("<input")
	buf.Write(paramsAsBytes(params))
	buf.WriteString("/>")
	return buf.String()
}

func paramsAsBytes(moreParams ...interface{}) []byte {
	params := make(map[string]string)
	params = paramsMerge(params, moreParams)
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	// 排序，保证输出的属性字段是固定顺序的
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	var buf bytes.Buffer

	for _, key := range keys {
		val := params[key]

		buf.WriteString(" ")
		buf.WriteString(key)
		buf.WriteString(`="`)
		buf.WriteString(template.HTMLEscapeString(val))
		buf.WriteString(`"`)
	}
	return buf.Bytes()
}

func paramsMerge(params map[string]string, moreParams []interface{}) map[string]string {
	for _, param := range moreParams {
		switch param.(type) {
		case map[string]string:
			for k, v := range param.(map[string]string) {
				params[k] = v
			}
		case string:
			_params := str_util.StringToMap(fmt.Sprint(param))
			for k, v := range _params {
				params[k] = v
			}
		default:
			panic(fmt.Sprintf("param type not supported: %T:%v", param, param))
		}
	}
	return params
}

func InputText(name string, value string, params ...interface{}) string {
	return InputTag("text", name, value, params...)
}

func InputHidden(name string, value string, params ...interface{}) string {
	return InputTag("hidden", name, value, params...)
}
func InputPassword(name string, value string, params ...interface{}) string {
	return InputTag("password", name, value, params...)
}
func InputEmail(name string, value string, params ...interface{}) string {
	return InputTag("email", name, value, params...)
}
func InputURL(name string, value string, params ...interface{}) string {
	return InputTag("url", name, value, params...)
}
func InputSearch(name string, value string, params ...interface{}) string {
	return InputTag("search", name, value, params...)
}
func InputFile(name string, value string, params ...interface{}) string {
	return InputTag("file", name, value, params...)
}
func InputSubmit(value string, params ...interface{}) string {
	return InputTag("submit", "", value, params...)
}
func InputReset(value string, params ...interface{}) string {
	return InputTag("reset", "", value, params...)
}

func Link(url string, text string, moreParams ...interface{}) string {
	params := make(map[string]string)
	params["href"] = url
	params["title"] = text

	var buf bytes.Buffer
	buf.WriteString("<a")
	buf.Write(paramsAsBytes(paramsMerge(params, moreParams)))
	buf.WriteString(">")
	buf.WriteString(template.HTMLEscapeString(text))
	buf.WriteString("</a>")

	return buf.String()
}

func CheckBox(name string, value string, label string, isChecked bool, params ...interface{}) string {
	allParams := make(map[string]string)
	if isChecked {
		allParams["checked"] = "checked"
	}
	allParams = paramsMerge(allParams, params)
	html := "<label>" + InputTag("checkbox", name, value, allParams) + template.HTMLEscapeString(label) + "</label>"
	return html
}

func DataList(id string, values []string) string {
	var buf bytes.Buffer
	buf.WriteString("<datalist id='")
	buf.WriteString(template.HTMLEscapeString(id))
	buf.WriteString("'>")

	for _, v := range values {
		buf.WriteString("<option value='")
		buf.WriteString(template.HTMLEscapeString(v))
		buf.WriteString("'>")
	}
	buf.WriteString("</datalist>")
	return buf.String()
}

func TextArea(name string, value string, params ...interface{}) string {
	allParams := make(map[string]string)
	allParams["name"] = name
	allParams["value"] = value

	var buf bytes.Buffer
	buf.WriteString("<textarea")
	buf.Write(paramsAsBytes(paramsMerge(allParams, params)))
	buf.WriteString(">")
	buf.WriteString(template.HTMLEscapeString(value))
	buf.WriteString("</textarea>")

	return buf.String()
}

type Options struct {
	Items []*htmlOption
}

type htmlOption struct {
	Value   interface{}
	Txt     string
	Checked bool
	Params  map[string]string
}

func NewOptions() *Options {
	return new(Options)
}

func (options *Options) AddOption(txt string, value interface{}, checked bool) {
	option := new(htmlOption)
	option.Txt = txt
	option.Value = value
	option.Checked = checked
	options.Items = append(options.Items, option)
}

func Select(name string, options *Options, params ...interface{}) string {
	allParams := make(map[string]string)
	allParams["name"] = name

	var buf bytes.Buffer
	buf.WriteString("<select")
	buf.Write(paramsAsBytes(paramsMerge(allParams, params)))
	buf.WriteString(">\n")

	for _, option := range options.Items {
		buf.WriteString("<option value='")

		value := template.HTMLEscapeString(fmt.Sprint(option.Value))
		buf.WriteString(value)
		buf.WriteString("'")
		if option.Checked {
			buf.WriteString(" selected='selected'")
		}
		if len(option.Params) > 0 {
			buf.Write(paramsAsBytes(option.Params))
		}
		buf.WriteString(">")
		buf.WriteString(template.HTMLEscapeString(option.Txt))
		buf.WriteString("</option>\n")
	}
	buf.WriteString("</select>")
	return buf.String()
}

var htmlTagReg = regexp.MustCompile(`>\s+<`)

// ReduceHTMLSpace 去除html tag 间的空格字符
func ReduceHTMLSpace(html string) string {
	return htmlTagReg.ReplaceAllString(html, "><")
}
