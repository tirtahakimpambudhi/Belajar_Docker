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
- Akan tetapi 'CMD' Instruksi hanya boleh satu jadi ketika ditulis lebih dari satu maka yang dijalankan yang paling akhir
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

## .dockerignore

- Fungsi dari .dockerignore persis sama dengan .gitignore yaitu mengignore file-file yang ditulis . Jadi saat build Dockerfile ke image sebenarnya yang di load pertama kali .dockerignore agar dilihat file file apa saja yang di ignore.
- Alasan ada nya .dockerignore untuk mengignore file tertentu saat melakukan ADD atau COPY instruksi agar file - file tertentu tidak ikut ke build jadi image

## Expose Instructions

- Berfungsi untuk menambah detail informasi atau dokumentasi bahwa image tersebut berjalan di port sekian
- Sebernarnya Expose instruksi tidak berpengaruh image tersebut harus berjalan di port itu hanya bersifat dokumentasi saja.
- Contoh Penggunaan

```Dockerfile
FROM golang:1.18-alpine

RUN mkdir app/
COPY main.go app/

EXPOSE 8081 tcp
CMD go run app/main.go
# EXPOSE port (tcp,udp)
```

## Env Instructions

- Berfungsi menambahkan enviroment variabel di image tersebut. Dan enviroment variabel tersebut bisa dirubah ketika saat pembuatan container memakai flag ```-e atau --env```
- Berfungsi ketika container sedang berjalan atau running.

```Dockerfile
FROM golang:1.18-alpine
# Single env
ENV APP_ENV="development"
# Multiple env
ENV APP_PORT="8081" APP_NAME="web_server"
RUN mkdir app/
COPY . app/

EXPOSE ${APP_PORT}
CMD go run app/main.go
```

## Volume Instructions

- Berfungsi untuk membuat volume secara otomatis saat membuat container
- Semua isi file yang di deklarasikan di volume Dockerfile akan di copy di docker volume yang dibuat secara otomatis tadi
- Contoh Pertama

```Dockerfile
FROM golang:1.18-alpine
# Single env
ENV APP_ENV="development"
# Multiple env
ENV APP_PORT="8081" APP_NAME="web_server" APP_LOG="app/log"

RUN mkdir app/
RUN mkdir -p ${APP_LOG}
COPY . app/
#Single file
VOLUME ${APP_LOG}
#Multiple
VOLUME ${APP_LOG} app/temp/log
EXPOSE ${APP_PORT}
CMD go run app/main.go
```

- ```VOLUME /path_container #untuk mounting```
- Contoh Kedua

```Dockerfile
FROM golang:1.18-alpine
# Single env
ENV APP_ENV="development"
# Multiple env
ENV APP_PORT="8081" APP_NAME="web_server" APP_LOG="app/log"

RUN mkdir app/
RUN mkdir -p ${APP_LOG}
COPY . app/
#Single file
VOLUME ["${APP_LOG}"]
#Multiple
VOLUME ["${APP_LOG}","app/temp/log"]
EXPOSE ${APP_PORT}
CMD go run app/main.go
```

## Workdir Instructions

- Seperti namanya 'WORKDIR' instruksi digunakan untuk mengubah direktori dimana program akan dijalankan. 'WORKDIR' instruksi juga berfungsi sebagai default direktori ketika kita masuk container pertama kali.
- Jika direktory tersebut tidak  ada maka autocreate jadi hati-hati dalam penggunaan nya
- Cara Penggunaan nya seperti perintah ```bash cd``` akan sedikit berbeda .

```Dockerfile
FROM golang:1.18-alpine
# Single env
ENV APP_ENV="development"
# Multiple env
ENV APP_PORT="8081" APP_NAME="web_server"
WORKDIR app
# workdir app
COPY . .

EXPOSE ${APP_PORT}
CMD go run main.go
```

- Jika kita perintah ```WORKDIR app/test``` maka kedua folder tersebut meski tidak ada tetap autocreate sehingga workdir nya app/test
- Jika kita perintah lagi ```WORKDIR /etc``` maka workdir nya berubah /etc

## User Instructions

- Berfungsi untuk menggunakan user atau ganti user .
- User Instruksi berguna ketika anda hanya menginjinkan folder tersebut di jalankan oleh user tertentu
- Contoh Penggunaan

```Dockerfile
FROM golang:1.22-alpine

RUN mkdir /app \
    && addgroup -S tirtahakimgroup \
    && adduser -S tirtahakim -G tirtahakimgroup \
    && chown -R tirtahakim:tirtahakimgroup /app

USER tirtahakim

WORKDIR /app
COPY . .

# beralih ke pengguna root untuk sementara agar perintah chown dapat dijalankan dengan benar.
USER root
RUN chown -R tirtahakim:tirtahakimgroup /app

# beralih ke pengguna tirtahakim 
USER tirtahakim
# Akses 755
#Membaca, menulis, dan mengeksekusi kepada pemilik (owner) file/direktori.
#Membaca dan mengeksekusi kepada grup pemilik (group) file/direktori.
#Membaca dan mengeksekusi kepada pengguna lain (others) di luar pemilik dan grup.
RUN chmod -R 755 /app

RUN go mod tidy

CMD go test ./test -v

```

## Arg Instructiions

- Berfungsi mendefinisikan variabel yang digunnakan oleh pengguna dan bisa di kirim saat build time menggunakann ```--build-arg key-value```
- Arg instruksi hanya berfungsi ketika build time, Jadi ketika container berjalan Arg instruksi tidak gunakan lagi alias sudah hilang . Kebalikan dari Env instruksi
- Cara penggunaan persis dengan Env instruksi

```Dockerfile
FROM golang:1.22.0-alpine
ARG APP_FILENAME="main"
ENV APP_PORT="8081" APP_NAME=${APP_FILENAME} APP_ENV="development" APP_LOG="/app/temp/log"

WORKDIR /app
COPY . .
RUN mv main.go ${APP_FILENAME}.go
RUN echo ${APP_NAME}
RUN go mod tidy

VOLUME ${APP_LOG}
CMD go run ${APP_NAME}.go
```

## Health Check Instructions

- Berfungsi memberitahukan bahwa container sedang baik baik saja atau tidak
- ```HEALTHCHECK [OPTIONS] CMD```
- OPTIONS
  - --interval=DURATION (30s) untuk jeda waktu setiap health check nya di jalankan
  - --timeout=DURATION (30s) batas waktu pengecek an health check
  - --start-period=DURATION (0s) jeda setelah jalan nya aplikasi  dan container kita
  - --retries=N (3) pengulangan health check ketika error berapa kali dijalankan untuk memastikan benar benar error
- Contoh Penggunaan

```Dockerfile
FROM golang:1.22.0-alpine
ARG APP_FILENAME="main"
ENV APP_PORT="8081" APP_NAME=${APP_FILENAME} APP_ENV="development" APP_LOG="/app/temp/log"

RUN apk --no-cache add curl
WORKDIR /app
COPY . .
RUN mv main.go ${APP_FILENAME}.go
RUN go mod tidy

VOLUME ${APP_LOG}
CMD go run ${APP_NAME}.go
HEALTHCHECK --interval=10s --timeout=60s --start-period=10s --retries=3 CMD curl http://localhost:${APP_PORT}/health
```

## Multi Stage

- Berfungsi ketika membutuh kan lebih dari 1 stage dan beda image
- Jadi Misal anda ingin menggunakan image alpine tetapi anda memiliki aplikasi golang maka anda butuh multi stage
- Bagaimana implementasi ? Mudah
  aplikasi anda yang ingin di jalankan di alpine atau image lain dibuat binary file terlebih dahulu yang stage nya memiliki image sesuai aplikasi anda setelah itu dipindahkan di di stage yang image nya alpine
- Contoh Penggunaan

```Dockerfile
FROM golang:1.22.0-alpine as build
ARG APP_FILENAME="main"
ENV APP_PORT="8081" APP_NAME=${APP_FILENAME} APP_ENV="development" APP_LOG="/app/temp/log"


WORKDIR /app
COPY . .
RUN mv main.go ${APP_FILENAME}.go
RUN go mod tidy

VOLUME ${APP_LOG}
RUN go build -o ./bin/${APP_FILENAME} ./${APP_FILENAME}.go

FROM alpine:latest
FROM build
ENV APP_FILENAME=${APP_NAME}
RUN apk --no-cache add curl
WORKDIR /bin
COPY --from=build /app/bin/${APP_FILENAME} .

CMD ./${APP_FILENAME}
HEALTHCHECK --interval=10s --timeout=60s --start-period=10s --retries=3 CMD curl http://localhost:${APP_PORT}/health
```
