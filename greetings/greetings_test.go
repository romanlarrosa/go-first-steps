package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name.
// Expected: msg matching want, no error
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string.
// Expected: empty string for msg, an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellosName calls greetings.Hello with a name.
// Expected: msg matching want, no error
func TestHellosName(t *testing.T) {
	names := []string{"Maria", "Ann"}
	msgs, err := Hellos(names)

	for name, msg := range msgs {
		want := regexp.MustCompile(`\b` + name + `\b`)
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hellos(["Maria", "Ann"]) = %q, %v, want match for format, nil`, msg, err)
		}
	}
}

// TestHellosEmpty calls greetings.Hello with at least an empty string.
// Expected: no map, an error
func TestHellosEmpty(t *testing.T) {
	names := []string{"Maria", ""}
	msg, err := Hellos(names)
	if msg != nil || err == nil {
		t.Fatalf(`Hellos(["...", ""]) = %q, %v, want nil, error`, msg, err)
	}
}
