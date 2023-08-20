package engine

import (
	"fmt"
	"sync"

	"joshuapare.com/fxdb/internal/storage"
)

type IndexType string

type IndexEngine struct {
	// Storage engine to use for the collection engine
	SE *storage.FSStorageEngine
	// In-memory Index Store
	IndexStore map[string]*Index
}

const (
	Unique IndexType = "UNIQUE"
)

type Index struct {
	// Type determines the type of index
	Type IndexType
	// Name is the name of the index
	Name string
	// Mu is the mutex for the Index
	Mu sync.Mutex
	// Store is the store that stores the index
	Store BTree
}

func NewIndexEngine(e *storage.FSStorageEngine) *IndexEngine {
	fmt.Println("setting up index engine")
	return &IndexEngine{
		SE:         e,
		IndexStore: make(map[string]*Index),
	}
}

// LoadIndexes loads indexes from the filesystem into memory
func (i *IndexEngine) LoadIndexes(collection string) error {
	// Loop through each index, and load the indexes into memory
	fmt.Println("loading indexes...")
	location := fmt.Sprintf("%s/%s", collection, i.SE.IndexFolderName)
	indexFiles, err := i.SE.ListDirectoryFiles(location)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range indexFiles {
		i.LoadIndex(fmt.Sprintf("%s/%s", location, file))
	}
	return nil
}

func (i *IndexEngine) LoadIndex(location string) {
	fmt.Printf("loading index at location %s into memory\n", location)
}

// TODO - Creates a new index with a name
func (i *IndexEngine) CreateIndex(itype IndexType, name string) error {
	fmt.Println("setting up %s type index '%s'\n", itype, name)
	return nil
}

// TODO - Deletes an existing index
func (i *IndexEngine) DeleteIndex(name string) error {
	fmt.Println("setting index with name '%s'\n", name)
	return nil
}
