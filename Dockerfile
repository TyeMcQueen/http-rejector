FROM golang

WORKDIR /app/build
COPY . .

RUN CGO_ENABLED=0 go build ./rejector.go


FROM scratch

WORKDIR /app

COPY --from=0 /app/build/rejector .

USER 101:101

EXPOSE 8000/tcp
ENTRYPOINT [ "/app/rejector" ]
