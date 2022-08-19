FROM mitjaziv/swagger-codegen-cli:latest

COPY square_connect_openapi.patch build/
COPY model_catalog_custom_attribute_value.patch build/
COPY model_catalog_item.patch build/
COPY model_item_variation_location_overrides.patch build/
COPY open-api-3_square build/

# Download swagger-codegen-cli 3.0.21 since I didn't find any Docker image
RUN apt-get update && apt install -y curl jq golang-go patch
WORKDIR build
RUN wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.21/swagger-codegen-cli-3.0.21.jar -O swagger-codegen-cli.jar
RUN cat open-api-3_square | jq '.data.attributes["json-spec"] | fromjson' > square-connect-openapi.json
RUN patch < square_connect_openapi.patch
RUN mkdir square-connect-sdk
RUN java -jar ./swagger-codegen-cli.jar generate -i square-connect-openapi.json -l go -o square-connect-sdk
WORKDIR square-connect-sdk
RUN patch --ignore-whitespace < ../model_catalog_custom_attribute_value.patch
RUN patch --ignore-whitespace < ../model_catalog_item.patch
RUN patch --ignore-whitespace < ../model_item_variation_location_overrides.patch

# ## fix some issues with Swagger. TODO: file an issue
RUN gofmt -s -w .

ENTRYPOINT ["/bin/sh"]
