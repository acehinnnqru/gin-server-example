from golang:1.21 as builder

### Private repository support
# ARG GITHUB_TOKEN

# the github namespace should be replaced with your own
# ENV CGOENABLE=0 GOOS=linux GO111MODULE=on GOPRIVATE="github.com/acehinnnqru/*"
# WORKDIR name//${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /app
COPY . .

RUN go mod tidy && go mod download
RUN go build -a -installsuffix nocgo -ldflags '-linkmode "external" -extldflags "-static"' -o ./app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
