package main

import "fmt"

func mergeArrays(arr1 []string, arr2 []string) []string {
	merged := make([]string, 0)
	merged = append(merged, arr1...)

	for _, name := range arr2 {
		found := false

		for _, existingName := range merged {
			if name == existingName {
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, name)
		}
	}

	return merged
}

func main() {

	// Contoh array input
	//array1 := []string{"John", "Bob", "Alice"}
	//array2 := []string{"Alice", "Eve", "Bob", "Michael"}

	// Menggabungkan dua array
	//mergedArray := mergeArrays(array1, array2)

	// Menampilkan array hasil penggabungan
	//fmt.Println("Array hasil penggabungan:")
	//fmt.Println(mergedArray)

	//test cases
	array1 := []string{"king", "devil jin", "akuma"}
	array2 := []string{"eddie", "geese"}
	mergedArray1 := mergeArrays(array1, array2)
	fmt.Println("Hasil Penggabungan Array Ke 1 : ", mergedArray1)
	//fmt.Println(mergedArray1)
	fmt.Println("")
	//["king","devil jin","akuma","eddie","steve","geese"]

	array3 := []string{"sergei", "jin"}
	array4 := []string{"jin", "steve", "bryan"}
	mergedArray2 := mergeArrays(array3, array4)
	fmt.Println("Hasil Penggabungan Array Ke 2 : ", mergedArray2)
	//fmt.Println(mergedArray2)
	fmt.Println("")
	//fmt.Println(mergeArrays([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	//{"sergei","jin","steve","bryan"}

	array5 := []string{"alisa", "yoshimitsu"}
	array6 := []string{"devil jin", "yoshimitsu", "alisa", "law"}
	mergedArray3 := mergeArrays(array5, array6)
	fmt.Println("Hasil Penggabungan Array Ke 3 : ", mergedArray3)
	//fmt.Println(mergedArray3)
	fmt.Println("")
	//fmt.Println(mergeArrays([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	//["alisa","yoshimitsu","deviljin","law"]

	array7 := []string{}
	array8 := []string{"devil jin", "sergei"}
	mergedArray4 := mergeArrays(array7, array8)
	fmt.Println("Hasil Penggabungan Array Ke 4 : ", mergedArray4)
	fmt.Println("")
	//fmt.Println(mergeArrays([]string{}, []string{"devil jin", "sergei"}))
	//["devil jin","sergei"]

	array9 := []string{"hwoarang"}
	array10 := []string{}
	mergedArray5 := mergeArrays(array9, array10)
	fmt.Println("Hasil Penggabungan Array Ke 5 : ", mergedArray5)
	fmt.Println("")
	//fmt.Println(mergeArrays([]string{"hwoarang"}, []string{}))
	//["hwoarang"]

	array11 := []string{}
	array12 := []string{}
	mergedArray6 := mergeArrays(array11, array12)
	fmt.Println("Hasil Penggabungan Array Ke 6 : ", mergedArray6)
	fmt.Println("")
	//fmt.Println(mergeArrays([]string{}, []string{}))
	//[]

}
