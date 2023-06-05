package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // Soal 1
    var nama string = "AswarAlam"
    var email string = "Aswaralam101@gmail.com"
    var sistemOperasi string = "Windows"
    var akunGit string = "aswralm"
    var akunTelegram string = "unkwnon"

    fmt.Println("Data Peserta Bootcamp Digital Skill - Golang Backend Development")
    fmt.Println("1. Nama:", nama)
    fmt.Println("2. Email:", email)
    fmt.Println("3. Sistem Operasi yang digunakan:", sistemOperasi)
    fmt.Println("4. Akun Gitlab/Github:", akunGit)
    fmt.Println("5. Akun Telegram:", akunTelegram)

    // Soal 2
    halo := "Halo Dunia"
    halo = strings.Replace(halo, "Dunia", "Golang", 1)
    fmt.Println(halo)

    // Soal 3
    var kataPertama string = "saya"
    var kataKedua string = "senang"
    var kataKetiga string = "belajar"
    var kataKeempat string = "golang"

    kataKedua = strings.Title(kataKedua)
    kataKetiga = strings.Replace(kataKetiga, "a", "A", 1)
    kataKeempat = strings.ToUpper(kataKeempat)

    fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

    // Soal 4
    var kataPertama string = "saya"
    var kataKedua string = "senang"
    var kataKetiga string = "belajar"
    var kataKeempat string = "golang"

    kataKedua = strings.Title(kataKedua)
    kataKetiga = strings.Replace(kataKetiga, "a", "A", 1)
    kataKeempat = strings.ToUpper(kataKeempat)

    fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

    // Soal 5
    kalimat := "halo halo bandung"
    kalimat = strings.Title(kalimat)
    
    kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
    result := fmt.Sprintf("%s - %d", kalimat, angka)

    fmt.Println(result)


    // Soal 6
    var panjangPersegiPanjang string = "8"
    var lebarPersegiPanjang string = "5"

    var alasSegitiga string = "6"
    var tinggiSegitiga string = "7"

    panjangPersegiPanjangInt, _ := strconv.Atoi(panjangPersegiPanjang)
    lebarPersegiPanjangInt, _ := strconv.Atoi(lebarPersegiPanjang)
    alasSegitigaInt, _ := strconv.Atoi(alasSegitiga)
    tinggiSegitigaInt, _ := strconv.Atoi(tinggiSegitiga)

    var luasPersegiPanjang int = panjangPersegiPanjangInt * lebarPersegiPanjangInt
    var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInt + lebarPersegiPanjangInt)
    var luasSegitiga int = (alasSegitigaInt * tinggiSegitigaInt) / 2

    fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
    fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
    fmt.Println("Luas Segitiga:", luasSegitiga)
    
    // Soal 7

    var nilaiJohn = 80
    var nilaiDoe = 50

    if nilaiJohn >= 80 {
        fmt.Println("John's index is A")
    } else if nilaiJohn >= 70 && nilaiJohn < 80 {
        fmt.Println("John's index is B")
    } else if nilaiJohn >= 60 && nilaiJohn < 70 {
        fmt.Println("John's index is C")
    } else if nilaiJohn >= 50 && nilaiJohn < 60 {
        fmt.Println("John's index is D")
    } else {
        fmt.Println("John's index is E")
    }

    if nilaiDoe >= 80 {
        fmt.Println("Doe's index is A")
    } else if nilaiDoe >= 70 && nilaiDoe < 80 {
        fmt.Println("Doe's index is B")
    } else if nilaiDoe >= 60 && nilaiDoe < 70 {
        fmt.Println("Doe's index is C")
    } else if nilaiDoe >= 50 && nilaiDoe < 60 {
        fmt.Println("Doe's index is D")
    } else {
        fmt.Println("Doe's index is E")
    }

    //Soal 8      
    var tanggal = 9
    var bulan = 10
    var tahun = 1999

    var namaBulan string

    switch bulan {
    case 1:
        namaBulan = "Januari"
    case 2:
        namaBulan = "Februari"
    case 3:
        namaBulan = "Maret"
    case 4:
        namaBulan = "April"
    case 5:
        namaBulan = "Mei"
    case 6:
        namaBulan = "Juni"
    case 7:
        namaBulan = "Juli"
    case 8:
        namaBulan = "Agustus"
    case 9:
        namaBulan = "September"
    case 10:
        namaBulan = "Oktober"
    case 11:
        namaBulan = "November"
    case 12:
        namaBulan = "Desember"
    }

    result := fmt.Sprintf("%d %s %d", tanggal, namaBulan, tahun)

    fmt.Println(result)

    //Soal 9

    var tahunKelahiran = 1999

    if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
        fmt.Println("Baby Boomer")
    } else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
        fmt.Println("Generasi X")
    } else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
        fmt.Println("Generasi Y (Millenials)")
    } else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
        fmt.Println("Generasi Z")
    } else {
        fmt.Println("Tidak termasuk dalam generasi yang terdaftar")
    }