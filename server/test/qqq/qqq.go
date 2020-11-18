package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CasbinRule struct {
	TablePrefix string `gorm:"-"`
	PType       string `gorm:"size:100"`
	V0          string `gorm:"size:100"`
	V1          string `gorm:"size:100"`
	V2          string `gorm:"size:100"`
	V3          string `gorm:"size:100"`
	V4          string `gorm:"size:100"`
	V5          string `gorm:"size:100"`
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=qmplus sslmode=disable password=123456")
	if err != nil {
		panic(err)
	}
	var lines []CasbinRule
	if err := db.Table("casbin_rule").Find(&lines).Error; err != nil {
		panic(err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
	defer db.Close()

}
