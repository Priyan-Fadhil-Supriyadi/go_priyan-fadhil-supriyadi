## (21) ORM and Code Structure (MVC)

1. Object Relational Mapping (ORM) akan membantu kita ketika mengonversi data antara tipe yang tidak kompatibel, sistemnya menggunakan objek oriented programming. Kelebihan ORM diantaranya adalah less repetitive query artinya tidak banyak menggunakan query, lalu ketika kita mendapatkan data maka akan otomatis tersimpan kedalam variable struct kita, kemudian validasi data lebih mudah, lalu ORM mempunyai feature cache query. Sedangkan kelemahan ORM antara lain adalah code nya lebih banyak juga proses nya memberatkan di bagian operasi kita, lalu ORM dapat mengeksekusi query yang tidak diinginkan sehingga menambah cost ketika operasi, kemudian kompleksiti query bisa terlalu panjang jika ditulis dengan ORM, lalu ketergantungan terhadap library.
2. GORM merupakan fungsi ORM yang merupakan RDBMS. GORM dapat dipasangkan dengan framework yang sudah dibuat sebelumnya. data migration penting dalam industri karena ketika melakukan update pada kolom database maka migrate akan melakukan tracking siapa yang menambahkan atau apa yang ditambahkan di versi yang baru. fungsi GORM akan memungkinkan kita melakukan auto migration. dan alasan mengapa kita melakukan db migration adalah ketika kita mengupdate database lebih simple dan ketika melakukan rollback lebih simple lagi, yang paling penting adalah kita bisa melakukan tracking perubahan database kita, terakhir kita akan selalu compatible jika menggunakan migration dari aplikasi terbaru dengan yang ada di database.
3. MVC adalah sebuah struktur yang terdiri dari model, view, controllers. Artinya nanti projek kita dipecah menjadi bagian bagian tersebut sehingga lebih mudah dibaca. Model itu seperti struct dimana model dari struktur data yang akan dimasukkan ke database. controller adalah isi dari logic bisnis nya kita, view adalah bagian awal seperti main.

### Task 1
Pertama-tama saya membuat 3 user untuk diuji coba apakah code yang saya buat telah berjalan dengan baik, 3 user tersebut sebagai berikut :
<img src="screenshots/part1create1.PNG">
<img src="screenshots/part1create2.PNG">
<img src="screenshots/part1create3.PNG">

kemudian saya mengubah isi data dari user dengan id 2 sehingga :
<img src="screenshots/part1update.PNG">

mengeluarkan semua user yang pernah dibuat :
<img src="screenshots/part1getall.PNG">

lalu menghapus user dengan id 11 sehingga :
<img src="screenshots/part1delete.PNG">

maka user dengan id 11 telah terhapus :
<img src="screenshots/part1buktiberhasildelete.PNG">

dan menampilkan user dengan id 9 :
<img src="screenshots/part1getuser.PNG">

### Task 2
karena isinya sama dengan database sebelumnya sehingga saya langsung mengeluarkan semua data user :
<img src="screenshots/part2getall.PNG">

dan menampilkan user dengan id 15 :
<img src="screenshots/part2getuser.PNG">

kemudian saya membuat user baru :
<img src="screenshots/part2create.PNG">

lalu menghapus user dengan id 15 sehingga :
<img src="screenshots/part2delete.PNG">

maka user dengan id 15 telah terhapus :
<img src="screenshots/part2buktidelete.PNG">

kemudian saya mengubah isi data dari user dengan id 17 sehingga :
<img src="screenshots/part2update.PNG">

karena logic dari user sama dengan logic dari books maka saya rasa saya cukup menampikan bahwa layered yang dibuat telah berhasil :
<img src="screenshots/part2books.PNG">