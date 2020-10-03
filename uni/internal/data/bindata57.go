package data

import(
	"os"
	"time"
)


func dictUnidictCfBytes() ([]byte, error) {
	return _dictUnidictCf, nil
}

func dictUnidictCf() (*asset, error) {
	bytes, err := dictUnidictCfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
