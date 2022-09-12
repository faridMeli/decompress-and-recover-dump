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

func (e *LayoutExecutor) RecoverDump() {
	var Layouts []model.Layout

	for _, data := range e.list {
		Layouts = append(Layouts, pkg.DecompressLayout(data.Item.CompressedValue.B))
	}

	removeDuplicateValues(&Layouts)
}

func removeDuplicateValues(layoutsSlice *[]model.Layout) {
	keys := make(map[string]bool)
	list := []model.Layout{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range *layoutsSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	*layoutsSlice = list
}
