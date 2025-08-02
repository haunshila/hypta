# Stage 1: Build
FROM node:22-alpine AS builder
WORKDIR /app
# Install build dependencies for Node.js
RUN apk add --no-cache libc6-compat python3 make g++
# Copy package files and install production dependencies
COPY applications/digital-wallet-app/package.json applications/digital-wallet-app/package-lock.json ./
RUN npm ci --production && npm audit fix
COPY applications/digital-wallet-app .
RUN npm run build

# Stage 2: Runtime
FROM nginx:alpine
# Copy the built React app
COPY --from=builder /app/build /usr/share/nginx/html
# Copy custom NGINX config for proxying API requests
COPY applications/digital-wallet-app/nginx.conf /etc/nginx/conf.d/default.conf
# Ensure non-root user permissions
RUN chown -R nginx:nginx /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]