package dao

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"superStar/models"
)

type SuperstarDao struct {
	db *gorm.DB
}

func NewSuperstarDao(db *gorm.DB) *SuperstarDao {
	return &SuperstarDao{db: db}
}

func (d *SuperstarDao) Get(id int) *models.StarInfo {
	//往starInfo对象里放入值,把id传进去
	//相当于java中new一个startInfo
	//data = new StarInfo(); data.setId(id)
	data := &models.StarInfo{}
	//d是superstar的dao,通过dao操作gorm的db
	//使用db里面的方法进行数据库操作(curd等等)
	db := d.db.First(data, id)
	if db.Error != nil {
		return nil
	} else {
		return data
	}
}

func (d *SuperstarDao) GetAll() []models.StarInfo {
	//定义返回数据dataList
	var dataList []models.StarInfo
	//查询全部结果
	db := d.db.Order("id desc").Find(dataList)
	if db.Error != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *SuperstarDao) Search(country string) []models.StarInfo {
	//定义返回数据dataList
	var dataList []models.StarInfo
	//通过country字段来进行查询,根据id来排序
	db := d.db.Where("country = ?", country).Order("id desc").Find(dataList)
	if db.Error != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *SuperstarDao) Delete(id int) error {
	//设置id和sysStatus为1,进行假删除
	data := &models.StarInfo{Id: id, SysStatus: 1}
	//通过更新sysStatus来进行假删除
	db := d.db.Model(&models.StarInfo{}).Updates(data)
	return db.Error

}

func (d *SuperstarDao) Upedate(info *models.StarInfo) error {
	//创建一个map,后面update使用map来更新,这样可以更新到为0,"",false的值的字段
	//因为struck无法更新值为0,"",false的字段
	data := make(map[string]interface{})
	//把struck结构转换成json
	j, err := json.Marshal(info)
	if err != nil {
		return err
	} else {
		//把json结构转换成map
		json.Unmarshal(j, &data)
		//使用map结构来更新
		db := d.db.Model(&models.StarInfo{}).Updates(data)
		return db.Error
	}
}

func (d *SuperstarDao) Create(data *models.StarInfo) error {
	//创建记录
	db := d.db.Create(data)
	return db.Error
}
