package service

import "learn-wire/dao/user"

type UserService struct {
	UserRepo *user.UserRepo
}
