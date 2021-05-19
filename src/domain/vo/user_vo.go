package vo

type UserVO struct {
	UserId   int64  `gorm:"primaryKey; not null" json:"userId" description:"用户id"`
	UserName string `gorm:"varchar(128); not null" json:"userName" description:"用户名"`
	Age      int8   `gorm:"varchar(128); not null" json:"age" description:"年龄"`
	Gender   int8   `gorm:"varchar(256); not null" json:"gender" description:"性别"`
}
