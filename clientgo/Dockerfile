FROM alpine
RUN apk add --update ca-certificates
WORKDIR /opt/app
ADD clientgo clientgo
RUN chmod +x clientgo
CMD [ "/opt/app/clientgo" ]
