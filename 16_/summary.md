## (16) MongoDB : Advanced Query - Array - Aggregation

1. dalam penggunaan advance query, kita bisa menggunakan comparison dimana __$ne = not equal = "!="__ kemudian ada __$gt = greater than = ">"__ kemudian __$lt = lower than = "<"__ kemudian __$gte = greater than equal = ">="__ kemudian __$lte = lower than equal = "<="__. adapun logical expression seperti $or-$and, dan $in-$nin bisa juga $not. bisa juga menggunakan evaluator dengan $regex.
2. advance query juga dapat mencari dengan $all yaitu mengharuskan data memiliki semua kriteria yang dicari, kemudian $size yang mengharuskan jumlah array memiliki kriteria yang dicari, kita juga bisa melakukan update dengan cara $push(untuk append data) dan $pop(untuk delete data).
3. No Sql bisa melakukan join, contohnya $lookup untuk melakukan left join, $match untuk mem-filter dokument, $project untuk merubah susunan dokumen yang akan ditampilkan, $cond merupakan filter dari dokumen tsb, $group melakukan summarize dari dokument. Kemudian ada $project yaitu untuk melakukan reshape apa saja yang akan ditampilkan. Ada juga $unwind untuk melakukan expand data, $sort untuk melakukan sorting data, $limit untuk membatasi data yang keluar. $skip untuk men skip data sesuai dengan masukan.

### Task
1. untuk mengerjakan soal saya menerapkan materi yang dipelaari pada video pembelajaran dan juga materi dari power point dan semuanya telah saya masukkan dalam pdf di praktikum. sehingga hasil akhirnya menjadid seperti berikut
<img src="screenshots/hasil akhir database setelah praktikum section 16.PNG">