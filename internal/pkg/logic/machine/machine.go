package machine

import (
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/rao"
)

func GetMachineList(ctx *gin.Context, req rao.GetMachineListParam) ([]*rao.MachineList, int64, error) {
	// 查询机器列表
	tx := dal.GetQuery().Machine

	conditions := make([]gen.Condition, 0)
	if req.Name != "" {
		conditions = append(conditions, tx.Name.Eq(req.Name))
	}
	if req.ServerType != 0 {
		conditions = append(conditions, tx.ServerType.Eq(req.ServerType))
	}

	// 排序
	sort := make([]field.Expr, 0, 5)
	if req.SortTag == 0 { // 默认排序(创建时间)
		sort = append(sort, tx.CreatedAt.Desc())
	}
	if req.SortTag == 1 { // 内存使用率升序
		sort = append(sort, tx.MemUsage)
	}
	if req.SortTag == 2 { // 内存使用率降序
		sort = append(sort, tx.MemUsage.Desc())
	}
	if req.SortTag == 3 { // 磁盘使用率升序
		sort = append(sort, tx.DiskUsage)
	}
	if req.SortTag == 4 { // 磁盘使用率降序
		sort = append(sort, tx.DiskUsage.Desc())
	}
	// 查询数据库
	limit := req.Size
	offset := (req.Page - 1) * req.Size
	machineList, count, err := tx.WithContext(ctx).Where(conditions...).Order(sort...).FindByPage(offset, limit)
	if err != nil {
		proof.Errorf("机器列表--获取机器列表数据失败，err:", err)
		return nil, 0, err
	}

	res := make([]*rao.MachineList, 0, len(machineList))
	for _, machineInfo := range machineList {
		machineTmp := &rao.MachineList{
			Name:              machineInfo.Name,
			CPUUsage:          machineInfo.CPUUsage,
			CPULoadOne:        machineInfo.CPULoadOne,
			CPULoadFive:       machineInfo.CPULoadFive,
			CPULoadFifteen:    machineInfo.CPULoadFifteen,
			MemUsage:          machineInfo.MemUsage,
			DiskUsage:         machineInfo.DiskUsage,
			MaxGoroutines:     machineInfo.MaxGoroutines,
			CurrentGoroutines: machineInfo.CurrentGoroutines,
			ServerType:        machineInfo.ServerType,
			Status:            machineInfo.Status,
			CreatedAt:         machineInfo.CreatedAt,
		}
		res = append(res, machineTmp)
	}

	return res, count, nil
}
