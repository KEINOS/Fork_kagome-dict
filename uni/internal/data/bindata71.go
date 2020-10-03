package data

import(
	"os"
	"time"
)


func dictUnidictCtBytes() ([]byte, error) {
	return _dictUnidictCt, nil
}

func dictUnidictCt() (*asset, error) {
	bytes, err := dictUnidictCtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ct", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
