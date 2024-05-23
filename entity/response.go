package entity

type Response struct {
	ID        uint   `gorm:"primaryKey"`
	RequestID uint   `gorm:"not null"`
	Status    int    `gorm:"not null"`
	Headers   string `gorm:"size:2048;not null"`
	Body      string `gorm:"size:2048;not null"`
}
