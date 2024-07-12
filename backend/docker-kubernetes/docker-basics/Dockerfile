FROM node:20-alpine3.20

# Set the working directory
WORKDIR /usr/app

# Copy package.json and possibly pnpm-lock.yaml if you have it
COPY package.json pnpm-lock.yaml* ./

# Install pnpm and install dependencies with it
RUN npm install -g pnpm && pnpm install

# Copy the rest of your application code
COPY . .

# Command to run your application
CMD ["pnpm", "start"]