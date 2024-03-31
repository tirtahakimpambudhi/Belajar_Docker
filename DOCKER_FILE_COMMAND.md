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
RUN echo APP_ENV="deveploment" > .env
```
- Contoh Penggunaan yang Kedua
```Dockerfile
FROM alpine:latest
# RUN ["command","argument","argument"]
RUN ["mkdir","app"]
RUN ["cd","app"]
RUN ["echo","APP_ENV='development'",">",".env"]
```
## CMD Instructions 
- Fungsi nya sama dengan 'RUN' Instruksi akan tetapi 'CMD' dijalankan ketika container sedang berjalan atau setiap perintah jalan nya container
- Akan tetapi 'CMD' Instruksi hanya boleh satu jadi ketika ditulis lebih dari satu maka yang dijalankan yang paling akhit
- Contoh Penggunaan Pertama
```Dockerfile
FROM alpine:latest
# RUN command argument
RUN mkdir app #untuk membuat directory
RUN cd app/ #untuk masuk directory app
RUN echo APP_ENV="deveploment" > .env
# CMD command argument
CMD cat .env
```
- Contoh Penggunaan Kedua
```Dockerfile
FROM alpine:latest
# RUN ["command","argument","argument"]
RUN ["mkdir","app"]
RUN ["cd","app"]
RUN ["echo","APP_ENV='development'",">",".env"]
# CMD ["command","argument","argument"]
CMD ["cat",".env"]
```