#how to install set up http server with metrics

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm -n prometheus-stack install  kube-prometheus-stack prometheus-community/kube-prometheus-stack
kubectl get prometheuses -n prometheus-stack
kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml -n  prometheus-stack
kubectl edit prometheuses -n prometheus-stack -oyaml
  spec:    
		additionalScrapeConfigs:      
			key: prometheus-additional.yaml      
			name: additional-configs
kubectl apply -f deploy/rbac.yml
kubectl apply -f deploy/service.yml
kubectl apply -f deploy/deployment.yml
kubectl port-forward -n prometheus-stack  --address 0.0.0.0 svc/kube-prometheus-stack-prometheus 9090:9090 
