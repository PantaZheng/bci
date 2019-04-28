package main

import(
	"fmt"
	"github.com/pantazheng/bci/models"
	"testing"
)

func TestMission(t *testing.T) {
	mission:=new(models.MissionJson)
	mission.Name="mission_test"
	mission.StartTime=mission.Name+"StartTime"
	mission.EndTime=mission.Name+"EndTime"
	mission.Content=mission.Name+"Content"
	mission.Participants =[]*models.UserBriefJson{{ID: 1},
		{ID:2},}
	res,err:=models.MissionCreate(mission)
	fmt.Println(res)
	fmt.Println(err)
}