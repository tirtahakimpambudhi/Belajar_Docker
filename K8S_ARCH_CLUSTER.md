## Pengertian Kubernetes

[Kubernetes](https://kubernetes.io/) adalah platform open-source yang dirancang untuk mengotomatiskan proses pengelolaan, penskalaan, dan penerapan aplikasi container. Kubernetes menyediakan kerangka kerja untuk menjalankan aplikasi containerized di lingkungan yang terdistribusi, membantu memastikan ketersediaan tinggi, skalabilitas, dan efisiensi sumber daya.

## Arsitektur Cluster Kubernetes
![architecture kubernetes](https://kubernetes.io/images/docs/kubernetes-cluster-architecture.svg)
Cluster Kubernetes terdiri dari beberapa komponen utama yang bekerja bersama untuk mengelola aplikasi yang berjalan di dalamnya. Berikut adalah komponen-komponen utama dari arsitektur Kubernetes:

### 1. *Node*

- *Master Node*: Mengontrol dan mengelola seluruh cluster. Master node terdiri dari beberapa komponen seperti API Server, Scheduler, Controller Manager, dan etcd.
  - *API Server*: Komponen utama yang berinteraksi dengan komponen lain dalam Kubernetes. Semua perintah ke cluster Kubernetes dikirim ke API Server.
  - *Scheduler*: Bertanggung jawab untuk menugaskan Pod ke Node berdasarkan sumber daya yang tersedia dan kebutuhan aplikasi.
  - *Controller Manager*: Mengelola kontroler yang menangani berbagai tugas seperti replikasi Pod, pengaturan endpoint, dan lainnya.
  - *etcd*: Penyimpanan key-value yang menyimpan semua data konfigurasi cluster.

- *Worker Node*: Menjalankan aplikasi dalam bentuk Pod dan dikendalikan oleh Master Node. Worker Node terdiri dari beberapa komponen:
  - *Kubelet*: Agen yang berjalan di setiap Node dan memastikan container dijalankan sebagaimana mestinya.
  - *Kube-proxy*: Mengelola jaringan dalam cluster, memastikan komunikasi antar Pod dan layanan berjalan lancar.
  - *Container Runtime*: Perangkat lunak yang menjalankan container, misalnya Docker atau containerd.

### 2. *Pod*

- Unit terkecil dalam Kubernetes yang berisi satu atau lebih container. Setiap Pod memiliki IP unik dan berbagi penyimpanan serta jaringan. Pod adalah cara Kubernetes mengelola aplikasi containerized.

### 3. *Service*

- Abstraksi untuk mendefinisikan kumpulan Pod dan menyediakan cara yang stabil untuk mengaksesnya. Service memungkinkan Anda mengakses aplikasi di dalam Pod tanpa perlu tahu IP Pod secara langsung.

### 4. *Namespace*

- Mekanisme untuk mengisolasi dan mempartisi sumber daya dalam cluster. Namespace memungkinkan pengguna untuk mengelola beberapa lingkungan seperti dev, staging, dan production dalam satu cluster.

### 5. *Deployment*

- Abstraksi yang mendefinisikan bagaimana aplikasi di-deploy dan di-manage di dalam cluster. Deployment mengelola proses scaling dan rolling update Pod secara otomatis.

### 6. *ConfigMap dan Secret*

- *ConfigMap*: Menyimpan data konfigurasi dalam bentuk key-value yang bisa digunakan oleh Pod.
- *Secret*: Menyimpan data sensitif seperti password dan token yang juga dalam bentuk key-value, tetapi lebih aman dari ConfigMap.

## Cara Kerja Cluster Kubernetes

1. *Deployment Aplikasi*: Anda membuat dan mendefinisikan file konfigurasi untuk Deployment yang mendeskripsikan aplikasi, jumlah replica, dan strategi update.
2. *Scheduling dan Penempatan*: Scheduler Kubernetes menempatkan Pod pada Node yang memiliki sumber daya yang cukup sesuai dengan kebutuhan aplikasi.
3. *Networking dan Service Discovery*: Kubernetes mengatur jaringan untuk memungkinkan komunikasi antar Pod dan mengelola layanan agar dapat diakses oleh pengguna atau aplikasi lain.
4. *Monitoring dan Logging*: Kubernetes memantau keadaan Pod dan Node, dan menyediakan log untuk analisis lebih lanjut.
5. *Skalabilitas*: Kubernetes secara otomatis menyesuaikan jumlah Pod berdasarkan beban kerja dan kebijakan yang telah ditentukan.
6. *Pengelolaan Update*: Kubernetes mendukung rolling updates untuk meminimalkan downtime dan memastikan kontinuitas layanan saat aplikasi diperbarui.
