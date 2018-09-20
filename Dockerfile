FROM scratch
EXPOSE 8080
ENTRYPOINT ["/serviceslist"]
COPY ./bin/ /