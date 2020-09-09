package connector

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type GormConnector struct {
	DB     *gorm.DB
	Config Config
}

func NewGorm(config Config) *GormConnector {
	config.Compact()

	return &GormConnector{
		Config: config,
	}
}

func (t *GormConnector) Connect(dialector gorm.Dialector) {
	var (
		err   error
		count = 0
	)

	for t.DB == nil {
		t.DB, err = gorm.Open(dialector, &gorm.Config{})

		if err != nil {
			count++

			if count == t.Config.MaxRetries {
				panic(err)
			}
		}

		time.Sleep(t.Config.IntervalMs * time.Millisecond)
	}

	log.Println("GORM connected..")
}

func (t GormConnector) Find(params map[string]interface{}, result interface{}) (success bool) {
	var info *gorm.DB

	if len(params) > 0 {
		info = t.DB.Where(params).Find(result)
	} else {
		info = t.DB.Find(result)
	}

	if info.Error != nil {
		panic(info.Error)
	}

	if info.RowsAffected == 0 {
		return false
	}

	return true
}

func (t GormConnector) Create(items interface{}) {
	info := t.DB.Create(items)

	if info.Error != nil {
		panic(info.Error)
	}
}

func (t GormConnector) Delete(items interface{}, result interface{}) {
	info := t.DB.Where(items).Delete(result)

	if info.Error != nil {
		panic(info.Error)
	}
}
