## (10) Dynamic Programming

1. Dynamic programming merupakan algoritma teknik untuk menyelesaikan problem secara optimal. dengan memecahkan problem tersebut kedalam problem problem yang lebih mudah atau gampang. 
2. karakter dari dynamic programming diantaranya adalah overlapping subproblem, yaitu ketika kita mencari solusi maka subproblem akan dipanggil beberapa kali maka kita mengetahui bahwa problem tersebut akan dilakukan berkali kali yang akan digambarkan dengan tree. yang kedua adalah optimal substructure property, jadi setiap problem harus menggunakan problem itu sendiri atau harus optimal.
3. metode pada dynamic programming adalah, top down with memoization, yang merupakan pemecahan masalah menggunakan recursive jadi ketika telah memecahkan subproblem tertentu dia akan menyimpan hasilnya ke subdata yang temporary, dengan begitu tidak perlu menyelesaikan masalah tersebut berulang-ulang. method berikutnya adalah botton-up tabulation, yang merupakan opposite dari top-down yang menghindari rekursif jadi dia akan menyelesaikan masalah dari yang terkecil.

### Task
1. untuk membuat program Fibonacci Number Top-down, saya memerlukan memory dikarenakan metode tersebut merupakan pemecahan masalah menggunakan recursive jadi ketika telah memecahkan subproblem tertentu dia akan menyimpan hasilnya ke subdata yang temporary. sehingga program akan menjadi seperti berikut
<img src="screenshots/SS problem 1.PNG">

2. untuk membuat program Fibonacci Number Bottom-up, saya tidak menggunakan rekursif yang dilakukan sebelumnya dengan metode Number Top-down tetapi saya membuat array yang akan menyimpan nilai fibonacci yang terus dilakukan looping sebanyak n, sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 2.PNG">

3. untuk dapat mencari nilai minimum yang harus dipilih oleh katak untuk sampai ke N maka saya membutuhkan library math dan array untuk menyimpan hasil paling minimal yang dicari dari perhitungannya, sehingga program akan menjadi seperti berikut.
<img src="screenshots/SS problem 3.PNG">