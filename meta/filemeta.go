package meta

import (
	"mime/multipart"
	"time"
)

// 文件元数据
type FileMeta struct {
	Id        string // id 可以使用 sha1 md5
	Name      string // 文件名
	Size      int64  // 大小
	Localtion string // 文件路径
	UploadAt  string // 时间戳
	Type      int64  // 类型
}

var fileMetas map[string]*FileMeta

func init() {
	fileMetas = make(map[string]*FileMeta)
}

/*
	新增或者更新元数据
*/
func UpdaeFileMeta(fm *FileMeta) {
	fileMetas[fm.Id] = fm
}

// 通过sha1获取文件元数据对象
func GetFileMeta(hex string) *FileMeta {
	return fileMetas[hex]
}

func CreateFileMetaByHeader(h *multipart.FileHeader) *FileMeta {
	return &FileMeta{
		Id:        "",
		Name:      h.Filename,
		Size:      h.Size,
		Localtion: "/Users/silence/Downloads/goProject/src/tmp/" + h.Filename,
		UploadAt:  time.Now().Format("2006-01-02 15:04:05"),
		Type:      0,
	}
}
