package main

import (
	"net/http"
	"testing"

	"github.com/danangamw/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlpath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlpath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlpath:  "/snippet/view/3",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlpath:  "/snippet/view/-3",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlpath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlpath:  "/snippet/view/bar",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlpath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlpath)
			assert.Equal(t, code, tt.wantCode)

			if tt.wantBody != "" {
				assert.StripContains(t, body, tt.wantBody)
			}
		})
	}
}

// func TestUserSignup(t *testing.T) {
// 	app := newTestApplication(t)
// 	ts := newTestServer(t, app.routes())
// 	defer ts.Close()

// 	_, _, body := ts.get(t, "/user/signup")
// 	validCSRFToken := extractCSRFToken(t, body)

// 	const (
// 		validName     = "Bob"
// 		validPassword = "validPa$$word"
// 		validEmail    = "bob@example.com"
// 		formTag       = "<form action='/user/signup' method='POST' novalidate>"
// 	)

// 	tests := []struct {
// 		name         string
// 		userName     string
// 		userEmail    string
// 		userPassword string
// 		csrfToken    string
// 		wantCode     int
// 		wantFormTag  string
// 	}{
// 		{
// 			name:         "Valid submission",
// 			userName:     validName,
// 			userEmail:    validEmail,
// 			userPassword: validPassword,
// 			csrfToken:    validCSRFToken,
// 			wantCode:     http.StatusSeeOther,
// 		},
// 		{
// 			name:         "Invalid CSRF Token",
// 			userName:     validName,
// 			userEmail:    validEmail,
// 			userPassword: validPassword,
// 			csrfToken:    "wrongToken",
// 			wantCode:     http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Empty Name",
// 			userName:     "",
// 			userEmail:    validEmail,
// 			userPassword: validPassword,
// 			csrfToken:    validCSRFToken,
// 			wantCode:     http.StatusUnprocessableEntity,
// 			wantFormTag:  formTag,
// 		},
// 		{
// 			name:         "Empty Password",
// 			userName:     validName,
// 			userEmail:    validEmail,
// 			userPassword: "",
// 			csrfToken:    validCSRFToken,
// 			wantCode:     http.StatusUnprocessableEntity,
// 			wantFormTag:  formTag,
// 		},
// 		{
// 			name:         "Invalid email",
// 			userName:     validName,
// 			userEmail:    "bob@example.",
// 			userPassword: validPassword,
// 			csrfToken:    validCSRFToken,
// 			wantCode:     http.StatusUnprocessableEntity,
// 			wantFormTag:  formTag,
// 		},
// 		{
// 			name:         "Short Password",
// 			userName:     validName,
// 			userEmail:    validEmail,
// 			userPassword: "pa$$",
// 			csrfToken:    validCSRFToken,
// 			wantCode:     http.StatusUnprocessableEntity,
// 			wantFormTag:  formTag,
// 		},
// 		{
// 			name:         "Duplicate Email",
// 			userName:     validName,
// 			userEmail:    "dupe@example.com",
// 			userPassword: validPassword,
// 			csrfToken:    validCSRFToken,
// 			wantCode:     http.StatusUnprocessableEntity,
// 			wantFormTag:  formTag,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			form := url.Values{}
// 			form.Add("name", tt.userName)
// 			form.Add("email", tt.userEmail)
// 			form.Add("password", tt.userPassword)
// 			form.Add("csrf_token", tt.csrfToken)

// 			code, _, body := ts.postForm(t, "/user/signup", form)

// 			assert.Equal(t, code, tt.wantCode)

// 			if tt.wantFormTag != "" {
// 				assert.StripContains(t, body, tt.wantFormTag)
// 			}
// 		})
// 	}
// }
