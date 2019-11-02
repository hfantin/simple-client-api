FROM alpine
RUN apt-get update && apt-get install libpq-dev -y
WORKDIR /opt/app
ADD /target/release/clientrs clientrs
RUN chmod +x clientrs
CMD ["/opt/app/clientrs"]
