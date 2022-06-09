package setenv

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetEnv(t *testing.T) {
	dirName := ".hoge"
	err := SetEnv(dirName)
	if err != nil {
		t.Errorf("Error:Failed get current user%v", err)
	}

	user, _ := user.Current()
	dstDir := filepath.Join(user.HomeDir, dirName)
	key := os.Getenv("Path")
	if !strings.Contains(key, dstDir) {
		t.Errorf("Not set Env\n%s", key)
	}
	os.Remove(dstDir)
}

func TestRemoveEnv(t *testing.T) {
	dirName := ".hoge"
	SetEnv(dirName)
	err := RemoveEnv(dirName)
	if err != nil {
		t.Errorf("Error:Failed get current user%v", err)
	}

	user, _ := user.Current()
	dstDir := filepath.Join(user.HomeDir, dirName)
	key := os.Getenv("Path")
	if strings.Contains(key, dstDir) {
		t.Errorf("Not set Env\n%s", key)
	}

	os.Remove(dstDir)
}
