package data

import(
	"os"
	"time"
)


func dictUnidictBxBytes() ([]byte, error) {
	return _dictUnidictBx, nil
}

func dictUnidictBx() (*asset, error) {
	bytes, err := dictUnidictBxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
