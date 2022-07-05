package folder

import (
	"io/fs"
	"os"
)

func createFolder(folderRelativePath string, permission fs.FileMode) error {
	err := os.MkdirAll(folderRelativePath, permission)
	if err != nil {
		return err
	}
	return nil
}

func removeFolder(folderRelativePath string) error {
	err := os.RemoveAll(folderRelativePath)
	if err != nil {
		return err
	}
	return nil
}

// CreateFolders creates all folder passed as string array
//
// If path is already a directory, CreateFolders returns nil (no error).
// If there is an error, it will be of type *fs.PathError.
func CreateFolders(neededFolderRelPath []string, permission fs.FileMode) error {
	for _, folderRelPath := range neededFolderRelPath {
		err := createFolder(folderRelPath, permission)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveFolders removes all folder passed as string array
//
// If the path does not exist, RemoveFolders returns nil (no error).
// If there is an error, it will be of type *fs.PathError.
func RemoveFolders(neededFolderRelPath []string) error {
	for _, folderRelPath := range neededFolderRelPath {
		err := removeFolder(folderRelPath)
		if err != nil {
			return err
		}
	}
	return nil
}
