FROM centos as buildstage

ADD Frontend /
RUN chmod +x /Frontend

FROM scratch

COPY  --from=buildstage Frontend /

CMD ["/Frontend"]