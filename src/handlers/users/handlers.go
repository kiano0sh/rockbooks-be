package users

import (
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"golang.org/x/crypto/bcrypt"
)

func (user *User) Register() error {
	hashedPassword, err := EncryptPassword(user.Password)
	if err != nil {
		return grapherrors.ReturnGQLError("Something went wrong on registration!", err)
	}
	user.Password = hashedPassword
	result := database.DB.Create(&user)
	if result.Error != nil {
		return grapherrors.ReturnGQLError("شما قبلا ثبت نام کرده اید", result.Error)
	}
	return nil
}

func (user *User) Authenticate() error {
	var userResult *User
	result := database.DB.Where("email = ?", user.Email).First(&userResult)
	if result.Error != nil {
		return grapherrors.ReturnGQLError("نام کاربری یا رمز عبور صحیح نیست", result.Error)
	}
	passwordIsOk := CheckPassword(userResult.Password, user.Password)
	if !passwordIsOk {
		return grapherrors.ReturnGQLError("نام کاربری یا رمز عبور صحیح نیست", result.Error)
	}
	return nil
}

func GetUserIdByEmail(email string) (int64, error) {
	var userResult *User
	result := database.DB.Where("email = ?", email).First(&userResult)
	if result.Error != nil {
		panic("No user has been found!")
	}
	return userResult.ID, nil
}

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
