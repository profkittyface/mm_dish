package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LocationPost struct {
	UserId   int
	AuthKey  string
	Location string
}

type LocationAuth struct {
	Id      int    `gorm:"column:id"`
	UserId  int    `gorm:"column:userid"`
	AuthKey string `gorm:"location_key"`
}

func (LocationAuth) TableName() string {
	return "location_key"
}

type UserLocation struct {
	Id           int       `gorm:"column:id"`
	Location     Location  `gorm:"column:location"`
	UserId       int       `gorm:"column:userid"`
	Last_updated time.Time `gorm:"column:last_updated"`
}

type Tabler interface {
	TableName() string
}

func (UserLocation) TableName() string {
	return "location"
}

type Location struct {
	X, Y float64
}

// // Scan implements the sql.Scanner interface
// func (loc *Location) Scan(v interface{}) error {
// 	// Scan a value into struct from database driver
// }

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.X, loc.Y)},
	}
}
