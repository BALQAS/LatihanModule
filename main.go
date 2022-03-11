package latihanmodule

import "fmt"

//Fungsi Parameter

func Judul(jenisKue string, rasaKue string) {
	fmt.Println("Cara Membuat", jenisKue, rasaKue)

}

//Fungsi Sebagai Paramater
// func DaftarBelanja(bahanMixing string, bahanTopping func(jenisCoklat)string)){
// 	if jenisCoklat=="coklat cair"{
// 		fmt.Println(bahanMixing, jenisCoklat, "belanja beres")
// 	}else{
// 		fmt.Println("segera belanja!")
// 	}
// }

//Fungsi Return Value
func CekBahan(BahanUtama string) string {
	if BahanUtama != "telur" {
		return "segera beli dahulu"
	} else {
		return "Pecahkan" + BahanUtama
	}
}

//Fungsi Multiple Return Value

func PersiapanMixing() (string, string, string) {
	return "Campur semua bahan", "Mixer selama 15 menit", "masukan coklat cair"
	//step1,step2,step3:=persiapanmixing
}

//Function Anonymous
