package brickDump

import (
	"github.com/faridMeli/decompress-and-recover-dump/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/internal/model"
	pkg "github.com/faridMeli/decompress-and-recover-dump/pkg"
)

type BrickExecutor struct {
	list []data.DataCompressed
}

func NewBrickExecutor(list []data.DataCompressed) *BrickExecutor {
	return &BrickExecutor{
		list: list,
	}
}

func (e *BrickExecutor) RecoverDump() {
	var bricks []model.Brick
	//var collections []model.Collection

	for _, data := range e.list {
		// if strings.Contains(data.Item.Key.S, "collection") {
		// 	collections = append(collections, pkg.DecompressCollection(data.Item.CompressedValue.B))
		// } else {
		bricks = append(bricks, pkg.DecompressBrick(data.Item.CompressedValue.B))
		//}
	}
	return
}
