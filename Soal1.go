//Soal 1

package main

import (
	"fmt"
	"math"
)

var LuasLingkaran float64
var KelilingLingkaran float64

func HitungLingkaran(jariJari *float64) {
	LuasLingkaran = math.Pi * math.Pow(*jariJari, 2)
	KelilingLingkaran = 2 * math.Pi * *jariJari
}

func main() {
	var jariJari float64 = 7
	HitungLingkaran(&jariJari)
	fmt.Printf("Luas Lingkaran: %.2f\n", LuasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", KelilingLingkaran)
}

//Soal 2
package main

import (
	"fmt"
)

func Introduce(sentence *string, name string, gender string, occupation string, age string) {
	var title string
	if gender == "laki-laki" {
		title = "Pak"
	} else if gender == "perempuan" {
		title = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, occupation, age)
}

func main() {
	var sentence string
	Introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	Introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

//Soal 3
package main

import (
	"fmt"
)

func TambahBuah(buah *[]string, namaBuah ...string) {
	*buah = append(*buah, namaBuah...)
}

func main() {
	var buah = []string{}

	TambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}
}
//Soal 4
type Buah struct {
    Nama string
    Warna string
    AdaBijinya bool
    Harga int
}

func main() {
    nanas := Buah{"Nanas", "Kuning", false, 9000}
    jeruk := Buah{"Jeruk", "Oranye", true, 8000}
    semangka := Buah{"Semangka", "Hijau & Merah", true, 10000}
    pisang := Buah{"Pisang", "Kuning", false, 5000}

    fmt.Println(nanas)
    fmt.Println(jeruk)
    fmt.Println(semangka)
    fmt.Println(pisang)
}

//Soal 5
package main

import "fmt"

type Segitiga struct {
    Alas, Tinggi int
}

func (s Segitiga) Luas() int {
    return (s.Alas * s.Tinggi) / 2
}

type Persegi struct {
    Sisi int
}

func (p Persegi) Luas() int {
    return p.Sisi * p.Sisi
}

type PersegiPanjang struct {
    Panjang, Lebar int
}

func (pp PersegiPanjang) Luas() int {
    return pp.Panjang * pp.Lebar
}

func main() {
    s := Segitiga{3, 4}
    fmt.Println(s.Luas())

    p := Persegi{5}
    fmt.Println(p.Luas())

    pp := PersegiPanjang{2, 6}
    fmt.Println(pp.Luas())
}

//Soal 6

package main

import "fmt"

type Segitiga struct {
    Alas, Tinggi int
}

func (s Segitiga) Luas() int {
    return (s.Alas * s.Tinggi) / 2
}

type Persegi struct {
    Sisi int
}

func (p Persegi) Luas() int {
    return p.Sisi * p.Sisi
}

type PersegiPanjang struct {
    Panjang, Lebar int
}

func (pp PersegiPanjang) Luas() int {
    return pp.Panjang * pp.Lebar
}

func main() {
    s := Segitiga{3, 4}
    fmt.Println(s.Luas())

    p := Persegi{5}
    fmt.Println(p.Luas())

    pp := PersegiPanjang{2, 6}
    fmt.Println(pp.Luas())
}

//Soal 7
package main

import "fmt"

type Movie struct {
    Title    string
    Genre    string
    Duration int
    Year     int
}

func TambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]Movie) {
    *dataFilm = append(*dataFilm, Movie{title, genre, duration, year})
}

func main() {
    var dataFilm = []Movie{}

    TambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
    TambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
    TambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
    TambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

    for _, film := range dataFilm {
        fmt.Printf("%s %d %s %d\n", film.Title, film.Duration, film.Genre, film.Year)
    }
}

//Soal 8
package main

import (
    "fmt"
    "math"
)

type SegitigaSamaSisi struct {
    Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
    return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
    return s.Alas * 3
}
