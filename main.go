package main

import (
	"fmt"
	TugasPekan3 "TugasPekan3"
)

func main() {
	var jariJari float64 = 7
	TugasPekan3.HitungLingkaran(&jariJari)
	fmt.Printf("Luas Lingkaran: %.2f\n", TugasPekan3.LuasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", TugasPekan3.KelilingLingkaran)

	var sentence string
	TugasPekan3.Introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	TugasPekan3.Introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	var buah = []string{}

	TugasPekan3.TambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	nanas := TugasPekan3.Buah{"Nanas", "Kuning", false, 9000}
	jeruk := TugasPekan3.Buah{"Jeruk", "Oranye", true, 8000}
	semangka := TugasPekan3.Buah{"Semangka", "Hijau & Merah", true, 10000}
	pisang := TugasPekan3.Buah{"Pisang", "Kuning", false, 5000}

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	s := TugasPekan3.Segitiga{3, 4}
	fmt.Println(s.Luas())

	p := TugasPekan3.Persegi{5}
	fmt.Println(p.Luas())

	pp := TugasPekan3.PersegiPanjang{2, 6}
	fmt.Println(pp.Luas())

	var dataFilm = []TugasPekan3.Movie{}

	TugasPekan3.TambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	TugasPekan3.TambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	TugasPekan3.TambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	TugasPekan3.TambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for _, film := range dataFilm {
		fmt.Printf("%s %d %s %d\n", film.Title, film.Duration, film.Genre, film.Year)
	}
}
