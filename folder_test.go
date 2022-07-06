package folder

import (
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

var validFolders = []string{"exampleFolder", "anotherFolder"}
var invalidFolders = []string{".", ""}
var permission fs.FileMode = 0755

func TestCreateAndRemoveValidFolders(t *testing.T) {
	err := CreateFolders(validFolders, permission)
	assert.Equal(t, nil, err, "expected no error")
	err = RemoveFolders(validFolders)
	assert.Equal(t, nil, err, "expected no error")
}

func TestCreateAlreadyExistingFolders(t *testing.T) {
	err := CreateFolders(validFolders, permission)
	assert.Equal(t, nil, err, "expected no error")
	err = CreateFolders(validFolders, permission)
	assert.Equal(t, nil, err, "expected no error")
	err = RemoveFolders(validFolders)
	assert.Equal(t, nil, err, "expected no error")
}

func TestCreateInvalidFolders(t *testing.T) {
	_, err := os.Create(validFolders[0])
	assert.Equal(t, nil, err, "expected no error")
	err = CreateFolders(validFolders, permission)
	assert.NotEqual(t, nil, err, "expected error")
	err = os.Remove(validFolders[0])
	assert.Equal(t, nil, err, "expected no error")
}

func TestRemoveInvalidFolders(t *testing.T) {
	err := RemoveFolders(invalidFolders)
	assert.NotEqual(t, nil, err, "expected error")
}
