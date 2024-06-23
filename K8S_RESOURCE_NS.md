# Namespace

## Pengertian

Namespace adalah sebuah mekanisme untuk mengelompokkan dan memisahkan sumber daya dalam sebuah klaster. Ini berguna untuk mengelola sumber daya dengan lebih efisien dan untuk mencegah konflik nama antara berbagai aplikasi atau tim yang menggunakan klaster yang sama. Namespace juga bisa digunakan dalam selector sebuah sumber daya. Dengan ada nya Namespace bisa mencegah Name Collision (Tabrakan Nama).

## Membuat Namespace dalam file YAML atau JSON

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: namespace-app
```

## Membuat Namespace dengan Perintah CLI

```bash
kubectl create namespace namespace-app
```

## Menempatkan Pods dalam Namespace

Dengan Menempatkan Pods dalam berbeda Namespace bisa mencegah Name Collision (Tabrakan Name). Tapi ingat Namespace hanya berguna untuk mencegah konflik nama antara berbagai aplikasi atau sumber daya , meskipun berbeda Namespace Pods tetap bisa saling berkomunikasi antara Pods lainya. Jadi Namespace bukan untuk mengisolasi sumber daya hanya mencegah konflik name.

## Membuat Namespace

### Dengan file config YAML atau JSON

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: namespace-app
```

### Dengan CLI

```bash
kubectl create namespace namespace-app
```

## Menempatkan Pods kedalam Namespace

### Dengan file config YAML atau JSON

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
  namespace: namespace-app1
spec:
  containers:
  - name: mycontainer
    image: myimage
```

## Selector Pods Menggunakan Namespace

- Secara default saat kita command ```bash kubectl get po``` akan mengarah ke namespace bawaan dan namanya biasa nya default.
- Lalu bagaiman melihat pods yang berbeda namespace ? . Dengan menambahkan flags --namespace namanya atau -n. Contoh nya

```bash
kubectl get po -n namespace 
#or 
kubectl get po -namespace namespace
```

## Menghapus Namespace

- Ingat saat menghapus namespace maka sumber daya yang ada didalam nya pun ikut terhapus

```bash
kubectl delete ns namespace
```

## Melihat List Namespaces

```bash
kubectl get ns
kubectl get namespace
kubectl get namespaces
```
