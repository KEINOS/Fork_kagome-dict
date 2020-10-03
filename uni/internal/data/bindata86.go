package data

import(
	"os"
	"time"
)


func dictUnidictDiBytes() ([]byte, error) {
	return _dictUnidictDi, nil
}

func dictUnidictDi() (*asset, error) {
	bytes, err := dictUnidictDiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.di", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
