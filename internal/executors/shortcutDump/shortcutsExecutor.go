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

func (e *ShortcutExecutor) RollbackDump() {
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
}
