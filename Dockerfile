FROM busybox:ubuntu-14.04

ENTRYPOINT ["/lattice-app"]

COPY lattice-app /lattice-app
RUN chmod a+x /lattice-app