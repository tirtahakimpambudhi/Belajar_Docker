# PROBE

## Pengertian

PROBE dalam konteks Kubernetes adalah mekanisme yang digunakan untuk memeriksa kesehatan dari container yang berjalan dalam pod. Kubernetes menyediakan tiga jenis probe, yaitu:

- **Liveness Probe**: Memeriksa apakah container masih berjalan. Jika probe ini gagal, Kubernetes akan merestart container tersebut.
- **Readiness Probe**: Memastikan bahwa container siap menerima trafik. Jika gagal, container tersebut tidak akan menerima permintaan jaringan.
- **Startup Probe**: Digunakan untuk mengetahui kapan aplikasi dalam container telah siap berjalan. Startup Probe cocok untuk Pod yang membutuhkan proses startup lama, ini dapat digunakan untuk memastikan Pod tidak mati oleh kubelet sebelum selesai berjalan dengan sempurna.

Probe-probe ini dapat dikonfigurasi dengan berbagai cara pengecekan nya, seperti HTTP GET, TCP Socket, dan Exec Command, tergantung pada kebutuhan aplikasi yang dijalankan.

## Melihat Probe

```bash
kubectl get po
kubectl describe po namepods
```

## Konfigurasi Probe

Untuk mengkonfigurasi probe, beberapa parameter yang dapat diatur meliputi:

- **initialDelaySeconds**: Waktu tunggu sebelum probe pertama kali dijalankan setelah container dimulai. Defaultnya adalah 0.
- **periodSeconds**: Frekuensi (dalam detik) probe untuk menjalankan pengecekan. Defaultnya adalah 10.
- **timeoutSeconds**: Waktu maksimal (dalam detik) untuk menunggu respons dari probe. Jika probe tidak mendapat respons dalam waktu ini, maka dianggap gagal. Defaultnya adalah 1.
- **successThreshold**: Jumlah minimal keberhasilan berturut-turut yang diperlukan untuk probe dianggap berhasil setelah mengalami kegagalan. Defaultnya adalah 1.
- **failureThreshold**: Jumlah kegagalan yang diizinkan sebelum container dianggap gagal. Defaultnya adalah 3.

```yaml
# Contoh konfigurasi Liveness Probe dengan parameter-parameter di atas
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 5
  timeoutSeconds: 2
  successThreshold: 1
  failureThreshold: 3

# Contoh konfigurasi Readiness Probe dengan TCP Socket
readinessProbe:
  tcpSocket:
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
  timeoutSeconds: 1
  successThreshold: 1
  failureThreshold: 3

# Contoh konfigurasi Startup Probe dengan Exec Command
startupProbe:
  exec:
    command:
    - cat
    - /tmp/healthy
  initialDelaySeconds: 15
  periodSeconds: 10
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 3
```

## Cara melihat hasil Probe dari Pods

```bash
kubectl logs -n namespace name_pods
```
