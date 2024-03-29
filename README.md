# sarc

A text transformation CLI.

## Installation

### Go

```bash
go install github.com/tom-on-the-internet/sarc@latest
```

### Nix

Use the flake.nix found in the root of this repository.

```bash
nix shell github:tom-on-the-internet/sarc
```

### Download the binary

Binaries for Mac, Linux and Windows are available in the [releases section](https://github.com/tom-on-the-internet/sarc/releases).

## Usage

```txt
sarc [-f formatter] [-i interactive]
```

### Usage Examples

```bash
echo 'Hello World!' | sarc # hElLo WoRlD!
echo 'Hello World!' | sarc -f cursive # 𝓗𝓮𝓵𝓵𝓸 𝓦𝓸𝓻𝓵𝓭!
echo 'Hello World!' | sarc -i # Enter interactive mode
```

## Formatters

Sample text `Hello World!`

```txt
ball          🅗🅔🅛🅛🅞 🅦🅞🅡🅛🅓!
cursive       𝓗𝓮𝓵𝓵𝓸 𝓦𝓸𝓻𝓵𝓭!
lowercase     hello world!
nomodify      Hello World!
reverse       !dlroW olleH
rot13         Uryyb Jbeyq!
sarcastic     hElLo WoRlD!
smallcaps     ʜᴇʟʟᴏ ᴡᴏʀʟᴅ!
stencil       ℍ𝕖𝕝𝕝𝕠 𝕎𝕠𝕣𝕝𝕕!
swapcase      hELLO wORLD!
uppercase     HELLO WORLD!
upsidedown    !pʅɹoM oʅʅǝH
```

## Video Explanation

[![IMAGE ALT TEXT](http://img.youtube.com/vi/V9dOyqtm3vc/0.jpg)](http://www.youtube.com/watch?v=V9dOyqtm3vc "Sarc Video")
