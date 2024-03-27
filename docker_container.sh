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


#Logging di Container
docker container logs name_container
# Realtime Logging di Container
docker container logs -f name_container
