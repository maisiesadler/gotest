package mbot

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var world = []byte("world")

var fileLoc = "/db/testing.db"

// var fileLoc = "/Users/maisiesadler/test/testing.db"

func PersistAdd(name string, value string) error {
	db, err := bolt.Open(fileLoc, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		key := []byte(name)
		val := []byte(value)

		err = bucket.Put(key, val)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func PersistGet(name string) (string, error) {
	db, err := bolt.Open(fileLoc, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var val string

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found!", world)
		}

		key := []byte(name)
		val = string(bucket.Get(key))
		fmt.Println(val)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return val, nil
}

func PersistKeyExists(name string) (bool, error) {
	db, err := bolt.Open("/Users/maisiesadler/test/bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found!", world)
		}

		key := []byte(name)
		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return false, nil
}
