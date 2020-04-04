package html_util

import (
	"testing"
)

func TestInputText(t *testing.T) {
	type args struct {
		name   string
		value  string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				name:   "abc",
				value:  "dev",
				params: nil,
			},
			want: `<input name="abc" type="text" value="dev"/>`,
		},
		{
			name: "case 2",
			args: args{
				name:  "abc",
				value: "dev",
				params: []interface{}{
					`style="color:red"`,
					"a='b'",
				},
			},
			want: `<input a="b" name="abc" style="color:red" type="text" value="dev"/>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InputText(tt.args.name, tt.args.value, tt.args.params...); got != tt.want {
				t.Errorf("InputText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	type args struct {
		name    string
		options *Options
		params  []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				name:    "abc",
				options: &Options{Items: []*htmlOption{}},
				params:  nil,
			},
			want: `<select name="abc">
</select>`,
		},
		{
			name: "case 2",
			args: args{
				name: "abc",
				options: &Options{Items: []*htmlOption{
					{
						Value:   "123",
						Txt:     "你好",
						Checked: false,
						Params:  nil,
					},
					{
						Value:   "234",
						Txt:     "你好2",
						Checked: true,
						Params: map[string]string{
							"style": "color:red",
						},
					},
				}},
				params: nil,
			},
			want: `<select name="abc">
<option value='123'>你好</option>
<option value='234' selected='selected' style="color:red">你好2</option>
</select>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.name, tt.args.options, tt.args.params...); got != tt.want {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink(t *testing.T) {
	type args struct {
		url        string
		text       string
		moreParams []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				url:        "golang.org",
				text:       "gogogo",
				moreParams: nil,
			},
			want: `<a href="golang.org" title="gogogo">gogogo</a>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Link(tt.args.url, tt.args.text, tt.args.moreParams...); got != tt.want {
				t.Errorf("Link() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextArea(t *testing.T) {
	type args struct {
		name   string
		value  string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				name:   "abc",
				value:  "text",
				params: nil,
			},
			want: `<textarea name="abc" value="text">text</textarea>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextArea(tt.args.name, tt.args.value, tt.args.params...); got != tt.want {
				t.Errorf("TextArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
