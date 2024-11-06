# Frontend Dockerfile
FROM node:18-alpine

WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy the application code
COPY . .

# Expose port (adjust if necessary)
EXPOSE 3000

# Run the application
CMD ["npm", "run", "prod"]
