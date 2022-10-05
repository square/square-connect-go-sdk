FROM swaggerapi/swagger-codegen-cli:2.4.15

COPY square_connect_openapi.patch build/
COPY model_catalog_custom_attribute_value.patch build/
COPY model_catalog_item.patch build/
COPY model_item_variation_location_overrides.patch build/
COPY open-api-3_square build/

# Download swagger-codegen-cli 3.0.21 since I didn't find any Docker image
RUN apk add curl jq go && \
    cd build && \
    wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.21/swagger-codegen-cli-3.0.21.jar -O swagger-codegen-cli.jar && \
    cat open-api-3_square | jq '.content["content"] | fromjson' > square-connect-openapi.json    && \
    patch < square_connect_openapi.patch && \
    mkdir square-connect-sdk && \
    java -jar ./swagger-codegen-cli.jar generate -i square-connect-openapi.json -l go -o square-connect-sdk && \
    cd square-connect-sdk && \
#    && \
#    patch < ../model_catalog_custom_attribute_value.patch  && \
    patch < ../model_catalog_item.patch
#    patch < ../model_item_variation_location_overrides.patch

# fix some issues with Swagger. TODO: file an issue
RUN cd build/square-connect-sdk && \
    gofmt -s -w . && \
    cd ../..

ENTRYPOINT ["/bin/sh"]
