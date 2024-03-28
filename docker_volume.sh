# Jadi Docker Volume Secara default akan membuat sendiri ketika kita membuat container karena sebuah container harus punya volume
# Melihat Docker Volume
docker volume ls
# Membuat Volume
docker volume create name_volume
# Menghapus Volume
docker volume rm name_volume # Container Harus Berhenti

#Container Volume
docker container create --name name_container -p port_host:port_container --mount "type=volume,source=name_volume,destination=folder_container,readonly(optional)"  --memory --cpus image:tag

# Backup Volume 
# Flow Volume -> bind mount folder host
1. docker container stop # berhentikan container terlebih dahulu agar tidak ada perubahan
2. docker create --name backup_container --mount "type=volume,source=name_volume,destination=folder_container" --mount "type=bind,source=folder_host,destination=folder_container" image:tag # buat container baru dengan 2 mount volume yang ingin kita backup dan bind sistme host
 run --rm #agar setelah berhenti langsung kehapus container nya
# Jadi cara kerja sistem nya  volume yang kita backup itu destination nya di taruh folder container untuk mount pertaman
# mount kedua bind mount destination nya di folder container bebas namanya dan sumber nya di taruh ke host agar ke sharing 
3. tar cvf folder_containar_bind_dest/backup_name.tar.gz /folder_container_volume_det
# masuk ke terminal dan masuk folder destination dari mount volume jadikan tar.gz dan di pindahkan ke destination mount bind agar ke sharing keluar
Short Cut
docker container run --rm --name backup_mongo_container  --mount "type=volume,source=mongodata,destination=/data_mongodb" --mount "type=bind,source=/home/tirtahakim/Documents/DOCKER_COMMAND/backup,destination=/backup_mongodb" ubuntu:latest tar cvf /backup_mongodb/backup_inline.tar.gz /data_mongodb


# Restore Volume
# Flow Bind mount folder_backup_host -> volume
1. docker container stop # berhentikan container terlebih dahulu agar tidak ada perubahan
2. docker container run --rm --name restore_mongo_container --mount "type=bind,source=folder_host_backup,destination=/backup" --mount "type=volume,source=new_volume,destination=/data" ubuntu:latest bash -c "cd /data && tar xfv backup/backup_file.tar.gz --strip 1"
