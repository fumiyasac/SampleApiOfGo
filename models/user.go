package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"username" xorm:"'username'"`
	Password string `json:"password" xorm:"'password'"`
}

var engine *xorm.Engine

//
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:fumiya0921@/sample_api")
	if err != nil {
		panic(err)
	}
}

// UserRepository ...
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	result, _ := engine.Get(&user)
	if result {
		return &user
	}
	return nil
}
