package main

import (
	"strings"
	"testing"
)

// TestChangeTag
func TestChangeTag(t *testing.T) {
	vals := []string{"bla-bla", "image: repo/srv:tag", "goo-goo"}
	r := Request{Service: "srv", Namespace: "test", Tag: "123", Repository: "repo/srv"}
	res := changeTag(strings.Join(vals, "\n"), r)
	if !strings.Contains(res, "123") {
		t.Errorf("Fail TestChangeTag, %s\n", res)
	}
}

// TestChangeTagNoTag
func TestChangeTagNoTag(t *testing.T) {
	vals := []string{"bla-bla", "image: repo/srv", "goo-goo"}
	r := Request{Service: "srv", Namespace: "test", Tag: "123", Repository: "repo/srv"}
	res := changeTag(strings.Join(vals, "\n"), r)
	if !strings.Contains(res, "123") {
		t.Errorf("Fail TestChangeTagNoTag, %s\n", res)
	}
}

// TestToken
func TestToken(t *testing.T) {
	r := Request{Service: "srv", Namespace: "test", Tag: "123", Repository: "repo/srv"}
	token, err := genToken(r)
	if err != nil {
		t.Errorf("Fail TestToken, error %v\n", err)
	}
	req, err := decodeToken(token)
	if err != nil {
		t.Errorf("Fail TestToken, error %v\n", err)
	}
	if !compareRequests(req, r) {
		t.Errorf("Fail TestToken, %+v != %+v\n", req, r)
	}
}
