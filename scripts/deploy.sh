#!/bin/bash
set -e

# Make scripts executable
chmod +x scripts/deploy-backend.sh
chmod +x scripts/deploy-frontend.sh

# Deploy both backend and frontend in parallel
./scripts/deploy-backend.sh &
./scripts/deploy-frontend.sh &

wait
  