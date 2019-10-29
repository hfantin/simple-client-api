FROM node:10-alpine
WORKDIR /opt/app
COPY package*.json ./
RUN npm install --only=prod
COPY . .
CMD [ "npm", "start" ]