package main

import "testing"

var emailTest string
var email string
var gopher bool

func BenchmarkIsGopher(b *testing.B) {
	emailTest = "foo@bar.com"
	for i := 0; i < b.N; i++ {
		email, gopher = isGopher(emailTest)
	}
}
