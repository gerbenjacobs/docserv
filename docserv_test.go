package docserv

import "testing"

func TestNewDocServ(t *testing.T) {
	d := NewDocServ([]string{"README.md"})
	if d.Config.Port != "9000" {
		t.Fail()
	}
	if d.Config.Docs[0] != "README.md" {
		t.Fail()
	}
}
