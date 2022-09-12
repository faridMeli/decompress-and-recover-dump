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
	var pages []model.Page

	for _, data := range e.list {
		pages = append(pages, pkg.DecompressPage(data.Item.CompressedValue.B))
	}

	removeDuplicateValues(&pages)
}

func removeDuplicateValues(pagesSlice *[]model.Page) {
	keys := make(map[string]bool)
	list := []model.Page{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range *pagesSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	*pagesSlice = list
}
