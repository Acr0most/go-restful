package main

import (
	"github.com/Acr0most/go-restful/connector"
	"github.com/Acr0most/go-restful/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/*
CREATE TABLE IF NOT EXISTS `examples` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `name` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/

type Example struct {
	CommonModelFields
	Name string `json:"name"`
}

type CommonModelFields struct {
	ID        uint           `gorm:"primarykey;autoIncrement:true"`
	CreatedAt time.Time      `json:"created_at" time_format:"sql_datetime" time_location:"UTC"`
	UpdatedAt *time.Time     `json:"updated_at" time_format:"sql_datetime" time_location:"UTC"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" time_format:"sql_datetime" time_location:"UTC" gorm:"index"`
}

func main() {
	handle := handler.RestfulHandler{}

	connection := connector.NewGorm(connector.Config{MaxRetries: 10, IntervalMs: 1000})
	connection.Connect(mysql.Open("root:rootpw@tcp(127.0.0.1:6743)/example?charset=utf8mb4&parseTime=True&loc=Local"))

	handle.InitRouter(handler.Config{
		"example": handler.HandlerConfig{
			Handler: handler.ConnectorHandler{Connector: connection},
			Dummy: handler.Dummy{
				Single:   &Example{},
				Multiple: &[]Example{},
			},
		},
	}, 80)

	err := handle.Handle()

	if err != nil {
		panic(err)
	}
}
