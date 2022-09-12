package shortcutDump

import (
	"strings"

	"github.com/faridMeli/decompress-and-recover-dump/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/internal/model"
	pkg "github.com/faridMeli/decompress-and-recover-dump/pkg"
)

type ShortcutExecutor struct {
	list []data.DataCompressed
}

func NewShortcutExecutor(list []data.DataCompressed) *ShortcutExecutor {
	return &ShortcutExecutor{
		list: list,
	}
}

func (e *ShortcutExecutor) RecoverDump() {
	count := 0
	var shortcuts []model.Shortcut
	var collections []model.Collection

	for _, data := range e.list {
		count++
		if strings.Contains(data.Item.Key.S, "collection") {
			collections = append(collections, pkg.DecompressCollection(data.Item.CompressedValue.B))
		} else {
			shortcuts = append(shortcuts, pkg.DecompressShortcut(data.Item.CompressedValue.B))
		}
	}

	removeShortcutsDuplicateValues(&shortcuts)
	removeCollectionsDuplicateValues(&collections)
}

func removeShortcutsDuplicateValues(shortcutsSlice *[]model.Shortcut) {
	keys := make(map[string]bool)
	list := []model.Shortcut{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range *shortcutsSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	*shortcutsSlice = list
}

func removeCollectionsDuplicateValues(collectionsSlice *[]model.Collection) {
	keys := make(map[string]bool)
	list := []model.Collection{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, brick := range *collectionsSlice {
		if _, value := keys[brick.ID]; !value {
			keys[brick.ID] = true
			list = append(list, brick)
		}
	}
	*collectionsSlice = list
}
