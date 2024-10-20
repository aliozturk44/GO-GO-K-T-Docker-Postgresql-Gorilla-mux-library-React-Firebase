package config

import (
	"fmt"
	"source/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

// "host=localhost user=postgres password=depixen-pass dbname=postgres port=5439 sslmode=disable"
//const dbinfo = "postgres://postgres:depixen-pass@localhost/postgres?sslmode=disable"

const dbinfo = "host=localhost user=postgres password=depixen-pass dbname=postgres port=5439 sslmode=disable"

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	DB.AutoMigrate(&model.Tb_casestudy{})
	//DB.Exec("CREATE TABLE IF NOT EXISTS tb_casestudy (id SERIAL PRIMARY KEY NOT NULL, title TEXT NOT NULL, description TEXT NOT NULL, imageuri TEXT NOT NULL, createddate TEXT NOT NULL)")

}
