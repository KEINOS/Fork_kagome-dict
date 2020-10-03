package data

import(
	"os"
	"time"
)


func dictUnidictAiBytes() ([]byte, error) {
	return _dictUnidictAi, nil
}

func dictUnidictAi() (*asset, error) {
	bytes, err := dictUnidictAiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ai", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
