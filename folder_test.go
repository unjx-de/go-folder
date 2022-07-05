package folder

import (
	"io/fs"
	"os"
	"testing"
)

var validFolders = []string{"exampleFolder", "anotherFolder"}
var invalidFolders = []string{".", ""}
var permission fs.FileMode = 0755

func TestCreateAndRemoveValidFolders(t *testing.T) {
	err := CreateFolders(validFolders, permission)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	err = RemoveFolders(validFolders)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestCreateAlreadyExistingFolders(t *testing.T) {
	err := CreateFolders(validFolders, permission)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	err = CreateFolders(validFolders, permission)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	err = RemoveFolders(validFolders)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestCreateInvalidFolders(t *testing.T) {
	_, err := os.Create(validFolders[0])
	if err != nil {
		t.Errorf("%v\n", err)
	}
	err = CreateFolders(validFolders, permission)
	if err == nil {
		t.Errorf("expected Error\n")
	}
	err = os.Remove(validFolders[0])
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestRemoveInvalidFolders(t *testing.T) {
	err := RemoveFolders(invalidFolders)
	if err == nil {
		t.Errorf("expected Error\n")
	}
}
