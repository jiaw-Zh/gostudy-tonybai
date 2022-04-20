package main

import (
	"fmt"
	// "reflect"
	"unicode/utf8"
	// "unsafe"
)

// var s = "中国人"

// func main()  {
// 	fmt.Printf("the length of s is %d\n",len(s))

// 	for i:=0;i<len(s);i++{
// 		fmt.Printf("0x%x ",s[i])
// 	}
// 	fmt.Printf("\n")

// 	fmt.Printf("the character count in s is %d\n", utf8.RuneCountInString(s))

// 	for _,c := range s{
// 		fmt.Printf("0x%x ",c)
// 	}
// 	fmt.Printf("\n")
// }

// rune -> []byte
func encodeRune() {
	var r rune = 0x4e2d
	fmt.Printf("the unicode character is %c\n", r)
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r) //对rune进行utf-8编码
	fmt.Printf("utf-8 representation is 0x%X\n", buf)
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xad}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the utf-8 characters after decode is %s\n", string(r))
}

func dumpBytesArry(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func main() {
	var s = "hello"
	// hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// fmt.Printf("0x%x\n",hdr.Data)
	// p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	// dumpBytesArry((*p)[:])
	// for i :=0;i<len(s);i++{
	// 	fmt.Printf("%c\n",s[i])
	// }
	fmt.Printf("%c", s[0])
}
