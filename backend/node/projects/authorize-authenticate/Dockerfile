# Stage 1: Build the application
FROM node:20-alpine AS build

WORKDIR /app

# Install pnpm
RUN npm install -g pnpm

# Copy package files
COPY package.json pnpm-lock.yaml ./

# Install all dependencies (including devDependencies)
RUN pnpm install --frozen-lockfile

# Copy source files
COPY tsconfig.json ./
COPY src ./src

# Build the application
RUN pnpm run build

# Prune devDependencies
RUN pnpm prune --prod

# Stage 2: Create a smaller production image
FROM node:20-alpine

WORKDIR /app

# Copy built assets and production dependencies from build stage
COPY --from=build /app/dist ./dist
COPY --from=build /app/node_modules ./node_modules
COPY package.json ./

# Set environment variables
ENV NODE_ENV=production \
    PORT=3000

# Expose the application port
EXPOSE 3000

# Start the application
CMD ["node", "dist/index.js"]