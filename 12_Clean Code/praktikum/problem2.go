package main

type Kendaraan struct {
	kecepatan int
}

// menginisiasi kecepatan dan menambahkannya di functuon akselerasi
func (k *Kendaraan) berjalan() {
	k.akselerasi(10)
}

// menambahkan kecepatan dengan pertambahan sebanyak function dipanggil
func (k *Kendaraan) akselerasi(pertambahan int) {
	k.kecepatan = k.kecepatan + pertambahan
}

func main() {
	mobilCepat := Kendaraan{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Kendaraan{}
	mobilLamban.berjalan()
}
