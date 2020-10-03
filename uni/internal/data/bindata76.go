package data

import(
	"os"
	"time"
)


func dictUnidictCyBytes() ([]byte, error) {
	return _dictUnidictCy, nil
}

func dictUnidictCy() (*asset, error) {
	bytes, err := dictUnidictCyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cy", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
