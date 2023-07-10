package main

import "fmt"

func Mapping(slice []string, targets []string) map[string]int {
	occurrences := make(map[string]int)

	for _, target := range targets {
		occurrences[target] = 0

		for _, str := range slice {
			if str == target {
				occurrences[target]++
			}
		}
	}

	return occurrences
}

func main() {

	perintah1()
	perintah2()
}

func perintah1() {
	
	// slice input2 target string
	slice := []string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}

	// String-target yang ingin dihitung kemunculannya
	targetStrings := []string{"asd", "qwe"}

	// Menghitung kemunculan string-target dalam slice
	occurrences := Mapping(slice, targetStrings)

	// Menampilkan hasil
	for target, count := range occurrences {
		//fmt.Printf("Jumlah kemunculan string '%s' dalam slice: %d\n", target, count)
		fmt.Printf("[Map '%s' : %d]\n", target, count)
	}
	fmt.Println(" ")
	fmt.Println("[Map : ]")
	//fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	//map[asd:2 qwe:3]
	//-----------------------------------------------------------------------------------------
}

func perintah2() {
	slice := []string{}
	targetStrings := []string{}
	occurrences := Mapping(slice, targetStrings)
	for target, count := range occurrences {
		fmt.Printf("[Map '%s' : %d]\n", target, count)
	}
	//fmt.Println(Mapping([]string{}))
	//map[]
}
