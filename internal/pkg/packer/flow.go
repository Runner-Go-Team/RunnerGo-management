package packer

import (
	"github.com/go-omnibus/proof"
	"go.mongodb.org/mongo-driver/bson"

	"kp-management/internal/pkg/dal/mao"
	"kp-management/internal/pkg/dal/rao"
)

func TransSaveFlowReqToMaoFlow(req *rao.SaveFlowReq) *mao.Flow {

	nodes, err := bson.Marshal(mao.Node{Nodes: req.Nodes})
	if err != nil {
		proof.Errorf("flow.nodes bson marshal err %w", err)
	}

	edges, err := bson.Marshal(mao.Edge{Edges: req.Edges})
	if err != nil {
		proof.Errorf("flow.edges bson marshal err %w", err)
	}

	return &mao.Flow{
		SceneID: req.SceneID,
		TeamID:  req.TeamID,
		Version: req.Version,
		Nodes:   nodes,
		Edges:   edges,
		//MultiLevelNodes: req.MultiLevelNodes,
	}
}

func TransMaoFlowToRaoGetFowResp(f *mao.Flow) *rao.GetFlowResp {

	var n mao.Node
	if err := bson.Unmarshal(f.Nodes, &n); err != nil {
		proof.Errorf("flow.nodes bson unmarshal err %w", err)
	}

	var e mao.Edge
	if err := bson.Unmarshal(f.Edges, &e); err != nil {
		proof.Errorf("flow.edges bson unmarshal err %w", err)
	}

	return &rao.GetFlowResp{
		SceneID: f.SceneID,
		TeamID:  f.TeamID,
		Version: f.Version,
		Nodes:   n.Nodes,
		Edges:   e.Edges,
		//MultiLevelNodes: f.MultiLevelNodes,
	}
}

func TransMaoFlowsToRaoFlows(flows []*mao.Flow) []*rao.Flow {
	ret := make([]*rao.Flow, 0)
	for _, f := range flows {
		var n mao.Node
		if err := bson.Unmarshal(f.Nodes, &n); err != nil {
			proof.Errorf("flow.nodes bson unmarshal err %w", err)
		}

		var e mao.Edge
		if err := bson.Unmarshal(f.Edges, &e); err != nil {
			proof.Errorf("flow.edges bson unmarshal err %w", err)
		}

		ret = append(ret, &rao.Flow{
			SceneID: f.SceneID,
			TeamID:  f.TeamID,
			Version: f.Version,
			Nodes:   n.Nodes,
			Edges:   e.Edges,
			//MultiLevelNodes: nil,
		})
	}
	return ret
}
