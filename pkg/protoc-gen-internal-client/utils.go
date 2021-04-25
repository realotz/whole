package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

func MarshalStr(fullname string) protoreflect.FullName {
	name := string(fullname)
	name = strings.ReplaceAll(name, "_", ".")
	if name == "" {
		return ""
	}
	temp := strings.Split(name, ".")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	return protoreflect.FullName(s)
}

func Marshal(fullname protoreflect.FullName) protoreflect.FullName {
	name := string(fullname)
	name = strings.ReplaceAll(name, "_", ".")
	if name == "" {
		return ""
	}
	temp := strings.Split(name, ".")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	return protoreflect.FullName(s)
}

func commentDesc(comment protogen.CommentSet) string {
	str := comment.Leading.String()
	str = strings.ReplaceAll(str, "//", "")
	str = strings.ReplaceAll(str, "\n", "")
	return strings.TrimSpace(str)
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
