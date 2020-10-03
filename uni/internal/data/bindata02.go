package data

import(
	"os"
	"time"
)


func dictUnidictAcBytes() ([]byte, error) {
	return _dictUnidictAc, nil
}

func dictUnidictAc() (*asset, error) {
	bytes, err := dictUnidictAcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ac", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
