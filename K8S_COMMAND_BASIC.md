# K8s Command CLI BASIC

## Untuk Melihat Informasi Cluster

```bash
kubectl cluster info
```

## Untuk Melihat Config

```bash
kubectl config view
```

## Untuk Melihat Deskripsi Sebuah Resource

- Jenis Jenis Resource ada pods,nodes,namespaces,replication controller,deployment,service dll
- short cut
  - pods -> po
  - namespaces -> ns
  - replication controller -> rc
  - service -> svc
  - nodes -> no
  - deployment -> deploy

```
kubectl describer type_resource name_resource
```

## Untuk Melihat Sebuah List Resource

- Jenis Jenis Resource ada pods,nodes,namespaces,replication controller,deployment,service dll
- short cut
  - pods -> po
  - namespaces -> ns
  - replication controller -> rc
  - service -> svc
  - nodes -> no
  - deployment -> deploy
- setiap resource memiliki flag tambahan

```
kubectl get type_resource
```

## Untuk Menciptakan atau Memperbarui Resource yang Didefinisikan Dalam File YAML Atau JSON

```bash
kubectl apply -f name_file.yaml
```

## Untuk Menciptakan Resource yang Didefinisikan Dalam File YAML Atau JSON

- Akan Error jika dilakukan kedua kali

```bash
kubectl create -f name_file.yaml
```

## Untuk Menciptakan Resource yang Tanpa Didefinisikan File YAML Atau JSON

```bash
kubectl create type_resource name_resource --requirment-flag
```
