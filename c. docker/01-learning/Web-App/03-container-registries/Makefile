.PHONY: build
build:
	docker build --tag my-scratch-image .

# Have to authenticate to dockerhub and create repo first:
# https://docs.docker.com/engine/reference/commandline/login/
.PHONY: push-dockerhub
push-dockerhub:
	docker tag my-scratch-image ashishkhanagwal/scratch-image # defaults to latest
	docker push ashishkhanagwal/scratch-image

	docker tag my-scratch-image ashishkhanagwal/scratch-image:v-1.0.0
	docker push ashishkhanagwal/scratch-image:v-1.0.0

# Have to authenticate to ghcr.io first
# https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-to-the-container-registry
.PHONY: push-github-packages
push-github-packages:
	docker tag my-scratch-image ghcr.io/ashish-khanagwal/scratch-image # defaults to latest
	docker push ghcr.io/ashish-khanagwal/scratch-image

	docker tag my-scratch-image ghcr.io/ashish-khanagwal/scratch-image:v-1.0.0 # defaults to latest
	docker push ghcr.io/ashish-khanagwal/scratch-image:v-1.0.0
