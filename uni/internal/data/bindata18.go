package data

import(
	"os"
	"time"
)


func dictUnidictAsBytes() ([]byte, error) {
	return _dictUnidictAs, nil
}

func dictUnidictAs() (*asset, error) {
	bytes, err := dictUnidictAsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.as", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
