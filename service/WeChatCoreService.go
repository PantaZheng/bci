package service

import (
	"../config"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/mp/user/tag"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"log"
	"strconv"
)

var (
	wechatConfigTree =config.Conf.Get("wechat").(*toml.Tree)
	wechatOriId = wechatConfigTree.Get("OriId").(string)
	wechatAppId = wechatConfigTree.Get("AppId").(string)
	wechatAppSecret = wechatConfigTree.Get("AppSecret").(string)
	wechatToken = wechatConfigTree.Get("Token").(string)
	wechatEncodedAESKey = wechatConfigTree.Get("EncodedAESKey").(string)
	defaultClt = wechatClient()
	tagTeacher = 0
	tagStudent = 0
)

func TextMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	_=ctx.RawResponse(resp)
}



func MenuClickEventHandler(ctx *core.Context) {
	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "请先登记个人信息")
	if err:=ctx.RawResponse(resp);err!=nil{
		fmt.Printf("MenuClickEventHandlerERR:%v", err)
	}
}

func SubscribeEventHandler(ctx *core.Context){
	event := request.GetSubscribeEvent(ctx.MixedMsg)
	clt := wechatClient()
	info,_:=user.Get(clt,event.FromUserName,"")
	resp := response.NewText(event.FromUserName,event.ToUserName,event.CreateTime, UserInit(info))
	if err:=ctx.RawResponse(resp);err!=nil{
		fmt.Printf("SubscribeEventHandlerERR:%v",err)
	}

}

func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	_=ctx.NoneResponse()
}

func WechatServer(ctx iris.Context) {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(DefaultEventHandler)
	mux.DefaultEventHandleFunc(DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)
	mux.EventHandleFunc(request.EventTypeSubscribe,SubscribeEventHandler)
	msgHandler := mux

	msgServer := core.NewServer(wechatOriId, wechatAppId, wechatToken, wechatEncodedAESKey,msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)
}

func wechatClient() *core.Client{
	accessTokenTokenServer :=core.NewDefaultAccessTokenServer(wechatAppId,wechatAppSecret,nil)
	return core.NewClient(accessTokenTokenServer,nil)
}

//向微信服务器创建Tag
func CreateTag(){
	_,err:=tag.Create(defaultClt,"student")
	if err!=nil{
		fmt.Printf("\nCreateTag%v\n",err)
	}
	_,err=tag.Create(defaultClt,"teacher")
	if err!=nil{
		fmt.Printf("\nCreateTag%v\n",err)
	}
}

//获得Tag的ID
func GetTagList(){
	tagList,err:=tag.List(defaultClt)
	if err!=nil{
		fmt.Printf("%v",err)
	}
	for _,v :=range tagList{
		if v.Name == "student"{
			tagStudent=v.Id
		}
		if v.Name == "teacher"{
			tagTeacher=v.Id
		}
		fmt.Printf("\nid: "+strconv.Itoa(v.Id)+"\tname: "+v.Name+""+"\tcount: "+strconv.Itoa(v.UserCount)+"\n")
	}
}

func AddRoleTag(openIdList []string, tagId int){
	if err:=tag.BatchTag(defaultClt,openIdList,tagId);err!=nil{
		fmt.Printf("AddRoleTagError:%v\n",err)
	}
	GetTagList()
}

func DelRoleTag(openIdList []string, tagId int){
	if err:=tag.BatchUntag(defaultClt,openIdList,tagId);err!=nil{
		fmt.Printf("AddRoleTagError:%v\n",err)
	}
	GetTagList()
}

func DefaultMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsClickButton("架构","RelationShip")
	btnProjectMission:=menu.Button{}
	btnProjectMission.SetAsClickButton("项目/任务","ProjectMission")
	btnEnroll:=menu.Button{}
	btnEnroll.SetAsViewButton("登记","https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx2203c68c9311ea43&redirect_uri=http://bci.renjiwulian.com/test&response_type=code&scope=snsapi_base&state=12#wechat_redirect")
	defaultButtons:= []menu.Button{btnRelationShip,btnProjectMission,btnEnroll}
	defaultMenu:=menu.Menu{}
	defaultMenu.Buttons= defaultButtons
	err:=menu.Create(defaultClt,&defaultMenu)
	if err!=nil{
		fmt.Printf("DefaultMenu%v\n",err)
	}
}

func TeacherMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsClickButton("架构","RelationShip")
	btnProject:=menu.Button{}
	btnProject.SetAsClickButton("项目","Project")
	btnPersonal:=menu.Button{}
	btnPersonal.SetAsClickButton("个人","Personal")
	teacherButtons := []menu.Button{btnRelationShip, btnProject,btnPersonal}
	teacherRule :=menu.MatchRule{}
	teacherRule.TagId=strconv.Itoa(tagTeacher)
	teacherMenu :=menu.Menu{}
	teacherMenu.Buttons= teacherButtons
	teacherMenu.MatchRule=&teacherRule
	_,err:=menu.AddConditionalMenu(defaultClt,&teacherMenu)
	if err!=nil{
		fmt.Printf("\nTeacherMenu:%v\n",err)
	}
}

func StudentMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsClickButton("架构","RelationShip")
	btnMission :=menu.Button{}
	btnMission.SetAsClickButton("任务","Project")
	btnPersonal:=menu.Button{}
	btnPersonal.SetAsClickButton("个人","Personal")
	studentButtons := []menu.Button{btnRelationShip, btnMission,btnPersonal}
	studentRule := menu.MatchRule{}
	studentRule.TagId=strconv.Itoa(tagStudent)
	studentMenu := menu.Menu{}
	studentMenu.Buttons= studentButtons
	studentMenu.MatchRule = &studentRule
	_,err:=menu.AddConditionalMenu(defaultClt,&studentMenu)
	if err!=nil{
		fmt.Printf("\nStudentMenu%v\n",err)
	}
}

func TestMenu(){
	M,err:=menu.TryMatch(defaultClt,"oPKFh5lM9MA6_Svd39Km-84no7c8")
	if err!=nil{
		fmt.Printf("%v\n",err)
	}else{
		for _, v:=range M.Buttons{
			fmt.Printf( v.Name+"\t")
		}
	}
	fmt.Printf("\n")
}