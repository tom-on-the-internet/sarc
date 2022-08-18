package main

import (
	"testing"
)

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
			want: "TREASURE",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "THIS IS QUITE COMPLICATED. WOULDN'T YOU AGREE?",
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
THIS TEXT SPANS
OVER MULTIPLE
LINES!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Uppercase(tt.args.str); got != tt.want {
				t.Errorf("Uppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowercase(t *testing.T) {
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
			want: "this is quite complicated. wouldn't you agree?",
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
this text spans
over multiple
lines!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Lowercase(tt.args.str); got != tt.want {
				t.Errorf("Lowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
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
			want: "erusaert",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "?eerga uoy t'ndluoW .detacilpmoc etiuq SI sihT",
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
!senil
ELPITLUM revo
snaps txet sihT
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Reverse(tt.args.str); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapcase(t *testing.T) {
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
			want: "TREASURE",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "tHIS is QUITE COMPLICATED. wOULDN'T YOU AGREE?",
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
tHIS TEXT SPANS
OVER multiple
LINES!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Swapcase(tt.args.str); got != tt.want {
				t.Errorf("Swapcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallcaps(t *testing.T) {
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
			want: "ᴛʀᴇᴀsᴜʀᴇ",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "ᴛʜɪs ɪs ꞯᴜɪᴛᴇ ᴄᴏᴍᴘʟɪᴄᴀᴛᴇᴅ. ᴡᴏᴜʟᴅɴ'ᴛ ʏᴏᴜ ᴀɢʀᴇᴇ?",
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
ᴛʜɪs ᴛᴇxᴛ sᴘᴀɴs
ᴏᴠᴇʀ ᴍᴜʟᴛɪᴘʟᴇ
ʟɪɴᴇs!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Smallcaps(tt.args.str); got != tt.want {
				t.Errorf("Smallcaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpsideDown(t *testing.T) {
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
			want: "ǝɹnsɐǝɹʇ",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "?ǝǝɹƃɐ noʎ ʇ'upʅnoM .pǝʇɐɔᴉʅdɯoɔ ǝʇᴉnb SI sᴉɥ┴",
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
!sǝuᴉʅ
Ǝ˥ԀI┴˥∩W ɹǝʌo
suɐds ʇxǝʇ sᴉɥ┴
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := UpsideDown(tt.args.str); got != tt.want {
				t.Errorf("UpsideDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStencil(t *testing.T) {
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
			want: "𝕥𝕣𝕖𝕒𝕤𝕦𝕣𝕖",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "𝕋𝕙𝕚𝕤 𝕀𝕊 𝕢𝕦𝕚𝕥𝕖 𝕔𝕠𝕞𝕡𝕝𝕚𝕔𝕒𝕥𝕖𝕕. 𝕎𝕠𝕦𝕝𝕕𝕟'𝕥 𝕪𝕠𝕦 𝕒𝕘𝕣𝕖𝕖?",
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
𝕋𝕙𝕚𝕤 𝕥𝕖𝕩𝕥 𝕤𝕡𝕒𝕟𝕤
𝕠𝕧𝕖𝕣 𝕄𝕌𝕃𝕋𝕀ℙ𝕃𝔼
𝕝𝕚𝕟𝕖𝕤!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Stencil(tt.args.str); got != tt.want {
				t.Errorf("Stencil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBall(t *testing.T) {
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
			want: "🅣🅡🅔🅐🅢🅤🅡🅔",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "🅣🅗🅘🅢 🅘🅢 🅠🅤🅘🅣🅔 🅒🅞🅜🅟🅛🅘🅒🅐🅣🅔🅓. 🅦🅞🅤🅛🅓🅝'🅣 🅨🅞🅤 🅐🅖🅡🅔🅔?",
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
🅣🅗🅘🅢 🅣🅔🅧🅣 🅢🅟🅐🅝🅢
🅞🅥🅔🅡 🅜🅤🅛🅣🅘🅟🅛🅔
🅛🅘🅝🅔🅢!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Ball(tt.args.str); got != tt.want {
				t.Errorf("Ball() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursive(t *testing.T) {
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
			want: "𝓽𝓻𝓮𝓪𝓼𝓾𝓻𝓮",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "𝓣𝓱𝓲𝓼 𝓘𝓢 𝓺𝓾𝓲𝓽𝓮 𝓬𝓸𝓶𝓹𝓵𝓲𝓬𝓪𝓽𝓮𝓭. 𝓦𝓸𝓾𝓵𝓭𝓷'𝓽 𝔂𝓸𝓾 𝓪𝓰𝓻𝓮𝓮?",
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
𝓣𝓱𝓲𝓼 𝓽𝓮𝔁𝓽 𝓼𝓹𝓪𝓷𝓼
𝓸𝓿𝓮𝓻 𝓜𝓤𝓛𝓣𝓘𝓟𝓛𝓔
𝓵𝓲𝓷𝓮𝓼!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Cursive(tt.args.str); got != tt.want {
				t.Errorf("Cursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRot13(t *testing.T) {
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
			want: "gernfher",
		},
		{
			name: "complicated",
			args: args{
				str: "This IS quite complicated. Wouldn't you agree?",
			},
			want: "Guvf VF dhvgr pbzcyvpngrq. Jbhyqa'g lbh nterr?",
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
Guvf grkg fcnaf
bire ZHYGVCYR
yvarf!
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Rot13(tt.args.str); got != tt.want {
				t.Errorf("Rot13() = %v, want %v", got, tt.want)
			}
		})
	}
}
