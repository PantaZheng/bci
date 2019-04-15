package models

import (
	"../database"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	OpenId string `gorm:"primary_key;unique" json:"openid"`
	Code string `gorm:"not null" json:"code"`
	Name         string `gorm:"not null" json:"name"`
	Sex          string `gorm:"not null" json:"sex"`
	Role         string `gorm:"not null " json:"role"`
	School       string `gorm:"not null" json:"school"`
	Supervisor	string `gorm:"not null" json:"supervisor"`
	//HduId        string `gorm:"VARCHAR(191)"`
	//Level        string `gorm:"VARCHAR(191)"`
	//TagId        int `gorm:"VARCHAR(191)"`
	//WeChatAccount string `gorm:"VARCHAR(191)"`
	//WechatNickname string `gorm:"not null VARCHAR(255)"`
	//QQ string `gorm:"VARCHAR(191)"`
	//Telephone string `gorm:"VARCHAR(191)"`
	//Email string `gorm:"VARCHAR(191)"`
	//EduStartDate string `gorm:"VARCHAR(191)"`
	//Graduate string `gorm:"VARCHAR(191)"`
}

type MemberInfo struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

func CheckTableUser() {
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&User{})
	MakeTestData()
}

func DropTableUsers(){
	database.DB.DropTable("users")
	log.Printf("删除用户表\n")
}

func MakeTestData(){
	_= EnrollUser(&User{OpenId: "test1",Role:"unEnrolled"})
	_= EnrollUser(&User{OpenId: "test2",Role:"unEnrolled"})
	_= EnrollUser(&User{OpenId: "test3",Role:"unEnrolled"})
	_= EnrollUser(&User{OpenId: "student1",Name:"student1",Role:"student",Supervisor:"teacher1"})
	_= EnrollUser(&User{OpenId: "student2",Name:"student2",Role:"student",Supervisor:"teacher1"})
	_= EnrollUser(&User{OpenId: "student3",Name:"student3",Role:"student",Supervisor:"teacher2"})
	_= EnrollUser(&User{OpenId: "teacher1",Name:"戴国骏",Role:"teacher"})
	_= EnrollUser(&User{OpenId: "teacher2",Name:"张桦",Role:"teacher"})
	_= EnrollUser(&User{OpenId: "teacher_unknown",Name:"其他导师",Role:"teacher"})
	log.Printf("创建测试用户数据")
}


//根据Role获得成员信息
func GetMembersByRole(role string) ( memberList [] MemberInfo) {
	var userList [] User
	database.DB.Find(&userList,&User{Role:role}).Select("id","name")
	memberList=make([]MemberInfo,len(userList))
	for i,v := range userList{
		memberList[i].Id=v.ID
		memberList[i].Name=v.Name
	}
	log.Printf("GetAllMembers,role:\t"+role+"\n")
	return memberList
}

//登记信息
func EnrollUser( user *User) (err error){
	recordUser:=User{}
	database.DB.FirstOrCreate(&recordUser,&User{OpenId:user.OpenId})
	database.DB.Model(&recordUser).Updates(user)
	log.Printf("EnrollUser\trole:"+user.Role+"\topenid:"+user.OpenId)
	return
}

