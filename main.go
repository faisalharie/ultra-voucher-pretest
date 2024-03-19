package main

import (
	"fmt"
)

//mengubah karakter string ke dalam rune, lalu urutkan sesuai abjad 
func sortStringRune(str string) string {
	s := []rune(str)
	for i := 1; i < len(s); i++ {
		key := s[i]
		j := i - 1
		for j >= 0 && s[j] > key {
			s[j+1] = s[j]
			j -= 1
		}
		s[j+1] = key
	}
	return string(s)
}

// membuat map of slice string, dengan sorted string sebagai key
// setiap kata yang berisi huruf sama setelah disort, diappend ke dalam slice
func groupAnagrams(arr []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range arr {
		sortedWord := sortStringRune(word)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func main() {
	arr := []string{"cook", "coko", "save", "taste", "aves", "vase", "state", "map"}
	groupedAnagrams := groupAnagrams(arr)
	fmt.Println(groupedAnagrams)
}
