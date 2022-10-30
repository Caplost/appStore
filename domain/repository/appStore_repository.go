package repository
import (
	"github.com/jinzhu/gorm"
	"git.imooc.com/coding-535/appStore/domain/model"
)
//创建需要实现的接口
type IAppStoreRepository interface{
    //初始化表
    InitTable() error
    //根据ID查处找数据
    FindAppStoreByID(int64) (*model.AppStore, error)
    //创建一条 appStore 数据
	CreateAppStore(*model.AppStore) (int64, error)
    //根据ID删除一条 appStore 数据
	DeleteAppStoreByID(int64) error
    //修改更新数据
	UpdateAppStore(*model.AppStore) error
    //查找appStore所有数据
	FindAll()([]model.AppStore,error)

}
//创建appStoreRepository
func NewAppStoreRepository(db *gorm.DB) IAppStoreRepository  {
	return &AppStoreRepository{mysqlDb:db}
}

type AppStoreRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *AppStoreRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.AppStore{}).Error
}

//根据ID查找AppStore信息
func (u *AppStoreRepository)FindAppStoreByID(appStoreID int64) (appStore *model.AppStore,err error) {
	appStore = &model.AppStore{}
	return appStore, u.mysqlDb.First(appStore,appStoreID).Error
}

//创建AppStore信息
func (u *AppStoreRepository) CreateAppStore(appStore *model.AppStore) (int64, error) {
	return appStore.ID, u.mysqlDb.Create(appStore).Error
}

//根据ID删除AppStore信息
func (u *AppStoreRepository) DeleteAppStoreByID(appStoreID int64) error {
	return u.mysqlDb.Where("id = ?",appStoreID).Delete(&model.AppStore{}).Error
}

//更新AppStore信息
func (u *AppStoreRepository) UpdateAppStore(appStore *model.AppStore) error {
	return u.mysqlDb.Model(appStore).Update(appStore).Error
}

//获取结果集
func (u *AppStoreRepository) FindAll()(appStoreAll []model.AppStore,err error) {
	return appStoreAll, u.mysqlDb.Find(&appStoreAll).Error
}

