FROM node:latest
MAINTAINER Murilo Locatelli

COPY . /var/www

WORKDIR /var/www

EXPOSE 3000

RUN npm install

ENTRYPOINT ["npm", "run", "dev"]
