FROM alpine
RUN apk add --update ca-certificates
WORKDIR /opt/app
ADD clientgo clientgo
RUN chmod +x clientgo
# CMD ["clientgo"]
# CMD ["pwd"]
# CMD ["ls"]
CMD [ "/opt/app/clientgo" ]
