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
      MONGO_INITDB_DATABASE: admin
```

## Mount
- source adalah path dari host atau nama volume
- target adalah path dari container
- mode ada 2 (readonly dan readwrite) default readwrite
### Bind
#### Short Syntax
```yaml
version: "3.9"

services:
  mongodb_example_bind:
    container_name: "mongodb_example_bind"
    image: mongo:latest
    ports:
      - "8081:27017/tcp"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: t1rt@h4k1m
      MONGO_INITDB_DATABASE: admin
    volumes:
    #source:target/mode
      - ../../../../../db/mongodb:/data/db:rw
```
#### Long Syntax
```yaml
version: "3.9"

services:
  mongodb_example_bind:
    container_name: "mongodb_example_bind"
    image: mongo:latest
    ports:
      - "8081:27017/tcp"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: t1rt@h4k1m
      MONGO_INITDB_DATABASE: admin
    volumes:
      - type: bind
        source: "../../../../../db/mongodb"
        target: "/data/db"
        read_only: false
```

### Volume

#### Short Syntax
```yaml
version: "3.9"

services:
  mongodb_example_bind:
    container_name: "mongodb_example_bind"
    image: mongo:latest
    ports:
      - "8081:27017/tcp"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: t1rt@h4k1m
      MONGO_INITDB_DATABASE: admin
    volumes:
    #source:target/mode
      - mongo_data:/data/db:rw
volumes:
  mongo_data:
    name: "mongo_data"
```
#### Long Syntax
```yaml
version: "3.9"

services:
  mongodb_example_bind:
    container_name: "mongodb_example_bind"
    image: mongo:latest
    ports:
      - "8081:27017/tcp"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: t1rt@h4k1m
      MONGO_INITDB_DATABASE: admin
    volumes:
      - type: volume
        source: mongo_data
        target: "/data/db"
        read_only: false
volumes:
  mongo_data:
    name: "mongo_data"
```
