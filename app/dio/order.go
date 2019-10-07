package dio

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FoodOrder struct {
	ID         string `gorm:"primary_key"`
	Menu       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleatedAt *time.Time
}

type FoodOrderRepository interface {
	Upsert(FoodOrder) (FoodOrder, error)
	FindByID(id string) (FoodOrder, error)
	DeleteByID(id string) error
}

type FoodOrderPersistance struct {
	db *gorm.DB
}

func (p FoodOrderPersistance) Upsert(order FoodOrder) (FoodOrder, error) {
	var result FoodOrder

	dbRes := p.db.Raw(upsertFoodOrderSQL, order.ID, order.Menu, order.CreatedAt, order.UpdatedAt).Scan(&result)
	if dbRes.Error != nil {
		return FoodOrder{}, dbRes.Error
	}

	return result, nil
}

func (p FoodOrderPersistance) FindByID(id string) (FoodOrder, error) {
	var result FoodOrder

	dbRes := p.db.Raw(getFoodOrderSQL, id).Scan(&result)
	if dbRes.Error != nil {
		return FoodOrder{}, dbRes.Error
	}

	return result, nil
}

func (p FoodOrderPersistance) DeleteByID(id string) error {
	dbRes := p.db.Raw(deleteFoodOrderSQL, id)
	if dbRes.Error != nil {
		return dbRes.Error
	}
	return nil
}

var upsertFoodOrderSQL =
`insert into food_order (id, menu, created_at, updated_at)
values ($1, $2, $3, $4)
on conflict (id)
do
	update
	set menu= EXCLUDED.menu
returning *;`

var getFoodOrderSQL =
`select * from food_order
    where id = $1;`

var deleteFoodOrderSQL =
`delete from food_order
	where food_order.id = $1;`
