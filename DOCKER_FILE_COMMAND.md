# Dockerfile

## Docker Build
- untuk membuat image dari docker file
```bash
docker build -t name_image folder_in_dockerfile_exist
```
- Untuk Menampikan Output dari sebuah Dockerfile ketika jadi image 
```bash
docker build -t name_image folder_in_dockerfile_exist --progress=plain
```
- Saat Proses build docker secara default menyimpan perubahan Dockerfile di cache jadi ketika anda mencoba build ulang tanpa ada perubahan docker tidak akan build ulang secara manual 
- Akan tetapi tetapi ketika ingin build ulang secara manual pakai command 
```bash
docker build -t name_image folder_in_dockerfile_exist --no-cache
```

## From Instructions
- Digunakan untuk statment awal ketika ingin menggunakan image docker dari docker registry (docker hub)
- Sebenarnya bisa tanpa 'FROM' Instruksi tapi jarang sekali orang yang memulai benar benar dari nol tanpa image dari docker hub
- Contoh Penggunaan
```Dockerfile
FROM image:tag
```
- Ketika Penulisan tanpa default nya latest

## Run Insturctions
- Digunakan ketika anda ingin menjalankan perintah saat proses build Dockerfile jadi image 
- Contoh Penggunaan yang Pertama
```Dockerfile
FROM alpine:latest
# RUN command argument
RUN mkdir app #untuk membuat directory
RUN cd app/ #untuk masuk directory app
RUN touch .env
```
- Contoh Penggunaan yang Kedua
```Dockerfile
FROM alpine:latest
# RUN ["command","argument","argument"]

```
