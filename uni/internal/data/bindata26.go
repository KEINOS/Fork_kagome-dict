package data

import(
	"os"
	"time"
)


func dictUnidictBaBytes() ([]byte, error) {
	return _dictUnidictBa, nil
}

func dictUnidictBa() (*asset, error) {
	bytes, err := dictUnidictBaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ba", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
