package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/agelinazf/egb"
)

//相册表
type Gallery struct {
	//主键
	Id          int `orm:"column(Id);pk"`
	//栏目Id
	CateId      int    `orm:"column(CateId)"`
	//名字
	Title       string `orm:"column(Title)"`
	//描述
	Description string `orm:"column(Description);null"`
	//创建时间
	CreateTime  string `orm:"column(CreateTime)"`
	//更新时间
	UpdateTime  string `orm:"column(UpdateTime)"`
	//排序值
	Sort        int `orm:"column(Sort);default(0)"`
	//缩略图
	Thumb	string 	`orm:"column(Thumb)"`
	//点击量
	Hit         int64    `orm:"column(Hit);default(0)"`
	//下载量
	PicsDownload    int64    `orm:"column(PicsDownload);default(0)"`
}

func (t *Gallery) TableName() string {
	return "gallery"
}

//GetOneGalleryById 获取一个相册
//@params	Id
//@return	*Gallery
func GetOneGalleryById(Id int) (*Gallery, error) {
	gallery := new(Gallery)
	gallery.Id = Id

	if err := ormer().Read(gallery,"Id"); err != nil {
		beego.Error("GetOneGalleryById : " + err.Error())
		return nil, fmt.Errorf(ErrInfo[DataBaseGetError])
	}
	return gallery,nil
}

//GetGallerysNum 获取相册的数量
//@params	cateId keyword(搜索title的关键词)
//@return	int
func GetGallerysNum(catId int, keyword string) int {
	var data []orm.Params
	//todo count
	sql := "SELECT Id FROM gallery WHERE CateId = ? AND Title LIKE ? "
	keyword = "%" + keyword + "%"
	ormer().Raw(sql, catId, keyword).Values(&data)
	return len(data)
}

//GetGallerys 获取相册
//@params	catId keyword(搜索title的关键词) pagesize offset
//@return	[]orm.Params
func GetGallerys(catId int, keyword string, pagesize, offset int) []orm.Params {

	var data []orm.Params
	sql := `SELECT * FROM gallery WHERE CateId = ? AND Title LIKE ? ORDER BY gallery.Sort DESC,UpdateTime DESC LIMIT ?,?`
	keyword = "%" + keyword + "%"
	ormer().Raw(sql, catId, keyword, offset, pagesize).Values(&data)
	return data
}

//CreateOneGallery 新建一篇相册
//@params	catId title thumb source description content
//@return	error
func CreateOneGallery(cateId,modelId int, title, description string) error {
	gallery := new(Gallery)
	gallery.CateId = cateId
	gallery.Title = title
	gallery.Description = description
	gallery.CreateTime = egb.TimeNowUnix()
	gallery.UpdateTime = egb.TimeNowUnix()

	if _,err := ormer().Insert(gallery); err != nil {
		beego.Error("CreateOneGallery : " + err.Error())
		return fmt.Errorf(ErrInfo[SystemError])
	}
	return nil
}

//UpdateGallery 更新相册
//@params	id title description
//@return	error
func UpdateGallery(id, cateId, modelId int, title, description string) error {
	gallery,err := GetOneGalleryById(id)
	if err != nil {
		return err
	}

	gallery.CateId = cateId
	gallery.Title = title
	gallery.Description = description
	gallery.UpdateTime = egb.TimeNowUnix()

	if _,err := ormer().Update(gallery); err != nil {
		beego.Error("UpdateGallery : " + err.Error())
		return fmt.Errorf(ErrInfo[SystemError])
	} else {
		return nil
	}
}

//UpdateGallerySort 更新相册排序
//@params	id sort
//@return	error
func UpdateGallerySort(id, sort int) error {
	gallery,err := GetOneGalleryById(id)
	if err != nil {
		return fmt.Errorf(ErrInfo[SystemError])
	}
	gallery.Sort = sort
	gallery.UpdateTime = egb.TimeNowUnix()
	if _,err := ormer().Update(gallery,"Sort"); err != nil {
		beego.Error("UpdateGallerySort : " + err.Error())
		return fmt.Errorf(ErrInfo[SystemError])
	}
	return nil
}

//DeleteOneGallery 删除一个相册
//@params	id
//@return	error
func DeleteOneGallery(id int) error {
	gallery,err := GetOneGalleryById(id)
	if err != nil {
		return err
	}
	if _,err := ormer().Delete(gallery); err != nil {
		beego.Error("DeleteOneGallery : " + err.Error())
		return fmt.Errorf(ErrInfo[SystemError])
	}
	return nil
}
