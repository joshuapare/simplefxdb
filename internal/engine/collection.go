package engine

import (
	"fmt"

	"joshuapare.com/fxdb/internal/storage"
)

type CollectionEngine struct {
	// Storage engine to use for the collection engine
	SE *storage.FSStorageEngine
}

func NewCollectionEngine(e *storage.FSStorageEngine) *CollectionEngine {
	fmt.Println("setting up collections engine")
	return &CollectionEngine{
		SE: e,
	}
}

func (c *CollectionEngine) CreateCollection(name string) error {
	if c.collectionExists(name) {
		return fmt.Errorf("collection with name %s already exists", name)
	}

	fmt.Printf("creating collection with name %s\n", name)
	if err := c.SE.CreateStorageDir(name); err != nil {
		return err
	}
	return nil
}

func (c *CollectionEngine) DeleteCollection(name string) error {
	fmt.Printf("removing collection with name %s\n", name)
	return nil
}

func (c *CollectionEngine) GetCollections() []string {
	fmt.Printf("listing collections")
	collections, err := c.SE.ListDirectories("")
	if err != nil {
		fmt.Printf("Error listing collections")
	}
	return collections
}

func (c *CollectionEngine) collectionExists(name string) bool {
	return c.SE.DirExists(name)
}
