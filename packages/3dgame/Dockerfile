# Build etherbase-client
FROM node:20-slim AS hooks-builder

WORKDIR /app
WORKDIR /app/packages/etherbase-client
COPY packages/etherbase-client/package*.json ./
COPY packages/etherbase-client/tsconfig.json ./
COPY packages/etherbase-client/src ./src
RUN npm install
RUN npm run build

# Build 3dgame
FROM node:20-slim AS app-builder

WORKDIR /app
WORKDIR /app/packages/3dgame
COPY packages/3dgame/package*.json ./
COPY packages/3dgame .

# Create the node_modules/@msquared/etherbase-client directory and copy the built files
RUN mkdir -p ./node_modules/@msquared/etherbase-client
COPY --from=hooks-builder /app/packages/etherbase-client/dist ./node_modules/@msquared/etherbase-client/dist
COPY --from=hooks-builder /app/packages/etherbase-client/package.json ./node_modules/@msquared/etherbase-client/package.json

# Create public directory if it doesn't exist
RUN mkdir -p public

# Install dependencies and build
RUN npm install
ENV NEXT_TELEMETRY_DISABLED=1
RUN npm run build

# Production image
FROM node:20-slim AS runner

WORKDIR /app

# Create public directory
RUN mkdir -p public

# Copy necessary files from builder
COPY --from=app-builder /app/packages/3dgame/package*.json ./
COPY --from=app-builder /app/packages/3dgame/.next ./.next
COPY --from=app-builder /app/packages/3dgame/node_modules ./node_modules
COPY --from=app-builder /app/packages/3dgame/public ./public

# Start the application
ENV NODE_ENV=production
ENV PORT=8080
EXPOSE 8080

# Ensure proper entrypoint
ENTRYPOINT []
CMD ["npm", "run", "start"]