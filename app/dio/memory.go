package dio

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var memory *gorm.DB

func NewMemory() (*gorm.DB, error) {
	// FIXME: use env variable if go public repo
	var err error
	config := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", "34.87.91.33", "5432", "postgres", "dio_memory", "[vdw,j5^d")

	memory, err = gorm.Open("postgres", config)
	if err != nil {
		return nil, err
	}
	memory.DB().SetMaxIdleConns(10)
	memory.LogMode(true)

	return memory, nil
}

func GetMemory() *gorm.DB {
	return memory
}
