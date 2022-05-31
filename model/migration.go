package model

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid", "user(id)", "CASCADE", "CASCADE")
}
