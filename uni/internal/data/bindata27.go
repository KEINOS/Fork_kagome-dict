package data

import(
	"os"
	"time"
)


func dictUnidictBbBytes() ([]byte, error) {
	return _dictUnidictBb, nil
}

func dictUnidictBb() (*asset, error) {
	bytes, err := dictUnidictBbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
