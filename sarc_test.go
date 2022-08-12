package main

import "testing"

func TestSarcastic(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				str: "treasure",
			},
			want: "tReAsUrE",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "tHiS iS qUiTe CoMpLiCaTeD. wOuLdN't YoU aGrEe?",
		},
		{
			name: "multiline",
			args: args{
				str: `
This text spans
over MULTIPLE
lines!
`,
			},
			want: `
tHiS tExT sPaNs
OvEr MuLtIpLe
LiNeS!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Sarcastic(tt.args.str); got != tt.want {
				t.Errorf("Sarcastic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoModify(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				str: "treasure",
			},
			want: "treasure",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "This IS quite complicated. Wouldn't you agree?",
		},
		{
			name: "multiline",
			args: args{
				str: `
This text spans
over MULTIPLE
lines!
`,
			},
			want: `
This text spans
over MULTIPLE
lines!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NoModify(tt.args.str); got != tt.want {
				t.Errorf("NoModify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUppercase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uppercase(tt.args.str); got != tt.want {
				t.Errorf("Uppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowercase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lowercase(tt.args.str); got != tt.want {
				t.Errorf("Lowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.str); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapcase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Swapcase(tt.args.str); got != tt.want {
				t.Errorf("Swapcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallcaps(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Smallcaps(tt.args.str); got != tt.want {
				t.Errorf("Smallcaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpsideDown(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpsideDown(tt.args.str); got != tt.want {
				t.Errorf("UpsideDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStencil(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stencil(tt.args.str); got != tt.want {
				t.Errorf("Stencil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBall(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ball(tt.args.str); got != tt.want {
				t.Errorf("Ball() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursive(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cursive(tt.args.str); got != tt.want {
				t.Errorf("Cursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
