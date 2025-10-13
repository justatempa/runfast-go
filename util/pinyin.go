package util

import (
	"strings"

	"github.com/Lofanmi/pinyin-golang/pinyin"
)

var dict = pinyin.NewDict()

func PinYin(ch string, pinyinStr string) bool {
	// wo shi yige zhong guo ren
	// quan := dict.Sentence("我是yige中国人").None()
	// wsyzgr
	// jian := dict.Abbr("我是yige中国人", "")
	if strings.Contains(ch, pinyinStr) {
		return true
	}
	res := dict.Sentence(ch).None()
	quan := strings.ReplaceAll(res, " ", "")
	jian := dict.Abbr(ch, "")
	return strings.Contains(quan, pinyinStr) || strings.Contains(jian, pinyinStr)

}
