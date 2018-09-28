FROM openjdk:8-jre-alpine

RUN mkdir /app

WORKDIR /app

ADD build/libs/config-1.0.jar /app

ENTRYPOINT [ "java", "-jar", "config-1.0.jar" ]
