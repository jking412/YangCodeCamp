package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"column:email;type:varchar(255);unique"`
	Password string `gorm:"column:password;type:varchar(255)"`
	NickName string `gorm:"column:nickname;type:varchar(255)"`
}

type UserCourseProgress struct {
	ID       int `gorm:"primaryKey"`
	UserID   int
	CourseID int
	Status   Status
}

type UserChapterProgress struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	ChapterID int
	Status    Status
}
