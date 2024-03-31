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

## Label Insturctions
- Label insturksi ini berfungsi menambahkan detail informasi tentang image nya seperti author,company,email,dll.
- Label instruksi tidak terlalu berpengaruh di docker image nya. Dan Jika ingin melihat label perlu command 
``` docker image inspect name_image:tag```
- Contoh Penggunaan
```Dockerfile
FROM alpine:latest
# LABEL key="value"
LABEL author="example"
LABEL github="https://github.com/example" email="example@gmail.com"

CMD echo "Label Instruksi"
```


## Add Instructions
- Add Instruksi digunakan untuk menambahkan file dari source ke destination 
- source nya bisa berupa file dari host , url. Dan jika file nya berupa archive akan di ekstrak secara otomatis di destionation nya. Dan jika berupa url akan download secara otomatis juga
```Dockerfile
FROM alpine:latest
# ADD source destination
RUN mkdir document
ADD folder_host/file.txt document
ADD folder_host/*.txt document
CMD echo file.txt
```

## Copy Instructions
- Fungsi dari Copy instruksi persis dengan add tapi hanya bisa benar-benar copy file source ke destination dan tidak memiliki fitur lain nya
- Best Practice Programmer Docker biasa nya pakai COPY. Sebisa mungkin Pakai COPY
```Dockerfile
FROM alpine:latest
# COPY source destination
RUN mkdir document
COPY folder_host/file.txt document
COPY folder_host/*.txt document
CMD echo file.txt
```