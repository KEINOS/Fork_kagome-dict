package data

import(
	"os"
	"time"
)


func dictUnidictCpBytes() ([]byte, error) {
	return _dictUnidictCp, nil
}

func dictUnidictCp() (*asset, error) {
	bytes, err := dictUnidictCpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
