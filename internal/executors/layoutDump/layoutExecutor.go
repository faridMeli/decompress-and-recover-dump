package layoutDump

import (
	"github.com/faridMeli/decompress-and-recover-dump/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/internal/model"
	pkg "github.com/faridMeli/decompress-and-recover-dump/pkg"
)

type LayoutExecutor struct {
	list []data.DataCompressed
}

func NewLayoutExecutor(list []data.DataCompressed) *LayoutExecutor {
	return &LayoutExecutor{
		list: list,
	}
}

func (e *LayoutExecutor) RevoverDump() {
	var Layouts []model.Layout

	for _, data := range e.list {
		Layouts = append(Layouts, pkg.DecompressLayout(data.Item.CompressedValue.B))
	}
}
