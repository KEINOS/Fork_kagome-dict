package data

import(
	"os"
	"time"
)


func dictUnidictAmBytes() ([]byte, error) {
	return _dictUnidictAm, nil
}

func dictUnidictAm() (*asset, error) {
	bytes, err := dictUnidictAmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.am", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
