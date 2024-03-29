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


# Docker Network

## Docker Driver
- Bridge Ini adalah driver jaringan default untuk Docker.
Bridge network memungkinkan kontainer berkomunikasi satu sama lain pada host yang sama.
Setiap kontainer di jaringan bridge mendapatkan alamat IP yang dapat diakses secara lokal di host.
Jaringan bridge otomatis dibuat ketika Anda menjalankan kontainer dan tidak menspesifikasi jaringan tertentu.

- Host menggunakan driver ini, kontainer tidak diisolasi dari host.
Kontainer berbagi alamat IP dan port host.
Kontainer dapat diakses melalui alamat IP host tanpa memerlukan pengalihan port.
Host network sering digunakan ketika kontainer perlu mendengarkan port secara langsung di host tanpa memerlukan pembuatan port forwarding.

- Overlay network digunakan untuk menghubungkan kontainer di beberapa host yang berbeda.
Ini memungkinkan kontainer di host yang berbeda untuk berkomunikasi satu sama lain.
Overlay network menggunakan teknologi overlay seperti VXLAN untuk mengemas paket data dan mengirimkannya melalui jaringan yang sudah ada.
Overlay network biasanya digunakan dalam lingkungan Kubernetes atau ketika Anda menjalankan Docker Swarm mode untuk melakukan orkestrasi kontainer pada infrastruktur yang terdistribusi.

- Macvlan memungkinkan Anda untuk memberikan alamat MAC yang unik kepada setiap kontainer, sehingga membuatnya terlihat sebagai perangkat fisik di jaringan yang terhubung.
Kontainer dengan macvlan dapat berkomunikasi langsung dengan perangkat di jaringan fisik yang sama dengan host Docker.
Biasanya digunakan dalam situasi di mana Anda perlu kontainer memiliki kehadiran di jaringan yang sama dengan perangkat fisik, seperti dalam skenario ketika Anda ingin kontainer terlihat sebagai mesin fisik dalam jaringan lokal.

- None (Default Container) Driver ini menghilangkan semua konektivitas jaringan dari kontainer.
Kontainer yang menggunakan driver none tidak memiliki antarmuka jaringan yang terpasang.
Biasanya digunakan untuk kasus uji coba atau skenario di mana Anda ingin mengisolasi kontainer sepenuhnya dari jaringan.

## Melihat Network
```bash
docker network ls
```
## Membuat Network
```bash
docker create --driver driver(bridge,host,none,etc) name_network
```
## Menghapus Network
```bash
docker network rm name_network # jika network sudah digunakan container maka hapus container nya terlebih dahulu
```

## Container Network Bridge
```bash
# Agar Container bisa saling berkomunikasi kita harus di hubungkan dalam satu network bridge
# Contoh Implementasi 
# Menghubungkan mongo_express dengan mongodb
# 1. Buat Network
docker network create --driver bridge mongodb_network

# 2. Buat container mongodb dan di sambungkan dengan network
docker container create --name mongodb_container --network mongodb_network -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=12345 --memory 200mb --cpus 0.5 mongo:latest

# 3. Buat container mongo express dan di sambungkan dengan network
docker container create --name mongo_express_admin_ct --network mongodb_network -p 8081:8081 \
-e ME_CONFIG_MONGODB_ADMINUSERNAME="root" \
-e ME_CONFIG_MONGODB_ADMINPASSWORD="12345" \
-e ME_CONFIG_MONGODB_URL="mongodb://root:12345@mongodb_container:27017/" \
-e ME_CONFIG_BASICAUTH_USERNAME="admin" \
-e ME_CONFIG_BASICAUTH_PASSWORD="admin" \
--memory 150mb --cpus 0.5 mongo-express:latest
# Nama host dari nama container nya 
# ketika container ingin menghubungkan network pakai flag --network name_network 
```

## Menghapus Koneksi container dengan network
```bash 
docker network disconnect name_network name_container
```

## Menambahkan Koneksi container ke network 
```bash
docker network connect name_network name_container
```


# Docker Inspect
### Untuk Melihat Detail dari network,container,image dan volume

```bash
docker enum(network,container,image,volume) inspect name_(network,container,image,volume)
```

# Docker Prune
### Untuk Penghapusan Suatu network,container,image dan volume yang sudah tidak digunakan secara otomatis
```bash
docker enum(network,container,image,volume) prune 
```

### Short Cut Penghapusan semua network,container,image dan volume yang tidak  digunakan 
```bash
docker system prune
```
