# square-connect-go-sdk
Square Connect SDK generated using our [open-api spec](https://docs.connect.squareup.com/v2/api/specifications/open-api-3_square) and Swagger

## Build

```
docker build -t square-connect-go-sdk .
docker run --rm -it -v $PWD:/mnt square-connect-go-sdk -c "rm -fr /mnt/swagger; cp -r /build/square-connect-sdk/ /mnt/swagger"
```
