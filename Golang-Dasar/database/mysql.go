package database

var connection string


// ini auto dijalankan ketika package di import
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}