package ontology_go_sdk

import (
	sdkcom "github.com/infinitete/idfor-go-sdk/common"
	"github.com/infinitete/idfor/common"
)

type didDocument struct {
	ontSdk *OntologySdk
	native *NativeContract
}

type didKey struct {
	Key []byte
}

type didValue struct {
	Key   []byte
	Value []byte
}

func (d *didDocument) Put(acc *Account, key []byte, value []byte) (common.Uint256, error) {
	param := didValue{
		Key:   key,
		Value: value,
	}
	imt, err := d.native.NewNativeInvokeTransaction(0, 0, byte(0), DIDDOC_CONTRACT_ADDRESS, "PutDocument", []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	err = d.ontSdk.SignToTransaction(imt, acc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return d.ontSdk.SendTransaction(imt)
}

func (d *didDocument) Get(key []byte) (*sdkcom.PreExecResult, error) {
	param := didKey{
		Key: key,
	}
	imt, err := d.native.NewNativeInvokeTransaction(0, 0, byte(0), DIDDOC_CONTRACT_ADDRESS, "GetDocument", []interface{}{param})
	if err != nil {
		return nil, err
	}

	return d.ontSdk.PreExecTransaction(imt)
}
