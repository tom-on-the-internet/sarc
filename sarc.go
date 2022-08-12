package main

import (
	"strings"
	"unicode"
)

func Sarcastic(str string) string {
	str = strings.ToLower(str)
	runes := []rune(str)
	upperCase := false

	for idx := range runes {
		if !unicode.IsLetter(runes[idx]) {
			continue
		}

		if upperCase {
			runes[idx] = unicode.ToUpper(runes[idx])
		}

		upperCase = !upperCase
	}

	return string(runes)
}

func NoModify(str string) string {
	return str
}

func Uppercase(str string) string {
	return strings.ToUpper(str)
}

func Lowercase(str string) string {
	return strings.ToLower(str)
}

func Reverse(str string) string {
	strs := strings.Split(str, "")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	str = strings.Join(strs, "")

	return str
}

func Swapcase(str string) string {
	runes := []rune(str)
	for idx, rune := range runes {
		if !unicode.IsLetter(rune) {
			continue
		}

		if unicode.IsUpper(rune) {
			runes[idx] = unicode.ToLower(rune)
			continue
		}

		runes[idx] = unicode.ToUpper(rune)
	}

	return string(runes)
}

func Smallcaps(str string) string {
	lookup := map[rune]rune{
		'a': 'ᴀ',
		'b': 'ʙ',
		'c': 'ᴄ',
		'd': 'ᴅ',
		'e': 'ᴇ',
		'f': 'ғ',
		'g': 'ɢ',
		'h': 'ʜ',
		'i': 'ɪ',
		'j': 'ᴊ',
		'k': 'ᴋ',
		'l': 'ʟ',
		'm': 'ᴍ',
		'n': 'ɴ',
		'o': 'ᴏ',
		'p': 'ᴘ',
		'q': 'ꞯ',
		'r': 'ʀ',
		's': 's',
		't': 'ᴛ',
		'u': 'ᴜ',
		'v': 'ᴠ',
		'w': 'ᴡ',
		'y': 'ʏ',
		'x': 'x',
		'z': 'ᴢ',
	}

	result := []rune{}

	for _, r := range str {
		if unicode.IsLetter(r) {
			result = append(result, lookup[unicode.ToLower(r)])
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func UpsideDown(str string) string {
	lookup := map[rune]rune{
		'a': 'ɐ',
		'b': 'q',
		'c': 'ɔ',
		'd': 'p',
		'e': 'ǝ',
		'f': 'ⅎ',
		'g': 'ƃ',
		'h': 'ɥ',
		'i': 'ᴉ',
		'j': 'ɾ',
		'k': 'ʞ',
		'l': 'ʅ',
		'm': 'ɯ',
		'n': 'u',
		'o': 'o',
		'p': 'd',
		'q': 'b',
		'r': 'ɹ',
		's': 's',
		't': 'ʇ',
		'u': 'n',
		'v': 'ʌ',
		'w': 'ʍ',
		'x': 'x',
		'y': 'ʎ',
		'z': 'z',
	}

	result := []rune{}

	str = Reverse(str)

	for _, r := range str {
		if unicode.IsLetter(r) {
			result = append(result, lookup[unicode.ToLower(r)])
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func Stencil(str string) string {
	lookup := map[rune]rune{
		'a': '𝕒',
		'b': '𝕓',
		'c': '𝕔',
		'd': '𝕕',
		'e': '𝕖',
		'f': '𝕗',
		'g': '𝕘',
		'h': '𝕙',
		'i': '𝕚',
		'j': '𝕛',
		'k': '𝕜',
		'l': '𝕝',
		'm': '𝕞',
		'n': '𝕟',
		'o': '𝕠',
		'p': '𝕡',
		'q': '𝕢',
		'r': '𝕣',
		's': '𝕤',
		't': '𝕥',
		'u': '𝕦',
		'v': '𝕧',
		'w': '𝕨',
		'y': '𝕪',
		'x': '𝕩',
		'z': '𝕫',
	}

	result := []rune{}

	for _, r := range str {
		if unicode.IsLetter(r) {
			result = append(result, lookup[unicode.ToLower(r)])
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func Ball(str string) string {
	lookup := map[rune]rune{
		'a': '🅐',
		'b': '🅑',
		'c': '🅒',
		'd': '🅓',
		'e': '🅔',
		'f': '🅕',
		'g': '🅖',
		'h': '🅗',
		'i': '🅘',
		'j': '🅙',
		'k': '🅚',
		'l': '🅛',
		'm': '🅜',
		'n': '🅝',
		'o': '🅞',
		'p': '🅟',
		'q': '🅠',
		'r': '🅡',
		's': '🅢',
		't': '🅣',
		'u': '🅤',
		'v': '🅥',
		'w': '🅦',
		'x': '🅧',
		'y': '🅨',
		'z': '🅩',
	}

	result := []rune{}

	for _, r := range str {
		if unicode.IsLetter(r) {
			result = append(result, lookup[unicode.ToLower(r)])
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func Cursive(str string) string {
	lookup := map[rune]rune{
		'a': '𝓪',
		'b': '𝓫',
		'c': '𝓬',
		'd': '𝓭',
		'e': '𝓮',
		'f': '𝓯',
		'g': '𝓰',
		'h': '𝓱',
		'i': '𝓲',
		'j': '𝓳',
		'k': '𝓴',
		'l': '𝓵',
		'm': '𝓶',
		'n': '𝓷',
		'o': '𝓸',
		'p': '𝓹',
		'q': '𝓺',
		'r': '𝓻',
		's': '𝓼',
		't': '𝓽',
		'u': '𝓾',
		'v': '𝓿',
		'w': '𝔀',
		'x': '𝔁',
		'y': '𝔂',
		'z': '𝔃',
	}

	result := []rune{}

	for _, r := range str {
		if unicode.IsLetter(r) {
			result = append(result, lookup[unicode.ToLower(r)])
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}
