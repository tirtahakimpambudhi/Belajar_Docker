# Replication Controller, Replication Set & Daemon Set

## Pengertian

Replication Controller (RC), ReplicaSet (RS), dan DaemonSet (DS) adalah komponen dalam Kubernetes yang bertugas untuk mengelola siklus hidup dari pod. Masing-masing memiliki peran yang sedikit berbeda:

- **Replication Controller (RC)** adalah fitur Kubernetes yang paling tua yang bertujuan untuk menjaga jumlah replika pod tetap sesuai dengan yang diinginkan. Jika jumlah pod kurang dari yang diharapkan, RC akan membuat pod baru. Jika lebih, maka akan menghapus pod yang berlebih.

- **ReplicaSet (RS)** adalah penerus dari Replication Controller yang mendukung selektor yang lebih ekspresif. RS memastikan bahwa jumlah replika pod tertentu tersedia di cluster setiap saat.

- **DaemonSet (DS)** memastikan bahwa semua (atau beberapa) Node menjalankan salinan dari sebuah pod. Saat Node baru ditambahkan ke cluster, pod dari DaemonSet akan otomatis ditambahkan ke Node tersebut. DaemonSet cocok untuk menjalankan service yang harus berjalan pada semua Node, seperti agent log atau monitoring.

## Perbedaan Utama

- **Replication Controller** tidak mendukung selektor label yang ekspresif seperti ReplicaSet.
- **ReplicaSet** adalah versi yang lebih baru dan lebih fleksibel dari Replication Controller dengan dukungan untuk selektor label yang lebih kompleks.
- **DaemonSet** berbeda dengan RC dan RS karena DS memastikan bahwa setiap Node memiliki salinan pod yang dijalankan, yang sangat berguna untuk tugas-tugas yang harus berjalan di setiap Node.

## Replication Controller

![Replication Controller](./assets/images/rc_arch.png)
*Gambar Arsitektur Replication Controller*

![Replication Controller](./assets/images/rc_arch_after.png)
*Gambar Arsitektur Replication Controller Jika Pods 1 bermasalah atau hilang*

Seperti yang kita jelaskan sebelum nya Replication Controller akan menambah Pods Sesuai Jumlah yang di konfigurasikan jika Pods nya kurang . Jika lebih maka akan di kurangi sehingga Sesuai Jumlah yang di konfigurasikan.

## Contoh Konfigurasi Replication Controller 

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: replication-controller
  labels:
    replication-controller: ltrue
  annotations:
    description: "to keep the number of web-server pods"
spec:
  replicas: 2
  selector:
    app: "web-server"
  template:
    metadata:
      name: "web-server-nginx"
      labels:
        app: "web-server"
      annotations:
        description: "Web Server Pods"
    spec:
      containers:
        - name: nginx
          image: "nginx:alpine"
          ports:
            - containerPort: 80
          startupProbe:
            httpGet:
              path: /
              port: 80
            initialDelaySeconds: 10
            periodSeconds: 30
            failureThreshold: 3
            successThreshold: 1
            timeoutSeconds: 5
```

Jika nilai selector dalam spesifikasi ReplicationController tidak cocok dengan labels yang didefinisikan pada pod yang dikelola olehnya, maka ReplicationController tidak akan dapat mengelola pod tersebut dengan benar. Dalam konteks Kubernetes, selector digunakan untuk menemukan pod mana yang harus dikontrol berdasarkan labels yang sesuai.

Kesimpulan nya nilai selector harus sama dengan labels yang di pods template agar bisa di selector untuk di jaga jumlah pods nya.