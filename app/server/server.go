package server

func GetConnectionString() string {
	dbName := "product_go"
	return "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}