FROM golang:1.17 AS compiler
WORKDIR /src/app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./a.out .

FROM scratch
COPY --from=compiler /src/app/a.out /server
ENTRYPOINT ["/server"]