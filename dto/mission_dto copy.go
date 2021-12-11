package dto

import "xietong.me/ginessential/model"

type MissionDto struct {
	Name string `json:"name"`
}

func ToMissionDto(mission model.Mission) MissionDto {
	return MissionDto{
		Name: mission.File,
	}
}
