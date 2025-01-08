package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	initDB()
	batchInsert()

	Query()

	fmt.Println("delete")

	Delete()

	Query()
}

func Delete() {
	db := db.WithContext(context.Background()).Session(&gorm.Session{})
	err := db.Table(User{}.TableName()).Delete(&User{
		ID: 1,
	}).Error
	if err != nil {
		fmt.Println("del fail. ")
		return
	}

	fmt.Println("del success. ")
	return
}

func Query() {
	db := db.WithContext(context.Background()).Session(&gorm.Session{})
	var userList []User

	var num, size int
	if num == 0 && size == 0 {
		num, size = 1, 50
	}

	db = db.Table(User{}.TableName())
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		fmt.Println("query count fail. err:", err)
		return
	}
	fmt.Println("count:", total)

	//db.Where("ID = (?)", 1)
	fmt.Println(*db.Statement.Dest.(*int64))
	err = db.Offset((num - 1) * size).Limit(size).Order("created_at DESC").
		Find(&userList).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err != nil {
		fmt.Println("query  fail. err:", err)
		return
	}

	//fmt.Println(*db.Statement.Dest.(*int64))
	err = db.Count(&total).Error
	if err != nil {
		fmt.Println("query count fail. err:", err)
		return
	}
	fmt.Println("count:", total)
	fmt.Println(*db.Statement.Dest.(*int64))

	fmt.Printf("Found user: %+v\n", userList)
}

func batchInsert() {
	// 创建用户
	userList := []User{
		{ID: 1, Name: "John Doe", Age: 25},
		{ID: 2, Name: "Bob", Age: 26},
	}
	db.CreateInBatches(&userList, len(userList))
}

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Age       int
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;default:NULL" json:"deleted_at"`
}

// TableName 指定 User 结构体对应的表名为 custom_users
func (User) TableName() string {
	return "custom_users"
}

var (
	db  *gorm.DB
	err error
)

func initDB() {
	// 创建一个开发环境下的日志记录器
	//logger, _ := zap.NewProduction()
	//defer logger.Sync()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接 SQLite 数据库
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
}

// GormLogger 自定义的 GORM 日志记录器，使用 zap 记录日志
type GormLogger struct {
	Logger *zap.Logger
}

func (l GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	//TODO implement me
	panic("implement me")
}

func (l GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	//panic("implement me")
	l.Logger.Info(s)
}

func (l GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	//panic("implement me")
	l.Logger.Warn(s)
}

func (l GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	//panic("implement me")
	l.Logger.Error(s)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	//panic("implement me")
}

// Print 实现 gorm.logger.Interface 的 Print 方法
func (l GormLogger) Print(v ...interface{}) {
	if len(v) > 1 {
		switch v[0] {
		case "sql":
			l.Logger.Info("SQL Query", zap.String("query", v[3].(string)), zap.Any("values", v[4]))
		case "error":
			l.Logger.Error("Database Error", zap.Error(v[1].(error)))
		}
	}
}
