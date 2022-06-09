package setenv

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func SetEnv(dirName string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}
	dstDir := filepath.Join(user.HomeDir, dirName)
	if _, err := os.Stat(dstDir); err != nil {
		os.Mkdir(dstDir, 0777)
	}

	key := os.Getenv("Path") + ";" + dstDir + ";"
	os.Setenv("Path", key)

	return nil
}

func RemoveEnv(dirName string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}
	dstDir := filepath.Join(user.HomeDir, dirName)

	key := strings.Replace(os.Getenv("Path"), dstDir+";", "", -1)
	os.Setenv("Path", key)

	return nil
}
