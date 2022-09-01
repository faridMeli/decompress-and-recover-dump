package filterDump

import (
	"github.com/faridMeli/decompress-and-recover-dump/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/internal/model"
	pkg "github.com/faridMeli/decompress-and-recover-dump/pkg"
)

type FilterExecutor struct {
	list []data.DataCompressed
}

func NewFilterExecutor(list []data.DataCompressed) *FilterExecutor {
	return &FilterExecutor{
		list: list,
	}
}

func (e *FilterExecutor) RevoverDump() {
	var Filters []model.Filter

	for _, data := range e.list {
		Filters = append(Filters, pkg.DecompressFilter(data.Item.CompressedValue.B))
	}
}
