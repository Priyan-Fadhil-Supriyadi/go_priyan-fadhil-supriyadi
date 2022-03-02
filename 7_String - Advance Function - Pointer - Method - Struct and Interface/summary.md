## (7) String - Advance Function - Pointer - Method - Struct and Interface

1. string merupakan salah satu variable penting apabila kita ingin membuat program. pada golang terdapat beberapa kegunaan library strings, diantaranya __len__ (untuk menghitung panjang string), __compare__ (di gambarkan dengan notasi "==" untuk membandingkan antar string), __contain__ (untuk membandingkan apakah string kedua adalah bagian dari string pertama atau tidak), __substing__ (cara untuk mengambil string tertentu), __replace__ (untuk mengganti string tertentu), __insert__ (untuk menambahkan kata pada variable string)
2. variadic function adalah ketika kita mau memanggil sebuah function tetapi parameter yang ingin di inputkan itu jumlahnya berbeda-beda, caranya dengan mengganti parameter dengan temporary slice. Selanjutnya ada anonymous function adalah sebuah function yang tidak mempunyai nama, ini untuk mengelompokkan sebuah kumpulan code didalam sebuah function. Selanjutnya closure, closure adalah sebuah anonymous function dimana ketika memakai anon function variable diluar dia akan melakukkan reference. Kemudian Defer function, yaitu dia akan melakukan running function tersebut diakhir program, dimana defer ini merujuk pada sistem LIFO (last in - first out).
3. pointer adalah sebuah variable yang dapat menyimpan memori address dari variable yang lain. Struct disebut juga call object dalam golang, struct merepresentasikan sifat OOP di golang, namun golang tidak mempunyai OOP. Method adalah function yang di tempel pada sebuah type, method membantu kita menulis code dengan gaya OOP pada G0, method menghindari naming conflict, dan akan memudahkan pembacaan code.

### Task
1. untuk membuat program compare string yang mengoutputkan string yang muncul di kedua teks yang diinputkan. Maka saya melakukan perulangan sebanyak kata dengan jumlaf huruf terbanyak disetiap perulangan saya melakukan pengecekan apakah kata dari panjang indeks 0-i merupakan bagian dari kata kedua yang lebih pendek dari kata yang lain, apabila kondisi tersebut true atau benar maka kata indeks 0-i akan masuk kedalam variable baru. Dimana variable tersebut akan direturn ke function main. Sehingga program akan menjadi seperti berikut.
<img src="screenshots/SS problem 1.PNG">

2. untuk membuat program yang dapat melakukan caesar cipher. idenya adalah dengan mengulang melakukan perulangan dari range ascii a-z yaitu 97 - 122. sehingga diperlukan bantuan variable untuk membuat kondisi tersebut terjadi. Sehingga run program akan menjadi seperti berikut
<img src="screenshots/SS problem 2.PNG">

3. program swap bertujuan untuk menukarkan nilai yang ada pada variable a dan b. Karena yang dilemppar pada function swap adalah nilai alamatnya maka idenya disini adalah dengan menukarkan nilainya langsung menggunakan bantuan variable. sehingga menjadi seperti berikut
<img src="screenshots/SS problem 3.PNG">

4. pada function getMinMax saya melakukan perulangan sebanyak numbers, kemudian membuat 2 branching untuk nilai min dan max. sehingga memiliki output seperti berikut.
<img src="screenshots/SS problem 4.PNG">

5. pada function avarage saya melakukan perulangan untuk menjumlahkan setiap score, setelah selesai saya membagi total tersebut dengan jumlah score yang ada dan mengembalikan nilai tersebut. pada function Min saya melakukan perulangan yang didalamnya terdapat branching untuk menentukan bilangan terkecil. pada function Max saya melakukan perulangan yang didalamnya terdapat branching untuk menentukan bilangan terbesar. sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 5.PNG">

6. pada function encode dan decode di program ini, idenya adalah apabila menyesuaikan dengan test case yang ada pada problem 6, z akan diubah dengan a, y akan diubah dengan y sehingga ini terlihat polanya terbalik, seperti a - z menjadi z - a. sehingga saya memasukkan setiap huruf yang diterima untuk kemudian dilakukan operasi mod dengan rune dari 'a' hasilnya akan mengurangi rune dari 'z' karena dibalik. sehingga akan menjadi program seperti berikut ini.
<img src="screenshots/SS problem 6.PNG">