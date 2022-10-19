package main

import (
	"Belajar-go/helpers"
	"os"
)

func main() {
	var bio = helpers.Data()
	var indeks = os.Args[1]
	helpers.Hasil(bio, indeks)

}
