package util

import "github.com/speps/go-hashids"

type TypeNumber interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

func NewHashID() (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.Salt = ""
	hd.MinLength = 10

	return hashids.NewWithData(hd)
}

func SetID[T TypeNumber](id T) string {
	h, _ := NewHashID()
	idStr, _ := h.Encode([]int{int(id)})

	return idStr
}

func GetID(str string) int {
	h, _ := NewHashID()
	ids, _ := h.DecodeWithError(str)

	if len(ids) > 0 {
		return ids[0]
	}

	return 0
}

func BatchGetID(ids []string) []int {
	var list []int
	if len(ids) == 0 {
		return list
	}
	for _, id := range ids {
		list = append(list, GetID(id))
	}
	return list
}
