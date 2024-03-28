# Jadi Docker Volume Secara default akan membuat sendiri ketika kita membuat container karena sebuah container harus punya volume
# Melihat Docker Volume
docker volume ls
# Membuat Volume
docker volume create name_volume
# Menghapus Volume
docker volume rm name_volume # Container Harus Berhenti

#Container Volume
docker container create --name name_container -p port_host:port_container --mount "type=volume,source=name_volume,destination=folder_container,readonly(optional)"  --memory --cpus image:tag

# Backup Volune
1. docker container stop # berhentikan container terlebih dahulu agar tidak ada perubahan
2. docker create  # buat container baru dengan 2 mount volume yang ingin kita backup dan bind sistme host
