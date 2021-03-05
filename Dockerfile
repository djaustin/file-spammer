FROM golang:alpine AS build

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o file-spammer

FROM scratch
COPY --from=build /src/file-spammer /
CMD ["/file-spammer"]

