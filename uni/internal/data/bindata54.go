package data

import(
	"os"
	"time"
)


func dictUnidictCcBytes() ([]byte, error) {
	return _dictUnidictCc, nil
}

func dictUnidictCc() (*asset, error) {
	bytes, err := dictUnidictCcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
