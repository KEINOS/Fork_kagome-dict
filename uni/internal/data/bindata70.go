package data

import(
	"os"
	"time"
)


func dictUnidictCsBytes() ([]byte, error) {
	return _dictUnidictCs, nil
}

func dictUnidictCs() (*asset, error) {
	bytes, err := dictUnidictCsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
