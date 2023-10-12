package main

import (
	"net/url"
	"testing"
)

func Test_isValidUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{
			url:  "http://localhost:8080/?coords=52.52+13.41",
			want: true,
		},
		{
			url:  "http://localhost:8080/favicon.ico",
			want: false,
		},
		{
			url:  "http://localhost:8080/",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, _ := url.Parse(tt.url)
			if got := isValidUrl(str); got != tt.want {
				t.Errorf("isValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
