FROM scratch
COPY microfest-server /
ENTRYPOINT ["/microfest-server"]
