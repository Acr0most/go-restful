package connector

import (
	"errors"
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
	err := t.DB.Transaction(func(tx *gorm.DB) error {
		var info *gorm.DB

		if len(params) > 0 {
			if info = tx.Where(params).Find(result); info.Error != nil {
				return info.Error
			}
		} else {
			if info = tx.Find(result); info.Error != nil {
				return info.Error
			}
		}

		if info.RowsAffected == 0 {
			return errors.New("CUSTOM:5 no rows effected")
		}

		return nil
	})

	if err != nil {
		log.Panic("err", err)
		return false
	}

	return true
}

func (t GormConnector) Create(items interface{}) {
	err := t.DB.Transaction(func(tx *gorm.DB) error {
		var info *gorm.DB

		info = tx.Create(items)

		if info.Error != nil {
			log.Printf("error %#v", info.Error)
			return info.Error
		}

		info = info.Find(items)

		if info.Error != nil {
			return info.Error
		}

		return nil
	})

	if err != nil {
		log.Panic("err", err)
	}
}

func (t GormConnector) Delete(items interface{}, result interface{}) {
	info := t.DB.Where(items).Delete(result)

	if info.Error != nil {
		log.Panic("err", info.Error)
	}
}

func (t GormConnector) Patch(items interface{}, result interface{}, model interface{}) {
	err := t.DB.Transaction(func(tx *gorm.DB) error {
		var info *gorm.DB

		info = tx.Where(items).Find(model)

		if info.Error != nil {
			return info.Error
		}

		if info.RowsAffected == 0 {
			return errors.New("no row affected")
		}

		info = info.Updates(result)

		if info.Error != nil {
			return info.Error
		}

		return nil
	})

	if err != nil {
		log.Panic("err", err)
	}

}
