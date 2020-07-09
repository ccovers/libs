package pinyin

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

var args pinyin.Args = pinyin.NewArgs()

// 返回所有字符，汉字则转换为拼音，且多音字取第一个
func Pinyin(words string) string {
	pySlice := make([]string, 0)
	runes := []rune(strings.TrimSpace(words))
	for _, word := range runes {
		pys := pinyin.Pinyin(string([]rune{word}), args)
		if len(pys) > 0 && len(pys[0]) > 0 && len(pys[0][0]) > 0 {
			pySlice = append(pySlice, pys[0][0])
		} else {
			pySlice = append(pySlice, string([]rune{word}))
		}
	}
	return strings.Join(pySlice, "")
}

var firstLetters = getArgs()

func getArgs() pinyin.Args {
	args := pinyin.NewArgs()
	args.Style = pinyin.FirstLetter
	return args
}

// 返回简称，汉字取拼音首字母，且多音字取第一个
func PinyinShort(words string) string {
	runes := []rune(strings.TrimSpace(words))
	pinyinSlice := make([]rune, 0, len(runes))
	for _, word := range runes {
		pys := pinyin.SinglePinyin(word, firstLetters)
		if len(pys) > 0 && len(pys[0]) > 0 {
			word = rune(pys[0][0])
		}
		if word >= 'a' && word <= 'z' {
			word -= ('a' - 'A')
		}
		pinyinSlice = append(pinyinSlice, word)
	}
	return string(pinyinSlice)
}
