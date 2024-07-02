package service

import (
	"fmt"
	"sort"
)

func SecondFilter() {
	var input string
	fmt.Print("Input: ")
	fmt.Scanln(&input)

	subsequences := findSubsequences(input)
	for _, subseq := range subsequences {
		fmt.Println(subseq)
	}
}

// HasUnicodeEven Questa funzione calcola la somma delle rune di una sottosequenza e verifica se è un numero pari.
func hasUnicodeEven(s string) bool {
	sum := 0
	for _, char := range s {
		sum += int(char)
	}
	return sum%2 == 0
}

// FindSubsequences Questa funzione genera tutte le sottosequenze di almeno 2 caratteri della stringa s.
// Ogni sottosequenza verifica se ha la somma delle rune è pari.
// Le sottosequenze vengono ordinate in ordine alfabetico e restituite.
func findSubsequences(s string) []string {
	var result []string
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 2; j <= n; j++ {
			subseq := s[i:j]
			if hasUnicodeEven(subseq) {
				result = appendIfNotExists(result, subseq)
			}
		}
	}

	sort.Strings(result)

	return result
}

// questa funzione aggiunge un elemento a una slice solo se non è già presente, evitando duplicati nelle sottosequenze trovate.
func appendIfNotExists(slice []string, elem string) []string {
	for _, s := range slice {
		if s == elem {
			return slice
		}
	}
	return append(slice, elem)
}
