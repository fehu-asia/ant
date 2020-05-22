package meta

import (
	"ant/db"
	"fmt"
	"mime/multipart"
	"time"
)

// 文件元数据
type FileMeta struct {
	Id string // id 可以使用 sha1 md5
	// Linux文件名的长度限制是255个字符(Byte)。
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
	db.SaveFileMetaData(fm.Id, fm.Name, fm.Size, fm.Localtion)
	//fileMetas[fm.Id] = fm
}

// 通过sha1获取文件元数据对象
func GetFileMeta(hash string) *FileMeta {
	file, e := db.GetFileMeta(hash)
	if e != nil {
		fmt.Println("GetFileMeta error : ", e)
		return nil
	}
	return &FileMeta{
		Id:        file.Sha256,
		Name:      file.Name.String,
		Size:      file.Size.Int64,
		Localtion: file.Addr.String,
	}
}

// 删除元数据
func DeleteFileMeta(hash string) *FileMeta {
	meta := GetFileMeta(hash)
	delete(fileMetas, hash)
	return meta
}

// 通过sha1获取文件元数据对象
func GetAllFileMeta() map[string]*FileMeta {
	return fileMetas
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
