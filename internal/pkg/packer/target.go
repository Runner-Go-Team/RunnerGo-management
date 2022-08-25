package packer

import (
	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/rao"
)

func TransFolderReqToTarget(folder *rao.SaveFolderReq, userID int64) *model.Target {
	return &model.Target{
		ID:            folder.TargetID,
		TeamID:        folder.TeamID,
		TargetType:    consts.TargetTypeFolder,
		Name:          folder.Name,
		ParentID:      folder.ParentID,
		Method:        folder.Method,
		Sort:          folder.Sort,
		TypeSort:      folder.TypeSort,
		Status:        1,
		Version:       folder.Version,
		CreatedUserID: userID,
		RecentUserID:  userID,
	}
}

func TransTargetReqToTarget(target *rao.CreateTargetReq, userID int64) *model.Target {
	return &model.Target{
		ID:            target.TargetID,
		TeamID:        target.TeamID,
		TargetType:    consts.TargetTypeAPI,
		Name:          target.Name,
		ParentID:      target.ParentID,
		Method:        target.Method,
		Sort:          target.Sort,
		TypeSort:      target.TypeSort,
		Status:        1,
		Version:       target.Version,
		CreatedUserID: userID,
		RecentUserID:  userID,
	}
}

func TransTargetToFolderAPI(targets []*model.Target) []*rao.FolderAPI {
	ret := make([]*rao.FolderAPI, 0)
	for _, t := range targets {
		ret = append(ret, &rao.FolderAPI{
			TargetID:      t.ID,
			TeamID:        t.TeamID,
			TargetType:    t.TargetType,
			Name:          t.Name,
			ParentID:      t.ParentID,
			Method:        t.Method,
			Sort:          t.Sort,
			TypeSort:      t.TypeSort,
			Version:       t.Version,
			CreatedUserID: t.CreatedUserID,
			RecentUserID:  t.RecentUserID,
		})
	}
	return ret
}

func TransTargetToGroupScene(targets []*model.Target) []*rao.GroupScene {
	ret := make([]*rao.GroupScene, 0)
	for _, t := range targets {
		ret = append(ret, &rao.GroupScene{
			TargetID:      t.ID,
			TeamID:        t.TeamID,
			TargetType:    t.TargetType,
			Name:          t.Name,
			ParentID:      t.ParentID,
			Method:        t.Method,
			Sort:          t.Sort,
			TypeSort:      t.TypeSort,
			Version:       t.Version,
			CreatedUserID: t.CreatedUserID,
			RecentUserID:  t.RecentUserID,
		})
	}
	return ret
}

func TransGroupReqToTarget(group *rao.SaveGroupReq, userID int64) *model.Target {
	return &model.Target{
		ID:            group.TargetID,
		TeamID:        group.TeamID,
		TargetType:    consts.TargetTypeGroup,
		Name:          group.Name,
		ParentID:      group.ParentID,
		Method:        group.Method,
		Sort:          group.Sort,
		TypeSort:      group.TypeSort,
		Status:        1,
		Version:       group.Version,
		CreatedUserID: userID,
		RecentUserID:  userID,
	}
}

func TransSceneReqToTarget(scene *rao.SaveSceneReq, userID int64) *model.Target {
	return &model.Target{
		ID:            scene.TargetID,
		TeamID:        scene.TeamID,
		TargetType:    consts.TargetTypeScene,
		Name:          scene.Name,
		ParentID:      scene.ParentID,
		Method:        scene.Method,
		Sort:          scene.Sort,
		TypeSort:      scene.TypeSort,
		Status:        1,
		Version:       scene.Version,
		CreatedUserID: userID,
		RecentUserID:  userID,
	}
}
