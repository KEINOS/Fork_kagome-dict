package data

import(
	"os"
	"time"
)


func dictUnidictChBytes() ([]byte, error) {
	return _dictUnidictCh, nil
}

func dictUnidictCh() (*asset, error) {
	bytes, err := dictUnidictChBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ch", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
