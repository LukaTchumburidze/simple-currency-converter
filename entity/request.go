package entity

type Request struct {
	ID      uint   `gorm:"primaryKey"`
	Headers string `gorm:"size:2048;not null"`
	Body    string `gorm:"size:2048;not null"`
}
