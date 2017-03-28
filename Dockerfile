FROM alpine

MAINTAINER TFG Co <backend@tfgco.com>

RUN mkdir /app
ADD ./bin/go-helloworld /app/go-helloworld

WORKDIR /app

EXPOSE 8080

CMD /app/go-helloworld
