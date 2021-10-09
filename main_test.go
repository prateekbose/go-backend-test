package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUserSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:3000/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := req.URL.Query()
	r.Add("Id", "1")
	r.Add("Name", "Prateek Bose")
	r.Add("Email", "prateekbose20011@gmail.com")
	r.Add("Password", "password")

	req.URL.RawQuery = r.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddPostSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:3000/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := req.URL.Query()
	r.Add("UserId", "1")
	r.Add("Id", "101")
	r.Add("Caption", "Some Caption")
	r.Add("ImageURL", "https://images.unsplash.com/photo-1632516152703-078685593865")

	req.URL.RawQuery = r.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddPostUnauthorized(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:3000/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := req.URL.Query()
	r.Add("UserId", "1001")
	r.Add("Id", "101")
	r.Add("Caption", "Some Caption")
	r.Add("ImageURL", "https://images.unsplash.com/photo-1632516152703-078685593865")

	req.URL.RawQuery = r.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetPost(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000/posts/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserPostsUnauthorised(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000/posts/users/11010", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUserAndPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
