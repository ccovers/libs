package pinyin

import (
	"fmt"
	"testing"
)

func TestPinyin(t *testing.T) {
	words := "你好吗 world?"
	fmt.Println(words, "--->", Pinyin(words))
	fmt.Println(words, "--->", PinyinShort(words))
}

func ExamplePinyin() {
	words := "你好吗 world?"
	fmt.Println(words, "--->", Pinyin(words))
	// Output:
	// 你好吗 world? ---> nihaoma world?
}
