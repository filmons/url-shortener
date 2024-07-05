# Use the official Node.js image as the base image
FROM node:14-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Install serve to serve the built application
RUN npm install -g serve

# Expose the port on which the application will run
EXPOSE 3000

# Command to run the application
CMD ["serve", "-s", "build"]
