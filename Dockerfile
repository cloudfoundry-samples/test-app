FROM busybox:ubuntu-14.04

ENTRYPOINT ["/test-app"]

COPY test-app /test-app
RUN chmod a+x /test-app