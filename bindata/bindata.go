// Code generated by go-bindata.
// sources:
// resources/template.html
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

var _resourcesTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x59\xdd\x72\xdb\xb8\x15\xbe\xcf\x53\x60\xb5\xdb\xb1\x9c\x88\xa4\x65\x25\x8e\x23\x4b\x9e\xcd\x3a\xf6\xa6\x33\x99\x24\x53\x67\xa7\xd3\xe9\xee\x05\x44\x82\x24\x6c\x88\x60\x09\x50\xb6\x92\x7a\xa6\xaf\xd1\xd7\xeb\x93\xf4\x1c\x40\x94\xc1\x3f\xc9\x99\x96\x99\x91\x2c\x00\xfc\xce\x2f\x0e\xbe\x83\xcc\x7e\x78\xf7\xe9\xe2\xcb\xdf\x3e\x5f\x92\x54\x2f\xc5\xf9\xb3\x19\x7e\x11\x41\xb3\x64\x3e\x60\xd9\x00\x07\x18\x8d\xce\x9f\x11\x78\x66\x4b\xa6\x29\x09\x53\x5a\x28\xa6\xe7\x83\xdf\xbe\x5c\x79\xa7\x83\xcd\x94\xe6\x5a\xb0\xf3\x48\x86\x8a\x15\xab\x59\x60\x7f\xda\x29\xa5\xd7\xd5\xdf\xf8\x04\xcf\xb7\x7f\x92\xe7\xe4\xd7\x82\x47\x5c\x2d\xdd\xa1\xb7\x44\xf1\x65\x2e\xd8\x88\x14\x4c\xe5\x32\x53\x7c\x05\x7f\xd3\x2c\x22\x29\x7c\xac\xc9\xc5\xf5\x35\x49\xe0\x35\xb2\x58\x93\x9f\x43\xb9\x58\x37\xde\x4f\xb5\xce\xd5\x34\x08\x12\xae\xd3\x72\xe1\x87\x72\x19\x6c\x56\x05\x49\x4b\x5a\xf0\xcc\x51\x8c\x7c\x06\x89\xa0\x3f\x23\x4a\x2e\xe1\x83\x66\x5c\xaf\x71\x4d\xb5\xc4\x47\x80\xd1\xe3\xcf\x12\x16\x90\x6f\x8f\x70\xf0\x78\x77\x6c\x71\xcb\xb5\xb7\x90\xf7\x9e\xe2\x5f\x79\x96\x4c\xc9\x42\x16\x11\x2b\x70\xe8\xac\xbe\x76\x29\xbf\x3e\x69\xe1\xbe\x35\x0f\x35\x2b\xae\x99\x26\x65\x6e\x6d\x28\x4a\xc1\x14\xd1\x92\x24\x72\xc5\x8a\x8c\xe8\x94\x59\xe7\x35\xad\x6a\x98\x01\x6e\xca\x05\x5d\x83\x38\x21\xc3\xdb\xba\x36\xa1\x60\xb4\x40\x45\x74\xda\xa9\x82\xc5\xeb\x72\x4e\x2c\x24\xd5\x53\x22\x58\xac\xeb\x90\x77\x3c\xd2\xe9\x94\x8c\x8f\x8e\xfe\x54\x9f\xc8\x69\x14\x19\xb3\xc7\x47\x79\xaf\xc1\x5f\x52\xae\x08\xcb\x54\x59\xa0\xad\x60\xa1\x2c\x35\x2b\x48\x52\x6a\xf8\x52\x84\x16\x8c\xb0\x7f\x94\x54\xa0\x1f\x70\x7a\x18\xc9\x72\x21\x58\x74\x48\x78\x96\x3d\x2e\xf4\xdb\x4e\x31\x46\x4c\x63\x5e\x28\xed\x85\x29\x17\x4d\x37\x6d\xd4\xf3\xd0\xa2\x29\x39\xee\xd3\xd1\x05\x13\x74\x0f\x56\xc1\x93\x74\x17\x18\x18\xfc\x91\x29\xcd\x22\x13\x48\x30\x4f\x14\xb0\x45\xd7\xb0\x3b\x20\x71\x37\x20\x60\xa6\x2c\x93\x74\x04\x49\x00\xde\xd6\xff\xf9\xd7\xbf\x15\xc9\xca\x5b\x46\x20\x24\xae\x91\x26\x46\xdf\x67\xe4\x51\xb7\x85\x0e\xd2\x53\x2d\xdc\x89\x84\xb6\xd5\x74\x3a\xef\xcc\xa8\x0a\x52\xcb\xfc\x29\x80\x8e\x6a\xbb\xf1\x20\xb9\xb5\x5c\xf6\x41\x42\x08\x3e\xc0\x26\xcb\x99\x84\x3a\x65\x1d\x6b\x36\x96\xcd\xa3\xa0\x0a\x02\x94\x1e\x98\xd7\x4c\xac\x21\xd1\x08\x94\x4e\x59\xe2\x7a\x19\x93\x3b\xba\x56\xb5\x40\x64\xd2\xab\xd2\xd5\xa8\xd5\x28\x32\xee\x7c\xa7\xc2\xa0\x29\xf9\x01\xca\xa6\x2c\x34\xcd\x74\x9f\xd2\x7f\x2d\x68\x9e\xa3\x66\x54\x83\x3a\x4b\x7a\xcf\x97\xe5\xd2\x6e\x3e\x02\x5b\x48\xe6\x9a\xcb\x0c\xf6\x89\xab\xd9\x1d\xbc\xd3\x2a\x7c\xf8\xd3\xce\xd4\xb5\x01\x48\x6f\xb3\x97\xdf\xbc\x3e\xcd\x1b\x65\x6c\x49\x8b\x84\x67\xa8\x2a\x2d\xb5\xec\x55\xd2\xa8\x13\x42\xa8\x14\xc3\xec\x86\x14\x36\xa9\xad\x52\x30\x0e\x8f\x00\x02\x95\x4c\x81\xa2\x98\xd2\xcb\x05\x2b\x60\x27\x50\x45\xe2\x82\x86\xa8\xbe\x72\xcf\x82\x2b\x59\x10\x76\x4f\x31\x0c\x53\x12\xc3\x0f\x6a\xab\x9f\x89\xfc\x38\x98\x90\xa1\xcc\x30\x76\xbc\x80\x62\x00\x71\xc1\x28\xe6\x50\x2f\x32\x6d\xbd\x32\x72\xc1\xcc\xb1\xb4\x26\xe0\x42\xf8\x34\xfa\xcd\x07\x77\xde\xd8\x9b\x0c\xaa\xc2\xc2\x04\x5b\xc2\xbb\x5d\x85\xe4\x2e\x95\x78\xa4\x55\xbf\xe0\xb5\x71\xc3\x77\x9d\x35\xb0\x5d\x44\x52\x2a\xe2\x1a\xce\x71\x37\xce\xab\xdd\x30\x60\xb7\x67\xec\xae\x61\x4d\xba\xb1\x26\x13\x7f\x32\x99\x1c\xef\x04\xd4\x77\xd2\x02\x2a\x07\xf1\xb8\x0f\xf1\xe4\xc4\x3f\x39\x39\x79\xb5\x57\x45\x28\xdb\x05\x64\xfd\xa8\x6b\x2e\x96\x25\x24\x44\x4d\xff\x97\xdd\xd2\x8e\x77\x0b\xd2\x69\xc1\xb6\xa2\x54\x4b\x96\x9d\xb6\xd2\x5c\xe3\x26\x7d\xe2\x5e\xef\xb7\x2b\xe6\x71\x43\xf5\x57\x3d\xaa\xef\x0e\x23\x7a\xdd\x60\xd5\xbd\xde\x03\xf6\x72\x0f\x98\x35\xb4\x09\x37\xe9\x83\x3b\xd9\x0d\x87\x1e\x6b\xa3\xbd\xec\x43\x3b\xdd\x8d\x96\x48\x11\xb1\xcc\x53\x4b\x2a\x84\x03\x97\x78\xcd\x8a\x58\xe5\xec\xa9\x7f\xfc\x7a\x7c\xd2\x8d\x09\x65\xe6\x57\x83\x47\x14\x33\x55\x63\x4a\x0c\x30\x90\x81\x9c\xb3\x90\x75\x6c\xe0\x8d\x7c\x01\x35\x8c\xd5\xe4\x8b\x1e\xe7\x8c\xfd\xd7\xc7\xa7\x93\x27\xcb\x37\xc0\x8e\x78\x77\xed\x05\x32\xae\x98\xdf\x13\x1a\x23\xb1\x61\x50\x00\xd7\x1b\x15\xf6\xf0\xb8\xe7\x5f\x25\x1e\x63\xe3\x7e\xcf\x4e\x17\x0c\x0a\x63\x65\xd2\xd4\x4a\xe8\x21\x83\x9a\x02\x6d\x6a\x90\x41\x99\x69\x28\x79\x53\x32\x18\xd4\x27\x04\x87\x34\x4f\xd9\xce\xa3\xbe\x57\xe2\x3e\x8a\x09\x3e\xf9\x4d\x73\x81\xf4\xbc\x3a\x28\x5c\x47\x50\xc1\x93\xcc\x0b\x41\xaf\x16\xb2\x66\xf7\xda\x33\xf3\x53\x62\x17\x74\xab\x66\x21\x90\xf2\xec\x00\xa8\x13\xd9\xf6\xeb\x86\xe8\xec\x78\xdf\xcc\x77\x03\xe4\xa5\x10\x5d\xe2\x3b\x29\x74\xeb\xc5\x2e\xc1\x9b\x37\xfb\x65\x06\xd8\x7b\xe5\x85\xcc\x59\x01\x7e\xb5\xa7\xe5\x82\x21\xe9\x80\x4e\x2c\x83\xa6\x03\x99\x03\x1c\x92\x7c\x49\x13\x70\x38\x50\x1a\x3c\x46\xd5\x14\xfe\x72\x05\x69\x64\xe3\x40\x70\xc8\x82\x27\x36\xa3\x43\x6d\x98\x39\xf2\xf0\x9b\x52\x69\xec\xeb\xf8\x57\x38\xb5\xb9\x3d\x70\x11\xc5\x05\x58\x30\x20\x4e\xb0\xbe\xda\x91\x2d\xc2\xca\x97\x49\x3f\xef\xe8\x3d\x3f\xc1\xbc\xf7\x3c\xda\x1e\xd2\x8a\x94\xca\x32\x65\x50\xd7\x24\x11\x36\x94\x11\x8b\x69\x29\xea\x24\x59\x66\x62\xed\xc9\x0c\xba\xb5\x05\xc7\x6e\xaa\x67\x73\x64\x48\x25\xf6\x73\xb0\xbf\x6c\x9b\x5a\x72\xad\xcb\x38\x76\x45\xfd\xbc\x64\x11\xa7\x44\x85\x50\x87\x33\xd3\xf2\x0e\x1d\xcb\x5e\x9d\x00\xa3\x3a\x6c\x88\xc7\x7e\x4f\xd3\xf0\x16\x56\xaf\xc1\x14\x63\x10\xb0\x3b\xae\x32\x20\xfd\x9a\xc4\x98\x0d\x96\xe0\xc9\x6c\xeb\x51\x2b\x40\xd5\x80\xe0\x41\x81\x91\x64\x2a\x3b\xd0\x98\x07\x2b\xf4\x16\x06\x08\xc8\xa7\x42\x19\x48\x8c\x1d\x3f\x58\xa7\x39\xea\x1b\x6f\x99\x3d\x9d\x49\x3d\xf4\xbb\x5f\x3b\xec\x61\xde\xf8\x38\x21\xec\xf4\x64\xf5\xec\x6a\xbb\x9a\x6b\x3a\xdb\xa9\x46\x58\x8c\xe2\x8f\x6d\x42\xaf\x7e\x8d\x4e\xe8\x09\x62\xf7\x49\xc5\xf0\x41\xb2\x6b\x0e\x09\x3f\x22\x6b\x59\xda\x1d\x72\x07\x66\x23\xa9\x5c\x30\xc8\xaf\x38\x66\x86\x91\x56\x01\xac\xc2\xd7\x72\xbd\xad\x67\xfd\xa9\x8a\x4f\xbb\xfe\xf5\x7a\xba\xe9\xa0\x14\xf2\x61\x37\xf8\xfe\xbd\xd0\x85\xbb\x7b\x7f\xd5\x70\xcd\x6d\xc4\x0e\xe0\x0e\x11\xe0\xe1\xcb\xfb\x1c\x53\x1b\x53\xd9\xf4\x2c\x50\xd6\x38\x6e\x0d\xe0\x73\x60\x3e\xb8\x75\x73\xf6\x76\xf8\xb5\x6b\x47\xf2\x6c\x5b\x6b\xc6\xa7\x47\xed\x2d\xe9\xc3\x2c\xa0\x35\x7a\x26\x33\x63\xfb\x26\x33\xdd\xb6\xd2\x2d\x62\x06\xb8\x9d\x5d\x7d\x1d\x54\xd3\x7a\xfc\x9c\x05\xce\xcd\xdb\x0c\x8e\xe3\x5b\x28\xbc\x62\x3e\x30\xa3\x2a\x65\x4c\x0f\x08\x10\xbe\x78\x3e\xa8\x6e\xcc\xc2\x28\xbb\x51\x7e\x28\x64\x19\xc5\xe0\x12\x66\xee\xce\xe8\x0d\xbd\x0f\x04\x5f\xa8\x20\x85\x94\x16\x98\xd6\xfe\x8d\x0a\xde\xf8\x13\xff\xc8\x4a\x50\xc1\xb7\x6f\xc4\x7f\x27\xc3\x6b\x56\xac\x2e\x64\x16\xf3\xc4\xbf\x5e\x67\x9a\xde\x5f\xe3\x34\x79\x78\xf0\xc1\x67\x7e\xa8\xd4\xa0\xfb\x4a\x10\xef\x1c\x47\x70\xd8\x47\xeb\x56\x59\xdf\x98\xdb\x73\x29\xd4\x18\x8f\x81\x8b\x78\x31\x5d\x72\x01\xa9\x32\x78\xcf\xc4\x0a\x36\x56\x48\xc9\x47\x56\x32\xf2\x01\x35\x1f\x8c\x9c\x71\x1c\xf6\xda\xc3\x66\x39\x8c\x5c\xc0\x2e\x59\x14\x7c\x44\xb6\x33\x23\xf2\xb6\xe0\x14\x54\x55\x34\x53\x9e\x82\x03\x31\xee\xac\xf2\x3f\xe2\x9d\x69\x9b\xd3\x48\x21\x81\xd3\xfc\x78\x69\x9e\xc6\x05\x1f\x94\xc9\xa4\x90\x65\x16\x79\xd5\xb2\xc9\xeb\x93\x37\x17\x6f\x9b\xf7\x80\x9b\xbb\x3f\x7b\x39\x31\xc9\xef\x89\x92\x02\x0a\xd6\x8f\xd1\x18\xff\x75\x6b\x93\xd1\x55\x43\x95\x0e\x71\x57\x27\x57\x93\xcb\x77\xdd\x94\xc2\x04\xdd\x7b\x7c\x69\x3f\x5a\x6c\x9e\x4e\xe5\x6d\xed\x7c\xa2\xea\x18\x69\x16\x79\x82\xab\x66\x39\xc6\x21\xcf\xe4\x91\xa7\xd7\x39\xb3\x25\x67\x3f\x88\xe0\x3d\x29\x86\xf7\x8b\xce\x87\xd7\x7b\xdf\xe8\x47\x32\x2c\x91\x40\x78\x66\x4b\x75\x90\xbb\x08\xf8\x4b\x41\x2d\xab\xaf\x6b\x85\xcf\xae\xf8\x6e\x13\xfb\xd5\x5e\xe1\x40\xc5\x77\x11\x91\xb3\xbe\x62\x30\x0b\xec\x85\xfe\x0c\xb7\x1b\x7c\x45\x7c\x55\xdd\x65\x60\x71\xaa\x76\x28\x0e\xf3\x08\x4a\x83\xc9\xe4\x41\xb5\xc4\x9c\x8b\xe6\x2a\x63\xf0\xb8\x7d\x67\xe9\xf8\xfc\xdd\x46\x2f\x63\x36\xc8\x18\x6f\x70\x02\x00\x42\xa1\xf6\x6b\x8f\x30\x48\xd4\xba\xa4\x6d\xa7\xec\x4a\x2b\x85\x59\xec\x44\xd5\x99\xc5\x07\xaa\x51\x41\xb3\x84\x91\x9f\x6e\x47\xe4\xa7\x15\x99\xce\x4d\x75\xba\x32\x87\xcb\xc3\x43\x6d\x2d\xd4\xc5\xf3\x19\xdd\xd4\xc1\x1b\xba\xa2\x50\xe9\x79\xae\xa7\x2b\xc9\xa3\xe1\xd1\xe1\x56\x9d\x5a\xd0\x07\x24\xa2\x9a\x7a\xd5\xd8\x7c\x00\x12\x7f\xba\x25\xff\x04\xbd\xf0\x9c\x7e\x78\x18\x9c\xdb\x91\x87\x87\x59\x40\xcf\x67\x50\x3e\x5b\x1a\x02\xa1\x76\x75\x99\x05\xa5\xa8\xb9\x6c\xeb\x18\xd7\x1f\x4e\x47\x4d\x9a\x9b\xd2\x71\xc2\x53\x1d\xe0\xe2\xbb\x89\x35\x30\x0e\x6e\x1b\xd5\xb4\x01\xa0\x6b\x26\x6c\x15\x6f\xdb\x58\x4f\x84\x67\x33\xeb\x66\xa2\x8a\xf0\x7f\x39\x7e\x1e\x87\xf0\x80\xb9\x81\xf3\x05\xf2\xdc\x20\x9f\xff\xbf\x44\xe0\x7f\x85\x95\xd8\xef\x04\x89\xec\x97\x72\x9e\x0a\xc0\xe5\x10\xa3\xf7\x15\x04\xec\xe3\x4f\xd9\x07\x49\xa3\xe1\xe1\x59\x7b\xbd\x71\xca\x8a\x16\xd8\x52\x2b\x32\x27\x95\xfb\xfd\x84\xe9\xcb\x4d\x83\xf2\xcb\xfa\x02\x63\xf3\x91\x2e\xd9\xb0\x1e\x9f\xc3\xb3\xed\xfb\x98\x8f\xdf\x07\x60\x32\x18\x10\x0c\x44\x10\x90\x4f\x39\x30\x1b\x73\xd3\x4e\x64\x18\x42\xa7\x96\x85\xcc\xcc\xa1\x6a\x7f\x3f\xfa\xc3\x37\x95\xc3\xdf\x94\x17\x10\x75\x60\x58\xd8\x81\xd5\xc1\xc8\xc7\x55\x61\x25\x8a\xbc\x98\x93\x41\x3b\x3d\x1f\x05\xbe\x8d\x22\xbc\xc2\xc8\xb4\x29\xe2\x2c\xc3\x4b\x6d\xe0\xba\x06\xa9\x66\xd7\x7b\xe0\x5b\xd8\xae\xcc\x81\xab\x65\xe6\x9a\x84\x0c\x5d\xb2\x15\xa6\x98\xe5\x1f\x60\xe5\x10\xbb\xb8\xc3\xb3\xc6\xcc\x17\xba\x70\x27\x1e\xec\x17\x36\xb6\x43\x14\xc1\x01\xf8\xe8\x0c\xbe\x66\x56\xb6\x2f\x58\x96\xe8\x14\x46\x5e\xbc\x70\xc5\x58\x13\xf9\x1f\x3e\x54\x9c\x4b\xd4\xfb\xc3\x46\xed\xe1\x41\x28\x38\x78\x62\xe4\x6a\x3b\x22\x31\x15\x8a\x55\x32\xb7\x56\x5f\x6d\x4c\xb0\x26\x6e\x0d\x7a\xd4\x94\xb9\x32\x8d\x7e\x11\x28\xc8\x30\xa0\x6f\xb5\x2e\xf8\xa2\xd4\x18\x46\xb7\xf2\x0c\x1c\x93\xbb\xcc\xc2\x08\xf6\x59\x85\x0f\x34\xe0\x43\xb3\x26\xa5\xea\xd3\x5d\xf6\x79\xd3\xfb\x0f\xf9\x61\x73\xe5\x36\x1f\x78\x47\x3e\xe0\x81\x73\xb0\x9b\x8c\xe2\xd3\x91\xa2\xbf\xac\xff\x1c\x0d\x79\x74\xb8\x3b\xc7\x36\x4e\x6c\xf8\xcc\xc4\xbd\xe6\xb4\xef\x8e\x6c\xe5\x03\xbb\xe8\x49\x4e\xd8\x26\xc3\x63\xbe\xcf\x3b\x06\xfd\x82\x81\x25\x21\x1b\x06\xbf\x2f\x9a\x5b\xe1\xf7\x45\x30\x22\x07\x07\x87\xfb\x3d\xc6\x9e\xb0\xab\xec\x1b\x4e\x85\x09\x36\x27\x7b\x60\xfe\x47\xff\xbf\x01\x00\x00\xff\xff\x68\x9f\x52\x58\xe1\x1f\x00\x00")

func resourcesTemplateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesTemplateHtml,
		"resources/template.html",
	)
}

func resourcesTemplateHtml() (*asset, error) {
	bytes, err := resourcesTemplateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/template.html", size: 8161, mode: os.FileMode(420), modTime: time.Unix(1462884674, 0)}
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
	"resources/template.html": resourcesTemplateHtml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"template.html": &bintree{resourcesTemplateHtml, map[string]*bintree{}},
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

