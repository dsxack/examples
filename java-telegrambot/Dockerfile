# build stage
FROM fabric8/java-alpine-openjdk8-jdk:latest AS build-env

WORKDIR /root/

ADD . .
RUN ./gradlew fatJar

FROM fabric8/java-alpine-openjdk8-jdk:latest

ENV JAVA_APP_JAR app.jar
ENV AB_OFF true

COPY --from=build-env /root/build/libs/*.jar /deployments/app.jar
