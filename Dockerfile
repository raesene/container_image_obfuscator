FROM scratch

COPY image.tar /
COPY main /

ENTRYPOINT ["/main","image.tar"]