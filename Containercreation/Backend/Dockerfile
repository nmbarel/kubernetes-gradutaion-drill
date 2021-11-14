FROM centos as buildstage

ADD Backend /
RUN chmod +x /Backend

FROM scratch

COPY  --from=buildstage Backend /

CMD ["/Backend"]