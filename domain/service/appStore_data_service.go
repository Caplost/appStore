package service

import (
	"git.imooc.com/coding-535/appStore/domain/model"
	"git.imooc.com/coding-535/appStore/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IAppStoreDataService interface {
	AddAppStore(*model.AppStore) (int64 , error)
	DeleteAppStore(int64) error
	UpdateAppStore(*model.AppStore) error
	FindAppStoreByID(int64) (*model.AppStore, error)
	FindAllAppStore() ([]model.AppStore, error)
}


//创建
//注意：返回值 IAppStoreDataService 接口类型
func NewAppStoreDataService(appStoreRepository repository.IAppStoreRepository,clientSet *kubernetes.Clientset) IAppStoreDataService{
	return &AppStoreDataService{ AppStoreRepository:appStoreRepository, K8sClientSet: clientSet,deployment:&v1.Deployment{}}
}

type AppStoreDataService struct {
    //注意：这里是 IAppStoreRepository 类型
	AppStoreRepository repository.IAppStoreRepository
	K8sClientSet  *kubernetes.Clientset
	deployment  *v1.Deployment
}


//插入
func (u *AppStoreDataService) AddAppStore(appStore *model.AppStore) (int64 ,error) {
	 return u.AppStoreRepository.CreateAppStore(appStore)
}

//删除
func (u *AppStoreDataService) DeleteAppStore(appStoreID int64) error {
	return u.AppStoreRepository.DeleteAppStoreByID(appStoreID)
}

//更新
func (u *AppStoreDataService) UpdateAppStore(appStore *model.AppStore) error {
	return u.AppStoreRepository.UpdateAppStore(appStore)
}

//查找
func (u *AppStoreDataService) FindAppStoreByID(appStoreID int64) (*model.AppStore, error) {
	return u.AppStoreRepository.FindAppStoreByID(appStoreID)
}

//查找
func (u *AppStoreDataService) FindAllAppStore() ([]model.AppStore, error) {
	return u.AppStoreRepository.FindAll()
}

