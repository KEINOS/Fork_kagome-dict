package data

import(
	"os"
	"time"
)


func dictUnidictDgBytes() ([]byte, error) {
	return _dictUnidictDg, nil
}

func dictUnidictDg() (*asset, error) {
	bytes, err := dictUnidictDgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.dg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
