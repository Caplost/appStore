package handler
import (
	"context"
    "git.imooc.com/coding-535/appStore/domain/service"
	log "github.com/asim/go-micro/v3/logger"
	appStore "git.imooc.com/coding-535/appStore/proto/appStore"
)
type AppStoreHandler struct{
     //注意这里的类型是 IAppStoreDataService 接口类型
     AppStoreDataService service.IAppStoreDataService
}


// Call is a single request handler called via client.Call or the generated client code
func (e *AppStoreHandler) AddAppStore(ctx context.Context,info *appStore.AppStoreInfo , rsp *appStore.Response) error {
	log.Info("Received *appStore.AddAppStore request")


	return nil
}

func (e *AppStoreHandler) DeleteAppStore(ctx context.Context, req *appStore.AppStoreId, rsp *appStore.Response) error {
	log.Info("Received *appStore.DeleteAppStore request")

	return nil
}

func (e *AppStoreHandler) UpdateAppStore(ctx context.Context, req *appStore.AppStoreInfo, rsp *appStore.Response) error {
	log.Info("Received *appStore.UpdateAppStore request")

	return nil
}

func (e *AppStoreHandler) FindAppStoreByID(ctx context.Context, req *appStore.AppStoreId, rsp *appStore.AppStoreInfo) error {
	log.Info("Received *appStore.FindAppStoreByID request")

	return nil
}

func (e *AppStoreHandler) FindAllAppStore(ctx context.Context, req *appStore.FindAll, rsp *appStore.AllAppStore) error {
	log.Info("Received *appStore.FindAllAppStore request")

	return nil
}


