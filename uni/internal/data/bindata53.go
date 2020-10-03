package data

import(
	"os"
	"time"
)


func dictUnidictCbBytes() ([]byte, error) {
	return _dictUnidictCb, nil
}

func dictUnidictCb() (*asset, error) {
	bytes, err := dictUnidictCbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
