package main

import (
	"fmt"
	"os"
	"strconv"
)

// deklarasi variabel lokasi file dan file yang dibuat
var path = "LATIHAN/invoice.txt"

// cetak error
func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)

}	

func createfile(){
    // deteksi apakah file sudah ada
    var _, err = os.Stat(path)
	fmt.Println("==> file sudah ada dengan nama", path)

    // buat file baru jika belum ada
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) { return }
        defer file.Close()
		fmt.Println("==> file berhasil dibuat", path)
    }
}

func cetakinvoice(barang string, jumlah int, harga int) {
    // buka file dengan level akses READ & WRITE
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) { return }
    defer file.Close()

    // tulis data ke file
    var invoice = barang + ", Sebanyak " + strconv.Itoa(jumlah) + " pcs, dengan total harga Rp." + strconv.Itoa(harga) +"\n"
    _, err = file.WriteString(invoice)
    if isError(err) { return }

    // simpan perubahan
    err = file.Sync()
    if isError(err) { return }

    fmt.Println("==> invoice berhasil ditambahkan")
}

// fungsi cetak cli
func cetakcli(barang string, kuantitas int, totalharga int){
	fmt.Printf("Barang yang anda pilih adalah %v dengan jumlah sebanyak %v\n", barang, kuantitas)
	fmt.Printf("Harga yang harus anda bayar adalah Rp. %v\n", totalharga)
}

func main() {
	// bikin file (panggil fungsi)
	createfile()

	// deklarasi variabel
	var totalharga, kuantitas, pilih int

	// print list barang dan harga
	fmt.Println("Program List Harga Barang")
	fmt.Print(`
	1. Ikan Bilis	Rp. 5.000
	2. Lemari	Rp. 10.000
	3. Sepeda	Rp. 15.000
	4. Laptop	Rp. 25.000
	`)

	// input pilihan
	fmt.Println("\nSilahkan pilih barang yang diinginkan =")
	fmt.Scanln(&pilih)
	fmt.Println("Silahkan masukkan jumlah barang yang dipilih =")
	fmt.Scanln(&kuantitas)

	// sistem pengkondisian jenis barang yang dipilih
	if pilih == 1 {
		barang1 := "ikan bilis"
		harga1 := 5000
		totalharga = kuantitas * harga1
		cetakcli(barang1, kuantitas, totalharga)
		cetakinvoice(barang1, kuantitas, totalharga) //panggil fungsi cetakinvoice

	} else if pilih == 2 {
		barang2 := "lemari"
		harga2 := 10000
		totalharga = kuantitas * harga2
		cetakcli(barang2, kuantitas, totalharga)
		cetakinvoice(barang2, kuantitas, totalharga) //panggil fungsi cetakinvoice

	} else if pilih == 3 {
		barang3 := "Sepeda"
		harga3 := 15000
		totalharga = kuantitas * harga3
		cetakcli(barang3, kuantitas, totalharga)
		cetakinvoice(barang3, kuantitas, totalharga) //panggil fungsi cetakinvoice

	} else if pilih == 4 {
		barang4 := "Laptop"
		harga4 := 25000
		totalharga = kuantitas * harga4
		cetakcli(barang4, kuantitas, totalharga)
		cetakinvoice(barang4, kuantitas, totalharga) //panggil fungsi cetakinvoice
	}
}
