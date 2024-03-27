# Untuk Melihat Semua Container
docker container ls
# Untuk Melihat Semua Container Yang Di run maupun tidak
docker container ls -a
# Membuat Container
docker container create --name name image:tag
# Menjalankan Container
docker container start name_container or container_ID
# Memberhentikan Container
docker container stop name_container or container_ID
# Menghapus Container
docker container rm name_container


# Logging di Container
docker container logs name_container
# Realtime Logging di Container
docker container logs -f name_container


# Container Exec untuk masuk ke Container dan Mengeksekusi kode program di dalam container
docker container exec -i -t name_container /bin/bash
-i # Untuk input tetap aktif
-t # Untuk mode terminal akses

# Container Port secara default hanya bisa di akses di dalam container yang bersifat isolation
# Agar Bisa Mengakses Port dari Container dengan port forwarding
docker container create --name name --publish or -p porthost:portcontainer image:tag
# Multiple Port Forwarding
docker container create --name name --publish or -p porthost:portcontainer --publish or -p porthost:portcontainer image:tag
# Port Container harus sesuai dengan di docs image registry nya karena kalau tidak sesuai tidak bisa di jalankan


# Membuat Container dengan memasukan environment variable
docker container create --name db_container_1 -p 3307:3306 --env or -e ENV_ROOT_PASSWORD=secret image:tag
# Multiple env
docker container create --name db_container_1 -p 3307:3306 --env or -e ENV_ROOT_USERNAME=user  --env or -e ENV_ROOT_PASSWORD=secret image:tag
