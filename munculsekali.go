package main

import (
	"fmt"
	"strings"
)

func findUniqueNumbers(input string) []int {
	counts := make(map[int]int)

	// Memisahkan angka-angka dalam input
	numbers := strings.Fields(input)

	// Menghitung kemunculan setiap angka dalam input
	for _, numStr := range numbers {
		var num int
		_, err := fmt.Sscanf(numStr, "%d", &num)
		if err == nil {
			counts[num]++
		}
	}

	// Mengumpulkan angka-angka yang hanya muncul 1 kali
	uniqueNumbers := make([]int, 0)
	for num, count := range counts {
		if count == 1 {
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	return uniqueNumbers
}

func main() {
	satu()
	dua()
	tiga()
	empat()
	lima()
}

func satu() {
	// Input berisi kumpulan angka
	input := "1 2 3 4 1 2 3 "

	// Mencari angka-angka yang hanya muncul 1 kali
	uniqueNumbers := findUniqueNumbers(input)

	// Menampilkan hasil
	fmt.Print("Muncul Sekali :")
	fmt.Println(uniqueNumbers)

}

func dua() {
	input := "7 6 5 2 3 7 5 2"
	uniqueNumbers := findUniqueNumbers(input)
	fmt.Print("Muncul Sekali :")
	fmt.Println(uniqueNumbers)

}

func tiga() {
	input := "1 2 3 4 5"
	uniqueNumbers := findUniqueNumbers(input)
	fmt.Print("Muncul Sekali :")
	fmt.Println(uniqueNumbers)

}

func empat() {
	input := "1 1 2 2 3 3 4 4 5 5"
	uniqueNumbers := findUniqueNumbers(input)
	fmt.Print("Muncul Sekali :")
	fmt.Println(uniqueNumbers)

}

func lima() {
	input := "0 8 7 2 5 0 4"
	uniqueNumbers := findUniqueNumbers(input)
	fmt.Print("Muncul Sekali :")
	fmt.Println(uniqueNumbers)

}
