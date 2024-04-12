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
- Terbagi 2 cara long syntax dan short syntax 
- Long Syntax 
```yaml
services:
    nginx_example_service:
    container_name: "nginx_example_service"
    image: nginx:latest
    ports:
      - target: 80
        published: 8081
        protocol: tcp
```
- Short Syntax
```yaml
services:
    nginx_example_service: 
    container_name: "nginx_example_service"
    image: nginx:latest
    ports:
    #host_port:container_port/protocol
      - "8081:80/tcp"
```

## Env
- Contoh Penggunaan 
```yaml
services:
  mongodb_example_env:
    container_name: "mongodb_example_env"
    image: mongo:latest
    ports:
      - "8081:27017/tcp"
    environment:
      MONGO_INITDB_ROOT_USERNAME: username
      MONGO_INITDB_ROOT_PASSWORD: password
      MONGO_INITDB_DATABASE: database
```