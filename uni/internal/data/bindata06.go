package data

import(
	"os"
	"time"
)


func dictUnidictAgBytes() ([]byte, error) {
	return _dictUnidictAg, nil
}

func dictUnidictAg() (*asset, error) {
	bytes, err := dictUnidictAgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ag", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
