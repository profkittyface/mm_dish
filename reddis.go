package main

import (
	"context"
)

type LocationKeyChecker interface {
	CheckLocationKey() bool
	CheckUserID()	bool	
}

var ctx = context.Background()

// func ReddisClient() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// 	return rdb
// }

type DatabaseLocationKeyChecker struct {
	UserId int
	LocationKey string
}
