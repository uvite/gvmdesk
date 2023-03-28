// Copyright 2021 rule101. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This file contains DB related functions
	It takes care of initializing the db as well as querying and processing DB entries
*/

package surge

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
	"time"

	"log"

	"github.com/nutsdb/nutsdb"
	"github.com/uvite/gvmdesk/server/models"
	"github.com/uvite/gvmdesk/server/platform"
)

const alertBucker = "alertsBucket"
const coinBucket = "coinBucket"

var db *nutsdb.DB

type FileFilterState int

const (
	All = iota
	Downloading
	Seeding
	Completed
	Paused
)

//InitializeDb initializes db
func InitializeDb() {
	var err error
	opt := nutsdb.DefaultOptions

	opt.Dir = platform.GetSurgeDir() + string(os.PathSeparator) + "db"
	db, err = nutsdb.Open(opt)
	if err != nil {
		log.Panic(err)
	}
}

//CloseDb .
func CloseDb() {
	db.Close()
}

// Gets all Files in the DB
func dbGetAllFiles() []models.File {
	files := []models.File{}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.GetAll(alertBucker)
			if err != nil {
				return err
			}

			for _, entry := range entries {

				newFile := &models.File{}
				json.Unmarshal(entry.Value, newFile)
				files = append(files, *newFile)
			}

			return nil
		}); err != nil {
		log.Println("Get all db files error:", err)
	} else {
		return files
	}
	return files
}

// Gets a File by providing the fileHash
func dbGetFile(Hash string) (*models.File, error) {
	result := &models.File{}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			fileKey := []byte(Hash)
			e, err := tx.Get(alertBucker, fileKey)
			if err != nil {
				return err
			}

			err = json.Unmarshal(e.Value, result)
			return err
		}); err != nil {
		return nil, err
	}

	return result, nil
}

// Inserts a File to the DB
func DbInsertFile(File models.File) {

	if File.DateTimeAdded == 0 {
		File.DateTimeAdded = time.Now().Unix()
	}

	if err := db.Update(
		func(tx *nutsdb.Tx) error {

			fileKey := []byte(File.FileHash)
			fileBytes, _ := json.Marshal(File)

			if err := tx.Put(alertBucker, fileKey, fileBytes, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Panic(err)
	}
}

// Deletes a File by providing the fileHash
func dbDeleteFile(Hash string) error {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(Hash)
			if err := tx.Delete(alertBucker, key); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Println("Db delete file", err)
		return err
	}
	return nil
}

//DbWriteSetting Stores or updates a key with a given value
func DbWriteSetting(Name string, value string) error {
	err := db.Update(
		func(tx *nutsdb.Tx) error {

			keyBytes := []byte(Name)
			valueBytes := []byte(value)

			if err := tx.Put(coinBucket, keyBytes, valueBytes, 0); err != nil {
				return err
			}
			return nil
		})
	return err
}

//DbReadSetting Reads a key and returns value
func DbReadSetting(Name string) (string, error) {
	result := ""
	key := []byte(Name)

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			bytes, err := tx.Get(coinBucket, key)
			if err != nil {
				return err
			}

			result = string(bytes.Value)

			return err
		}); err != nil {
		return "", err
	}

	return result, nil
}

//PagedQueryResult is a paging query result for file searches
type PagedQueryResult struct {
	Result []models.FileTransfer
	Count  int
}

//PagedQueryResult is a paging query result for file searches
type PagedQueryRemoteResult struct {
	Result []models.FileListing
	Count  int
}


//SearchLocalFile runs a paged query
func SearchLocalFile(Query string, filterState FileFilterState, OrderBy string, IsDesc bool, Skip int, Take int) PagedQueryResult {

	resultFiles := []models.File{}

	allFiles := dbGetAllFiles()
	for _, file := range allFiles {
		if strings.Contains(strings.ToLower(file.FileName), strings.ToLower(Query)) || strings.Contains(strings.ToLower(file.FileHash), strings.ToLower(Query)) {
			resultFiles = append(resultFiles, file)
		}
	}

	fileFilterFunc := func(f models.File) bool { return true }

	//Filter files on filter state
	switch filterState {
	/*case All: //added for clarity
	fileFilterFunc = func(f models.File) bool { return true }ec
	break*/
	case Downloading:
		fileFilterFunc = func(f models.File) bool { return f.IsDownloading }
	case Seeding:
		fileFilterFunc = func(f models.File) bool { return f.IsUploading }
	case Completed:
		fileFilterFunc = func(f models.File) bool { return f.IsAvailable }
	case Paused:
		fileFilterFunc = func(f models.File) bool { return f.IsPaused }
	}

	resultFiles = filterFile(resultFiles, fileFilterFunc)

	totalNum := len(resultFiles)

	switch OrderBy {
	default:
		if !IsDesc {
			sort.Sort(sortByFileNameAsc(resultFiles))
		} else {
			sort.Sort(sortByFileNameDesc(resultFiles))
		}
	}

	left := Skip
	right := Skip + Take

	if left > len(resultFiles) {
		left = len(resultFiles)
	}

	if right > len(resultFiles) {
		right = len(resultFiles)
	}

	//Subset
	resultFiles = resultFiles[left:right]
	resultListings := []models.FileTransfer{}

	for i := 0; i < len(resultFiles); i++ {
		listing := models.FileTransfer{
			FileHash:      resultFiles[i].FileHash,
			FileName:      resultFiles[i].FileName,
			FileSize:      resultFiles[i].FileSize,
			IsDownloading: resultFiles[i].IsDownloading,
			IsHashing:     resultFiles[i].IsHashing,
			IsMissing:     resultFiles[i].IsMissing,
			IsPaused:      resultFiles[i].IsPaused,
			IsUploading:   resultFiles[i].IsUploading,
			Topic:         resultFiles[i].Topic,
			NumSeeders:    len(GetSeeders(resultFiles[i].FileHash)),
		}

		//If file is downloading set progress
		if listing.IsDownloading || listing.IsPaused {
			//numChunksLocal := chunksDownloaded(resultFiles[i].ChunkMap, resultFiles[i].NumChunks)
			//listing.Progress = float32(float64(numChunksLocal) / float64(resultFiles[i].NumChunks))
			//numChunksLocal := chunksDownloaded(resultFiles[i].ChunkMap, resultFiles[i].NumChunks)
			listing.Progress = 1.0
		} else {
			listing.Progress = 1.0
		}

		resultListings = append(resultListings, listing)

	}

	return PagedQueryResult{
		Result: resultListings,
		Count:  totalNum,
	}
}
