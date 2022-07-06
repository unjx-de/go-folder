package folder

import (
	"github.com/go-playground/assert/v2"
	"io/fs"
	"os"
	"testing"
)

var validFolders = []string{"exampleFolder", "anotherFolder"}
var invalidFolders = []string{".", ""}
var permission fs.FileMode = 0755

func TestCreateAndRemoveValidFolders(t *testing.T) {
	err := CreateFolders(validFolders, permission)
	assert.Equal(t, nil, err)
	err = RemoveFolders(validFolders)
	assert.Equal(t, nil, err)
}

func TestCreateAlreadyExistingFolders(t *testing.T) {
	err := CreateFolders(validFolders, permission)
	assert.Equal(t, nil, err)
	err = CreateFolders(validFolders, permission)
	assert.Equal(t, nil, err)
	err = RemoveFolders(validFolders)
	assert.Equal(t, nil, err)
}

func TestCreateInvalidFolders(t *testing.T) {
	_, err := os.Create(validFolders[0])
	assert.Equal(t, nil, err)
	err = CreateFolders(validFolders, permission)
	assert.NotEqual(t, nil, err)
	err = os.Remove(validFolders[0])
	assert.Equal(t, nil, err)
}

func TestRemoveInvalidFolders(t *testing.T) {
	err := RemoveFolders(invalidFolders)
	assert.NotEqual(t, nil, err)
}
