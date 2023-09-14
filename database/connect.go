package database

import (
	"github.com/dilson-m/golang-fiber-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	v := "Nao conseguiu conectar ao banco de dados !!!"
	//		user  pass   Nome BD
	dsn := "admin:admin@/fluent_admin?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(v)
	}

	// fmt.Println(db)
	// fmt.Printf("Conexao no DB com sucesso... %v\n", db)

	connection.AutoMigrate(models.User{})

}
