package main

const utang = 20000

func cekUtang(jumlahUang int) (pesan string) {
	if jumlahUang == utang {
		pesan = "Uang pas utang lunas"
	} else if jumlahUang > utang {
		pesan = "Uang lebih ini kembaliannya"
	} else {
		pesan = "Uang kurang"
	}
	return
}

func cekWarna(warna string) (pesan string) {
	switch warna {
	case "hijau":
		pesan = "anda penyuka keindahan"
	case "merah":
		pesan = "anda penyuka pemberani"
	case "putih":
		pesan = "Putih itu suci"
	default:
		pesan = "warna anda belum tau artinya"
	}
	return
}
func main() {
	println(cekUtang(21000))
	println(cekWarna("putih"))
}
