package main

import (
	"encoding/json"
	"testing"

	"github.com/square/square-connect-go-sdk/swagger"
	"github.com/stretchr/testify/require"
)

func TestEcomPatch(t *testing.T) {
	s := swagger.CatalogItem{}
	// Ensure that Ecom fields exist. As long as the test compiles we are good and don't need any assertions.
	s.EcomUri = "xyz"
	s.EcomImageUris = []string{"xyz"}
	s.EcomAvailable = true
	s.EcomVisibility = "xyz"
}

func TestOAuthPermissionPatch(t *testing.T) {
	// Snippet API Oauth permissions are available. As long as the test compiles we are good and don't need any assertions.
	_ = swagger.ONLINE_STORE_SITE_READ_OAuthPermission
	_ = swagger.ONLINE_STORE_SNIPPETS_READ_OAuthPermission
	_ = swagger.ONLINE_STORE_SNIPPETS_WRITE_OAuthPermission
}

func TestModelCatalogItemPatch(t *testing.T) {
	strValue := ""
	s := swagger.CatalogItem{
		Description:  &strValue,
		Name: &strValue,
	}
	// Ensure that omitempty has been removed
	json, err := json.Marshal(s)
	require.NoError(t, err)
	require.Equal(t, "{\"name\":\"\",\"description\":\"\"}", string(json))
}

func TestModelCatalogCustomAttributeValuePatch(t *testing.T) {
	strValue := ""
	boolValue := false
	s := swagger.CatalogCustomAttributeValue{
		StringValue:  &strValue,
		BooleanValue: &boolValue,
	}
	// Ensure that omitempty has been removed
	json, err := json.Marshal(s)
	require.NoError(t, err)
	require.Equal(t, "{\"string_value\":\"\",\"boolean_value\":false}", string(json))
}
