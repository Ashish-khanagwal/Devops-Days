# FROM node:15 means that you are using the official Node.js Docker image with version 15 as the starting point for building your own Docker image.
# This base image already contains Node.js and related tools, which makes it a good foundation for creating containers that run Node.js applications.
FROM node:15

WORKDIR /app

# This command copies the package.json file from the local context (the directory where the Dockerfile resides) into the current working directory of the image.
# Since we set the working directory to /app earlier, this means the package.json file from your local directory will be copied into the /app directory within the image.
COPY package.json .

RUN npm install
COPY . ./ 

ENV PORT 3000
EXPOSE $PORT

# CMD ["node", "index.js"]
# It's telling Docker to run the node command with the argument index.js. The node command is used to execute JavaScript files in a Node.js runtime environment, and index.js is the file you want to run.

CMD ["node", "index.js"]