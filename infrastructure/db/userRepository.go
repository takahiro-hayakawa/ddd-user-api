package infrastructure

import (
	"github.com/jinzhu/gorm"
	mUser "github.com/takahiro-hayakawa/user-api-server/domain/model/user"
	sUser "github.com/takahiro-hayakawa/user-api-server/domain/service/user"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	userRepository := UserRepository{db}
	return userRepository
}

func (uRepo UserRepository) Save(user mUser.User) error {
	now := time.Now()
	var userDTO sUser.UserDTO
	if user.ID.Value == 0 {
		userDTO = sUser.UserDTO{Name: user.Name.Value, MailAddress: user.MailAddress.Value, CreatedTime: now, UpdatedTime: now}
	} else {
		userDTO = sUser.UserDTO{ID: user.ID.Value, Name: user.Name.Value, MailAddress: user.MailAddress.Value, CreatedTime: now, UpdatedTime: now}
	}
	result := uRepo.db.Create(&userDTO)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uRepo UserRepository) FindByUserID(userID mUser.UserID) *mUser.User {
	var userDTO sUser.UserDTO
	uRepo.db.First(&userDTO, userID.Value)
	user := mUser.NewUser(userID, mUser.NewUserName(userDTO.Name), mUser.NewMailAddress(userDTO.MailAddress))
	return &user
}

func (UserRepository) FindByUserName(userName mUser.UserName) *mUser.User {
	return nil
}

func (UserRepository) FindByMailAddress(mailAddress mUser.MailAddress) *mUser.User {
	user := mUser.NewUser(mUser.NewUserID(1), mUser.NewUserName("takahiro"), mailAddress)
	return &user
}

func (UserRepository) Delete(user mUser.User) {
	return
}
