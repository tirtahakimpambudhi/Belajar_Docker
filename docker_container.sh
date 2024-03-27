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


