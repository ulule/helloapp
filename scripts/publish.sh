source scripts/version.sh

make docker-build

docker build -t gcr.io/${GCLOUD_PROJECT_ID}/helloapp:${HELLOAPP_VERSION} .
gcloud docker -- push gcr.io/${GCLOUD_PROJECT_ID}/helloapp:${HELLOAPP_VERSION}
