package obfs

import "testing"

func TestNewobfs(t *testing.T) {
	a, err := NewObfs("plain")
	a.SetInfo(SetHost("www.google.com"))
	t.Log(a, err, "xxx")
}
