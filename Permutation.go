
package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk menghasilkan permutasi stok barang
func permute(items []string, start int, result *[]string) {
	if start == len(items)-1 {
		*result = append(*result, "["+strings.Join(items, ", ")+"]")
		return
	}

	for i := start; i < len(items); i++ {
		items[start], items[i] = items[i], items[start]
		permute(items, start+1, result)
		items[start], items[i] = items[i], items[start] // Backtracking
	}
}

func main() {
	stocks := []string{"Shampo", "Detergen", "Lotion", "Sabun"}
	var result []string

	permute(stocks, 0, &result)

	// Cetak jumlah total permutasi
	fmt.Printf("Total Permutasi: %d\n", len(result))

	// Cetak hasil permutasi stok barang
	fmt.Println("Permutasi Stok Barang:")
	for _, perm := range result {
		fmt.Println(perm)
	}

}
