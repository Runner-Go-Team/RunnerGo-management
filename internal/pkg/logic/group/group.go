package group

import (
	"context"
	"fmt"

	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/query"
	"kp-management/internal/pkg/dal/rao"
	"kp-management/internal/pkg/packer"

	"go.mongodb.org/mongo-driver/bson"
)

func Save(ctx context.Context, req *rao.SaveGroupReq, userID int64) error {
	target := packer.TransGroupReqToTarget(req, userID)
	group := packer.TransGroupReqToGroup(req)

	collection := dal.GetMongo().Database(dal.MongoD()).Collection(consts.CollectGroup)

	return query.Use(dal.DB()).Transaction(func(tx *query.Query) error {

		if target.ID == 0 {
			if err := tx.Target.WithContext(ctx).Create(target); err != nil {
				return err
			}

			group.TargetID = target.ID
			_, err := collection.InsertOne(ctx, group)

			tx.Operation.WithContext(ctx).Create(&model.Operation{
				TeamID:   target.TeamID,
				UserID:   userID,
				Category: consts.OperationCategoryCreate,
				Name:     fmt.Sprintf("创建分组 - %s", target.Name),
			})

			return err
		}

		if _, err := tx.Target.WithContext(ctx).Omit(tx.Target.CreatedUserID).Updates(target); err != nil {
			return err
		}

		_, err := collection.UpdateOne(ctx, bson.D{{"target_id", target.ID}}, bson.M{"$set": group})

		tx.Operation.WithContext(ctx).Create(&model.Operation{
			TeamID:   target.TeamID,
			UserID:   userID,
			Category: consts.OperationCategoryCreate,
			Name:     fmt.Sprintf("修改分组 - %s", target.Name),
		})

		return err
	})
}