# Setting awal

Bagian ini harus anda kerjakan sebelum mengirimkan kontribusi. Setting awal ini hanya dilakukan sekali saja.

* Install Git
* Buat account di Github
* Fork repo ini ke repo pada account anda, caranya: login kemudian `fork`.

* Pada komputer lokal anda, kerjakan langkah berikut:

```
$ git clone https://github.com/galihepj/blog-go.git

```

Pada kondisi saat ini, di komputer lokal anda sudah terdapat repo `blog-go` yang berada pada direktori dengan nama yang sama. Untuk keperluan berkontribusi, ada 2 nama repo yang harus anda setting:
  1. origin => menunjuk ke repo milik anda di github, hasil dari fork.
  2. upstream => menunjuk ke repo milik upstream (repo asli) di account galihepj
Repo `origin` sudah dituliskan konfigurasinya pada saat anda melakukan proses clone dari repo anda. Konfigurasi repo upstream harus dibuat.


Tambahkan remote upstream:

```
$ git remote add upstream https://github.com/galihepj/blog-go.git
```


* Selesai setting awal.


# Sync Repo Sebelum Pull Request (PR)

Bagian ini dikerjakan sebelum mengirimkan PR. Tujuannya untuk memastikan bahwa yang kita buat belum dibuat oleh orang lain.

```
$ git fetch upstream
$ git checkout master
$ git merge upstream/master
$ git status
$ git push origin master

$
```

Jika belum ada, maka anda bisa mulai menyiapkan dan mengirimkan perubahan.

# Membuat Perubahan dan Mengirim PR

## Membuat Perubahan

Untuk setiap perubahan, buat dalam branch, jangan buat dalam master karena akan mencampur adukkan dan membuat susah dilacak.

```
$ git checkout -b berkontribusi
$ git branch
* berkontribusi
  master
$
```

Setelah itu editlah seperlunya (bisa meliputi hanya satu file atau lebih banyak lagi), setelah selesai, kirim perubahan tersebut ke repo milik anda (bukan upstream tetapi origin).

Selanjutnya proses pull request dilakukan melalui web di github dengan memilih new pull request.
