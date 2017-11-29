package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Key struct {
	gorm.Model
	Value string
}

// Insert add a new element in the database
func Insert(db *gorm.DB, k *Key) error {
	if err := db.Create(k).Error; err != nil {
		return err
	}
	return nil
}

//ListAll returns all element in database
func ListAll(db *gorm.DB) ([]Key, error) {
	var ret []Key
	cursor := db.Find(&ret)
	if cursor.RecordNotFound() {
		return nil, fmt.Errorf("RecordNotFound: Nothing found in database")
	}
	return ret, nil
}

// InitDatabase create table scheme and an handler to the database
func InitDatabase(dbDsn string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Key{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
