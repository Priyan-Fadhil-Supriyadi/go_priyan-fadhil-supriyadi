## Data Structure

1. library reflect berfungsi untuk mengecek value dari variable yang ingin dilihat. nilai default dari int itu adalah 0, array memiliki index dan value dimana index adalah tempat menyimpan nilai pada array tersebut. Slice sama seperti array yaitu merupakan sebuah struktur data juga pada golang yang berisi group of elemenet, bedanya jika array alokasinya itu sifatnya tidak bisa berubah, kalo slice alokasi nya bersifat dinamik atau bisa berubah. Slice dapat diisi menggunakan array yang telah terbentuk yang disebut dengan refference type. Slice memiliki pointer(mereferensikan kondisi array), len(panjang dari segmen slice), dan capacity(jumlah maksimum).
2. map merupakan struktur data pada bahasa pemrograman golang, bedanya adalah map memiliki key dan value, dimana key-nya bersifat unik. Jadi kesimpulan perbedaan antara ketiganya adalah array itu sifatnya tetap atau fix jadi tidak dapat diubah lagi, slice lebih dinamis dan bisa berubah, yang terakhir adalah map dimana map mempunyai key dan value.
3. function adalah sekumpulan kode yang dikelompokkan atau dipanggil dengan nama tertentu. Function dibutuhkan ketika kita membuat sistem yang cukup besar dan dapat mempermudah unit testing. function memiliki bermacam macam jenis, ada function yang tidak melakukan return sama sekali, function yang me return nilai tunggal, dan function yang me return banyak nilai.

### Task

1. untuk membuat program yang menggabungkan 2 array tanpa data yang sama saya menggunakan struktur data map. dimana saya membuat map untuk melakukan pengecekan apakah nilai array yang akan dimasukkan kedalam array baru sebelumnya pernah dimasukkan atau belum sama sekali. sehingga program akan menjadi seperti berikut
<img src="screenshots/SS problem 1.PNG">

2. untuk membuat program yang meng-output-kan inputan user yang hanya muncul satu kali. saya melakukan cek dengan melakukan nested looping dan mengecek apabila terdapat bilangan yang sama dari kiri ke kanan maka bilangan tersebut tidak akan dimasukkan kedalam array baru. kemudian array baru tersebut akan dikembalikan. sebelum melakukan proses tersebut, function yang saya buat akan melakukan pengubahan tipe data dari byte ke string, kemudian diproses, setelah selesai karena function me return int, maka tipe data string diubah ke int. Sehingga program akan menjadi seperti berikut
<img src="screenshots/SS problem 2.PNG">

3. untuk membuat program mencari target sum dari nomor yang telah dilakukan sort, maka saya melakukan cara yang hampir sama dengan algo problem 2 yaitu melakukan nested looping tetapi disini saya menjumlahkan tiap indeks satu persatu, apabila hasilnya sesuai dengan targer, maka indeks tersebut akan disimpan kedalam array baru. Sehingga program yang menjadi seperti berikut
<img src="screenshots/SS problem 3.PNG">
