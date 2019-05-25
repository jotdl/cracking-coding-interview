package permutation

import "testing"

func TestIsPermutation(t *testing.T) {
	type args struct {
		src string
		txt string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "aabcde = bcaade",
			want: true,
			args: args{
				src: "aabcde",
				txt: "bcaade",
			},
		},
		{
			name: "aabcde = aabcde",
			want: true,
			args: args{
				src: "aabcde",
				txt: "aabcde",
			},
		},
		{
			name: "aabcde = cedaaba",
			want: false,
			args: args{
				src: "aabcde",
				txt: "cedaaba",
			},
		},
		{
			name: "aabcde = aabcdez",
			want: false,
			args: args{
				src: "aabcde",
				txt: "aabcdez",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutation(tt.args.src, tt.args.txt); got != tt.want {
				t.Errorf("IsPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
