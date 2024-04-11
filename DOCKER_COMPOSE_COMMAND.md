# Docker Compose Command 
Docker compose merupakan sebuah alat yang dibuat di atas Docker yang digunakan untuk mendefinisikan maupun mengonfigurasi serta menjalankan aplikasi yang tersusun atas berbagai kontainer dengan cara mudah.
Jadi Inti nya Docker compose adalah tools atau alat untuk mengonfigurasi serta menjalankan aplikasi yang tersusun berbagai container.

## Cara Membuat Docker Compose 
- Untuk Melakukan Perintah Membuat Docker Compose Anda Berada di Dalam Folder yang ada file docker compose nya
- Nama Dari Docker Compose sesuai dengan folder nya 
```bash
docker compose create
```
## Cara Menjalankan Container yang Ada di Docker Compose
```bash
docker compose start
```
- Nah Pertanyaan nya kalau kita sebelum nya memiliki project yang banya dan sama sama pakai docker compose jadi akan semua akan di jalan kan ?
- Jawab Nya Tidak , karena setiap perintah docker compose hanya relative dengan satu file docker compose yang berada di workdir kita.
  
## Cara Melihat Container yang ada di Docker Compose
- Cara Pertama 
```bash
docker compose ls -a
```