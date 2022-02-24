## (4) Basic Programming

1. time complexity adalah bagaimana kita mengitung suatu algoritma atau suatu sistem yang berjalan dalam hal waktu. caranya yaitu pertama kita harus menentukan seberapa banyak operasi dominan yang dilakukan. operasi dominan merupakan operasi yang primitif seperti penambahan, perkalian, dsb. terdapat beberapa jenis time complexity. yang pertama adalah constant time - O(1) dimana operasi dominan hanya akan di eksekusi sebanyak 1 kali, yang kedua adalah linear time - O(n) dimana ketika melakukan input n maka operasi dominan akan di eksekusi sebanyak n kali, selanjutnya adalah logarithmic time - O(log n) dimana n akan dieksekusi sebanyak 2^x = n kali, kemudian ada quadratic time - O(n^2) yaitu ketika kita memiliki input n maka akan di eksekusi sebanyak n^2 kali. Selain itu ada juga exponential dan factorial time
2. sebuah komputer dapat melakukan pemrosesan data sebanyak 10^8 atau 100.000.000 operasi/1 detik. maka ketika kita sudah menentukan misalnya ternyata time complexity kita adalah O(n) lalu menginputkan 1 milyar data maka komputer akan memproses algoritma ini sebanyak 10 detik. namun berdasarkan beberapa penelitian usahakan kita tidak boleh memproses algoritma selama lebih dari 2 detik.
3. sebuah proses sebelum masuk kedalam cpu maka nilai yang dioperasikan akan disimpan didalam sebuah memory terlebih dahulu, memory ini adalah penyimpanan sementara. Space complexity berarti kita harus bisa menghitung berapa alokasi memori yang dibutuhkan dalam algoritma tersebut. cara menghitung space complexity adalah dengan menghitung seberapa banyak variable atau ruang yang dialokasikan ke memori. 

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