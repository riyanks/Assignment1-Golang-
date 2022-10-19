package helpers

func Data() []Person {
	var p1 = Person{
		Nama:      "Riyan",
		Alamat:    "Way Huwi",
		Pekerjaan: "programer",
		Alasan:    "menambah wawasan mengenai go-lang",
	}
	var p2 = Person{
		Nama:      "Hafiz",
		Alamat:    "Medan",
		Pekerjaan: "Koki",
		Alasan:    "menambah wawasan mengenai go-lang",
	}
	var p3 = Person{
		Nama:      "Dzakir",
		Alamat:    "Bandar Lampung",
		Pekerjaan: "pengusaha",
		Alasan:    "menambah wawasan mengenai go-lang",
	}
	var p4 = Person{
		Nama:      "Thor",
		Alamat:    "Jakarta timur",
		Pekerjaan: "programer",
		Alasan:    "menambah wawasan mengenai go-lang",
	}

	var p5 = Person{
		Nama:      "Nisa",
		Alamat:    "Medan",
		Pekerjaan: "Pelajar",
		Alasan:    "Ingin mendapatkan pekerjaan",
	}
	var biodata = []Person{p1, p2, p3, p4, p5}

	return biodata
}
