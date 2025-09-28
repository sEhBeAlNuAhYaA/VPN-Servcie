package api

import (
	"VPNService/internal/service/service"
	"net/http"
)

type apiImpl struct {
	vpnProc service.VpnProc
}

func NewapiImpl(vpnProc service.VpnProc) *apiImpl {
	apii := apiImpl{vpnProc: vpnProc}
	return &apii
}

func (*apiImpl) Get(w http.ResponseWriter, r *http.Request) {

}
func (*apiImpl) Post(w http.ResponseWriter, r *http.Request) {

}
func (*apiImpl) Delete(w http.ResponseWriter, r *http.Request) {

}
func (*apiImpl) Update(w http.ResponseWriter, r *http.Request) {

}
