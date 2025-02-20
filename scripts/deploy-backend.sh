#!/bin/bash
set -e

export PROJECT_ID=venator-unsepulcher-3029873
export REGION=europe-west2

# Build Go services
(cd backend && GOOS=linux GOARCH=amd64 go build -o build/reader ./cmd/reader) &
(cd backend && GOOS=linux GOARCH=amd64 go build -o build/writer ./cmd/writer) &

wait

gcloud auth configure-docker ${REGION}-docker.pkg.dev

# Build and push reader service
docker build -f backend/Dockerfile.reader -t etherbase-reader backend
docker tag etherbase-reader ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-reader:latest
docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-reader:latest

# Build and push writer service
docker build -f backend/Dockerfile.writer -t etherbase-writer backend
docker tag etherbase-writer ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-writer:latest
docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-writer:latest

# Deploy reader service
gcloud run deploy etherbase-reader \
  --image ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-reader:latest \
  --region ${REGION} \
  --platform managed \
  --allow-unauthenticated \
  --max-instances 4 \
  --min-instances 1 \
  --cpu 2 \
  --memory 1024Mi \
  --timeout 3600s \
  --concurrency 80 &

# Deploy writer service
gcloud run deploy etherbase-writer \
  --image ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-writer:latest \
  --region ${REGION} \
  --platform managed \
  --allow-unauthenticated \
  --max-instances 4 \
  --min-instances 1 \
  --cpu 2 \
  --memory 1024Mi \
  --timeout 3600s \
  --concurrency 80 \
  --set-env-vars WRITER_PORT=8080 &

wait 