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

func ModifyCommunity(community *models.Community) error {
	old, err := dao.CommunityDetailQuery(community.CommunityName)
	if err != nil {
		return dao.ErrorQueryFailed
	}
	if old != nil && old.CommunityId != community.CommunityId {
		return dao.ErrorNotExisted
	}
	if err = dao.CommunityModify(community); err != nil {
		return dao.ErrorModifyFailed
	}
	return nil
}
