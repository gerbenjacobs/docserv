package docserv

import "testing"

func TestNewDocServ(t *testing.T) {
	d := NewDocServ([]string{"README.md"})
	if d.DocServConfig.Port != "9000" {
		t.Fail()
	}
	if _, ok := d.DocServConfig.Docs["README.md"]; !ok {
		t.Fail()
	}
}

func TestNewStaticDocServ(t *testing.T) {
	d := NewStaticDocServ(map[string][]byte{"README.md": []byte("A Readme")})
	if d.DocServConfig.Port != "9000" {
		t.Fail()
	}
	if val, ok := d.DocServConfig.Docs["README.md"]; ok {
		if string(val) != "A Readme" {
			t.Fail()
		}
	}
	if _, ok := d.DocServConfig.Docs["README.md"]; !ok {
		t.Fail()
	}
	if d.DocServConfig.UseStatic != true {
		t.Fail()
	}
}
