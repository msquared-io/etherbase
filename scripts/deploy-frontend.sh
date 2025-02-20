#!/bin/bash
set -e

export PROJECT_ID=venator-unsepulcher-3029873
export REGION=europe-west2

# Build frontend services
docker buildx build --platform linux/amd64 -f packages/3dgame/Dockerfile -t etherbase-3dgame . &
docker buildx build --platform linux/amd64 -f packages/frontend/Dockerfile -t etherbase-frontend . &

wait

gcloud auth configure-docker ${REGION}-docker.pkg.dev

# Push 3D game image
docker tag etherbase-3dgame ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-demo-3dgame:latest
docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-demo-3dgame:latest

# Push frontend image (commented out)
docker tag etherbase-frontend ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-demo-frontend:latest
docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-demo-frontend:latest

# Deploy 3D game
gcloud run deploy etherbase-demo-3dgame \
  --image ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-demo-3dgame:latest \
  --region ${REGION} \
  --platform managed \
  --allow-unauthenticated &

# Deploy frontend (commented out)
gcloud run deploy etherbase-demo-frontend \
  --image ${REGION}-docker.pkg.dev/${PROJECT_ID}/cloud-run-source-deploy/etherbase-demo-frontend:latest \
  --region ${REGION} \
  --platform managed \
  --allow-unauthenticated &

wait 