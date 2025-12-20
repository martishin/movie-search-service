KUBE_NS := movie-search
ENV_FILE := server/.env.prod

k8s-ns:
	kubectl create namespace $(KUBE_NS) --dry-run=client -o yaml | kubectl apply -f -

k8s-secret: k8s-ns
	kubectl -n $(KUBE_NS) create secret generic movie-search-env \
		--from-env-file=$(ENV_FILE) \
		--dry-run=client -o yaml | kubectl apply -f -

k8s-apply: k8s-secret
	kubectl -n $(KUBE_NS) apply -f k8s/postgres.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/redis.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/alloy-config.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/server.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/ingress.yaml
	$(MAKE) k8s-wait

k8s-deploy: k8s-apply

k8s-release:
	kubectl -n $(KUBE_NS) rollout restart deploy/movie-search-server
	$(MAKE) k8s-wait

k8s-wait:
	kubectl -n $(KUBE_NS) rollout status statefulset/postgres --timeout=240s
	kubectl -n $(KUBE_NS) rollout status deploy/redis --timeout=180s
	kubectl -n $(KUBE_NS) rollout status deploy/movie-search-server --timeout=240s

k8s-status:
	kubectl -n $(KUBE_NS) get all
	kubectl -n $(KUBE_NS) get ingress,certificate,pvc,cm,secret

k8s-describe:
	kubectl -n $(KUBE_NS) describe deploy/movie-search-server
	kubectl -n $(KUBE_NS) describe statefulset/postgres
	kubectl -n $(KUBE_NS) describe deploy/redis
	kubectl -n $(KUBE_NS) describe cm/alloy-config
	kubectl -n $(KUBE_NS) describe ingress/ms-api
