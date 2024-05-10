package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id_user"`
	NamaLengkap string `gorm:"varchar(300)" json:"nama_user"`
	Username    string `gorm:"varchar(300)" json:"username_user"`
	Password    string `gorm:"varchar(300)" json:"password_user"`
}
