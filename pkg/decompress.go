package pkg

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"github.com/faridMeli/decompress-and-recover-dump/internal/model"
)

func DecompressShortcut(strCompress string) model.Shortcut {
	strDescompress := stringDecompress(strCompress)
	res := model.Shortcut{}
	json.Unmarshal([]byte(strDescompress), &res)
	return res
}

func DecompressBrick(strCompress string) model.Brick {
	strDescompress := stringDecompress(strCompress)
	res := model.Brick{}
	json.Unmarshal([]byte(strDescompress), &res)
	return res
}

func DecompressCollection(strCompress string) model.Collection {
	strDescompress := stringDecompress(strCompress)
	res := model.Collection{}
	json.Unmarshal([]byte(strDescompress), &res)
	return res
}

func stringDecompress(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return string(s)
}