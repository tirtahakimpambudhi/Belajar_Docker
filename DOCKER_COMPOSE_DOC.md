# Documentasi Docker Compose File



## Services
- Digunakan untuk menampung satu container atau lebih
- Contoh Penggunaan
```yaml
services:
    name_container_1:
     ....
    name_container_2:
     ....
```

## Container
- Contoh Penggunaan 
```yaml
services:
    name_container_1:
        container_name: "name_container_1"
        image: image_name:version
```

## Ports
- 