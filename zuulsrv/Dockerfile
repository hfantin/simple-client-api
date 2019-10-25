FROM openjdk:8-jdk-alpine
RUN  apk update && apk upgrade && apk add netcat-openbsd
ARG JAR_FILE
COPY ${JAR_FILE} app.jar
ADD run.sh run.sh
RUN chmod +x run.sh
CMD ./run.sh