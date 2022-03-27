package dao

import (
	"goProject/models"
)

func InsertCommunity(community *models.Community) (err error) {
	err = dbConn.Create(community).Error
	return
}

func IfCommunityExisted(communityName string) (bool, error) {
	temp := makeStruct(map[string]interface{}{
		"community_name": communityName,
		"table_name":     "community",
	})
	return IfIsExisted(temp)
}

func BatchCommunityQuery(b *models.BatchCommunities) (rs []models.Community, err error) {
	rs = make([]models.Community, 0)
	if err = dbConn.Order(b.Order).Offset(b.PageSize * (b.Page - 1)).Limit(b.PageSize).Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func CommunityDetailQuery(communityName string) (community *models.Community, err error) {
	community = &models.Community{}
	if err = dbConn.Where("CommunityName=?", communityName).Find(community).Error; err != nil {
		return
	}
	return
}
