package data

import(
	"os"
	"time"
)


func dictUnidictCaBytes() ([]byte, error) {
	return _dictUnidictCa, nil
}

func dictUnidictCa() (*asset, error) {
	bytes, err := dictUnidictCaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ca", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
