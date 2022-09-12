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

func (e *FilterExecutor) RecoverDump() {
	var Filters []model.Filter

	for _, data := range e.list {
		Filters = append(Filters, pkg.DecompressFilter(data.Item.CompressedValue.B))
	}

	removeDuplicateValues(&Filters)
}

func removeDuplicateValues(filtersSlice *[]model.Filter) {
	keys := make(map[string]bool)
	list := []model.Filter{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range *filtersSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	*filtersSlice = list
}
