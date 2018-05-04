source scripts/version.sh

DEPLOYMENT=$(envsubst < scripts/kubernetes/manifests.yml)
echo "Applying deployment manifest..."
echo ${DEPLOYMENT}
cat <<EOF | kubectl apply -f -
${DEPLOYMENT}
EOF
