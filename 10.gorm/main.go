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

	// insert
	// testInsert(gormDb)

	// testInsertBatch(gormDb)

	// testHooks(gormDb)

	// select
	// testSelect(gormDb)
	// testSelectWithConditions(gormDb)
	// testSelectAdvanced(gormDb)

	// update
	// testSave(gormDb)
	// testUpdate(gormDb)

	// delete
	// testDelete(gormDb)
	// testDeleteByPrimaryKey(gormDb)
	testSoftDelete(gormDb)
}

type User struct {
	ID       uint `gorm:"primaryKey"`
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

func testSelect(db *gorm.DB) {
	user := User{}
	db.First(&user) // select * from user order by id limit 1;
	fmt.Println(user)

	user = User{}
	db.Take(&user) // select * from user limit 1;
	fmt.Println(user)

	user = User{}
	result := db.Last(&user) // select * from user order by id desc limit 1;
	fmt.Println(user)
	fmt.Println(result.RowsAffected) // 1
	fmt.Println(result.Error)        // nil

	user = User{ID: 9999999}
	result = db.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Can’t found the data")
	}

	user = User{}
	db.First(&user, 10) // select * from user where id = 10
	fmt.Println(user)

	user = User{}
	db.First(&user, "10") // select * from user where id = 10
	fmt.Println(user)

	user = User{}
	users := []User{}
	db.Find(&users, []int{1, 2, 3}) // select * from user where id in (1, 2, 3)
	fmt.Println(users)

	user = User{ID: 10}
	db.Find(&user) // select * from user where id = 10 limit 1;
	fmt.Println(user)

	// todo this sample might be incorrect
	user = User{}
	db.Debug().Model(User{ID: 10}).First(&user)
	fmt.Println(user)

	result = db.Find(&users)
	fmt.Println(result.RowsAffected, result.Error) // users,
}

func testSelectWithConditions(db *gorm.DB) {
	user := User{}
	// String condition
	db.Where("name like ?", "%Zenkie%").First(&user)
	fmt.Println(user)

	// Struct condition
	var users []User
	db.Where(&User{Name: "John Doe"}).Find(&users)
	fmt.Println(users)

	// Slice condition
	db.Where([]int{1, 2, 3}).Find(&users) // select .... where id in (1, 2, 3)
	fmt.Println(users)

	// Specify the query fields
	user = User{}
	db.Debug().Where(&User{ID: 5, Name: "John Doe"}, "id").Find(&user)
	// the name condition will be ignored
	fmt.Println(user)

	// Inline condition
	user = User{}
	db.Debug().First(&user, "id > ?", 5)
	fmt.Println(user)

	// Select fields
	user = User{}
	db.Select("name", "email").First(&user)
	fmt.Println(user)

	// Scan
	db.Debug().Where("id = ?", 1).Scan(&user)
	fmt.Println(user)
}

func testSelectAdvanced(db *gorm.DB) {
	// imagine there're hundreds fields in User struct
	type APIUser struct {
		ID   int
		Name string
	}
	var apiUsers []APIUser
	db.Debug().Model(&User{}).Limit(10).Find(&apiUsers)
	fmt.Println(apiUsers)

	// Subquery
	var users []User
	db.Debug().Table("(?) as u", db.Model(&User{}).
		Select("name", "birthday")).
		Where("birthday is not null").
		Find(&users)
	fmt.Println(users)
}

func testSave(db *gorm.DB) {
	var user User
	db.First(&user)

	user.Name = "Max Black"
	maxsBirthday, _ := time.Parse("2006-01-02", "1987-06-09")
	fmt.Println(maxsBirthday)
	user.Birthday = &maxsBirthday
	db.Save(&user)
}

func testUpdate(db *gorm.DB) {
	// Update single column
	db.Model(&User{}).Where("id = ?", 1).Update("name", "Caroline Channing")
	// UPDATE users SET name='hello' WHERE id=1;

	var user User
	db.First(&user)
	db.Model(&user).Update("name", "hello")
	fmt.Println(user)

	db.Debug().Model(&user).Update("name", "hola")

	// Update multiple columns
	carolinesBirth, _ := time.Parse("2006-01-02", "1987-05-28")
	db.Debug().Model(&user).Updates(User{Name: "Caroline Channing", Birthday: &carolinesBirth})
}

type Email struct {
	ID       int
	Name     string
	Address  string
	DeleteAt gorm.DeletedAt
}

func (e Email) TableName() string {
	return "email"
}
func testDelete(db *gorm.DB) {
	email := Email{
		ID: 1,
	}
	db.First(&email)
	fmt.Println(email)
	db.Debug().Delete(&email)
}
func testDeleteByPrimaryKey(db *gorm.DB) {
	db.Debug().Delete(&User{}, 5)
}
func testSoftDelete(db *gorm.DB) {
	db.Debug().Delete(&Email{ID: 1})
	var email Email
	db.First(&email, 1)
	fmt.Println(email)
}
