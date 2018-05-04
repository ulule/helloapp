source scripts/version.sh

make docker-build

docker build -t gcr.io/ulule-services/helloapp:${HELLOAPP_VERSION} .
gcloud docker -- push gcr.io/ulule-services/helloapp:${HELLOAPP_VERSION}
