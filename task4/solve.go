package main

import (
	//"fmt"
	"unicode"
)

func RemoveEven(arr []int) []int{
	slice := make([]int, 0)
	for i := range arr {
		if arr[i] % 2 == 1 {
			slice = append(slice, arr[i])
		}
	}
	return slice
}

func PowerGenerator(a int) func() int {
	b := a
	return func() (pow int) {
		pow = b
		b = b * a
		return
	}
}

func DifferentWordsCount(str string) int{
	vocab := make(map[string] int)
	word := ""
	for i := range  str {
		if unicode.IsLetter(rune(str[i])) {
			word = word + string(unicode.ToLower(rune(str[i])))
		} else {
			if word != "" {
				vocab[word] += 1
			}
			word = ""
		}
	}
	if word != "" {
		vocab[word] += 1
	}
	word = ""
	return len(vocab)
}

/*
func main() {
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result)
	gen := PowerGenerator(3)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8
	fmt.Println(DifferentWordsCount("Hi, HHHHi"))
}*/
