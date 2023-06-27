// Code generated by gen. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"github.com/sixwaaaay/shauser/user"
)

type (
	FollowActionReply    = user.FollowActionReply
	FollowActionRequest  = user.FollowActionRequest
	GetFollowersReply    = user.GetFollowersReply
	GetFollowersRequest  = user.GetFollowersRequest
	GetFollowingsReply   = user.GetFollowingsReply
	GetFollowingsRequest = user.GetFollowingsRequest
	GetFriendsReply      = user.GetFriendsReply
	GetFriendsRequest    = user.GetFriendsRequest
	GetUserReply         = user.GetUserReply
	GetUserRequest       = user.GetUserRequest
	GetUsersReply        = user.GetUsersReply
	GetUsersRequest      = user.GetUsersRequest
	LoginReply           = user.LoginReply
	LoginRequest         = user.LoginRequest
	RegisterReply        = user.RegisterReply
	RegisterRequest      = user.RegisterRequest
	UpdateUserReply      = user.UpdateUserReply
	UpdateUserRequest    = user.UpdateUserRequest
	User                 = user.User
)

var NewClient = user.NewUserServiceClient