FROM scratch

COPY bin/piemapping /var/www/

WORKDIR /var/www

EXPOSE 80

CMD ["./piemapping", "run"]
