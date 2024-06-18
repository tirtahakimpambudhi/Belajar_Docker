# Container

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
