package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestFormRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/any", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows doesnot have required fields when it does")
	}
}

func TestFormHas(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	form := New(r.PostForm)

	has := form.Has("someone")

	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Error("shows form doesnot have a field when it should have")
	}
}

func TestMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)

	if form.Valid() {
		t.Error("form shows min-length for non existent field")
	}

	isError := form.Error.Get("x")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "somw value")

	form = New(postedValues)

	form.MinLength("some_field", 100)

	if form.Valid() {
		t.Error("shows min length of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "sdkjfbksdjfnj")

	form = New(postedValues)

	form.MinLength("another_field", 5)

	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}

	isError = form.Error.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

}

func TestFormIsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email when it should not")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "mail@mail.com")

	form = New(postedValues)

	form.IsEmail("email")

	if !form.Valid() {
		t.Error("Got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "mail")

	form = New(postedValues)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("Got a valid email when we should not have")
	}
}
