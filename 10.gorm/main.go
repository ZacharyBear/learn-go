package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

const (
	user     = "go"
	pass     = "1234"
	database = "go"
	host     = "localhost"
	port     = 3306
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8&loc=Local", user, pass, host, port, database)
	sqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("An error occured when connect to the database")
	}
	// connect pool
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// testInsert(gormDb)

	// testInsertBatch(gormDb)

	testHooks(gormDb)
}

type User struct {
	ID       uint
	Name     string
	Email    string
	Birthday *time.Time
}

func (u User) TableName() string {
	return "user"
}

func testInsert(db *gorm.DB) {
	now := time.Now()
	user := User{
		Name:     "John Doe",
		Email:    "doe@zenkie.cn",
		Birthday: &now,
	}

	result := db.Create(&user)

	fmt.Println(user)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

func testInsertBatch(db *gorm.DB) {
	users := []User{
		{Name: "Zenkie Bear"},
		{Name: "Taylor Swift"},
		{Name: "John Cena"},
	}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}

// hooks
func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	emailRegex, err := regexp.Compile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	if !emailRegex.Match([]byte(u.Email)) {
		return errors.New("invalid email")
	}
	return
}

func testHooks(db *gorm.DB) {
	user := User{
		Name:  "Zenkie Bear",
		Email: "zq$zenkie.cn",
	}
	result := db.Create(&user)

	fmt.Println(result.Error)

	result = db.Session(&gorm.Session{
		SkipHooks: true,
	}).Create(&user)

	fmt.Print(user)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
