## (12) Clean Code

1. Clean code merupakan istilah untuk kode yang mudah dibaca, dipahami, dan diubah oleh programmer. kenapa clean code? karena untuk beekolaborasi dalam bekerja, untuk pengembangan fitur, dan agar  pengembangan lebih cepat. Adapun karakter-karakter clean code ini antara lain, pertama penamaan mudah dipahami, kedua adalah mudah dieja dan dicari, ketiga adalah singkat namun mendeskripsikan konteks, keempat adalah konsisten, kelima adalah hindari penambahan konsteks yang tidak perlu, yang keenam adalah komentar, kemudian ketujuh adalah function. kedelapan adalah mengikuti konvensi, kesembilan adalah formating.
2. clean code principle yang pertama adalah KISS (Keep It So Simple) yaitu menghindari membuat fungsi yang dibuta untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dan seterusnya. Fungsi atau class harus kecil, jangan gunakan terlalu banyak argumen, seimbang dan minimal. kemudian DRY (Dont Repeat Yourself) yaitu duplikasi code terjadi karena sering copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan berulang ulang. 
3. kemudian tujuan clean code adalah Refactoring yang merupakan proses refakturisasi kode yang dibuat, dengan cara mengubah struk tur internal tanpa mengubah perilaku eksternal prinsip KISS dan DRY bisa dicapai dengan cara refactor. Teknik dari refactoring adalah membuat sebuah abstraksi, memecah kode dengan fungsi/class, memperbaiki penamaan dan lokasi kode, dan deteksi kode yang memiliki duplikasi

### Task
1. untuk menganalisis kode pada problem 1 saya mencari karakter karakter dalam clean code yang tidak ada pada problem 1 tersebut sehingga saya menemukan beberapa diantaranya tidak diterapkan. Maka hasil analysis saya terhadap kode menjadi seperti berikut
<img src="screenshots/SS problem 1.PNG">

2. pada rewrite code saya menerapkan karakter karakter dalam clean code sehinggaa saya menghapus beberapa variable dan struct yang tidak diperlukan, merubah penamaan variable, merubah nama function, serta menambahkan komentar sehingga program menjadi seperti berikut
<img src="screenshots/SS problem 2.PNG">