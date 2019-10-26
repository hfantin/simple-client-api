FROM adoptopenjdk/openjdk11-openj9
#RUN  apk update && apk upgrade && apk add netcat-openbsd
ARG JAR_FILE
COPY ${JAR_FILE} app.jar
#ADD run.sh run.sh
#RUN chmod +x run.sh
#CMD ./run.sh
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom", "-Xms256m", "-Xmx1G", "-Xshareclasses", "-Xquickstart", "-jar","app.jar"]
