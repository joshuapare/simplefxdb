package storage

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type FSStorageEngine struct {
	// The parentmost directory holding all of the data for the database
	DataDir string
	// The name of the folder in a collection holding the index files
	IndexFolderName string
	// The name of the folder in a collection holding the data files
	DocumentFolderName string
	// Map of engine file mutexes for data read consistency
	MutexMap map[string]*sync.Mutex
}

// NewFSStorageEngine creates a new filesystem based storage engine
func NewFSStorageEngine() *FSStorageEngine {
	return &FSStorageEngine{
		DataDir:            "_data",
		IndexFolderName:    "idx",
		DocumentFolderName: "doc",
		MutexMap:           make(map[string]*sync.Mutex),
	}
}

// SetupFS ensures that the necessary directories are setup for the database
// engine storage
func SetupFS(e *FSStorageEngine) error {
	if !e.DirExists("") {
		err := os.Mkdir(e.DataDir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateStorageDir creates a new storage directory for documents
func (e *FSStorageEngine) CreateStorageDir(name string) error {
	dataLocation := e.pathInDataDir(name)
	dirsToCreate := []string{e.IndexFolderName, e.DocumentFolderName}

	// Make the parent directory
	err := os.Mkdir(dataLocation, 0755)
	if err != nil {
		return err
	}

	// Make the subdirs
	for _, subdir := range dirsToCreate {
		location := fmt.Sprintf("%s/%s", dataLocation, subdir)
		fmt.Printf("creating storage dir with name %s\n", location)
		err := os.Mkdir(location, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteToFS writes a new file to disk
func (e *FSStorageEngine) WriteToFS(path string, filename string, contents string, lock bool) error {
	if lock {
		file := fmt.Sprintf("%s/%s", path, filename)
		e.lockFile(file)
		defer e.unlockFile(file)
	}

	location := e.locationInDataDir(path, filename)
	if !e.DirExists(path) {
		return errors.New("directory does not exist")
	}

	f, err := os.Create(location)
	if err != nil {
		return errors.New("could not open file")
	}
	defer f.Close()

	f.Write([]byte(contents))

	return nil
}

// ReadFromFS reads the contents of a file on disk
func (e *FSStorageEngine) ReadFromFS(path string, filename string, lock bool) (string, error) {
	if lock {
		file := fmt.Sprintf("%s/%s", path, filename)
		e.lockFile(file)
		defer e.unlockFile(file)
	}

	contents, err := e.getFileContents(path, filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(*contents), nil
}

// ListDirectory lists the directories
func (e *FSStorageEngine) ListDirectories(path string) ([]string, error) {
	var dirNames []string

	files, err := os.ReadDir(e.pathInDataDir(path))
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
		}
	}

	return dirNames, nil
}

// ListDirectoryFiles lists the files in a directory
func (e *FSStorageEngine) ListDirectoryFiles(path string) ([]string, error) {
	var filePaths []string
	files, err := os.ReadDir(e.pathInDataDir(path))
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			break
		}
		filePaths = append(filePaths, file.Name())
	}

	return filePaths, nil
}

// getFileContents gets the contents of a file to a string
func (e *FSStorageEngine) getFileContents(path string, filename string) (*[]byte, error) {
	var contents []byte

	if e.FileExists(path, filename) {
		return &contents, errors.New("file already exists")
	}

	contents, err := os.ReadFile(e.locationInDataDir(path, filename))
	if err != nil {
		return &contents, err
	}

	return &contents, nil
}

// checkDirExists checks if a directory exists on the filesystem
func (e *FSStorageEngine) DirExists(path string) bool {
	info, err := os.Stat(e.pathInDataDir(path))
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// FileExists checks if a file exists on the filesystem
func (e *FSStorageEngine) FileExists(path string, filename string) bool {
	_, err := os.Stat(e.locationInDataDir(path, filename))
	return !os.IsNotExist(err)
}

// locationInDataDir Returns the location of the provided file and path within the data dir
func (e *FSStorageEngine) locationInDataDir(path string, filename string) string {
	return fmt.Sprintf("%s/%s", e.pathInDataDir(path), filename)
}

// pathInDataDir Returns the parent directory path of the provided path in the data dir
func (e *FSStorageEngine) pathInDataDir(path string) string {
	return fmt.Sprintf("%s/%s", e.DataDir, path)
}

// lockFile Locks a file from operations
func (e *FSStorageEngine) lockFile(filename string) {
	if _, ok := e.MutexMap[filename]; !ok {
		e.MutexMap[filename] = &sync.Mutex{}
	}
	e.MutexMap[filename].Lock()
}

// unlockFile unlocks a file for operations
func (e *FSStorageEngine) unlockFile(filename string) {
	e.MutexMap[filename].Unlock()
}
