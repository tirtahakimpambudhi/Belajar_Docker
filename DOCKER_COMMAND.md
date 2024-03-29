# Dokumentasi Docker Command

## Melihat Semua Image
```bash
docker image ls
```

## Cara Menginstall Image dari Docker Registry Bersifat Append
```bash
docker image pull name:tag
```

## Cara Menghapus Image
```bash
docker image rm name:tag
```

## Melihat Semua Container hanya yang di jalankan
```bash
docker container ls
```

## Melihat Semua Container yang Di-run maupun Tidak
```bash
docker container ls -a
```
## Membuat Container
```bash
docker container create --name name image:tag
```

## Menjalankan Container
```bash
docker container start name_container or container_ID
```
## Memberhentikan Container
```bash
docker container stop name_container or container_ID
```

## Menghapus Container
```bash
docker container rm name_container
```
## Logging di Container
```bash
docker container logs name_container
```
## Realtime Logging di Container
```bash
docker container logs -f name_container
```

## Container Exec
```bash
docker container exec -i -t name_container /bin/bash
# -i untuk input
# -t mode akses terminal
# shortcut -it
```
## Port Forwarding
```bash
docker container create --name name --publish or -p porthost:portcontainer image:tag
```

## Multiple Port Forwarding
```bash
docker container create --name name --publish or -p porthost:portcontainer --publish or -p porthost:portcontainer image:tag
```

## Membuat Container dengan Environment Variable
```bash
docker container create --name db_container_1 -p 3307:3306 --env or -e ENV_ROOT_PASSWORD=secret image:tag
```
## Multiple Environment Variable
```bash
docker container create --name db_container_1 -p 3307:3306 --env or -e ENV_ROOT_USERNAME=user  --env or -e ENV_ROOT_PASSWORD=secret image:tag
```
## Container Stats
```bash
docker container stats #Untuk melihat pemakai cpu dan storage setiap container
```

## Container Resource Limit
```bash
docker container create --name name_container_1 --memory 100m --cpus 0.5 image:tag
```
```bash
# Jika Mau Update jangan lupa atur memory swap kembali (memory swap > memory)
docker container update --memory 200m --memory-swap 400m --cpus 0.7 name_container_1
```
## Bind Mount
```bash
docker container create --name name_container_1 --mount "type=bind,source=folder_host,destination=folder_container,readonly(optional)" image:tag
```
## Docker Volume
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