FROM ubuntu
RUN mkdir /opt/app
ADD clientgo /opt/app/clientgo
RUN chmod +x /opt/app/clientgo
WORKDIR /opt/app
CMD ["/opt/app/clientgo"]
CMD ["./clientgo"]