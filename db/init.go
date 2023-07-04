/*
Package db
包含大部分对数据库的直接操作
函数书写顺序按照增删改查的顺序，同样结构的操作放在一起
Create -> Delete -> Update -> Read -> other
*/
package db

import (
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Mysql *gorm.DB

// MysqlConfig mysql config
type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// print mysql dsn
func (c MysqlConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database)
}

func Init() {
	c := MysqlConfig{
		Host:     config.GetString("mysql.host"),
		Port:     config.GetString("mysql.port"),
		Username: config.GetString("mysql.username"),
		Password: config.GetString("mysql.password"),
		Database: config.GetString("mysql.database"),
	}
	connect(&c)
}

// connect mysql
func connect(c *MysqlConfig) {
	var err error
	// Get Mysql connection
	if config.GetString("mysql.host") == "db" {
		for i := 0; i < 10; i++ {
			Mysql, err = gorm.Open(mysql.Open(c.String()), &gorm.Config{})
			if err == nil {
				fmt.Println("mysql connect success")
				break
			}
			time.Sleep(10 * time.Second)
		}
	} else {
		Mysql, err = gorm.Open(mysql.Open(c.String()), &gorm.Config{})
		if err != nil {
			fmt.Println("mysql connect error")
		}
	}

	// Auto migrate
	err = autoMigrate()
	if err != nil {
		fmt.Println("mysql auto migrate error")
	}
	Mysql = Mysql.Debug()
}

// autoMigrate 自动化建表
func autoMigrate() error {
	// 自动创建需要的表，需要注意的是，这里的顺序不能乱，因为有外键约束
	err := Mysql.AutoMigrate(&model.Course{},
		&model.Chapter{},
		&model.Question{},
		&model.QuestionDetail{})
	if err != nil {
		return err
	}
	return nil
}
