FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD bin/tipbackend /

EXPOSE 8080

ENTRYPOINT ["/tipbackend"]
