## (9) Brute Force - Greedy and Divide and Conquer

1. problem solving paradigma adalah bagaimana kita menyelesaikan sebuah masalah. macam macam pendekatan untuk problem solving paradigms adalah Brute force(complete search), divide and conquer, Greedy, dan Dynamic programming. complete search adalah metode yang memproses data dari awal sampai akhir, cost ini besar dan kurang efektif namun mudah dalam pembuatannya, complete search digunakan ketika tidak ada time complexity yang lebih cepat untuk ini.
2. Greedy merupakan algoritma yang mana ketika menyelesaikan masalah kita mencari nilai local optimal. hasil dari algoritma greedy bisa jadi optimal atau tidak, namun cepat.
3. secara prinsip apabila kita menggunakan algoritma divide and conquer akan memecah kelompok data menjadi kecil kecil biasanya dibagi 2. Jika dibandungkan antara sequential dan binary maka jelas lebih efektif apabila kita menggunakan binary search.

### Task
1. untuk mencari 3 bilangan yang membentuk suatu relasi dengan kondisi tertentu, saya memerlukan nested looping sebanyak 3 kali yaitu untuk mencari ketiga bilangan x,y,z yang tepat untuk membentuk suatu relasi dengan kondisi tertentu sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 1.PNG">

2. untuk dapat membagi input bilangan menjadi beberapa bagian berdasarkan pecahan tertentu, saya membutuhkan nested looping untuk melakukan perulangan sebanyak pecahan yang ada untuk kemudian di looping kedua, inputan bilangan yang akan dipecah akan dikurangi dari pecahan yang terbesar samapi terkecil dengan sarat tidak min. sehingga program akan menjadi seperti berikut
<img src="screenshots/SS problem 2.PNG">

3. untuk dapat mencari tinggi minimum knight yang ada untuk dapat melawan dragon dengan diameter leher tertentu maka saya branching untuk menentukan apakah D>H maka knight fall atau knight akan tumbang, apabila tidak maka saya melakukan sorting terhadap D dab H terlebih dahulu kemudian saya melakukan looping sebanyak jumlah H yang didalamnya dilakukan branching untuk mencari jumlah H minimum yang diperlukan untuk mengalahkan dragon. sehingga program akan menjadi seperti berikut.
<img src="screenshots/SS problem 3.PNG">

4. untuk melakukan binary search saya memerlukan 3 pointer yang di representasikan dengan 3 buah variable yaitu min, max, dan mid yang mempunyai peran seperti namanya. kemudian untuk merealisasikan metode binary saya menggunakan looping dengan kompleksitas sebesar O(log n) karena selalu dibagi menjadi kelompok kelompok kecil dan didalamnya akan dilakukan branching untuk menentukan kelompok mana yang akan dipilih dan dicari berdasarkan variable mid, sampai mid sama dengan yang dicari. sehingga program akan menjadi seperti berikut.
<img src="screenshots/SS problem 4.PNG">