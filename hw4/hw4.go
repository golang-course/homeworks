package hw4

import (
	"fmt"
	"sort"
	"strings"
)

func MostPopWords(text string) string {
	slice := strings.Split(strings.ToLower(text), " ")
	mapwords := make(map[string]int)
	for _, value := range slice {
		mapwords[value] = mapwords[value] + 1
	}
	type maptostruct struct {
		Key   string
		Value int
	}
	var sorted_str []maptostruct
	for key, value := range mapwords {
		sorted_str = append(sorted_str, maptostruct{key, value})
	}
	sort.Slice(sorted_str, func(i, j int) bool {
		return sorted_str[i].Value > sorted_str[j].Value
	})
	return fmt.Sprintf("Top 3 \n1. %s meets %v \n2. %s meets %v \n3. %s meets %v \n", sorted_str[0].Key, sorted_str[0].Value, sorted_str[1].Key, sorted_str[1].Value, sorted_str[2].Key, sorted_str[2].Value)
}
