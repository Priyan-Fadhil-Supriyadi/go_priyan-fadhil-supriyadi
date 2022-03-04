## (8) Recursive - Number Theory - Sorting - Searching

1. rekursif merupakan function yang dapat memanggil dirinnya sendiri. rekursif dapat mempermudah menyelesaikan problem dan kode yang dihasilkan juga lebih sederhana. rekursif sama dengan iterasi karena cara kerjanya sama. terdapat 2 strategi pada rekursif, pertama base case yaitu masalah yang dapat diselesaikan dalam sekali perulangan, jika tidak kita dapat menggunakan strategi kedua yaitu recurrence relations. number theory merupakan cabang dari ilmu matematika yang mempelajari studi integer.
2. searching bertujuan untuk memproses atau menemukan posisi dari value yang diberikan. terdapat beberapa cara, yang pertama ada linear search, dan ada builtins search juga pada go. selanjutnya Sorting, sorting akan mengurutkan angka yang kita punya, baik secara ascending(kecil ke besar) ataupun descending (besar ke kecil). terdapat beberapa teknik sorting, yang pertama ada selection sort - O(n^2), kemudian counting sort - O(n+k). builtins sort in golang metupakan package yang ada dalam golang untuk melakukan sorting.
3. stack memiliki ciri khas bahwa proses pada stack adalah LIFO (Last-In-First-Out), maksudnya data yang terakhir masuk adalah data yang pertama keluar. terdapat beberapa properti seperti push(element) yaitu untuk memasukkan element, pop() untuk mengeluarkan elemen pertama, dan top() untuk melihat nilai pada element pertama. selanjutnya Queue yang mengimplementasikan FIFO (First-In-First-Out). terdapat 3 pproperti dalam queue yaitu enqueue(element) untuk memasukkan element, dequeue() untuk menghapus element pertama, dan front() mengecek nilai pada first element. Terakhir ada deque, deque mempunyai 2 tempat untuk bisa menghilangkan element (depan dan belakang), terdapat 4 properti dalam deque yaitu push_front(element) untuk memasukkan nilai pada element terdepan, push_back(element) yaitu memasukkan element paling belakang, dequeue() yaitu menghilangkan element pertama dalam queue, terakhir front() untul mengecek element pertama. Set adalah jenis array yang tidak membolehkan adanya value duplicate sehingga akan menghilangkan value yang sama. Map merupakan element yang dapat berisi lebih dari 1 komponen.

### Task
1. saya mengerjakan problem 1 dengan menggunakan algo sebelumnya dan melakukan perulangan dengan kondisi tertentu sebanyak inputan user. sehingga problem menjadi seperti berikut.
<img src="screenshots/SS problem 1.PNG">

2.  saya melakukan teknik rekursif untuk mengerjakan algo fibonacci. sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 2.PNG">

3. saya menggunakan algo sebelumnya untuk mencari bilangan prima dan melooping sebanyak wide untuk membentuk bangun datar nya. sehingga program menjadi seperti berikut.
<img src="screenshots/SS problem 3.PNG">

4. saya melakukan looping sebanyak inputan array kemudian didalamnya menggunakan beberapa branching sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 4.PNG">

5. saya melakukan looping yang didalamnya terdapat iterasi sehingga perogram menjadi seperti berikut
<img src="screenshots/SS problem 5.PNG">

6. saya melakukan looping dan melakukan operasi pengurangan sehingga menjadi seperti berikut
<img src="screenshots/SS problem 6.PNG">

7. saya melakukan nested looping sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 7.PNG">

8. program belum selesai