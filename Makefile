.PHONY: deploy
deploy:
	kubectl apply -f deploy/crds/cache.example.com_memcacheds_crd.yaml
	operator-sdk build docker.io/kuritka/memcached-operator:v0.0.1 && docker push docker.io/kuritka/memcached-operator:v0.0.1
	kubectl scale deployment memcached-operator -n default `` --replicas=0 1>/dev/null || true && kubectl apply -f deploy/operator.yaml
	sleep 3 && watch kubectl logs `kubectl get pod -l app=oper -n default -o jsonpath="{.items[0].metadata.name}"`

.PHONY: light
light:
	kubectl scale deployment memcached-operator -n default `` --replicas=0 1>/dev/null || true && kubectl apply -f deploy/operator.yaml
	sleep 3 && watch kubectl logs `kubectl get pod -l app=oper -n default -o jsonpath="{.items[0].metadata.name}"`

.PHONY: generate
generate:
	operator-sdk generate k8s && operator-sdk generate crds
