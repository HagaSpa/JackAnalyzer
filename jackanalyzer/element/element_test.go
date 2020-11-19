package element

import "testing"

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
