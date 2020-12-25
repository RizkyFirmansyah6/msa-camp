package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// deklarasi slice global untuk menampung size file
var fsize1, fsize2 []int64

// fungsi untuk membandingkan 2 antara direktori
func compare(path1, path2 string) {
	// slice untuk membandingkan file antar direktori
	var files1, files2 []string
	// fungsi untuk membaca semua file dari direktori source
	err := filepath.Walk(path1,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				//fmt.Println(path, info.Size())
				pat := strings.ReplaceAll(path, path1, "")
				files1 = append(files1, pat)
				fsize1 = append(fsize1, info.Size())
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	// fungsi untuk membaca semua file dari direktori target
	err1 := filepath.Walk(path2,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				//fmt.Println(path, info.Size())
				pat2 := strings.ReplaceAll(path, path2, "")
				files2 = append(files2, pat2)
				fsize2 = append(fsize2, info.Size())
			}
			return nil
		})
	if err1 != nil {
		log.Println(err1)
	}
	// membandingkan file source terhadap file target jika ada yang dihapus
	for i, e1 := range files1 {
		if !contains(files2, e1) {
			fmt.Println(e1, "DELETED")
		} else {
			// membandingkan size antar file souce dengan file target
			if checkSize(files2, e1, i) {
				fmt.Println(e1, "MODIFIED")
			}
		}
	}
	// membandingkan file target terhadap file source jika ada yang baru
	for _, e2 := range files2 {
		if !contains(files1, e2) {
			fmt.Println(e2, "NEW")
		}
	}
}

// fungsi untuk membandingkan antara file source dan destination begitu juga sebaliknya
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// fungsi untuk membandingkan ukuran file pada source dengan target
func checkSize(s []string, e string, idx int) bool {
	for i, a := range s {
		if a == e {
			if fsize1[idx] != fsize2[i] {
				return true
			}
		}
	}
	return false
}

// fungs main yang mengambil 2 parameter kemudian membadingkannya
func main() {
	var args = os.Args[1:]
	if len(args) != 2 {
		fmt.Println("program harus memiliki 2 parameter")
	} else {
		compare(args[0], args[1])
	}
}
