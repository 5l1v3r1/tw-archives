// +build appengine

package main

import (
	"google.golang.org/appengine/aetest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowNotFoundPage(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	HandleNotFound(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("Status is not NotFound: %d", res.Code)
	}
}

func TestIndexHandler(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	url := "/"
	req, err := inst.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	IndexHandler(res, req, nil)

	if res.Code != http.StatusOK {
		t.Errorf("Status is not OK: %d", res.Code)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	url := "/healthcheck"
	req, err := inst.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	HealthCheckHandler(res, req, nil)
	if res.Code != http.StatusNoContent {
		t.Errorf("Status is not NoContent: %d", res.Code)
	}
}
