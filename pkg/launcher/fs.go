package launcher

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// DefaultTokenFile is deprecated, and will be only used for migration.
const DefaultTokenFile = "credentials"

// DefaultConfigsFile stores cli credentials and hosts.
const DefaultConfigsFile = "configs"

// InfluxDir retrieves the influxdbv2 directory.
func InfluxDir() (string, error) {
	var dir string
	// By default, store meta and data files in current users home directory
	u, err := user.Current()
	if err == nil {
		dir = u.HomeDir
	} else if home := os.Getenv("HOME"); home != "" {
		dir = home
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		dir = wd
	}
	dir = filepath.Join(dir, ".gvmdesk")

	return dir, nil
}

// BoltFile returns the path to the bolt file for influxdb
func BoltFile() (string, error) {
	dir, err := InfluxDir()
	if err != nil {
		return "", err
	}
	var file string
	filepath.WalkDir(dir, func(p string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if file != "" {
			return fmt.Errorf("bolt file found")
		}

		if strings.Contains(p, ".bolt") {
			file = p
		}

		return nil
	})

	if file == "" {
		return "", fmt.Errorf("bolt file not found")
	}

	return file, nil
}
