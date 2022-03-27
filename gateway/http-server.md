### Deploy httpserver

```sh
kubectl create ns myhttpserver
kubectl label ns myhttpserver istio-injection=enabled
kubectl create -f httpserver.yml -n myhttpserver
```

```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=myhttpserver Inc./CN=*.myhttpserver.io' -keyout myhttpserver.io.key -out myhttpserver.io.crt
kubectl create -n istio-system secret tls myhttpserver-credential --key=myhttpserver.io.key --cert=myhttpserver.io.crt
kubectl apply -f httpgw.yml -n myhttpserver
```

### Check ingress ip

```sh
k get svc -n istio-system

istio-ingressgateway   LoadBalancer   $INGRESS_IP
```

### Access the httpserver via ingress

```sh
curl -s -I -X HEAD http://httpserver-mesh.myhttpserver
```
