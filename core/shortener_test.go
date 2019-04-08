package core

import (
	"testing"
)

var shortener *UrlShortener
var url1 string
var url2 string

func init() {
	db := NewUrlShortenerArrayDb(10)
	const idSize = 6
	shortener = NewUrlShortener(db, idSize)
	url1 = "https://www.google.com"
	url2 = "https://yahoo.co.jp"
}

// TODO add more test cases

func TestCreateSameOrigin(t *testing.T) {
	shortened_a, err := shortener.Create(url1)
	if err != nil {
		t.Fatalf("Error occured: %s\n", err)
	}
	t.Logf("shortened_a: %s\n", shortened_a)

	shortened_b, err := shortener.Create(url1)
	if err != nil {
		t.Fatalf("Error occured: %s\n", err)
	}
	t.Logf("shortened_b: %s\n", shortened_b)

	if *shortened_a != *shortened_b {
		t.Fatal("shortened_a and shortened_b must be same because they have same original url")
	}
}

func TestCreateDifferentOrigin(t *testing.T) {
	shortened_a, err := shortener.Create(url1)
	if err != nil {
		t.Fatalf("Error occured: %s\n", err)
	}
	t.Logf("shortened_a: %s\n", shortened_a)

	shortened_b, err := shortener.Create(url2)
	if err != nil {
		t.Fatalf("Error occured: %s\n", err)
	}
	t.Logf("shortened_b: %s\n", shortened_b)

	if *shortened_a == *shortened_b {
		t.Fatal("shortened_a and shortened_b must be different because they have different original url")
	}
}

func TestResolve(t *testing.T) {
	shortened_a, err := shortener.Create(url1)
	if err != nil {
		t.Fatalf("Error occured: %s\n", err)
	}
	t.Logf("shortened_a: %s\n", shortened_a)

	shortened_b, err := shortener.Resolve(shortened_a.id)
	if err != nil {
		t.Fatalf("Error occured: %s\n", err)
	}
	t.Logf("shortened_b: %s\n", shortened_b)

	if url1 != shortened_b.original {
		t.Fatal("not resolved properly")
	}
}
