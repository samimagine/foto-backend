package models

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"size:100;not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role     string `gorm:"size:10;not null"` // "admin", "foto", or "tourist"
    Bio      string `gorm:"size:255"`
    Social   string `gorm:"size:255"`
    Price    float64
	Latitude  float64 // User's latitude
	Longitude float64 // User's longitude
}

type Rating struct {
    ID        uint   `gorm:"primaryKey"`
    Comment   string `gorm:"size:255;not null"`
    Rating    int    `gorm:"not null"` // e.g., 1-5 stars
    TouristID uint   `gorm:"not null"`
    FotoID    uint   `gorm:"not null"`
}