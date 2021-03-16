package connector

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
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
		t.DB, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: 500 * time.Millisecond,
				LogLevel:      logger.Warn,
				Colorful:      true,
			}),
		})

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

func (t GormConnector) Find(params map[string]interface{}, result interface{}) (err error) {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		var info *gorm.DB

		if len(params) > 0 {
			joins := make(map[string]interface{}, len(params))

			for key := range params {
				if strings.Contains(key, ".") {
					joins[strings.Split(key, ".")[0]] = 1
				}
			}

			for asso := range joins {
				tx = tx.Joins(asso)
			}

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
}

func (t GormConnector) Create(items interface{}) (err error) {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		var info *gorm.DB

		info = tx.Create(items)

		if info.Error != nil {
			return info.Error
		}

		info = tx.Find(items)

		if info.Error != nil {
			return info.Error
		}

		return nil
	})
}

func (t GormConnector) Delete(items interface{}, result interface{}) (err error) {
	return t.DB.Where(items).Delete(result).Error
}

func (t GormConnector) Patch(items interface{}, result interface{}, model interface{}) (err error) {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		var info *gorm.DB

		info = tx.Where(items).Find(model)

		if info.Error != nil {
			return info.Error
		}

		if info.RowsAffected == 0 {
			return errors.New("no row affected")
		}

		info = tx.Updates(result)

		if info.Error != nil {
			return info.Error
		}

		return nil
	})
}
