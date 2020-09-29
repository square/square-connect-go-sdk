# square-connect-go-sdk
Square Connect SDK generated using our open-api spec and Swagger

## Building

```
docker build -t square-connect-go-sdk .
docker run --rm -it -v $PWD:/mnt square-connect-go-sdk -c "cp -r /build/square-connect-sdk/ /mnt/swagger"
```
