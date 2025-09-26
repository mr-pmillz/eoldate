FROM ghcr.io/mr-pmillz/alpine-bash-tini:latest

COPY eoldate /usr/local/bin/eoldate
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh /usr/local/bin/eoldate

ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]
