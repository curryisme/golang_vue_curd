package handler

import (
	"github.com/curryisme/golang_vue_curd/enum"
	"github.com/curryisme/golang_vue_curd/model"
	"github.com/curryisme/golang_vue_curd/query"
	"github.com/curryisme/golang_vue_curd/resp"
	"github.com/curryisme/golang_vue_curd/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) GetEntity(result model.User) resp.User {
	return resp.User{
		UserId:   result.UserId,
		UserName: result.UserName,
		UserAge:  result.UserAge,
	}
}


func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}

	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	list, err := h.UserSrv.List(&q)
	total, err := h.UserSrv.GetTotal(&q)

	if err != nil {
		panic(err)
	}
	if q.PageSize == 0 {
		q.PageSize = 5
	}

	ret := int(total % q.PageSize)
	ret2 := int(total / q.PageSize)
	totalPage := 0

	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}

	var newList []*resp.User
	for _, item := range list {
		r := h.GetEntity(*item)
		newList = append(newList, &r)
	}

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     total,
		TotalPage: totalPage,
		Data:      newList,
	}

	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) GetUserByIdHandler(c *gin.Context) {

	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Data:      nil,
	}

    userId, error := strconv.Atoi(c.Param("id"))

	if error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	u := model.User{
		UserId: userId,
	}

	result, err := h.UserSrv.GetUserByIdHandler(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	r := h.GetEntity(*result)
	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}


func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.UserSrv.AddUserHandler(u)
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if r.UserId != 0 {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.UserSrv.EditUserHandler(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}

	id := c.Param("id")
	_id,_err := strconv.Atoi(id)
	if(_err != nil){
		c.JSON(http.StatusOK,gin.H{"entity": entity})
		return
	}

	b, err := h.UserSrv.DeleteUserHandler(_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}