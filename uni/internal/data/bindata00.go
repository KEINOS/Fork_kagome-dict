package data

import(
	"os"
	"time"
)


func dictUnidictAaBytes() ([]byte, error) {
	return _dictUnidictAa, nil
}

func dictUnidictAa() (*asset, error) {
	bytes, err := dictUnidictAaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.aa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
