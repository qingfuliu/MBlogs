package logic

import (
	"goProject/dao"
	"goProject/generate"
	"goProject/models"
)

func CreateCommunity(community *models.CommunityDetail) error {
	if ok, err := dao.IfCommunityExisted(community.CommunityName); ok || err.Error() == "query failed" {
		return err
	}
	newCommunity := &models.Community{
		CommunityId:     generate.SnowFlakeUID(),
		CommunityDetail: community,
	}
	if err := dao.InsertCommunity(newCommunity); err != nil {
		return err
	}
	return nil
}
