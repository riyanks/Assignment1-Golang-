package helpers

import (
	"fmt"
	"strconv"
)

func Hasil(b []Person, c string) {
	var d, err = strconv.Atoi(c) //convert string to int
	_ = err

	if d <= len(b) {
		fmt.Println("Data ditemukan atas ")
		fmt.Println("Nama : ", b[d-1].Nama)
		fmt.Println("Alamat : ", b[d-1].Alamat)
		fmt.Println("Pekerjaan : ", b[d-1].Pekerjaan)
		fmt.Println("Alasan : ", b[d-1].Alasan)
	} else {
		fmt.Printf("Data ke %d tidak ditemukan", d)
	}

}
