package hashutil

import (
	"github.com/Tom7481079/hashutil/pkg"
	"testing"
)

func TestMD5Hash(t *testing.T) {
	result := hashutil.MD5Hash("hello")
	expected := "5d41402abc4b2a76b9719d911017c592"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSHA256Hash(t *testing.T) {
	result := hashutil.SHA256Hash("hello")
	expected := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
