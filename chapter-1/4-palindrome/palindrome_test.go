package palindrome

import "testing"

func TestIsPermutationOfPalindrom(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example",
			want: true,
			args: args{"Tact Coa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutationOfPalindrome(tt.args.src); got != tt.want {
				t.Errorf("IsPermutationOfPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
