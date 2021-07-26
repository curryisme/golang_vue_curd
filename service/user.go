package service

import (
	"errors"
	"fmt"
	"github.com/curryisme/golang_vue_curd/config"
	"github.com/curryisme/golang_vue_curd/model"
	"github.com/curryisme/golang_vue_curd/query"
	"github.com/curryisme/golang_vue_curd/repository"
	"github.com/curryisme/golang_vue_curd/utils"
)

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	GetUserByIdHandler(user model.User) (*model.User, error)
	AddUserHandler(user model.User) (*model.User, error)
	EditUserHandler(u model.User) (bool, error)
	DeleteUserHandler(id int) (bool, error)
}
type UserService struct {
	Repo repository.UserRepoInterface
}

func (srv *UserService) List(req *query.ListQuery) (users []*model.User, err error) {
	if req.PageSize < 1 {
		req.PageSize = config.PAGE_SIZE
	}
	return srv.Repo.List(req)
}

func (srv *UserService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *UserService) GetUserByIdHandler(user model.User) (*model.User, error) {
	return srv.Repo.GetUserByIdHandler(user)
}

func (srv *UserService) AddUserHandler(user model.User) (*model.User, error) {
	//根据用户名判断是否存在用户
	result := srv.Repo.ExistByUserName(user.UserName)
	if result != nil {
		fmt.Println("用户已经存在")
		return nil, errors.New("用户已经存在")
	}
	if user.UserPwd == "" {
		user.UserPwd = utils.Md5("123456")
	}else{
		user.UserPwd = utils.Md5(user.UserPwd)
	}
	return srv.Repo.AddUserHandler(user)
}

func (srv *UserService) EditUserHandler(user model.User) (bool, error) {
	//根据Id判断是否存在用户
	result := srv.Repo.ExistByUserId(user.UserId)
	if result != nil {
		return srv.Repo.EditUserHandler(user)
	}
	return false, errors.New("用户不存在")
}

func (srv *UserService) DeleteUserHandler(id int) (bool, error) {

	user := srv.Repo.ExistByUserId(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	return srv.Repo.DeleterHandler(*user)
}

