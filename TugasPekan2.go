//Soal 1

package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        if i%3 == 0 && i%2 != 0 {
            fmt.Println(i, "- I Love Coding")
        } else if i%2 == 0 {
            fmt.Println(i, "- Berkualitas")
        } else {
            fmt.Println(i, "- Santai")
        }
    }
}

//Soal 2

package main

import "fmt"

func main() {
    for i := 0; i < 7; i++ {
        for j := 0; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}

//Soal 3
package main

import (
    "fmt"
    "strings"
)

func main() {
    kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
    result := strings.Join(kalimat[2:], " ")
    fmt.Println("[" + result + "]")
}

//Soal 4
package main

import "fmt"

func main() {
    sayuran := []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
    for i, sayur := range sayuran {
        fmt.Printf("%d. %s\n", i+1, sayur)
    }
}

//Soal 5
package main

import "fmt"

func main() {
    satuan := map[string]int{
        "panjang": 7,
        "lebar": 4,
        "tinggi": 6,
    }
    for k, v := range satuan {
        fmt.Printf("%s: %d\n", k, v)
    }
}
//Soal  6
package main

import "fmt"

func luasPersegiPanjang(panjang int, lebar int) int {
    return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
    return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
    return panjang * lebar * tinggi
}

func main() {
    panjang := 12
    lebar := 4
    tinggi := 8

    luas := luasPersegiPanjang(panjang, lebar)
    keliling := kelilingPersegiPanjang(panjang, lebar)
    volume := volumeBalok(panjang, lebar, tinggi)

    fmt.Println(luas)
    fmt.Println(keliling)
    fmt.Println(volume)
}

//Soal 7
package main

import (
	"fmt"
)

func introduce(name string, gender string, occupation string, age string) string {
	var title string
	if gender == "laki-laki" {
		title = "Pak"
	} else if gender == "perempuan" {
		title = "Bu"
	}
	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, occupation, age)
}

func main() {
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

//Soal 8
package main

import (
	"fmt"
	"strings"
)

func buahFavorit(name string, buah ...string) string {
	buahQuoted := make([]string, len(buah))
	for i, b := range buah {
		buahQuoted[i] = fmt.Sprintf("%q", b)
	}
	return fmt.Sprintf("Halo nama saya %s dan buah favorit saya adalah %s", name, strings.Join(buahQuoted, ", "))
}

func main() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
}

//Soal 9
package main

import (
	"fmt"
)

func main() {
	var dataFilm = []map[string]string{}

	tambahDataFilm := func(judul string, durasi string, genre string, tahun string) {
		dataFilm = append(dataFilm, map[string]string{
			"judul":  judul,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
