package model

import (
	"time"
)

type Order struct {
	Id_Order       string
	Customer_Id    string
	Tanggal_Masuk  time.Time
	Tanggal_Keluar time.Time
}
