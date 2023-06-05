//
package deferErrorGoPackages

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"
)

func PrintKalimat(kalimat string, tahun int) {
	defer fmt.Printf("%s %d\n", kalimat, tahun)
}
//
func KelilingSegitigaSamaSisi(sisi int, format bool) (string, error) {
	if sisi == 0 {
		if format {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}

	keliling := sisi * 3

	if format {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}

	return fmt.Sprintf("%d", keliling), nil
}
//
func CetakAngka(angka *int) {
	fmt.Println(*angka)
}

func TambahAngka(tambah int, angka *int) {
	*angka += tambah
}

func TambahDataPhone(phones *[]string, phone ...string) {
	*phones = append(*phones, phone...)
}

func UrutkanDanTampilkan(phones []string) {
	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}
}

func LuasLingkaran(jariJari float64) float64 {
	return math.Round(math.Pi * math.Pow(jariJari, 2))
}

func KelilingLingkaran(jariJari float64) float64 {
	return math.Round(2 * math.Pi * jariJari)
}
