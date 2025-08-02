# Stage 1: Build
FROM node:22-alpine AS builder

WORKDIR /app

# Install build dependencies needed for Node.js native modules
RUN apk add --no-cache libc6-compat make g++

# Copy only dependency manifest files first for better layer caching
COPY applications/admin-ui/package.json applications/admin-ui/package-lock.json ./

# Install only production dependencies (skip devDep for smaller size)
RUN npm ci --production && npm audit fix --omit=dev

# Copy the full app source
COPY applications/admin-ui .

# Build the production-optimized React static files
RUN npm run build

# Stage 2: Serve React app with NGINX
FROM nginx:alpine

# Copy built static assets from builder stage to NGINX web root
COPY --from=builder /app/build /usr/share/nginx/html

# Set appropriate file ownership for NGINX to serve the files securely
RUN chown -R nginx:nginx /usr/share/nginx/html

# Expose port 80 for HTTP
EXPOSE 80

# Start NGINX in foreground
CMD ["nginx", "-g", "daemon off;"]
