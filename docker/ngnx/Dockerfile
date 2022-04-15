FROM ubuntu

MAINTAINER Murilo Locatelli <murilolocatelli@gmail.com>

RUN apt-get update

RUN apt-get install -y nginx

ADD exemplo /etc/nginx/sites-enabled/default

RUN echo "daemon off;" >> /etc/nginx/nginx.conf

RUN ln -sf /dev/stdout /var/log/nginx/access.log
RUN ln -sf /dev/stderr /var/log/nginx/error.log

VOLUME ["/usr/share/nginx/html"]

EXPOSE 80

CMD service nginx start