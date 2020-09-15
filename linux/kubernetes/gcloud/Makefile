CLUSTER-NAME=flux-demo
COMPUTE-ZONE=us-east1-b
PROJECT-NUMBER=1034691625917
PROJECT-NAME=anthos-playground
PROJECT-ID=anthos-playground-286117

list-projects:
	gcloud projects list --filter="name=$(PROJECT-NAME)"

cluster-create:
	gcloud container clusters create $(CLUSTER-NAME) --zone $(COMPUTE-ZONE) \
		--enable-ip-alias \
		--no-enable-autoupgrade \
		--enable-autorepair

cluster-delete:
	gcloud container clusters delete $(CLUSTER-NAME) --zone $(COMPUTE-ZONE)

get-credentials:
	gcloud container clusters get-credentials $(CLUSTER-NAME) --zone $(COMPUTE-ZONE)

clone-hello:
	git clone https://github.com/GoogleCloudPlatform/kubernetes-engine-samples

build-hello:
	cd kubernetes-engine-samples/hello-app; \
	docker build -t gcr.io/$(PROJECT-ID)/hello-app:v1 .

push-hello:
	podman push gcr.io/$(PROJECT-ID)/hello-app:v1

list-hello:
	gcloud container images list-tags gcr.io/$(PROJECT-ID)/hello-app

delete-hello:
	gcloud container images delete gcr.io/$(PROJECT-ID)/hello-app:v1

httpd-list:
	podman images list

httpd-tag:
	podman tag docker.io/library/httpd gcr.io/$(PROJECT-ID)/httpd-server:v1

httpd-push:
	podman push gcr.io/$(PROJECT-ID)/httpd-server:v1

