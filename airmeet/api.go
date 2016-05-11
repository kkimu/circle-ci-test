package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v8"
)

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}

// RegisterEvent イベントを新規登録
func RegisterEvent(c echo.Context) error {
	fmt.Println("RegisterEvent")
	event := getPostEvent(c)
	pp.Println(event)

	if errs := validate.Struct(event); errs != nil {
		fmt.Println(errs)
		return c.JSON(http.StatusOK, NewError(400, fmt.Sprintf("%s", errs)))
	}

	event.Major = GenerateMajor()
	CreateEvent(event)
	return c.JSON(http.StatusOK, NewSuccess(event))
}

// GetEventInfo イベント情報を取得
func GetEventInfo(c echo.Context) error {
	fmt.Println("GetEventInfo")
	major, err := strconv.Atoi(c.Param("major"))
	fmt.Println(major)
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}

	event, err := GetEvent(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(event))
}

// RemoveEvent イベントを削除
func RemoveEvent(c echo.Context) error {
	fmt.Println("RemoveEvent")
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}

	event, err := DeleteEvent(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(event))
}

// Parse the request body, check input data
func getPostEvent(c echo.Context) *Event {
	// リクエストボディをパースして代入
	en, rn, desc, items := c.FormValue("eventName"), c.FormValue("roomName"), c.FormValue("description"), c.FormValue("items")

	return &Event{
		EventName:   en,
		RoomName:    rn,
		Description: desc,
		Items:       items,
		CreatedAt:   time.Now(),
	}
}

// RegisterUser ユーザを新規登録
func RegisterUser(c echo.Context) error {
	// UUID生成
	u := uuid.NewV4().String()
	fmt.Printf("UUIDv4: %s\n", u)

	user := getPostUser(c)

	user.ID = u

	err1 := imageSave(c, "image", u+".jpg")
	if err1 != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err1)))
	}

	err2 := imageSave(c, "image_header", u+"_header.jpg")
	if err2 != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err2)))
	}
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}
	if err := EventExist(major); err != nil {
		return c.JSON(http.StatusBadRequest, NewError(400, fmt.Sprintf("%s", err)))
	}
	user.Major = major
	pp.Println(user)

	if errs := validate.Struct(user); errs != nil {
		return c.JSON(http.StatusOK, NewError(400, fmt.Sprintf("%s", errs)))
	}

	CreateUser(user)
	return c.JSON(http.StatusOK, NewSuccess(user))
}

func imageSave(c echo.Context, input string, fname string) error {
	img, err := c.FormFile(input)
	if err != nil {
		return err
	}
	src, err := img.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("image/" + fname)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

// GetParticipants ユーザを取得
func GetParticipants(c echo.Context) error {
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}
	if err := EventExist(major); err != nil {
		return c.JSON(http.StatusBadRequest, NewError(400, fmt.Sprintf("%s", err)))
	}

	users, err := GetUsers(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(users))
}

// RemoveUser ユーザを削除
func RemoveUser(c echo.Context) error {
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}
	if err := EventExist(major); err != nil {
		return c.JSON(http.StatusBadRequest, NewError(400, fmt.Sprintf("%s", err)))
	}
	fmt.Println("delete")
	id := c.Param("id")

	user, err := DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(user))
}

// Parse the request body, check input data
func getPostUser(c echo.Context) *User {
	// リクエストボディをパースして代入
	un, prof, items := c.FormValue("name"), c.FormValue("profile"), c.FormValue("items")

	return &User{
		UserName:  un,
		Profile:   prof,
		Items:     items,
		CreatedAt: time.Now(),
	}
}
