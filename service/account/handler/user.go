package handler

import (
	"context"
	proto "k8s/service/account/proto"
)

func Signup(ctx context.Context, in *proto.ReqSignup, out *proto.ReqSignup) error{

}

func Signin(ctx context.Context, in *proto.ReqSignin, out *proto.RespSignin) error{

}
