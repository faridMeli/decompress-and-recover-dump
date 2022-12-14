package brickDump

import (
	"encoding/json"
	"log"

	"github.com/faridMeli/decompress-and-recover-dump/pkg/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/pkg/internal/decompress"
	"github.com/faridMeli/decompress-and-recover-dump/pkg/internal/model"
)

type BrickExecutor struct {
	list []data.DataCompressed
}

func NewBrickExecutor(list []data.DataCompressed) *BrickExecutor {
	return &BrickExecutor{
		list: list,
	}
}

func (e *BrickExecutor) RecoverDump() map[string][][]byte {
	var bricks []model.Brick
	var otherBricks []model.BrickTabbar

	for _, data := range e.list {
		brick := decompress.DecompressBrick(data.Item.CompressedValue.B)
		if !isTabbarBrick(brick) {
			bricks = append(bricks, brick)
		} else {
			convertAndAppendBrickTababr(&otherBricks, brick)
		}
	}

	bricks = removeBricksDuplicateValues(bricks)
	otherBricks = removeTabbarBricksDuplicateValues(otherBricks)

	finalResult := make(map[string][][]byte)

	for _, brick := range bricks {
		json, err := json.Marshal(brick)
		if err != nil {
			return nil
		}
		finalResult["Brick"] = append(finalResult["Brick"], json)
	}
	for _, brick := range otherBricks {
		json, err := json.Marshal(brick)
		if err != nil {
			return nil
		}
		finalResult["TabbarBrick"] = append(finalResult["TabbarBrick"], json)
	}

	return finalResult
}

func convertAndAppendBrickTababr(bricksTabbar *[]model.BrickTabbar, brick model.Brick) {
	tb := convertBricksToTabbarBricks(brick)
	*bricksTabbar = append(*bricksTabbar, tb)
}

func convertBricksToTabbarBricks(brick model.Brick) model.BrickTabbar {
	var tb model.BrickTabbar
	tabbarJson, err := json.Marshal(brick)
	if err != nil {
		log.Fatal("Error in Marshal")
	}

	json.Unmarshal(tabbarJson, &tb)

	return tb
}

func isTabbarBrick(brick model.Brick) bool {
	a := brick.Variants[0].BrickDTO.Data["tab_id"]
	return a != nil
}

func removeBricksDuplicateValues(bricksSlice []model.Brick) []model.Brick {
	keys := make(map[string]bool)
	list := []model.Brick{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range bricksSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	return list
}

func removeTabbarBricksDuplicateValues(bricksSlice []model.BrickTabbar) []model.BrickTabbar {
	keys := make(map[string]bool)
	list := []model.BrickTabbar{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range bricksSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	return list
}
