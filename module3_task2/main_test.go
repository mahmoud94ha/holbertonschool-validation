package main

import (
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"
)

func Test_HelloHandler(t *testing.T) {
  if !testing.Short() {
    t.Skip("Flag `-short` absent: skipping Unit Tests.")
  }

  tests := []struct {
    name         string
    queryString  string
    responseCode int
    body         string
  }{
    {
      name:         "Grace Hopper",
      queryString:  "name=Grace Hopper",
      responseCode: 200,
      body:         "Hello Grace Hopper!",
    },
    {
      name:         "Rosalind Franklin",
      queryString:  "name=Rosalind Franklin",
      responseCode: 200,
      body:         "Hello Rosalind Franklin!",
    },
    {
      name:         "Empty name",
      queryString:  "name=",
      responseCode: 400,
      body:         "",
    },
    {
      name:         "No name",
      queryString:  "",
      responseCode: 200,
      body:         "Hello there!",
    },
    {
      name:         "Multiple names",
      queryString:  "name=Foo&Name=Bar",
      responseCode: 200,
      body:         "Hello Foo!",
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      req, err := http.NewRequest("GET", "/hello", nil)
      if err != nil {
        t.Fatal(err)
      }

      queryParams, _ := url.ParseQuery(tt.queryString)
      req.URL.RawQuery = queryParams.Encode()
      rr := httptest.NewRecorder()
      handler := http.HandlerFunc(HelloHandler)
      handler.ServeHTTP(rr, req)

      expectedCode := tt.responseCode
      gotCode := rr.Code
      if gotCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %v want %v", gotCode, expectedCode)
      }

      expectedBody := tt.body
      gotBody := rr.Body.String()
      if gotBody != expectedBody {
        t.Errorf("handler returned unexpected body: got %v want %v", gotBody, expectedBody)
      }
    })
  }
}
