// Code generated by go-bindata.
// sources:
// README.md
// test/TEST.md
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x56\xdd\x6e\xdc\x36\x13\xbd\x8e\x9e\x62\x3e\xfb\x66\xd7\xf8\x42\xe5\xb6\x2e\x7a\x11\x34\x69\x13\xa0\x0e\x82\xd8\x41\x10\x2c\x0c\x8b\x2b\xcd\x4a\xb2\x25\x52\x20\x29\x6f\x17\x41\xde\xbd\x87\x3f\x5a\x29\xf1\x4f\x04\x18\xe0\x6a\xc4\xe1\x99\x33\x67\x0e\x7d\x4a\x95\x2e\x2d\x9b\xfb\x2c\x4b\x0b\x6a\x2d\x49\xfa\x5b\x77\x52\xd5\x34\xc8\xf2\x4e\xd6\x4c\xae\x91\x8e\x0e\x7a\xa4\x52\x2a\x6a\x95\xe3\xda\x48\xc7\x7e\xa5\xfd\x6b\x43\x83\xd1\xb7\x5c\x3a\xc2\xef\x9a\x15\x87\xa8\xf4\xb9\xc7\x9e\x95\x93\xae\xd5\x8a\xb6\x46\xef\x71\x82\xc8\x2e\x99\x69\x53\x6b\x44\xaf\x57\x8d\x73\x83\x3d\xcf\xf3\xf0\x53\x68\x53\xe7\x75\xeb\x9a\x71\x2b\x4a\xdd\xe7\x35\x9b\x2d\xab\x5b\x59\xea\xad\xcd\x13\xbe\x35\xed\xb4\xa1\x5e\x1b\x7f\x3c\x96\x7d\x48\x2e\xb2\xec\xbd\xa3\x46\x5a\x72\x7b\x8d\x68\xc5\xf6\x77\x32\x2c\xab\x16\x55\xb8\x86\x8f\x50\x2c\xed\x8c\xee\x69\xd7\x76\x6c\x0f\xd6\x71\x4f\xc8\x96\x20\xfb\x6f\xad\x07\x5b\x52\xa9\x87\x96\x2d\xed\x81\x25\x03\xd4\x97\xdb\x56\x55\xd2\xc9\x05\xde\x19\xe5\xad\x63\x1e\xf7\xac\xf2\xf9\xbb\x75\x96\x9d\x9e\xd2\x7b\x85\x6c\x5d\x97\x65\x45\xed\x69\x71\xf4\x72\xa4\x5f\x54\x57\xa0\x8e\x5d\x60\x7a\x40\x03\x08\xac\x8d\x76\x2a\x21\x41\x93\xdb\xb6\x6b\x1d\xd0\xfd\x9f\x64\x67\x51\xac\xbc\x43\x6c\x04\x1f\x20\xbf\x8d\x47\xd2\x0c\x45\x3c\x75\xfc\x23\xb0\x73\x21\x44\x11\xa0\x7f\xb6\x68\x7b\x96\x9d\x9d\xfd\x35\x13\xe5\x59\x3d\x3b\x43\xba\x02\x09\xb3\xca\xd2\xf9\x1f\x93\x7a\xc4\x07\xde\xbf\xd1\xe5\x25\x96\xab\xcd\xb5\x75\x06\x98\xbf\x9d\x7c\x7a\xfb\xfa\xcd\xc5\x5b\xd1\x57\x27\xdf\xd7\xf8\x5e\x7c\x1a\xd5\x6a\xed\xb7\xfb\xc4\x97\xb1\x9a\x63\x52\x91\xc7\xfa\x84\x6d\x68\xe3\xdb\x33\x48\xd7\x5c\x47\x40\x57\xa8\x5e\x6e\xf5\x3d\xea\x2c\x4d\x3b\x38\xf4\x05\x35\x96\x5a\xdd\xb3\x71\x51\x80\x73\x7f\xc1\x82\x5c\xd4\x0f\x29\x4f\xbb\x82\x8a\xc3\xd6\xb6\x6f\x5d\x94\x68\xa7\x4b\xd9\x2d\xe5\xb0\x6f\x58\x45\xa9\xeb\x7e\xc0\xeb\x98\x5e\x9b\xb6\x6e\x15\xbe\x84\xd0\xa1\xfd\x1e\xac\x7e\x4d\xe3\x00\xa9\x1c\x90\x1a\x63\x33\x1d\xf3\xd3\x4c\xac\xa6\xc1\x29\x1b\xcc\x14\x87\x5e\x4e\x83\xa5\x64\x0f\x14\xaa\x22\x3d\xba\x61\x74\x01\xc8\x3a\xbc\x18\x2d\x04\xee\x04\xdd\xdc\x5c\xfc\xd0\xe0\x7e\xd0\xc6\xc5\xf3\x52\x12\x71\x73\x33\x35\x25\x45\x3d\x75\xb9\xd3\x89\xd1\xec\x61\xaf\x22\xf9\x53\xc7\x7a\x39\x6c\x62\xcf\xae\x37\xd7\xdb\x83\xe3\x65\xeb\xce\x93\xee\xc4\xc5\x68\xdd\x6b\x6b\xd9\xad\x16\xd1\xf5\x83\xce\x42\x3c\x7f\x6a\xb5\x6b\xeb\x09\x93\x3b\x0c\x4c\xe9\xa8\x18\x41\x46\x33\x82\x98\x6f\xd9\x8b\xb7\xaa\x1a\x34\x6c\x84\xd2\x13\x61\x60\x91\xe7\x74\x8c\xa1\xec\xae\x45\x73\x94\x5f\xad\x2a\xde\xc9\xb1\x73\xe7\x14\x66\x66\x9d\xbd\xf8\x6c\x39\x89\x29\x3e\x5b\xad\xbb\xb0\x40\x8e\x2f\x0d\x83\x6e\xe3\x37\x7a\x42\x17\xa2\x5b\x24\xda\x61\x8e\x18\x89\x00\xd2\xd2\xe2\x99\x94\xec\x13\x79\x09\x4e\xaa\xb4\xc1\x82\x7e\x30\xb7\xec\xc5\x3b\x6d\xdd\x72\xf3\xb2\x12\xbf\xb9\xf1\xf1\x41\xa2\x3b\x7a\x17\x24\xf0\xee\xea\xea\x63\xaa\x0b\x00\x67\x34\x27\x27\x80\xf2\xd1\xb7\xf1\x99\x6c\xa9\xcd\xbf\xce\xf6\xdb\xab\x57\xaf\x22\x47\xef\xda\xba\xe9\xf0\x17\x6c\xee\x19\x8e\xec\x01\x35\xfd\x4b\xcd\xf2\xf3\x39\x1f\x3a\xe7\xa9\xba\x0c\x1f\x5d\xba\x03\x06\xe4\x31\x7c\x36\x44\x52\x46\xcf\xd6\xf3\x59\x2b\x69\xee\xc6\x0e\xb6\xf9\xfd\x28\xa2\x2b\xee\x61\x81\x0e\x26\xf4\xbf\x0d\x06\x8b\x59\xd9\x46\xbb\xd9\x80\x8d\xdc\x8b\xe8\x66\x38\xc2\xc0\x0b\x50\xb9\x7b\xd2\x57\xf3\x5e\x82\x1a\x93\xcf\x99\xc4\xa0\x6a\xd8\x74\x12\xa6\x9f\xf6\xe4\xf7\xfe\xde\x8a\xb8\xc8\x25\x0c\xb0\x55\x2a\x0c\x5b\x4c\x75\xc9\x36\x9f\x5e\x8b\xc6\xf5\x5d\x21\xe8\xeb\xcf\xf3\x8d\xd9\x9c\xb7\x46\x3b\xdf\xb7\xb6\x89\x97\x14\x82\xd1\xd3\x8b\xd9\xa4\x0a\xb2\xfa\x78\xbf\x9a\x51\x85\x96\x16\xdb\xb1\xed\x2a\xf8\x61\x91\x9c\x45\x78\x62\x4e\x29\x72\x4f\xcb\x7e\x66\xd9\x17\xd8\x08\x1c\x22\x66\xde\x1c\x69\x16\xb7\x76\xa6\xec\xf8\xf6\xd6\xfa\x9b\x76\x1d\xac\xb2\xaa\x1e\xed\xcd\x64\x62\x47\x5f\x15\xc7\x9b\x69\x2f\xd5\xfc\xcf\x00\xfc\x20\x82\x5d\x28\xa2\xf0\x7c\x85\x97\x89\x5d\x8e\x73\x5f\x84\x03\xd5\x21\xa9\x25\x9b\xfd\x18\x97\x78\xd9\x84\x50\x92\x73\xd0\x0f\x1c\xb5\xd1\xfb\x52\x5a\xae\x50\x12\x1b\x7e\xb2\x94\xe4\x74\x79\xc5\xbd\xce\xd7\x99\xf7\x4f\x3f\x6f\x0f\xf7\x95\x95\xc2\x0e\xaf\x92\xae\xdd\x1a\x69\x70\x8d\xe6\x4b\xb2\xd6\xb4\xfa\xa0\x1d\x9f\xd3\x3f\x7a\x0f\x55\xe1\xe8\x60\xc6\x16\x56\x0b\x38\x86\xd1\xd3\x12\x59\x83\x4e\x2a\x69\x1b\x86\x01\xfd\x17\x00\x00\xff\xff\xa8\xa0\x61\x5e\x44\x09\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.md",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.md", size: 2372, mode: os.FileMode(420), modTime: time.Unix(1463135742, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testTestMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcd\x41\x0a\xc2\x30\x10\x85\xe1\x7d\x4e\xf1\xc0\xad\xf4\x0e\x05\x0b\xba\x70\x23\x5e\x60\x6c\x9e\x6d\x31\x99\x91\x64\x24\x78\x7b\x03\xae\xba\x7c\xf0\x3e\xfe\x03\x46\x35\x5f\x59\xe0\xac\x8e\x2c\xe5\x15\xad\x29\xde\xb2\x30\x84\x33\x53\xb2\x23\x2e\x90\x0c\x41\xe5\x6c\x1a\xf7\x1f\x7c\x2a\x23\x9e\xf6\xf7\x9b\x2e\x43\x08\xf7\xf2\x45\x6d\x9b\xcf\x6b\xdf\x78\xd0\x1b\xa9\xc8\x84\x74\xdd\x5b\xb8\x4d\xe3\xe9\x3a\x0d\xbf\x00\x00\x00\xff\xff\x33\x2e\xeb\x5f\x7c\x00\x00\x00")

func testTestMdBytes() ([]byte, error) {
	return bindataRead(
		_testTestMd,
		"test/TEST.md",
	)
}

func testTestMd() (*asset, error) {
	bytes, err := testTestMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/TEST.md", size: 124, mode: os.FileMode(420), modTime: time.Unix(1460463742, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"README.md":    readmeMd,
	"test/TEST.md": testTestMd,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"README.md": &bintree{readmeMd, map[string]*bintree{}},
	"test": &bintree{nil, map[string]*bintree{
		"TEST.md": &bintree{testTestMd, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
