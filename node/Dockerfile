FROM node:lts-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8080

ENTRYPOINT [ "npm", "start" ]