package data

import(
	"os"
	"time"
)


func dictUnidictDcBytes() ([]byte, error) {
	return _dictUnidictDc, nil
}

func dictUnidictDc() (*asset, error) {
	bytes, err := dictUnidictDcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.dc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
