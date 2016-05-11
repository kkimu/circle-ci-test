package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"
)

var db *gorm.DB

func init() {
	ip := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	var err error
	if ip != "" {
		db, err = gorm.Open("mysql", "root:mysql@tcp("+ip+":3306)/airmeet?parseTime=True&loc=Japan")
	} else {
		db, err = gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/airmeet?parseTime=True&loc=Japan")
	}

	if err != nil {
		panic("failed to connect database")
	}

	db.DB()
}

// CreateEvent イベントを作成
func CreateEvent(event *Event) {
	//db.NewRecord(event)
	db.Create(&event)
	//db.Save(&event)
}

// GenerateMajor ユニークなmajorを生成し返す
func GenerateMajor() int {
	rand.Seed(time.Now().UnixNano())

	var major int
	for {
		major = rand.Intn(65535)
		fmt.Println("major = ", major)
		count := -1
		db.Model(&Event{}).Where("major = ?", major).Count(&count)
		fmt.Println("count = ", count)
		if count == 0 {
			break
		}
	}

	return major
}

// GetEvent 指定されたmajorのイベントの情報を取得
func GetEvent(major int) (*Event, error) {
	var event Event

	if err := db.Where("major = ?", major).First(&event).Error; err != nil {
		return nil, err
	}
	pp.Println(&event)
	return &event, nil
}

// DeleteEvent 指定されたmajorのイベントがあるか確認し、あれば削除
func DeleteEvent(major int) (*Event, error) {
	var event Event
	if err := db.Where("major = ?", major).First(&event).Error; err != nil {
		return nil, err
	}
	if err := db.Where("major = ?", major).Delete(&event).Error; err != nil {
		return nil, err
	}
	pp.Println(event)
	return &event, nil
}

// CreateUser ユーザを作成
func CreateUser(user *User) {
	//db.NewRecord(event)
	db.Create(&user)
	//db.Save(&event)
}

// GetUsers 指定されたmajorのイベントの参加者を取得
func GetUsers(major int) (*[]User, error) {
	var users []User

	if err := db.Where("major = ?", major).Find(&users).Error; err != nil {
		return nil, err
	}
	pp.Println(&users)
	return &users, nil
}

// EventExist 指定されたmajorのイベントが存在するか確認
func EventExist(major int) error {
	var event Event
	fmt.Println(major)
	if err := db.Where("major = ?", major).First(&event).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 指定されたmajorのイベントがあるか確認し、あれば削除
func DeleteUser(id string) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	pp.Println(user)
	return &user, nil
}
