FROM alpine:3.13.5

RUN apk update && \
    apk add --no-cache tzdata libc6-compat

RUN mkdir app
ADD apptemplate app/

ENTRYPOINT [ "./app/apptemplate"]
  
