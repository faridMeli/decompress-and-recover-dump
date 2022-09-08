package pageDump

import (
	"github.com/faridMeli/decompress-and-recover-dump/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/internal/model"
	pkg "github.com/faridMeli/decompress-and-recover-dump/pkg"
)

type PageExecutor struct {
	list []data.DataCompressed
}

func NewPageExecutor(list []data.DataCompressed) *PageExecutor {
	return &PageExecutor{
		list: list,
	}
}

func (e *PageExecutor) RecoverDump() {
	var Pages []model.Page

	for _, data := range e.list {
		Pages = append(Pages, pkg.DecompressPage(data.Item.CompressedValue.B))
	}
}
