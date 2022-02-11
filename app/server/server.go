package server

const dbName = "product_go"

func GetConnectionString() string {
	return "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
