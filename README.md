helloapp
========

A simple application to demonstrate a deployment on kubernetes.

Make sure your gcloud project is set and export it:

```console
$ export HELLOAPP_PROJECT_ID=$(gcloud config get-value project)
```

Create a reserved static address, if it's created yet:

```console
$ gcloud compute addresses create geoipfix-europe-west1 --region europe-west1
```

Retrieve the ip address:

```console
$ gcloud compute addresses list
NAME                   REGION        ADDRESS         STATUS
geoipfix-europe-west1  europe-west1  35.205.140.105  IN_USE
```

Export it:

```console
export HELLOAPP_SERVICE_STATIC_IP=35.205.140.105
```

Deploy!

```console
scripts/release.sh
```

Launch the pinger:


```console
scripts/test.sh
```
