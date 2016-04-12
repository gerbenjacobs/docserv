package docserv

import (
	"testing"
)

func TestParseAsIdentifierSimple(t *testing.T) {
	r := parseAsIdentifier("R3ADME.md")
	if r != "i-r3admemd" {
		t.Fail()
	}
}

func TestParseAsIdentifierRoot(t *testing.T) {
	r := parseAsIdentifier("/R3ADME.md")
	if r != "i-r3admemd" {
		t.Fail()
	}
}

func TestParseAsIdentifierNumbers(t *testing.T) {
	r := parseAsIdentifier("2016-04-12-README.md")
	if r != "i-20160412readmemd" {
		t.Fail()
	}
}

func TestParseAsIdentifierDirectory(t *testing.T) {
	r := parseAsIdentifier("doc/R3ADME.md")
	if r != "i-docr3admemd" {
		t.Fail()
	}
}

func TestParseAsIdentifierRelative(t *testing.T) {
	r := parseAsIdentifier("../R3ADME.md")
	if r != "i-r3admemd" {
		t.Fail()
	}
}

func TestParseAsIdentifierCraziness(t *testing.T) {
	r := parseAsIdentifier("R3AD?ME! Later...md")
	if r != "i-r3admelatermd" {
		t.Fail()
	}
}
