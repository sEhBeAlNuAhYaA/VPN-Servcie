package service

import (
	"VPNService/repository/repository"
)

type VpnProcImpl struct {
	rep repository.VpnRepository
}

func NewvpnProcImpl(rep repository.VpnRepository) *VpnProcImpl {
	vpnProc := VpnProcImpl{rep: rep}

	return &vpnProc
}

func (vpnProc *VpnProcImpl) Get() {

}

func (vpnProc *VpnProcImpl) Post() {

}

func (vpnProc *VpnProcImpl) Delete() {

}

func (vpnProc *VpnProcImpl) Update() {

}
