package controller

import (
	"../../service"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/kataras/iris"
)


// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func Wechat(ctx iris.Context) {
	wxAppId := "wx2203c68c9311ea43"
	//wxAppSecret := "40c40547e174ed99d1281b2890f7eeb3"
	wxOriId      := "gh_b4421d95f839"
	wxToken         := "HMIIoT"
	wxEncodedAESKey := "iesxoHBsnaKVry5E8xd8gavmJLTVVNcd8aS7w3KYOaU"

	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(service.DefaultEventHandler)
	mux.DefaultEventHandleFunc(service.DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, service.TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, service.MenuClickEventHandler)

	msgHandler := mux
	msgServer := core.NewServer(wxOriId, wxAppId, wxToken, wxEncodedAESKey, msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)




}





