FROM adoptopenjdk/openjdk11-openj9
ARG JAR_FILE
COPY ${JAR_FILE} app.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom", "-Xms256m", "-Xmx1G", "-Xshareclasses", "-Xquickstart", "-jar","app.jar"]
