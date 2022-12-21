package domain

type Example struct {
	ID   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}
