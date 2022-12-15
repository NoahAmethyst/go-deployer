package deploy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tristan-club/kit/chain_info"
	"go-deployer/constant"
	data_storage "go-deployer/contract/data"

	"os"
	"testing"
)

var chainDataStorage map[uint32]common.Address

func TestDataStorageAddAdmin(t *testing.T) {
	loadChainDataStorage()
	chainType := uint32(chain_info.ChainTypePolygon)
	pk := os.Getenv(constant.PK)
	client, auth, err := NewEthCliByChainType(chainType, pk)
	if err != nil {
		t.Errorf("new eth cli error:%s", err)
		return
	}
	//fromAddress := auth.From
	//load contract instance
	instance, err := data_storage.NewStore(chainDataStorage[chainType], client)
	if err != nil {
		t.Errorf("get instance error:%s", err)
		return
	}

	//Todo set app & admins
	appId := "GStar"
	admins := []common.Address{
		//default
		common.HexToAddress("0x19250A66F26a940004D934960FbdFF8dCC80C422"),
		common.HexToAddress("0xC14282949c8A946C9CCe9c3F0fA5BeACfcdD3609"),
		common.HexToAddress("0x62aADc6990694a8c9B8Ed7F286f6BB09251C2584"),
		//app
		common.HexToAddress("0x802a856Be73973f705948320E8704e4F131AcF62"),
	}

	for i, _admin := range admins {
		_, err = instance.AddAdmin(auth, appId, _admin)
		if err != nil {
			t.Errorf("%s", err)
			return
		} else {
			isAdmin, err := instance.CheckAdmin(nil, appId, _admin)
			if err != nil {
				t.Errorf("%s", err)
			} else {
				t.Logf("%s is admin:%v", _admin.String(), isAdmin)
			}
		}
		if i < len(admins)-1 {
			err := refreshTxCallData(client, auth)
			if err != nil {
				t.Errorf("%s", err)
				return
			}
		}
	}
}

func loadChainDataStorage() {
	chainDataStorage = map[uint32]common.Address{
		chain_info.ChainTypePolygon: common.HexToAddress("0x6a6F81cBaA4096e21db53d6cA16B0457F0134Cf2"),
		chain_info.ChainTypeOkc:     common.HexToAddress("0x105d1a706CE715AC7ab2FdA4b43A47C269A92D15"),
		chain_info.ChainTypeBsc:     common.HexToAddress("0xC47BD2C91Eaa0C54c39d54a70C81998fcb49C06A"),
		chain_info.ChainTypeKlaytn:  common.HexToAddress("0xC0e092354FcBC65A29EE3Fa22c9EA7CaFb740a28"),
		chain_info.ChainTypeEvmos:   common.HexToAddress("0xC0e092354FcBC65A29EE3Fa22c9EA7CaFb740a28"),
		chain_info.ChainTypeCronos:  common.HexToAddress("0x4BB0fb4736D82e20e4D187b23476De7ee847e3Cd"),
		chain_info.ChainTypeAurora:  common.HexToAddress("0x4BB0fb4736D82e20e4D187b23476De7ee847e3Cd"),
	}
}
