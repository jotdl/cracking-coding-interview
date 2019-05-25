package urlify

import "testing"

func TestURLify(t *testing.T) {
	type args struct {
		url       string
		curLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Mr John Smith",
			want: "Mr%20John%20Smith",
			args: args{"Mr John Smith    ", 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.args.url, tt.args.curLength); got != tt.want {
				t.Errorf("URLify() = %v, want %v", got, tt.want)
			}
		})
	}
}
