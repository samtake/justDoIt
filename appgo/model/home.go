package model

import (
	"github.com/jinzhu/gorm"
)

//BannerList  .
type BannerList struct {
	gorm.Model
	Icon string `gorm:"type:varchar(100);not null"json:"icon"`
	URL  string `gorm:"type:varchar(100);not null"json:"url"`
}

//LocalNavList  .
type LocalNavList struct {
	gorm.Model
	Icon           string `gorm:"type:varchar(100);not null"json:"icon"`
	Title          string `gorm:"type:varchar(20);not null"json:"title"`
	URL            string `gorm:"type:varchar(100);not null"json:"url"`
	StatusBarColor string `gorm:"type:varchar(20);not null"json:"statusBarColor"`
	HideAppBar     bool   `gorm:"-"json:"hideAppBar"`
}

// GridNavSubMainItem .
type GridNavSubMainItem struct {
	gorm.Model
	Icon           string `gorm:"type:varchar(100);not null"json:"icon"`
	Title          string `gorm:"type:varchar(20);not null"json:"title"`
	URL            string `gorm:"type:varchar(100);not null"json:"url"`
	StatusBarColor string `gorm:"type:varchar(20);not null"json:"statusBarColor"`
	TypeName       string `gorm:"type:varchar(20);not null"json:"type"`
}

//GridNavSub .
type GridNavSub struct {
	gorm.Model
	StartColor              string             `gorm:"type:varchar(20);not null"json:"startColor"`
	EndColor                string             `gorm:"type:varchar(20);not null"json:"endColor"`
	GridNavSubMainItem      GridNavSubMainItem `gorm:"foreignkey:GridNavSubMainItemRefer"` // 将 ItemRefer 作为外键
	GridNavSubMainItemRefer uint
	TypeName                string `gorm:"type:varchar(20);not null"json:"type"` //hotel、flight、travel、
}

//GridNav  .
type GridNav struct {
	gorm.Model
	Type            string     `gorm:"type:varchar(20);not null"`
	GridNavSub      GridNavSub `gorm:"foreignkey:GridNavSubRefer"` // 将 SubRefer 作为外键
	GridNavSubRefer uint
}

//SubNavList  .
type SubNavList struct {
	gorm.Model
	Icon       string `gorm:"type:varchar(100);not null"json:"icon"`
	Title      string `gorm:"type:varchar(20);not null"json:"title"`
	URL        string `gorm:"type:varchar(100);not null"json:"url"`
	HideAppBar bool   `gorm:"-"json:"hideAppBar"`
}

//SalesBox  .
type SalesBox struct {
	gorm.Model
	Icon     string `gorm:"type:varchar(100);not null"json:"icon"`
	Title    string `gorm:"type:varchar(20);not null"json:"title"`
	URL      string `gorm:"type:varchar(100);not null"json:"url"`
	TypeName string `gorm:"type:varchar(20);not null"json:"type"`
}
