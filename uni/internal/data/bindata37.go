package data

import(
	"os"
	"time"
)


func dictUnidictBlBytes() ([]byte, error) {
	return _dictUnidictBl, nil
}

func dictUnidictBl() (*asset, error) {
	bytes, err := dictUnidictBlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
