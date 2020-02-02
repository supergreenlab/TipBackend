FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD bin/tipserver /

EXPOSE 8080

ENTRYPOINT ["/tipserver"]
