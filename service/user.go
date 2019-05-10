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
	//LevelMap 用户权限管理
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
	u0 := &UserJSON{OpenID: "Stranger1", WeChatName: "神秘人", Code: "神秘代码", Name: "Stranger1", IDCard: "000", Level: 0, Telephone: "110"}
	u1 := &UserJSON{OpenID: "Emeritus1", WeChatName: "万磁王", Code: "九头蛇", Name: "无关教授1", IDCard: "001", Level: 1}
	u2 := &UserJSON{OpenID: "Student1", WeChatName: "逃学", IDCard: "2", Name: "student1", Level: 2}
	u3 := &UserJSON{OpenID: "Assistant1", WeChatName: "小秘书", IDCard: "003", Name: "Assistant1", Level: 3}
	u4 := &UserJSON{OpenID: "Senior1", WeChatName: "高级打工仔", IDCard: "4", Name: "Senior1", Level: 4}
	u5 := &UserJSON{OpenID: "Full1", WeChatName: "全职叫兽", IDCard: "5", Name: "Full1", Level: 5}
	_ = u0.Create()
	_ = u1.Create()
	_ = u2.Create()
	_ = u3.Create()
	_ = u4.Create()
	_ = u5.Create()
}

//UserJSON2User UserJSON转换到User.
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
	user.IDCard = userJSON.IDCard
	user.Level = userJSON.Level
	user.Telephone = userJSON.Telephone
	return
}

//User2UserJSON User表单转换到UserJSON.
func User2UserJSON(user *models.User) (userJSON UserJSON) {
	/**
	  @Author: PantaZheng
	  @Description:
	  @Date: 2019/5/9 12:04
	*/
	userJSON.ID = user.ID
	userJSON.OpenID = user.OpenID
	userJSON.WeChatName = user.WeChatName
	userJSON.Name = user.Name
	userJSON.IDCard = user.IDCard
	userJSON.Level = user.Level
	userJSON.Telephone = user.Telephone
	return
}

func (userJSON *UserJSON) exchangeOpenID() (err error) {
	/**
	@Author: PantaZheng
	@Description: 根据code换取openid
	@Date: 2019/5/9 12:32
	*/
	field = title + "exchangeOpenId:\t"
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
	field = title + "UserInitByWechat:\t"
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
	field = title + "Create:\t"
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
	*userJSON = User2UserJSON(&u)
	return
}

//Bind 用户绑定.
func (userJSON *UserJSON) Bind() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 2:31
	*/
	//检查openid和code
	field = title + "Bind:\t"
	if err = userJSON.exchangeOpenID(); err != nil {
		err = errors.New(field + err.Error())
		return
	}
	wechatUser := &models.User{OpenID: userJSON.OpenID}
	_ = wechatUser.First()
	//检查微信是否已绑定
	if wechatUser.Level != LevelMap["Stranger"] {
		err = errors.New(field + "用户" + wechatUser.Name + "已经绑定过,如有修改需要请联系管理员")
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
		if err = presortedUser.Updates(&models.User{OpenID: userJSON.OpenID, WeChatName: wechatUser.WeChatName}); err != nil {
			err = errors.New(field + err.Error())
		} else {
			*userJSON = User2UserJSON(presortedUser)
		}
	} else {
		//无预存信息,修改微信初始化的User，添加姓名和身份证号
		if err = wechatUser.Updates(&models.User{Name: userJSON.Name, IDCard: userJSON.IDCard}); err != nil {
			err = errors.New(field + err.Error())
		} else {
			*userJSON = User2UserJSON(wechatUser)
		}
	}
	return
}

//First 单用户查找的原子方法
func (userJSON *UserJSON) First() (err error) {
	field = title + "first:\t"
	u := userJSON.UserJSON2User()
	if err = u.First(); err != nil {
		err = errors.New(field + err.Error())
	} else {
		*userJSON = User2UserJSON(&u)
	}
	return
}

//UserFindByID 通过数据库ID查找单用户.
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

//UserFindByOpenID 通过微信OpenID查找单用户.
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

//UserFindByIDCard 通过身份证查找单用户.
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

//Find 多用户查找的原子方法.
func (userJSON *UserJSON) Find() (usersJSON []UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 13:49
	*/
	field = title + "Find:\t"
	u := userJSON.UserJSON2User()
	if users, err := u.Find(); err != nil {
		err = errors.New(field + err.Error())
	} else {
		usersJSON = make([]UserJSON, len(users))
		for i, v := range users {
			usersJSON[i] = User2UserJSON(v)
		}
	}
	return
}

//UsersFindByLevel 通过权限等级查找用户表.
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
	field = title + "Updates:\t"
	u := userJSON.UserJSON2User()
	newUser := new(models.User)
	newUser.ID = u.ID
	if err := u.Updates(newUser); err != nil {
		err = errors.New(field + err.Error())
	} else {
		*userJSON = User2UserJSON(&u)
	}
	return
}

//Delete 用户删除的原子方法.
func (userJSON *UserJSON) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:16
	*/
	field = title + "Delete:\t"
	u := userJSON.UserJSON2User()
	if err := u.Delete(); err != nil {
		err = errors.New(field + err.Error())
	} else {
		*userJSON = User2UserJSON(&u)
	}
	return
}

//UserDeleteByID 通过数据库ID删除用户.
func UserDeleteByID(id uint) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:28
	*/
	userJSON = UserJSON{ID: id}
	err = userJSON.Delete()
	return
}

//UserDeleteByIDCard 通过I身份证删除用户.
func UserDeleteByIDCard(idCard string) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:30
	*/
	userJSON = UserJSON{IDCard: idCard}
	err = userJSON.Delete()
	return
}
