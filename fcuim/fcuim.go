package fcuim

import (
	"time"

	SDK "git.fe-cred.com/idfor/idfor-go-sdk"
	"git.fe-cred.com/idfor/idfor-go-sdk/common"
	"git.fe-cred.com/idfor/idfor-go-sdk/utils"
)

var CONTRACT_ADDRESS, _ = utils.AddressFromHexString("0800000000000000000000000000000000000000")

func GetSchemes(sdk *SDK.OntologySdk, acc *SDK.Account) ([]*common.NotifyEventInfo, error) {

	tx, err := sdk.Native.InvokeNativeContract(0, 0, acc, byte(0), CONTRACT_ADDRESS, "getFcuimSchemes", []interface{}{})
	if err != nil {
		return nil, nil
	}

	time.Sleep(time.Second * 6)
	evt, err := sdk.GetSmartContractEvent(tx.ToHexString())

	if err != nil {
		return nil, nil
	}

	return evt.Notify, nil
}
