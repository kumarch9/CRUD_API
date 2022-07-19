package worker

import "gorm.io/gorm"

type WorkerInfo struct {
	gorm.Model

	Name        string  `bson:"name" json:"name"`
	Age         int     `bson:"age" json:"age"`
	Address     string  `bson:"address" json:"address"`
	State       string  `bson:"state" json:"state"`
	Country     string  `bson:"country" json:"country"`
	CompanyName string  `bson:"CompanyName" json:"CompanyName"`
	Salary      float64 `bson:"salary" json:"salary"`
	Email       string  `bson:"email" json:"email"`
}
