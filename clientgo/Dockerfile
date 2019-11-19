FROM alpine
RUN apk add --update ca-certificates
WORKDIR /opt/app
ADD build/app app app
RUN chmod +x app
CMD [ "/opt/app/app" ]
