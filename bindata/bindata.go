// Code generated by go-bindata.
// sources:
// etc/doc.html
// DO NOT EDIT!

package bindata

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

var _etcDocHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x59\xdd\x72\xdb\xb8\x15\xbe\xcf\x53\x60\xb5\xdb\xb1\x9c\x88\xa4\x65\x25\x8e\x23\xcb\x9e\xcd\x3a\xf6\xa6\x33\x99\x24\x53\x67\xa7\xd3\xe9\xee\x05\x44\x82\x22\x6c\x08\x60\x09\xd0\xb6\x92\x7a\xa6\xaf\xd1\xd7\xeb\x93\xf4\x1c\x40\x94\xc1\x3f\xc9\x99\x96\x99\x91\x2c\x00\xfc\xce\x2f\x0e\xbe\x83\xcc\x7e\x78\xf7\xe9\xfc\xcb\xdf\x3e\x5f\x90\xcc\x2c\xc5\xd9\xb3\x19\x7e\x11\x41\xe5\xe2\x74\xc0\xe4\x00\x07\x18\x4d\xce\x9e\x11\x78\x66\x4b\x66\x28\x89\x33\x5a\x68\x66\x4e\x07\xbf\x7d\xb9\x0c\x8e\x07\xeb\x29\xc3\x8d\x60\x67\x89\x8a\x35\x2b\x6e\x67\x91\xfb\xe9\xa6\xb4\x59\x55\x7f\xe3\x13\x3d\xdf\xfc\x49\x9e\x93\x5f\x0b\x9e\x70\xbd\xf4\x87\xde\x12\xcd\x97\xb9\x60\x23\x52\x30\x9d\x2b\xa9\xf9\x2d\xfc\x4d\x65\x42\x32\xf8\x58\x91\xf3\xab\x2b\xb2\x80\xd7\xc8\x7c\x45\x7e\x8e\xd5\x7c\xd5\x78\x3f\x33\x26\xd7\xd3\x28\x5a\x70\x93\x95\xf3\x30\x56\xcb\x68\xbd\x2a\x5a\xb4\xa4\x45\xcf\x3c\xc5\xc8\x67\x90\x08\xfa\x33\xa2\xd5\x12\x3e\xa8\xe4\x66\x85\x6b\xaa\x25\x21\x02\x8c\x1e\x7f\x96\xb0\x80\x7c\x7b\x84\x83\x27\xb8\x63\xf3\x1b\x6e\x82\xb9\xba\x0f\x34\xff\xca\xe5\x62\x4a\xe6\xaa\x48\x58\x81\x43\x27\xf5\xb5\x4b\xf5\xf5\x49\x0b\x77\xad\x79\xa8\x59\x71\xc5\x0c\x29\x73\x67\x43\x51\x0a\xa6\x89\x51\x64\xa1\x6e\x59\x21\x89\xc9\x98\x73\x5e\xd3\xaa\x86\x19\xe0\xa6\x5c\xd0\x15\x88\x13\x2a\xbe\xa9\x6b\x13\x0b\x46\x0b\x54\xc4\x64\x9d\x2a\x38\xbc\x2e\xe7\xa4\x42\x51\x33\x25\x82\xa5\xa6\x0e\x79\xc7\x13\x93\x4d\xc9\xf8\xe0\xe0\x4f\xf5\x89\x9c\x26\x89\x35\x7b\x7c\x90\xf7\x1a\xfc\x25\xe3\x9a\x30\xa9\xcb\x02\x6d\x05\x0b\x55\x69\x58\x41\x16\xa5\x81\x2f\x4d\x68\xc1\x08\xfb\x47\x49\x05\xfa\x01\xa7\x87\x89\x2a\xe7\x82\x25\xfb\x84\x4b\xf9\xb8\x30\x6c\x3b\xc5\x1a\x31\x4d\x79\xa1\x4d\x10\x67\x5c\x34\xdd\xb4\x56\x2f\x40\x8b\xa6\xe4\xb0\x4f\x47\x1f\x4c\xd0\x1d\x58\x05\x5f\x64\xdb\xc0\xc0\xe0\x8f\x4c\x1b\x96\xd8\x40\x82\x79\xa2\x80\x2d\xba\x82\xdd\x01\x89\xbb\x06\x01\x33\x55\xb9\xc8\x46\x90\x04\xe0\x6d\xf3\x9f\x7f\xfd\x5b\x13\x59\xde\x30\x02\x21\xf1\x8d\xb4\x31\xfa\x3e\x23\x0f\xba\x2d\xf4\x90\x9e\x6a\xe1\x56\x24\xb4\xad\xa6\xd3\x59\x67\x46\x55\x90\x46\xe5\x4f\x01\xf4\x54\xdb\x8e\x07\xc9\x6d\xd4\xb2\x0f\x12\x42\xf0\x01\x36\x59\xce\x14\xd4\x29\xe7\x58\xbb\xb1\x5c\x1e\x45\x55\x10\xa0\xf4\xc0\xbc\x61\x62\x05\x89\x46\xa0\x74\xaa\x12\xd7\xab\x94\xdc\xd1\x95\xae\x05\x42\xaa\xa0\x4a\x57\xab\x56\xa3\xc8\xf8\xf3\x9d\x0a\x83\xa6\xe4\x07\x28\x9b\xaa\x30\x54\x9a\x3e\xa5\xff\x5a\xd0\x3c\x47\xcd\xa8\x01\x75\x96\xf4\x9e\x2f\xcb\xa5\xdb\x7c\x04\xb6\x90\xca\x0d\x57\x12\xf6\x89\xaf\xd9\x1d\xbc\xd3\x2a\x7c\xf8\xd3\xcd\xd4\xb5\x01\xc8\x60\xbd\x97\xdf\xbc\x3e\xce\x1b\x65\x6c\x49\x8b\x05\x97\xa8\x2a\x2d\x8d\xea\x55\xd2\xaa\x13\x43\xa8\x34\xc3\xec\x86\x14\xb6\xa9\xad\x33\x30\x0e\x8f\x00\x02\x95\x4c\x83\xa2\x98\xd2\xcb\x39\x2b\x60\x27\x50\x4d\xd2\x82\xc6\xa8\xbe\xf6\xcf\x82\x4b\x55\x10\x76\x4f\x31\x0c\x53\x92\xc2\x0f\xea\xaa\x9f\x8d\xfc\x38\x9a\x90\xa1\x92\x18\x3b\x5e\x40\x31\x80\xb8\x60\x14\x73\xa8\x17\xd2\x38\xaf\x8c\x7c\x30\x7b\x2c\xad\x08\xb8\x10\x3e\xad\x7e\xa7\x83\xbb\x60\x1c\x4c\x06\x55\x61\x61\x82\x2d\xe1\xdd\xae\x42\x72\x97\x29\x3c\xd2\xaa\x5f\xf0\xda\xb8\xe1\xbb\xce\x1a\xd8\x2e\x22\x19\x15\x69\x0d\xe7\xb0\x1b\xe7\xd5\x76\x18\xb0\x3b\xb0\x76\xd7\xb0\x26\xdd\x58\x93\x49\x38\x99\x4c\x0e\xb7\x02\x9a\x3b\xe5\x00\xb5\x87\x78\xd8\x87\x78\x74\x14\x1e\x1d\x1d\xbd\xda\xa9\x22\x94\xed\x02\xb2\x7e\xd4\x35\x97\xaa\x12\x12\xa2\xa6\xff\xcb\x6e\x69\x87\xdb\x05\x99\xac\x60\x1b\x51\xba\x25\xcb\x4d\x3b\x69\xbe\x71\x93\x3e\x71\xaf\x77\xdb\x95\xf2\xb4\xa1\xfa\xab\x1e\xd5\xb7\x87\x11\xbd\x6e\xb1\xea\x5e\xef\x01\x7b\xb9\x03\xcc\x19\xda\x84\x9b\xf4\xc1\x1d\x6d\x87\x43\x8f\xb5\xd1\x5e\xf6\xa1\x1d\x6f\x47\x5b\x28\x91\x30\x19\xe8\x25\x15\xc2\x83\x5b\x04\xcd\x8a\x58\xe5\xec\x71\x78\xf8\x7a\x7c\xd4\x8d\x09\x65\xe6\x57\x8b\x47\x34\xb3\x55\x63\x4a\x2c\x30\x90\x81\x9c\xb3\x98\x75\x6c\xe0\xb5\x7c\x01\x35\x8c\xd5\xe4\x8b\x1e\xe7\x8c\xc3\xd7\x87\xc7\x93\x27\xcb\xb7\xc0\x9e\x78\x7f\xed\x39\x32\xae\x94\xdf\x13\x9a\x22\xb1\x61\x50\x00\x57\x6b\x15\x76\xf0\xb8\xe7\x5f\x15\x1e\x63\xe3\x7e\xcf\x4e\xe7\x0c\x0a\x63\x65\xd2\xd4\x49\xe8\x21\x83\x86\x02\x6d\x6a\x90\x41\x25\x0d\x94\xbc\x29\x19\x0c\xea\x13\x82\x43\x9a\x67\x6c\xeb\x51\xdf\x2b\x71\x17\xc5\x04\x9f\xfc\x66\xb8\x40\x7a\x5e\x1d\x14\xbe\x23\xa8\xe0\x0b\x19\xc4\xa0\x57\x0b\xd9\xb0\x7b\x13\xd8\xf9\x29\x71\x0b\xba\x55\x73\x10\x48\x79\xb6\x00\xd4\x89\x6c\xfb\x75\x4b\x74\xb6\xbc\x6f\xe7\xbb\x01\xf2\x52\x88\x2e\xf1\x9d\x14\xba\xf5\x62\x97\xe0\xf5\x9b\xfd\x32\x23\xec\xbd\xf2\x42\xe5\xac\x00\xbf\xba\xd3\x72\xce\x90\x74\x40\x27\x26\xa1\xe9\x40\xe6\x00\x87\x24\x5f\xd2\x05\x38\x1c\x28\x0d\x1e\xa3\x7a\x0a\x7f\xf9\x82\x0c\xb2\x71\x20\x38\x64\xce\x17\x2e\xa3\x63\x63\x99\x39\xf2\xf0\xeb\x52\x1b\xec\xeb\xf8\x57\x38\xb5\xb9\x3b\x70\x11\xc5\x07\x98\x33\x20\x4e\xb0\xbe\xda\x91\x2d\xc2\xca\x97\x8b\x7e\xde\xd1\x7b\x7e\x82\x79\xef\x79\xb2\x39\xa4\x35\x29\xb5\x63\xca\xa0\xae\x4d\x22\x6c\x28\x13\x96\xd2\x52\xd4\x49\xb2\x92\x62\x15\x28\x09\xdd\xda\x9c\x63\x37\xd5\xb3\x39\x24\x52\x89\xdd\x1c\xec\x2f\x9b\xa6\x96\x5c\x99\x32\x4d\x7d\x51\x3f\x2f\x59\xc2\x29\xd1\x31\xd4\x61\x69\x5b\xde\xa1\x67\xd9\xab\x23\x60\x54\xfb\x0d\xf1\xd8\xef\x19\x1a\xdf\xc0\xea\x15\x98\x62\x0d\x02\x76\xc7\xb5\x04\xd2\x6f\x48\x8a\xd9\xe0\x08\x9e\x92\x1b\x8f\x3a\x01\xba\x06\x04\x0f\x0a\x4c\x14\xd3\x72\xcf\x60\x1e\xdc\xa2\xb7\x30\x40\x40\x3e\x35\xca\x40\x62\xec\xf9\xc1\x39\xcd\x53\xdf\x7a\xcb\xee\x69\xa9\xcc\x30\xec\x7e\x6d\xbf\x87\x79\xe3\xe3\x85\xb0\xd3\x93\xd5\xb3\xad\xed\x6a\xae\xe9\x6c\xa7\x1a\x61\xb1\x8a\x3f\xb6\x09\xbd\xfa\x35\x3a\xa1\x27\x88\xdd\x25\x15\xc3\x07\xc9\x6e\x38\x24\xfc\x88\xac\x54\xe9\x76\xc8\x1d\x98\x8d\xa4\x72\xce\x20\xbf\xd2\x94\x59\x46\x5a\x05\xb0\x0a\x5f\xcb\xf5\xae\x9e\xf5\xa7\x2a\x3e\xed\xfa\xd7\xeb\xe9\xa6\x83\x32\xc8\x87\xed\xe0\xbb\xf7\x42\x17\xee\xf6\xfd\x55\xc3\xb5\xb7\x11\x5b\x80\x3b\x44\x80\x87\x2f\xee\x73\x4c\x6d\x4c\x65\xdb\xb3\x40\x59\xe3\xb8\x35\x80\xcf\x81\xf9\xe0\xd6\xf5\xd9\xdb\xe1\xd7\xae\x1d\xc9\xe5\xa6\xd6\x8c\x8f\x0f\xda\x5b\x32\x84\x59\x40\x6b\xf4\x4c\x76\xc6\xf5\x4d\x76\xba\x6d\xa5\x5f\xc4\x2c\x70\x3b\xbb\xfa\x3a\xa8\xa6\xf5\xf8\x39\x8b\xbc\x9b\xb7\x19\x1c\xc7\x37\x50\x78\xc5\xe9\xc0\x8e\xea\x8c\x31\x33\x20\x40\xf8\xd2\xd3\x41\x75\x63\x16\x27\xf2\x5a\x87\xb1\x50\x65\x92\x82\x4b\x98\xbd\x3b\xa3\xd7\xf4\x3e\x12\x7c\xae\xa3\x0c\x52\x5a\x60\x5a\x87\xd7\x3a\x7a\x13\x4e\xc2\x03\x27\x41\x47\xdf\xbe\x91\xf0\x5c\xc9\x94\x2f\xc2\xab\x95\x34\xf4\xfe\x0a\xc7\xc9\xc3\x43\x08\xce\x0a\x63\xad\x07\xdd\x77\x81\x78\xd9\x38\x82\x53\x3e\x59\xb5\xea\xf9\xda\xce\x9e\xdb\xa0\xc6\x78\x0a\x24\x24\x48\xe9\x92\x0b\xc8\x91\xc1\x7b\x26\x6e\x61\x47\xc5\x94\x7c\x64\x25\x23\x1f\x50\xe5\xc1\xc8\x1b\xc7\xe1\xa0\x3d\x6c\x97\xc3\xc8\x39\x6c\x8f\x79\xc1\x47\x64\x33\x33\x22\x6f\x0b\x4e\x41\x55\x4d\xa5\x0e\x34\x9c\x84\x69\x67\x79\xff\x11\x2f\x4b\xdb\x64\x46\x09\x05\x64\xe6\xc7\x0b\xfb\x34\x6e\xf6\xa0\x3e\x2e\x0a\x55\xca\x24\xa8\x96\x4d\x5e\x1f\xbd\x39\x7f\xdb\xbc\x00\x5c\x5f\xfa\xb9\x5b\x89\x49\x7e\x4f\xb4\x12\x50\xa9\x7e\x4c\xc6\xf8\xaf\x5b\x1b\x49\x6f\x1b\xaa\x74\x88\xbb\x3c\xba\x9c\x5c\xbc\xeb\xe6\x12\x36\xda\xc1\xe3\x4b\xbb\xd1\x52\xfb\x74\x2a\xef\x8a\xe6\x13\x55\xc7\x48\xb3\x24\x10\x5c\x37\xeb\x30\x0e\x05\x36\x8f\x02\xb3\xca\x99\xab\x35\xbb\x41\x04\xef\x49\x31\xbc\x58\xf4\x3e\x82\xde\x8b\xc6\x30\x51\x71\x89\xcc\x21\xb0\x7b\xa9\x83\xd5\x25\x40\x5c\x0a\xea\xe8\x7c\x5d\x2b\x7c\xb6\xc5\x77\x93\xd8\xaf\x76\x0a\x07\x0e\xbe\x8d\x81\x9c\xf4\x55\x81\x59\xe4\x6e\xf2\x67\xb8\xdd\xe0\x2b\xe1\xb7\xd5\x25\x06\x56\xa5\x6a\x87\xe2\x30\x4f\xa0\x26\xd8\x4c\x1e\x54\x4b\xec\x81\x68\xef\x30\x06\x8f\xdb\x77\x96\x8d\xcf\xde\xad\xf5\xb2\x66\x83\x8c\xf1\x1a\x27\x02\x20\x14\xea\xbe\x76\x08\x83\x44\xad\x4b\xda\xb4\xc8\xbe\xb4\x52\xd8\xc5\x5e\x54\xbd\x59\x7c\xa0\x0c\x15\x54\x2e\x18\xf9\xe9\x66\x44\x7e\xba\x25\xd3\x53\x12\x82\x7e\x97\xf6\x54\x79\x78\xa8\xad\x85\x82\x78\x36\xa3\xeb\x02\x78\x4d\x6f\x29\x94\x78\x9e\x9b\xe9\xad\xe2\xc9\xf0\x60\x7f\xa3\x4e\x2d\xe8\x03\x92\x50\x43\x83\x6a\xec\x74\x00\x12\x7f\xba\x21\xff\x04\xbd\xf0\x80\x7e\x78\x18\x9c\xb9\x91\x87\x87\x59\x44\xcf\x66\x50\x37\x5b\x1a\x02\x93\xf6\x75\x99\x45\xa5\xa8\xb9\x6c\xe3\x18\xdf\x1f\x5e\x2b\x4d\x9a\x9b\xd2\x73\xc2\x53\x1d\xe0\xe3\xfb\x89\x35\xb0\x0e\x6e\x1b\xd5\xb4\x01\xa0\x6b\x26\x6c\x14\x6f\xdb\x58\x4f\x84\x67\x33\xe7\x66\xa2\x8b\xf8\x7f\x39\x77\x1e\x87\xf0\x80\xb9\x86\xf3\x05\xf2\xdc\x22\x9f\xfd\x3f\x44\x1c\x82\x08\xfc\x3f\xb0\x12\x1b\x9d\x68\xa1\xfa\xa5\x9c\x65\x02\x70\x39\xc4\xe8\x7d\x05\x01\xfb\xf8\x93\xfc\xa0\x68\x32\xdc\x3f\x69\xaf\xb7\x4e\xb9\xa5\x05\xf6\xd2\x9a\x9c\x92\xca\xfd\xe1\x82\x99\x8b\x75\x67\xf2\xcb\xea\x1c\x63\xf3\x91\x2e\xd9\xb0\x1e\x9f\xfd\x93\xcd\xfb\x98\x8f\xdf\x07\x60\x33\x18\x10\x2c\x44\x14\x91\x4f\x39\x50\x1a\x7b\xc5\x4e\x54\x1c\x43\x8b\x26\x63\x66\xe7\x50\xb5\xbf\x1f\xfc\x11\xda\xca\x11\xae\xcb\x0b\x88\xda\xb3\xf4\x6b\xcf\xe9\x60\xe5\xe3\xaa\xb8\x12\x45\x5e\x9c\x92\x41\x3b\x3d\x1f\x05\xbe\x4d\x12\xbc\xbb\x90\xc6\x16\x71\x26\xf1\x36\x1b\x48\xae\x45\xaa\xd9\xf5\x1e\x88\x16\xf6\x29\xa7\x40\xd2\xa4\xbd\x1f\x21\x43\x9f\x65\xc5\x19\x66\xf9\x07\x58\x39\xc4\xf6\x6d\xff\xa4\x31\xf3\x85\xce\xfd\x89\x07\xf7\x85\x1d\xed\x10\x45\x70\x00\x3e\x38\x81\xaf\x99\x93\x1d\x0a\x26\x17\x26\x83\x91\x17\x2f\x7c\x31\xce\x44\xfe\x47\x08\x15\xe7\x02\xf5\xfe\xb0\x56\x7b\xb8\x17\x0b\x0e\x9e\x18\xf9\xda\x8e\x48\x4a\x85\x66\x95\xcc\x8d\xd5\x97\x6b\x13\x9c\x89\x1b\x83\x1e\x35\x65\xbe\x4c\xab\x5f\x02\x0a\x32\x0c\xe8\x5b\x63\x0a\x3e\x2f\x0d\x86\xd1\xaf\x3c\x03\xcf\xe4\x2e\xb3\x30\x82\x7d\x56\xe1\x03\x9d\xf7\xd0\xae\xc9\xa8\xfe\x74\x27\x3f\xaf\x9b\xfe\x21\xdf\x6f\xae\xdc\xe4\x03\xef\xc8\x07\x3c\x70\xf6\xb6\xb3\x50\x7c\x3a\x52\xf4\x97\xd5\x9f\x93\x21\x4f\xf6\xb7\xe7\xd8\xda\x89\x0d\x9f\xd9\xb8\xd7\x9c\xf6\xdd\x91\xad\x7c\xe0\x16\x3d\xc9\x09\x9b\x64\x78\xcc\xf7\xd3\x8e\xc1\xb0\x60\x60\x49\xcc\x86\xd1\xef\xf3\xe6\x56\xf8\x7d\x1e\x8d\xc8\xde\xde\xfe\x6e\x8f\xb1\x27\xec\x2a\xf7\x86\x57\x61\xa2\xf5\xc9\x1e\xd9\xff\xca\xff\x6f\x00\x00\x00\xff\xff\x39\xae\x7b\x5f\xda\x1f\x00\x00")

func etcDocHtmlBytes() ([]byte, error) {
	return bindataRead(
		_etcDocHtml,
		"etc/doc.html",
	)
}

func etcDocHtml() (*asset, error) {
	bytes, err := etcDocHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/doc.html", size: 8154, mode: os.FileMode(420), modTime: time.Unix(1461066317, 0)}
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
	"etc/doc.html": etcDocHtml,
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
	"etc": &bintree{nil, map[string]*bintree{
		"doc.html": &bintree{etcDocHtml, map[string]*bintree{}},
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

