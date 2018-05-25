docker build -t gcr.io/${GCLOUD_PROJECT_ID}/helloapp-python:0.1 .
gcloud docker -- push gcr.io/${GCLOUD_PROJECT_ID}/helloapp-python:0.1
