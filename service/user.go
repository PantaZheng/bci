package service

import (
	"errors"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
)

const title = "service.user."

var (
	LevelMap = map[string]int{
		//Stranger 未绑定
		"Stranger": 0,
		//Emeritus Professor emeritus 专家教授
		"Emeritus": 1,
		//Student 学生
		"Student": 2,
		//Assistant 助理
		"Assistant": 3,
		//Senior Senior lecturer 高级讲师
		"Senior": 4,
		//Full Full professor 全职教授
		"Full": 5,
	}
	field = ""
)

//UserJSON 用户Json原型
type UserJSON struct {
	/**
	@Author: PantaZheng
	@Description:用户JSON
	@Date: 2019/5/9 10:42
	*/
	ID         uint   `json:"id,omitempty"`
	OpenID     string `json:"openid,omitempty"`
	WeChatName string `json:"openid,omitempty"`
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	IDCard     string `json:"idCard,omitempty"`
	Level      int    `json:"level"`
	Telephone  string `json:"telephone,omitempty"`
}

func userTestData() {
	_ = UserJSON{OpenID: "Stranger1", IDCard: "0", Name: "Stranger1", Level: 0}.Create()
	_ = UserJSON{OpenID: "Emeritus1", IDCard: "1", Name: "Emeritus1", Level: 1}.Create()
	_ = UserJSON{OpenID: "student1", IDCard: "2", Name: "student1", Level: 2}.Create()
	_ = UserJSON{OpenID: "Assistant1", IDCard: "3", Name: "Assistant1", Level: 3}.Create()
	_ = UserJSON{OpenID: "Senior1", IDCard: "4", Name: "Senior1", Level: 4}.Create()
	_ = UserJSON{OpenID: "Full1", IDCard: "5", Name: "Full1", Level: 5}.Create()
}

//UserJSON2User
func (userJSON *UserJSON) UserJSON2User() (user models.User) {
	/**
	@Author: PantaZheng
	@Description: UserJSON转化为User
	@Date: 2019/5/9 13:14
	*/
	user.ID = userJSON.ID
	user.OpenID = userJSON.OpenID
	user.WeChatName = userJSON.WeChatName
	user.Name = userJSON.Name
	user.Code = userJSON.Code
	user.IDCard = userJSON.IDCard
	user.Level = userJSON.Level
	user.Telephone = userJSON.Telephone
	return
}

func (userJSON *UserJSON) exchangeOpenId() (err error) {
	/**
	@Author: PantaZheng
	@Description: 根据code换取openid
	@Date: 2019/5/9 12:32
	*/
	field = title + "exchangeOpenId" + ":\t\n"
	if userJSON.OpenID == "" {
		if userJSON.Code != "" {
			token := &oauth2.Token{}
			if err := ExchangeToken(token, userJSON.Code); err != nil {
				err = errors.New(field + err.Error())
			}
			userJSON.OpenID = token.OpenId
		} else {
			err = errors.New(field + "openid和code不可同时为空")
		}
	} else if userJSON.Code != "" {
		err = errors.New(field + "有openid就别传code了，我不会去换的")
	}
	return
}

//simplify 简化，仅保留：id、Name、Level.
func (userJSON *UserJSON) simplify() {
	userJSON.OpenID = ""
	userJSON.WeChatName = ""
	userJSON.Code = ""
	userJSON.IDCard = ""
	userJSON.Level = 0
	userJSON.Telephone = ""
}

//UserInitByWechat 用户登录微信初始化微信.
func UserInitByWechat(weChatInfo *user.UserInfo) string {
	/**
	@Author: PantaZheng
	@Description: 用户登录微信初始化微信,默认level等级为Stranger
	@Date: 2019/5/9 23:13
	*/
	field = title + "UserInitByWechat" + ":\t\n"
	u := new(UserJSON)
	u.OpenID = weChatInfo.OpenId
	if err := u.Create(); err != nil {
		err = errors.New(field + err.Error())
		return err.Error()
	}
	log.Printf("UserInit:\t" + weChatInfo.OpenId)
	return "欢迎关注"
}

//Create 创建User.
func (userJSON *UserJSON) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description: Service层 User创建
	@Date: 2019/5/9 23:11
	*/
	field = title + "Create" + ":\t\n"
	u := userJSON.UserJSON2User()
	if err = u.Create(); err != nil {
		err = errors.New(field + err.Error())
		return
	}
	//OpenID和IDCard零值设置为ID,并更新字段信息
	if u.OpenID == "" || u.IDCard == "" {
		newUser := &models.User{}
		if u.OpenID == "" {
			newUser.OpenID = string(u.ID)
		}
		if u.IDCard == "" {
			newUser.IDCard = string(u.ID)
		}
		if err = u.Updates(newUser); err != nil {
			err = errors.New(field + err.Error())
			return
		}
	}
	//userJSON接收数据库创建记录
	u.User2UserJSON(*userJSON)
	return
}

func (userJSON *UserJSON) Bind() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 2:31
	*/
	//检查openid和code
	field = title + "Bind" + ":\t\n"
	if err = userJSON.exchangeOpenId(); err != nil {
		err = errors.New(field + err.Error())
		return
	}
	wechatUser := &models.User{OpenID: userJSON.OpenID}
	_ = wechatUser.First()
	//检查微信是否已绑定
	if wechatUser.Level != LevelMap["Stranger"] {
		err = errors.New(field + "用户" + wechatUser.Name + "已经绑定过\n,如有修改需要请联系管理员")
		return
	}
	presortedUser := &models.User{IDCard: userJSON.IDCard}
	_ = presortedUser.First()
	//检查是否有预存信息
	if presortedUser.ID != 0 {
		//有预存信息，比对姓名
		if presortedUser.Name != userJSON.Name {
			err = errors.New(field + "用户名:" + userJSON.Name + "和身份证号:" + userJSON.IDCard + "不匹配,请检查输入信息")
			return
		}
		//TODO：进行绑定操作，1.删除微信初始化的User；2.预存User添加微信信息
		//软删除微信初始化创建的用户信息
		if err = wechatUser.Delete(); err != nil {
			err = errors.New(field + err.Error())
			return
		}
		//修改预存用户信息
		if err = presortedUser.Updates(&models.User{OpenID: userJSON.OpenID, Code: userJSON.Code, WeChatName: wechatUser.WeChatName}); err != nil {
			err = errors.New(field + err.Error())
			return
		}
		presortedUser.User2UserJSON(*userJSON)
	} else {
		//无预存信息,修改微信初始化的User，添加姓名和身份证号
		if err = wechatUser.Updates(&models.User{Name: userJSON.Name, IDCard: userJSON.IDCard}); err != nil {
			err = errors.New(field + err.Error())
			return
		}
		wechatUser.User2UserJSON(*userJSON)
	}
	return
}

func (userJSON *UserJSON) First() (err error) {
	field = title + "first" + ":\t\n"
	u := userJSON.UserJSON2User()
	if err = u.First(); err != nil {
		err = errors.New(field + err.Error())
	}
	u.User2UserJSON(*userJSON)
	return
}

func UserFindByID(id uint) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 1:52
	*/
	userJSON = UserJSON{ID: id}
	err = userJSON.First()
	return
}

func UserFindByOpenID(openid string) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 1:52
	*/
	userJSON = UserJSON{OpenID: openid}
	err = userJSON.First()
	return
}

func UserFindByIDCard(idCard string) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 1:52
	*/
	userJSON = UserJSON{IDCard: idCard}
	err = userJSON.First()
	return
}

func (userJSON *UserJSON) Find() (usersJSON []UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 13:49
	*/
	field = title + "Find" + ":\t\n"
	u := userJSON.UserJSON2User()
	if users, err := u.Find(); err != nil {
		err = errors.New(field + err.Error())
	} else {
		usersJSON = make([]UserJSON, len(users))
		for i, v := range users {
			v.User2UserJSON(usersJSON[i])
		}
	}
	return
}

func UsersFindByLevel(level int) (usersJSON []UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:01
	*/
	userJSON := UserJSON{Level: level}
	usersJSON, err = userJSON.Find()
	return
}

//Updates 更新用户数据，id定位用户记录.
func (userJSON *UserJSON) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:09
	*/
	field = title + "Updates" + ":\t\n"
	u := userJSON.UserJSON2User()
	newUser := new(models.User)
	newUser.ID = u.ID
	if err := u.Updates(newUser); err != nil {
		err = errors.New(field + err.Error())
		return
	}
	u.User2UserJSON(*userJSON)
	return
}

func (userJSON *UserJSON) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:16
	*/
	field = title + "Delete" + ":\t\n"
	u := userJSON.UserJSON2User()
	if err := u.Delete(); err != nil {
		err = errors.New(field + err.Error())
		return
	}
	u.User2UserJSON(*userJSON)
	return
}

func UserDeleteByID(id uint) (userJson UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:28
	*/
	userJson = UserJSON{ID: id}
	err = userJson.Delete()
	return
}

func UserDeleteByIDCard(idCard string) (userJson UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:30
	*/
	userJson = UserJSON{IDCard: idCard}
	err = userJson.Delete()
	return
}
