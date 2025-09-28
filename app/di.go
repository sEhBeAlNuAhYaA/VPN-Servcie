package app

import (
	apiImpl "VPNService/api"
	"VPNService/api/api"
	serv "VPNService/internal/service"
	"VPNService/internal/service/service"
	rep "VPNService/repository"
	"VPNService/repository/repository"
)

type di struct {
	repVpn     repository.VpnRepository
	serviceVpn service.VpnProc
	apiVpn     api.ApiInterface
}

func NewDI() *di { return &di{} }

func (diContainer *di) GetRepository() repository.VpnRepository {
	if diContainer.repVpn == nil {
		diContainer.repVpn = rep.NewvpnRepImpl()
	}
	return diContainer.repVpn
}

func (diContainer *di) GetService() service.VpnProc {
	if diContainer.serviceVpn == nil {
		diContainer.serviceVpn = serv.NewvpnProcImpl(diContainer.GetRepository())
	}
	return diContainer.serviceVpn
}

func (diContainer *di) GetApi() api.ApiInterface {
	if diContainer.apiVpn == nil {
		diContainer.apiVpn = apiImpl.NewapiImpl(diContainer.GetService())
	}

	return diContainer.apiVpn
}
