package element

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func Test_genContent(t *testing.T) {
	tests := []struct {
		name string
		s    interface{}
		want string
	}{
		{
			"string",
			"test",
			" test ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genContent(tt.s); got != tt.want {
				t.Errorf("genContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_class_MarshalXML(t *testing.T) {
	tests := []struct {
		name string
		cl   class
		want string
	}{
		{
			"test",
			class{
				modifier:  "class",
				className: "Main",
			},
			`
<class>
  <keyword> class </keyword>
  <identifier> Main </identifier>
</class>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, _ := xml.MarshalIndent(tt.cl, "", "  ")
			// trim \n
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(string(output), want) {
				t.Errorf("class.MarshalXML() = %v", string(output))
				t.Errorf("wantXml = %v", want)
			}
		})
	}
}
