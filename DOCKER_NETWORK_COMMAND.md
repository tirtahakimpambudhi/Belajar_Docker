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