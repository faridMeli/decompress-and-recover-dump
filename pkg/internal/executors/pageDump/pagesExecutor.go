package pageDump

import (
	"encoding/json"

	"github.com/faridMeli/decompress-and-recover-dump/pkg/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/pkg/internal/decompress"
	"github.com/faridMeli/decompress-and-recover-dump/pkg/internal/model"
)

type PageExecutor struct {
	list []data.DataCompressed
}

func NewPageExecutor(list []data.DataCompressed) *PageExecutor {
	return &PageExecutor{
		list: list,
	}
}

func (e *PageExecutor) RecoverDump() map[string][][]byte {
	var pages []model.Page

	for _, data := range e.list {
		pages = append(pages, decompress.DecompressPage(data.Item.CompressedValue.B))
	}

	removeDuplicateValues(&pages)

	finalResult := make(map[string][][]byte)

	for _, page := range pages {
		json, err := json.Marshal(page)
		if err != nil {
			return nil
		}
		finalResult["pages"] = append(finalResult["pages"], json)
	}

	return finalResult
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
