package models

import "golang.org/x/crypto/bcrypt"

//const DriverName  = "mysql"
//const MasterDataSourceName  = "root:666666@tcp(127.0.0.1:3306)/laravel_shop?charset=utf8"

type User struct {
	ID   int64   `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	Name string  `xorm:"not null comment('用户名') VARCHAR(50)" form:"user_name"`
	Pwd  string   `xorm:"not null comment('密码') VARCHAR(50)" form:"user_pwd"`
	Email string `xorm:"not null comment('邮箱') VARCHAR(50)" form:"user_email"`
	Active bool  `xorm:"not null comment('是否激活') default(false)" form:"active"`
	HashedPassword []byte `xorm:"-" form:"-"`
}

func (u User) IsValid() bool {
	return u.ID > 0
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

type Product struct {
	ID   int64   `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	Name string  `xorm:"not null comment('商品名') VARCHAR(50)" form:"productname"`
}

type Carts struct {

}