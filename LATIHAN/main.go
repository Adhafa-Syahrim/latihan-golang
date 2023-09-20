package main

import (
	"fmt"
	"os"
	"strconv"
)

var path = "LATIHAN/invoice.txt"

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

func main() {
	createfile()
	var totalharga, kuantitas, pilih int
	fmt.Println("Program List Harga Barang")
	fmt.Print(`
	1. Ikan Bilis	Rp. 5.000
	2. Lemari	Rp. 10.000
	3. Sepeda	Rp. 15.000
	4. Laptop	Rp. 25.000
	`)
	fmt.Println("\nSilahkan pilih barang yang diinginkan =")
	fmt.Scanln(&pilih)
	fmt.Println("Silahkan masukkan jumlah barang yang dipilih =")
	fmt.Scanln(&kuantitas)

	if pilih == 1 {
		var barang1 = "ikan bilis"
		fmt.Printf("Barang yang anda pilih adalah %v dengan jumlah sebanyak %v\n", barang1, kuantitas)
		totalharga = kuantitas * 5000
		fmt.Printf("Harga yang harus anda bayar adalah Rp. %v\n", totalharga)
		cetakinvoice(barang1, kuantitas, totalharga)

	} else if pilih == 2 {
		var barang2 = "lemari"
		fmt.Printf("Barang yang anda pilih adalah %v dengan jumlah sebanyak %v\n", barang2, kuantitas)
		totalharga = kuantitas * 10000
		fmt.Printf("Harga yang harus anda bayar adalah Rp. %v", totalharga)
		cetakinvoice(barang2, kuantitas, totalharga)

	} else if pilih == 3 {
		var barang3 = "Sepeda"
		fmt.Printf("Barang yang anda pilih adalah %v dengan jumlah sebanyak %v\n", barang3, kuantitas)
		totalharga = kuantitas * 15000
		fmt.Printf("Harga yang harus anda bayar adalah Rp. %v", totalharga)
		cetakinvoice(barang3, kuantitas, totalharga)

	} else if pilih == 4 {
		var barang4 = "Laptop"
		fmt.Printf("Barang yang anda pilih adalah %v dengan jumlah sebanyak %v\n", barang4, kuantitas)
		totalharga = kuantitas * 25000
		fmt.Printf("Harga yang harus anda bayar adalah Rp. %v", totalharga)
		cetakinvoice(barang4, kuantitas, totalharga)

	}

}
