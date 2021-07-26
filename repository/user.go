package repository

import (
	"fmt"
	"github.com/curryisme/golang_vue_curd/model"
	"github.com/curryisme/golang_vue_curd/query"
	"github.com/curryisme/golang_vue_curd/utils"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	GetUserByIdHandler(user model.User) (*model.User, error)
	ExistByUserName(userName string) *model.User
	AddUserHandler(user model.User) (*model.User, error)
	ExistByUserId(userId int) *model.User
	EditUserHandler(user model.User) (bool, error)
	DeleterHandler(user model.User) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []*model.User, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页

	if err := db.Order("user_id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var users []model.User
	db := repo.DB
	int64_ := int64(total)
	if err := db.Find(&users).Count(&int64_).Error; err != nil {
		return int(int64_), err
	}
	return int(int64_), nil
}

func (repo *UserRepository) GetUserByIdHandler(user model.User) (*model.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) ExistByUserName(userName string) *model.User {
	var user model.User
	result := repo.DB.First(&user,"user_name = ?" , userName)
	if result.RowsAffected > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByUserId(userId int) *model.User {
	var user model.User
	result := repo.DB.First(&user,"user_id = ?" , userId)
	if result.RowsAffected > 0 {
		return &user
	}
	return nil
}


func (repo *UserRepository) AddUserHandler(user model.User) (*model.User, error) {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}

func (repo *UserRepository) EditUserHandler(user model.User) (bool, error) {
	user.UserPwd =  utils.Md5(user.UserPwd)
	err := repo.DB.Model(&user).Where("user_id = ?", user.UserId).Select( "UserAge","UserPwd").Updates(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) DeleterHandler(user model.User) (bool, error) {
	err := repo.DB.Where("user_id=?",user.UserId).Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
