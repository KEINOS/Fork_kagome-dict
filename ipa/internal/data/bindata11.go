package data

import(
	"os"
	"time"
)


func dictIpadictAlBytes() ([]byte, error) {
	return _dictIpadictAl, nil
}

func dictIpadictAl() (*asset, error) {
	bytes, err := dictIpadictAlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.al", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600787007, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
