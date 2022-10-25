package idforlic_lm

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"
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
	ID, idReadErr := ioutil.ReadAll(LinuxIDFile)
	if idReadErr != nil {
		return "", idReadErr
	}
	return string(ID), nil
}

func GetID() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return linuxGetID()
		break
	}
	return "", errors.New("Unknown OS")
}
