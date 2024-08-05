package waw

import (
	"testing"
)

func TestGetAPIKeyFromArg(t *testing.T) {
	input := "example"
	result, _ := GetAPIKey(input)
	if result != input {
		t.Fatal("GetAPIKey(input) did not return the value from param")
	}
}

func TestGetAPIKeyFromEnv(t *testing.T) {
	input := "example"
	t.Setenv("GOWAW_KEY", input)
	result, _ := GetAPIKey()
	if result != input {
		t.Fatal("GetAPIKey() did not return the value from env, even though it was set")
	}
}

func TestGetAPIKeyNoKey(t *testing.T) {
	t.Setenv("GOWAW_KEY", "")
	_, err := GetAPIKey()
	if err == nil || err.Error() != "API Key could not be found in GOWAW_KEY env var or in the parameters" {
		t.Fatal("GetAPIKey() did not return expected errors when both arguments and env variables were missing")
	}
}
