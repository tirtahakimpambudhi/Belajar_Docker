# Docker Volume
```bash
docker volume ls #Melihar Volume
docker volume create name_volume #Membuat Volume
docker volume rm name_volume #Menghapus Volume 
```
## Container Volume
```bash
docker container create --name name_container -p port_host:port_container --mount "type=volume,source=name_volume,destination=folder_container,readonly(optional)" --memory --cpus image:tag
```
## Backup Volume
### Flow Volume -> bind mount folder host
```bash
docker container stop # berhentikan container terlebih dahulu agar tidak ada perubahan
```
```bash
docker create --name backup_container --mount "type=volume,source=name_volume,destination=folder_container" --mount "type=bind,source=folder_host,destination=folder_container" image:tag # buat container baru dengan 2 mount volume yang ingin kita backup dan bind sistme host
```
```bash
docker container exec -it name_container /bin/bash
tar cvf folder_containar_bind_dest/backup_name.tar.gz /folder_container_volume_dest
```
### Short cut
```bash
docker container run --rm --name backup_mongo_container  --mount "type=volume,source=mongodata,destination=/data_mongodb" --mount "type=bind,source=/home/tirtahakim/Documents/DOCKER_COMMAND/backup,destination=/backup_mongodb" ubuntu:latest tar cvf /backup_mongodb/backup_inline.tar.gz /data_mongodb
 #run --rm #agar setelah dijalankan langsung kehapus container nya
```
## Restore Volume
### Flow Bind mount folder_backup_host -> volume
```bash
docker container stop # berhentikan container terlebih dahulu agar tidak ada perubahan
```
```bash 
docker container run --rm --name restore_mongo_container --mount "type=bind,source=folder_host_backup,destination=/backup" --mount "type=volume,source=new_volume,destination=/data" ubuntu:latest bash -c "cd /data && tar xfv /backup/backup_file.tar.gz --strip 1"
```