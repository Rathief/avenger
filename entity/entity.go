package entity

type Order struct {
	ID           int     `json:"id"`
	CustomerName string  `json:"customer_name"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	TotalPrice   float64 `json:"total_price"`
}

type OrderInput struct {
	CustomerName string  `json:"customer_name"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	TotalPrice   float64 `json:"total_price"`
}

type User struct {
	ID            uint    `json:"id"`
	Username      string  `json:"username" validate:"required"`
	Password      string  `json:"password" validate:"required"`
	DepositAmount float64 `json:"deposit_amount"`
}

type Store struct {
	ID         uint    `json:"id" gorm:""`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TotalSales int     `json:"total_sales"`
	Rating     int     `json:"rating"`
}

type Weather struct {
	CloudPct    int     `json:"cloud_pct"`
	Temp        int     `json:"temp"`
	FeelsLike   int     `json:"feels_like"`
	Humidity    int     `json:"humidity"`
	MinTemp     int     `json:"min_temp"`
	MaxTemp     int     `json:"max_temp"`
	WindSpeed   float64 `json:"wind_speed"`
	WindDegrees int     `json:"wind_degrees"`
	Sunrise     int64   `json:"sunrise"`
	Sunset      int64   `json:"sunset"`
}
