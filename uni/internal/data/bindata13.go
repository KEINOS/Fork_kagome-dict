package data

import(
	"os"
	"time"
)


func dictUnidictAnBytes() ([]byte, error) {
	return _dictUnidictAn, nil
}

func dictUnidictAn() (*asset, error) {
	bytes, err := dictUnidictAnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.an", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
