## (4) Basic Programming

1. Variable digunakan untuk mennyimpam informasi di dalam program komputer, variable menyediakan cara untuk melakukan pelabelan data dengan nama yang deskriptif, kemudian variable juga memiliki tipe data diantaranya boolean, numeric, dan string. Cara mendeklarasikan variable dengan tipe data numeric contohnya: `var age int`. Constant juga berfungsi untuk menyimpan informasi namun nilai constant tidak bisa diubah seperti variable.
2. salah satu operators pada golang adalah arithmatic, jenis jenisnya yaitu addition, subtraction, division, multiplication, modulo, increment, dan decrement. Selain arithmatic, operator pada golang lainnya yaitu, comparison, logical, bitwise, assignment, miscellaneous.
3. terdapat beberapa struktur kontrol pada Go. untuk melakukan percabangan bisa menggunakan if dan switch. untuk melakukan looping atau perulangan bisa menggunakan for. loops adalah statement yang membuat code untuk dapat dieksekusi berkali-kali

### Task
1. untuk membuat program yang menghitunga luas permukaan tabung, pertama tema saya harus mengetahui terlebih dahulu rumus tabungnya. Kemudian saya meminta 2 inputan user yaitu tinggi tabung dan jari jari menggunakan scanf yang pada terminal input keduanya dipisahkan oleh satu kali spasi. kemudian inputan tersebut akan dioperasikan sesuai dengan rumus yang diberikan pada problem1, lalu print hasilnya
<img src="screenshots/SS problem 1.PNG">

2. untuk mendeskripsikan nilai siswa, saya meminta inputan nama dan nilai siswa terlebih dahulu, untuk kemudian dilakukan proses branching sesuai dengan kategori yang diberikan dari indeks nilai A - E, lalu outputkan hasilnya berupa nama dan nilai siswa
<img src="screenshots/SS problem 2.PNG">

3. untuk mengeluarkan faktor bilangan, saya membuatnya dengan cara melakukan perulangan terlebih dahulu. Kemudian didalam perulangan tersebut saya melakukan branching apabila inputan user habis dibagi dengan indeks dimana perulangan tersebut berjalan, maka indeks tersebut merupakan faktor dari inputan user tadi sehingga saya print indeksnya.
<img src="screenshots/SS problem 3.PNG">

4. untuk mengeluarkan bilangan prima, caranya hampir sama dengan algoritma pada problem 3 dimana saya membuat perulangan terlebih dahulu untuk didalamnya melakukan branching apabila inputan user habis dibagi dengan indeks dimana perulangan tersebut berjalan, maka inputan user tersebut bukanlah bilangan prima karena dapat dibagi selain dari 1 dan dirinya sendiri, sehingga akan me-return false, jika sampai perulangan terakhir kondisi tersebut tidak pernah true, maka akan me-return true.
<img src="screenshots/SS problem 4.PNG">

5. untuk mengetahui apakah kata tersebut merupakan palindrom atau bukan, saya melakukan perulangan dimana akan melakukan looping tiap char dari katanya. Didalam looping tersebut akan dilakukan branching apabila char pertama sama dengan char terakhir dalam kata tersebut maka variable sama akan bertambah 1, begitu seterusnya sampai char pertama berpindah satu demi satu kebelakang char terakhir berpindah satu demi satu kedepan. apabila variable sama bernilai sama dengan jumlah char yang di cek maka kata tersebut merupakan bilangan palindrom dan akan me-return true.
<img src="screenshots/SS problem 5.PNG">

6. untuk melakukan operasi eksponensial, saya membutuhkan variable tambahan yaitu satuan yang berfungsi untuk menjadi peran utama dalam function dimana didalam perulangan, variable satuan ini akan dijumlahkan dengan bilangan base sebanyak pangkat.
<img src="screenshots/SS problem 6.PNG">

7. untuk mengeluarkan output piramida, idenya adalah membuat spasi semakin berkurang dan * semakin bertambah, maka dari itu saya membutuhkan looping didalam looping untuk mencetak spasi dan bintang.
<img src="screenshots/SS problem 7.PNG">

8. untuk mengeluarkan operasi pada problem 8, idenya adalah saya harus menjumlahkan baris 1 dengan kolom 1 secara bergantian. Maka dari itu saya membutuhkan looping didalam looping untuk merealisasikan ide tersebut. Selain itu saya juga menambahkan branching untuk menyesuaikan persyaratan dan mempercantik output, sehingga apabila inputan user 30 masih bisa diterima dengan baik
<img src="screenshots/SS problem 8.PNG">