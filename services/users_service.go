package services

import (
	"fmt"
	"github.com/steven7/bookstore_utils-go/rest_errors"
	"github.com/steven7/go-microservces_users_api/domain/users"
	"github.com/steven7/go-microservces_users_api/utils/crypto_utils"
	"github.com/steven7/go-microservces_users_api/utils/date_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
	usersServiceInterface
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User,  rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	fmt.Println("in create user")
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetCurrentDateTimeDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	fmt.Println("in create user after save")
	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	dao := &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *usersService) UpdateUser(isPartial bool , user users.User) (*users.User, rest_errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get() ; err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func  (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func  (s *usersService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: request.Password,
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}