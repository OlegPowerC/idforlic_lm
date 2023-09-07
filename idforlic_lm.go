package idforlic_lm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

const (
	linuxMfile  = `/var/lib/dbus/machine-id`
	linuxRHfile = `/etc/machine-id`
)

func linuxGetID() (linuxId string, linuxIdErr error) {
	filecheck := linuxMfile
	_, fileExistError := os.Stat(filecheck)
	if os.IsNotExist(fileExistError) {
		filecheck = linuxRHfile
		_, fileExistError = os.Stat(filecheck)
		if os.IsNotExist(fileExistError) {
			return "", errors.New("Can not get ID on this OS")
		}
	}

	LinuxIDFile, LinuxIDFileErr := os.Open(filecheck)
	if LinuxIDFileErr != nil {
		return "", LinuxIDFileErr
	}
	defer LinuxIDFile.Close()
	IDb, idReadErr := ioutil.ReadAll(LinuxIDFile)
	ID := string(IDb)
	ID = strings.Replace(ID, "\n", "", -1)
	ID = strings.Replace(ID, "\r", "", -1)
	if idReadErr != nil {
		return "", idReadErr
	}
	return ID, nil
}

func GetID() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return linuxGetID()
		break
	}
	return "", fmt.Errorf("Unknown OS: %s", runtime.GOOS)
}
