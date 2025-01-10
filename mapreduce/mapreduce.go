package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	filename := "./mapreduce/pg-dorian_gray.txt"
	content, err := os.ReadFile(filename)
	check(err)
	words1 := mapF(filename, string(content))

	filename = "./mapreduce/pg-being_ernest.txt"
	content, err = os.ReadFile(filename)
	check(err)
	words2 := mapF(filename, string(content))

	f, err := os.Create("./reduce.txt")
	check(err)
	defer f.Close()

	for k, v := range reduceF(words1, words2) {
		fmt.Printf("%s = %d\n", k, v)
		_, err := f.WriteString(fmt.Sprintf("%s = %d\n", k, v))
		check(err)
	}
}

func mapF(filename, content string) map[string]int {
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(content, f)

	m := make(map[string]int)
	for _, w := range words {
		m[w] = m[w] + 1
	}
	return m
}

func reduceF(map1, map2 map[string]int) map[string]int {
	m := make(map[string]int)
	for k, v1 := range map1 {
		v2 := map2[k]
		m[k] = v1 + v2
	}
	return m
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
