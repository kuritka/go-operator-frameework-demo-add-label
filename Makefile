.PHONY: deploy
deploy:
	kubectl apply -f deploy/crds/cache.example.com_memcacheds_crd.yaml
	operator-sdk build docker.io/kuritka/memcached-operator:v0.0.1 && docker push docker.io/kuritka/memcached-operator:v0.0.1
	kubectl scale deployment memcached-operator -n default `` --replicas=0 1>/dev/null || true && kubectl apply -f deploy/operator.yaml

.PHONY: debug
debug: create-test-ns
	kubectl apply -f ./deploy/crds/ohmyglb.absa.oss_gslbs_crd.yaml
	kubectl apply -f ./deploy/crds/ohmyglb.absa.oss_v1beta1_gslb_cr.yaml
	operator-sdk run --local --namespace=test-gslb --enable-delve

.PHONY: logs
logs:
	kubectl logs `kubectl get pod -l app=oper -n default -o jsonpath="{.items[0].metadata.name}"`

.PHONY: light
light:
	kubectl scale deployment memcached-operator -n default `` --replicas=0 1>/dev/null || true && kubectl apply -f deploy/operator.yaml


.PHONY: generate
generate:
	operator-sdk generate k8s && operator-sdk generate crds

.PHONY: operator-deploy
operator-deploy:
	kubectl apply -f deploy/service_account.yaml
	kubectl apply -f deploy/role.yaml
	kubectl apply -f deploy/role_binding.yaml
	kubectl apply -f deploy/operator.yaml

.PHONY: memcaheds-deploy
memcaheds-deploy:
	kubectl apply -f deploy/crds/cache.example.com_v1alpha1_memcached_cr.yaml