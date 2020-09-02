package main

import (
    "testing"
    "regexp"
)

func TestGreeting(t *testing.T) {
    message     := "Hello"
    boldMessage := greeting("Hello")
    if ( len( message ) == 0 ) {
        t.Error("message empty")
    }
    match, _ := regexp.MatchString("<b>"+message+"</b>", boldMessage)
    if (!match) {
        t.Error("greeting is not bolded")
    }
}