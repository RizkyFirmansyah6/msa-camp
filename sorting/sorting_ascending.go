package main

import (
	"fmt"
	"math"
	"sort"
)

// fungsi untuk menggambar vertical barchart
func DrawChart(slice []int) {
	max := math.MinInt64
	// mencari nilai tertinggi dari slice inputan
	for _, e := range slice {
		if e > max {
			max = e
		}
	}
	var (
		row = max
		col = len(slice)
	)
	// membuat array 2 dimensi [row * col]
	// baris
	chart := make([][]string, row)
	for i := 0; i < row; i++ {
		// kolom tiap baris
		chart[i] = make([]string, col)
	}
	// code untuk menggambar chart
	// mengisi dengan sting kosong jika tidak memiliki nilai ("   ") dan bar jika memiliki nilai " | "
	for i, e := range slice {
		for j := row - 1; j >= 0; j-- {
			if j >= row-e {
				chart[j][i] = " | "
			} else {
				chart[j][i] = "   "
			}
		}
	}
	// code untuk print chart
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%s", chart[i][j])
		}
		fmt.Printf("\n")
	}
	// print nilai pada sumbu horizontal chart
	for i := 0; i < col; i++ {
		fmt.Printf(" %d ", slice[i])
	}
	fmt.Printf("\n")
}

func main() {
	// deklarasi slice input dengan tipe data integer
	input := []int{1,4,5,6,8,2}
	ascending := make([]int, len(input))
	copy(ascending,input)
	// menampilkan slice setelah sorting ascending
	fmt.Println("Sorted array (ascending)")
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	fmt.Println(input)
	// sorting ascending menggunakan package sort dengan fungsi [i] < [j]
	// mencari nilai terkecil pada slice dan meletakkanya pada index awal slice
	fmt.Println("Step visualization")
	sort.Slice(ascending, func(i, j int) bool {
		DrawChart(ascending)
		return ascending[i] < ascending[j]
	})
}
