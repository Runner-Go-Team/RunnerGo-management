package plan

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/mao"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/query"
)

func ImportScene(ctx context.Context, userID, planID int64, targetIDList []int64) ([]*model.Target, error) {
	retID := make([]*model.Target, 0)

	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectFlow)
	collectionTask := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectTask)

	err := dal.GetQuery().Transaction(func(tx *query.Query) error {

		targets, err := tx.Target.WithContext(ctx).Where(tx.Target.ID.In(targetIDList...)).Find()
		if err != nil {
			return err
		}

		memo := make(map[int64]int64)
		for len(targets) > 0 {
			for i, t := range targets {
				if t.ParentID == 0 {
					oldID := t.ID
					t.ID = 0
					t.PlanID = planID
					t.CreatedUserID = userID
					t.RecentUserID = userID
					t.Source = consts.TargetSourcePlan
					if err := tx.Target.WithContext(ctx).Create(t); err != nil {
						return err
					}

					memo[oldID] = t.ID

					if t.TargetType == consts.TargetTypeScene {
						var flow mao.Flow
						err := collection.FindOne(ctx, bson.D{{"scene_id", oldID}}).Decode(&flow)
						if err != nil && err != mongo.ErrNoDocuments {
							return err
						}
						if err != mongo.ErrNoDocuments {
							var ns *mao.Node
							if err := bson.Unmarshal(flow.Nodes, &ns); err != nil {
								return err
							}
							for _, n := range ns.Nodes {
								n.Data.From = "plan"
							}
							nodes, err := bson.Marshal(ns)
							if err != nil {
								return err
							}

							flow.SceneID = t.ID
							flow.Nodes = nodes
							if _, err := collection.InsertOne(ctx, flow); err != nil {
								return err
							}
						}

						var task mao.Task
						err = collectionTask.FindOne(ctx, bson.D{{"scene_id", oldID}}).Decode(&t)
						if err != nil && err != mongo.ErrNoDocuments {
							return err
						}
						if err != mongo.ErrNoDocuments {
							task.PlanID = planID
							task.SceneID = t.ID
							if _, err := collection.InsertOne(ctx, task); err != nil {
								return err
							}
						}

						retID = append(retID, t)
					}

					if i >= len(targets) {
						targets = targets[:len(targets)-1]
					} else if i == 0 {
						targets = targets[1:]
					} else {
						targets = append(targets[:i], targets[i+1:]...)
					}
				}

				if newID, ok := memo[t.ParentID]; ok {
					oldID := t.ID
					t.ID = 0
					t.ParentID = newID
					t.PlanID = planID
					t.CreatedUserID = userID
					t.RecentUserID = userID
					t.Source = consts.TargetSourcePlan
					if err := tx.Target.WithContext(ctx).Create(t); err != nil {
						return err
					}
					memo[oldID] = t.ID

					if t.TargetType == consts.TargetTypeScene {
						var flow mao.Flow
						err := collection.FindOne(ctx, bson.D{{"scene_id", oldID}}).Decode(&flow)
						if err != nil && err != mongo.ErrNoDocuments {
							return err
						}
						if err != mongo.ErrNoDocuments {
							var ns *mao.Node
							if err := bson.Unmarshal(flow.Nodes, &ns); err != nil {
								return err
							}
							for _, n := range ns.Nodes {
								n.Data.From = "plan"
							}
							nodes, err := bson.Marshal(ns)
							if err != nil {
								return err
							}

							flow.SceneID = t.ID
							flow.Nodes = nodes
							if _, err := collection.InsertOne(ctx, flow); err != nil {
								return err
							}
						}

						var task mao.Task
						err = collectionTask.FindOne(ctx, bson.D{{"scene_id", oldID}}).Decode(&t)
						if err != nil && err != mongo.ErrNoDocuments {
							return err
						}
						if err != mongo.ErrNoDocuments {
							task.PlanID = planID
							task.SceneID = t.ID
							if _, err := collection.InsertOne(ctx, task); err != nil {
								return err
							}
						}

						retID = append(retID, t)
					}

					if i >= len(targets) {
						targets = targets[:len(targets)-1]
					} else if i == 0 {
						targets = targets[1:]
					} else {
						targets = append(targets[:i], targets[i+1:]...)
					}
				}

			}
		}

		var sceneIDs []int64
		for oldSceneID, _ := range memo {
			sceneIDs = append(sceneIDs, oldSceneID)
		}

		v, err := tx.Variable.WithContext(ctx).Where(tx.Variable.SceneID.In(sceneIDs...)).Find()
		if err != nil {
			return err
		}
		var variables []*model.Variable
		for _, variable := range v {
			if newSceneID, ok := memo[variable.SceneID]; ok {
				variable.ID = 0
				variable.SceneID = newSceneID
				variables = append(variables, variable)
			}
		}
		if len(variables) > 0 {
			if err := tx.Variable.WithContext(ctx).CreateInBatches(variables, 5); err != nil {
				return err
			}
		}

		vi, err := tx.VariableImport.WithContext(ctx).Where(tx.VariableImport.SceneID.In(sceneIDs...)).Find()
		if err != nil {
			return err
		}
		var variablesImports []*model.VariableImport
		for _, variableImport := range vi {
			if newSceneID, ok := memo[variableImport.SceneID]; ok {
				variableImport.ID = 0
				variableImport.SceneID = newSceneID
				variablesImports = append(variablesImports, variableImport)
			}
		}
		if len(variablesImports) > 0 {
			if err := tx.VariableImport.WithContext(ctx).CreateInBatches(variablesImports, 5); err != nil {
				return err
			}
		}

		return nil

	})

	return retID, err
}
