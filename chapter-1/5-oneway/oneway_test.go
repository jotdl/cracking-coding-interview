package oneway

import "testing"

func TestIsOneWay(t *testing.T) {
	type args struct {
		original string
		after    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"pale", "ple"}, true},
		{"2", args{"pales", "pale"}, true},
		{"3", args{"pale", "bale"}, true},
		{"4", args{"pale", "bake"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOneWay(tt.args.original, tt.args.after); got != tt.want {
				t.Errorf("IsOneWay() = %v, want %v", got, tt.want)
			}
		})
	}
}
