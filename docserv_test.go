package docserv

import "testing"

func TestNewDocServ(t *testing.T) {
	d := NewDocServ([]string{"README.md"})
	if d.DocServConfig.Port != "9000" {
		t.Fail()
	}
	if d.DocServConfig.Docs[0] != "README.md" {
		t.Fail()
	}
}

func TestNewStaticDocServ(t *testing.T) {
	d := NewStaticDocServ([]string{"README.md"})
	if d.DocServConfig.Port != "9000" {
		t.Fail()
	}
	if d.DocServConfig.Docs[0] != "README.md" {
		t.Fail()
	}
	if d.DocServConfig.UseStatic != true {
		t.Fail()
	}
}