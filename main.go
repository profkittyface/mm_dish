package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("location_update", locationUpdate)

	GenerateLocationKey("ahunt")

	r.Run()
}

func locationUpdate(c *gin.Context) {
	var json LocationPost
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}

	if CheckLocationKey(json.AuthKey, json.UserId) == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Location key unknown"})
		return
	}

	coarray := json.Location
	trimmed := strings.ReplaceAll(coarray, "[", "")
	trimmed_1 := strings.ReplaceAll(trimmed, "]", "")
	s := strings.Split(trimmed_1, ",")
	x := s[0]
	y := s[1]
	gps_x, _ := strconv.ParseFloat(x, 32)
	gps_y, _ := strconv.ParseFloat(y, 32)

	location := UserLocation{
		Location:     Location{X: gps_x, Y: gps_y},
		UserId:       json.UserId,
		Last_updated: time.Now(),
	}
	InsertEntry(location)

}

func CheckLocationKey(locationKey string, userId int) bool {
	db := GetCursor()
	l := LocationAuth{}
	db.First(&l, "key=?", locationKey)
	return l.UserId == userId
}

func InsertEntry(user_location UserLocation) {
	db := GetCursor()
	db.Create(&user_location)
}

func GetCursor() *gorm.DB {
	dsn := "host=localhost user=mm password=mm dbname=mysterymeeting port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database")
	}
	return db
}

func GenerateLocationKey(username string) {
	s := []byte("mysterymeeting-" + username)
	h := sha1.New()
	h.Write(s)
	bs := hex.EncodeToString(h.Sum(nil))
	fmt.Println(bs)
}
