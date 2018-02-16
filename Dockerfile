FROM scratch
ADD /dist/main /
CMD ["/main"]

EXPOSE 8000