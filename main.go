package latihanmodule

import "fmt"

//Fungsi Parameter

func Judul(jenisKue string, rasaKue string) {
	fmt.Println("Cara Membuat", jenisKue, rasaKue)

}

//Fungsi Sebagai Paramater & fungsi anonymous

func CekDaftarBahan(BahanMixing string, BahanTopping string, func(JenisCoklat)string){
	if JenisCoklat=="Coklat cair"{
		fmt.Println("Bahan sudah siap")
	}else{
		fmt.Println("Bahan belum Siap")
	}
}
func JenisCoklat(string)

//Fungsi Return Value

func SiapkanBahan(BahanUtama string) string {
	if BahanUtama != "telur" {
		return "balik belanja lagi"
	} else {
		return "Pecahkan" + BahanUtama
	}
}

//Fungsi Multiple Return Value

func PersiapanMixing() (string, string, string,) {
	return "Campur semua bahan", "Mixer selama 15 menit", "masukan coklat cair"
	//step1,step2,step3:=persiapanmixing
}
//struct

type DataCostumer struct{
Nama string
JumlahPesanan int
}

//struct method
func ( Costumer DataCostumer) KartuUcapan() {
	fmt.Println("Terima kasih sudah membeli", DataCostumer.Nama)
}

//anonymous struct
DataPengiriman := struct {NomerHP int; Alamat string}{12345, "Bandung"}


// interface
type Kue interface {
	GetWarna()
}

type Rasa struct {
	Kopi string
	Buah string
}

func (m Rasa) GetWarna() {
	fmt.Println(m)
}
