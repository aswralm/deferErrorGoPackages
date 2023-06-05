//Soal 1

package main

import (
	"fmt"
	"math"
)

var luasLingkaran float64
var kelilingLingkaran float64

func hitungLingkaran(jariJari *float64) {
	luasLingkaran = math.Pi * math.Pow(*jariJari, 2)
	kelilingLingkaran = 2 * math.Pi * *jariJari
}

func main() {
	var jariJari float64 = 7
	hitungLingkaran(&jariJari)
	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)
}

//Soal 2
package main

import (
	"fmt"
)

func introduce(sentence *string, name string, gender string, occupation string, age string) {
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
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

//Soal 3
package main

import (
	"fmt"
)

func tambahBuah(buah *[]string, namaBuah ...string) {
	*buah = append(*buah, namaBuah...)
}

func main() {
	var buah = []string{}

	tambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

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

type segitiga struct {
    alas, tinggi int
}

func (s segitiga) luas() int {
    return (s.alas * s.tinggi) / 2
}

type persegi struct {
    sisi int
}

func (p persegi) luas() int {
    return p.sisi * p.sisi
}

type persegiPanjang struct {
    panjang, lebar int
}

func (pp persegiPanjang) luas() int {
    return pp.panjang * pp.lebar
}

func main() {
    s := segitiga{3, 4}
    fmt.Println(s.luas())

    p := persegi{5}
    fmt.Println(p.luas())

    pp := persegiPanjang{2, 6}
    fmt.Println(pp.luas())
}

//Soal 6

package main

import "fmt"

type segitiga struct {
    alas, tinggi int
}

func (s segitiga) luas() int {
    return (s.alas * s.tinggi) / 2
}

type persegi struct {
    sisi int
}

func (p persegi) luas() int {
    return p.sisi * p.sisi
}

type persegiPanjang struct {
    panjang, lebar int
}

func (pp persegiPanjang) luas() int {
    return pp.panjang * pp.lebar
}

func main() {
    s := segitiga{3, 4}
    fmt.Println(s.luas())

    p := persegi{5}
    fmt.Println(p.luas())

    pp := persegiPanjang{2, 6}
    fmt.Println(pp.luas())
}

//Soal 7
package main

import "fmt"

type movie struct {
    title    string
    genre    string
    duration int
    year     int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
    *dataFilm = append(*dataFilm, movie{title, genre, duration, year})
}

func main() {
    var dataFilm = []movie{}

    tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
    tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
    tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
    tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

    for _, film := range dataFilm {
        fmt.Printf("%s %d %s %d\n", film.title, film.duration, film.genre, film.year)
    }
}

//Soal 8
package main

import (
    "fmt"
    "math"
)

type segitigaSamaSisi struct {
    alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
    return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
    return s.alas * 3
}

type persegiPanjang struct {
    panjang, lebar int
}

func (pp persegiPanjang) luas() int {
    return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
    return 2 * (pp.panjang + pp.lebar)
}

type tabung struct {
    jariJari, tinggi float64
}

func (t tabung) volume() float64 {
    return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
    return 2*math.Pi*t.jariJari*(t.jariJari+t.tinggi)
}

type balok struct {
    panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
    return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
    return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

type hitungBangunDatar interface {
    luas() int
    keliling() int
}

type hitungBangunRuang interface {
    volume() float64
    luasPermukaan() float64
}

func main() {
    s := segitigaSamaSisi{3, 4}
    fmt.Println(s.luas())
    fmt.Println(s.keliling())

    pp := persegiPanjang{2, 6}
    fmt.Println(pp.luas())
    fmt.Println(pp.keliling())

    t := tabung{3, 4}
    fmt.Println(t.volume())
    fmt.Println(t.luasPermukaan())

    b := balok{2, 3, 4}
    fmt.Println(b.volume())
    fmt.Println(b.luasPermukaan())
}

//Soal 9
package main

import (
    "fmt"
    "strings"
)

type phone struct {
    name, brand string
    year        int
    colors      []string
}

func (p phone) display() string {
    return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %s\n", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
}

type displayData interface {
    display() string
}

func main() {
    p := phone{"iPhone 14", "Apple", 2022, []string{"Black", "White"}}
    fmt.Println(p.display())
}

//Soal 10
package main

import "fmt"

func luasPersegi(sisi int, showDetail bool) interface{} {
    if sisi == 0 {
        if showDetail {
            return "Maaf anda belum menginput sisi dari persegi"
        } else {
            return nil
        }
    }

    luas := sisi * sisi
    if showDetail {
        return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
    } else {
        return luas
    }
}

func main() {
    fmt.Println(luasPersegi(4, true))
    fmt.Println(luasPersegi(8, false))
    fmt.Println(luasPersegi(0, true))
    fmt.Println(luasPersegi(0, false))
}

//Soal 11
package main

import (
    "fmt"
    "strings"
)

func main() {
    var prefix interface{} = "hasil penjumlahan dari "
    var kumpulanAngkaPertama interface{} = []int{6, 8}
    var kumpulanAngkaKedua interface{} = []int{12, 14}

    p := prefix.(string)
    kap := kumpulanAngkaPertama.([]int)
    kak := kumpulanAngkaKedua.([]int)

    angka := append(kap, kak...)
    total := 0
    for _, a := range angka {
        total += a
    }

    angkaStr := make([]string, len(angka))
    for i, a := range angka {
        angkaStr[i] = fmt.Sprint(a)
    }

    fmt.Printf("%s%s = %d\n", p, strings.Join(angkaStr, " + "), total)
}
