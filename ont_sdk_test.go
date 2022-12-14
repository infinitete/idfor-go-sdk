/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package ontology_go_sdk

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	common2 "github.com/infinitete/idfor-go-sdk/common"
	"github.com/infinitete/idfor/common"
	"github.com/infinitete/idfor/core/payload"
	"github.com/infinitete/idfor/core/utils"
	"github.com/infinitete/idfor/core/validation"
	"github.com/infinitete/idfor/smartcontract/event"
	"github.com/infinitete/idfor/smartcontract/service/native/ont"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	"github.com/stretchr/testify/assert"
	"github.com/tyler-smith/go-bip39"
)

var (
	testOntSdk   *OntologySdk
	testWallet   *Wallet
	testPasswd   = []byte("123456")
	testDefAcc   *Account
	testGasPrice = uint64(0)
	testGasLimit = uint64(20000)
)

func TestOntId_NewRegIDWithAttributesTransaction(t *testing.T) {
	testOntSdk = NewOntologySdk()
}
func TestParseNativeTxPayload(t *testing.T) {
	testOntSdk = NewOntologySdk()
	pri, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb7cf")
	assert.Nil(t, err)
	acc, err := NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)
	state := &ont.State{
		From:  acc.Address,
		To:    acc.Address,
		Value: uint64(100),
	}
	transfers := make([]*ont.State, 0)
	for i := 0; i < 1; i++ {
		transfers = append(transfers, state)
	}
	_, err = testOntSdk.Native.Ont.NewMultiTransferTransaction(500, 20000, transfers)
	assert.Nil(t, err)
	_, err = testOntSdk.Native.Ont.NewTransferFromTransaction(500, 20000, acc.Address, acc.Address, acc.Address, 20)
	assert.Nil(t, err)
}

func TestParsePayload(t *testing.T) {
	testOntSdk = NewOntologySdk()
	//transferMulti
	payloadHex := "00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c0114c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"
	//one transfer
	payloadHex = "00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0164c86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"

	//one transferFrom
	payloadHex = "00c66b6a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a14d2c124dd088190f709b684e0bc676d70c41b3776c86a0114c86c0c7472616e7366657246726f6d1400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"

	payloadBytes, err := common.HexToBytes(payloadHex)
	assert.Nil(t, err)
	_, err = ParsePayload(payloadBytes)
	assert.Nil(t, err)
}

func TestParsePayloadRandom(t *testing.T) {
	testOntSdk = NewOntologySdk()
	pri, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb7cf")
	assert.Nil(t, err)
	acc, err := NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)
	assert.Nil(t, err)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000000; i++ {
		amount := rand.Intn(1000000)
		state := &ont.State{
			From:  acc.Address,
			To:    acc.Address,
			Value: uint64(amount),
		}
		param := []*ont.State{state}
		invokeCode, err := utils.BuildNativeInvokeCode(ONT_CONTRACT_ADDRESS, 0, "transfer", []interface{}{param})
		res, err := ParsePayload(invokeCode)
		assert.Nil(t, err)
		if res["param"] == nil {
			fmt.Println("amount:", amount)
			fmt.Println(res["param"])
			return
		} else {
			stateInfos := res["param"].([]common2.StateInfo)
			assert.Equal(t, uint64(amount), stateInfos[0].Value)
		}
		tr := ont.TransferFrom{
			Sender: acc.Address,
			From:   acc.Address,
			To:     acc.Address,
			Value:  uint64(amount),
		}
		invokeCode, err = utils.BuildNativeInvokeCode(ONT_CONTRACT_ADDRESS, 0, "transferFrom", []interface{}{tr})
		res, err = ParsePayload(invokeCode)
		assert.Nil(t, err)
		if res["param"] == nil {
			fmt.Println("amount:", amount)
			fmt.Println(res["param"])
			return
		} else {
			stateInfos := res["param"].(common2.TransferFromInfo)
			assert.Equal(t, uint64(amount), stateInfos.Value)
		}
	}
}
func TestParsePayloadRandomMulti(t *testing.T) {
	testOntSdk = NewOntologySdk()
	pri, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb7cf")
	assert.Nil(t, err)
	acc, err := NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)
	assert.Nil(t, err)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		amount := rand.Intn(10000000)
		state := &ont.State{
			From:  acc.Address,
			To:    acc.Address,
			Value: uint64(amount),
		}
		paramLen := rand.Intn(20)
		if paramLen == 0 {
			paramLen += 1
		}
		params := make([]*ont.State, 0)
		for i := 0; i < paramLen; i++ {
			params = append(params, state)
		}
		invokeCode, err := utils.BuildNativeInvokeCode(ONT_CONTRACT_ADDRESS, 0, "transfer", []interface{}{params})
		res, err := ParsePayload(invokeCode)
		assert.Nil(t, err)
		if res["param"] == nil {
			fmt.Println(res["param"])
			fmt.Println(amount)
			fmt.Println("invokeCode:", common.ToHexString(invokeCode))
			return
		} else {
			stateInfos := res["param"].([]common2.StateInfo)
			for i := 0; i < paramLen; i++ {
				assert.Equal(t, uint64(amount), stateInfos[i].Value)
			}
		}
	}
}

func TestOntologySdk_TrabsferFrom(t *testing.T) {
	testOntSdk = NewOntologySdk()
	payloadHex := "00c66b1421ab6ece5c9e44fa5e35261ef42cc6bc31d98e9c6a7cc814c1d2d106f9d2276b383958973b9fca8e4f48cc966a7cc80400e1f5056a7cc86c51c1087472616e736665721400000000000000000000000000000000000000020068164f6e746f6c6f67792e4e61746976652e496e766f6b65"
	payloadBytes, err := common.HexToBytes(payloadHex)
	assert.Nil(t, err)
	res, err := ParsePayload(payloadBytes)
	assert.Nil(t, err)
	fmt.Println("res:", res)

	//java sdk,  transferFrom
	//amount =100
	payloadHex = "00c66b14d2c124dd088190f709b684e0bc676d70c41b37766a7cc8149018fbdfe16d5b1054165ab892b0e040919bd1ca6a7cc8143e7c40c2a2a98e3f95adace19b12ef4a1d7a35066a7cc801646a7cc86c0c7472616e7366657246726f6d1400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"
	//amount =10
	//payloadHex = "00c66b14d2c124dd088190f709b684e0bc676d70c41b37766a7cc8149018fbdfe16d5b1054165ab892b0e040919bd1ca6a7cc8143e7c40c2a2a98e3f95adace19b12ef4a1d7a35066a7cc85a6a7cc86c0c7472616e7366657246726f6d1400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"

	//amount = 1000000000
	payloadHex = "00c66b14d2c124dd088190f709b684e0bc676d70c41b37766a7cc8149018fbdfe16d5b1054165ab892b0e040919bd1ca6a7cc8143e7c40c2a2a98e3f95adace19b12ef4a1d7a35066a7cc80400ca9a3b6a7cc86c0c7472616e7366657246726f6d1400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"

	//java sdk, transfer
	//amount = 100
	payloadHex = "00c66b14d2c124dd088190f709b684e0bc676d70c41b37766a7cc814d2c124dd088190f709b684e0bc676d70c41b37766a7cc801646a7cc86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"

	//amount = 10
	payloadHex = "00c66b14d2c124dd088190f709b684e0bc676d70c41b37766a7cc814d2c124dd088190f709b684e0bc676d70c41b37766a7cc85a6a7cc86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"
	//amount = 1000000000
	payloadHex = "00c66b14d2c124dd088190f709b684e0bc676d70c41b37766a7cc814d2c124dd088190f709b684e0bc676d70c41b37766a7cc80400ca9a3b6a7cc86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65"

	payloadBytes, err = common.HexToBytes(payloadHex)
	assert.Nil(t, err)
	res, err = ParsePayload(payloadBytes)
	assert.Nil(t, err)
	fmt.Println("res:", res)
}

//transferFrom
func TestOntologySdk_ParseNativeTxPayload2(t *testing.T) {
	testOntSdk = NewOntologySdk()
	var err error
	assert.Nil(t, err)
	pri, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb7cf")
	acc, err := NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)

	pri2, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb8cf")
	assert.Nil(t, err)

	pri3, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb9cf")
	assert.Nil(t, err)
	acc, err = NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)

	acc2, err := NewAccountFromPrivateKey(pri2, signature.SHA256withECDSA)

	acc3, err := NewAccountFromPrivateKey(pri3, signature.SHA256withECDSA)
	amount := 1000000000
	txFrom, err := testOntSdk.Native.Ont.NewTransferFromTransaction(500, 20000, acc.Address, acc2.Address, acc3.Address, uint64(amount))
	assert.Nil(t, err)
	tx, err := txFrom.IntoImmutable()
	assert.Nil(t, err)
	invokeCode, ok := tx.Payload.(*payload.InvokeCode)
	assert.True(t, ok)
	code := invokeCode.Code
	res, err := ParsePayload(code)
	assert.Nil(t, err)
	assert.Equal(t, acc.Address.ToBase58(), res["sender"].(string))
	assert.Equal(t, acc2.Address.ToBase58(), res["from"].(string))
	assert.Equal(t, uint64(amount), res["amount"].(uint64))
	assert.Equal(t, "transferFrom", res["functionName"].(string))
	fmt.Println("res:", res)
}
func TestOntologySdk_ParseNativeTxPayload(t *testing.T) {
	testOntSdk = NewOntologySdk()
	var err error
	assert.Nil(t, err)
	pri, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb7cf")
	acc, err := NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)

	pri2, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb8cf")
	assert.Nil(t, err)

	pri3, err := common.HexToBytes("75de8489fcb2dcaf2ef3cd607feffde18789de7da129b5e97c81e001793cb9cf")
	assert.Nil(t, err)
	acc, err = NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)

	acc2, err := NewAccountFromPrivateKey(pri2, signature.SHA256withECDSA)

	acc3, err := NewAccountFromPrivateKey(pri3, signature.SHA256withECDSA)
	y, _ := common.HexToBytes(acc.Address.ToHexString())

	fmt.Println("acc:", common.ToHexString(common.ToArrayReverse(y)))
	assert.Nil(t, err)

	amount := uint64(1000000000)
	tx, err := testOntSdk.Native.Ont.NewTransferTransaction(500, 20000, acc.Address, acc2.Address, amount)
	assert.Nil(t, err)

	tx2, err := tx.IntoImmutable()
	assert.Nil(t, err)
	res, err := ParseNativeTxPayload(tx2.ToArray())
	assert.Nil(t, err)
	fmt.Println("res:", res)
	assert.Equal(t, acc.Address.ToBase58(), res["from"].(string))
	assert.Equal(t, acc2.Address.ToBase58(), res["to"].(string))
	assert.Equal(t, amount, res["amount"].(uint64))
	assert.Equal(t, "transfer", res["functionName"].(string))

	transferFrom, err := testOntSdk.Native.Ont.NewTransferFromTransaction(500, 20000, acc.Address, acc2.Address, acc3.Address, 10)
	transferFrom2, err := transferFrom.IntoImmutable()
	r, err := ParseNativeTxPayload(transferFrom2.ToArray())
	assert.Nil(t, err)
	fmt.Println("res:", r)
	assert.Equal(t, r["sender"], acc.Address.ToBase58())
	assert.Equal(t, r["from"], acc2.Address.ToBase58())
	assert.Equal(t, r["to"], acc3.Address.ToBase58())
	assert.Equal(t, r["amount"], uint64(10))

	ongTransfer, err := testOntSdk.Native.Ong.NewTransferTransaction(uint64(500), uint64(20000), acc.Address, acc2.Address, 100000000)
	assert.Nil(t, err)
	ongTx, err := ongTransfer.IntoImmutable()
	assert.Nil(t, err)
	res, err = ParseNativeTxPayload(ongTx.ToArray())
	assert.Nil(t, err)
	fmt.Println("res:", res)
}

func TestOntologySdk_GenerateMnemonicCodesStr2(t *testing.T) {
	mnemonic := make(map[string]bool)
	testOntSdk := NewOntologySdk()
	for i := 0; i < 100000; i++ {
		mnemonicStr, err := testOntSdk.GenerateMnemonicCodesStr()
		assert.Nil(t, err)
		if mnemonic[mnemonicStr] == true {
			panic("there is the same mnemonicStr ")
		} else {
			mnemonic[mnemonicStr] = true
		}
	}
}

func TestOntologySdk_GenerateMnemonicCodesStr(t *testing.T) {
	testOntSdk := NewOntologySdk()
	for i := 0; i < 1000; i++ {
		mnemonic, err := testOntSdk.GenerateMnemonicCodesStr()
		assert.Nil(t, err)
		private, err := testOntSdk.GetPrivateKeyFromMnemonicCodesStrBip44(mnemonic, 0)
		assert.Nil(t, err)
		acc, err := NewAccountFromPrivateKey(private, signature.SHA256withECDSA)
		assert.Nil(t, err)
		si, err := signature.Sign(acc.SigScheme, acc.PrivateKey, []byte("test"), nil)
		boo := signature.Verify(acc.PublicKey, []byte("test"), si)
		assert.True(t, boo)

		tx, err := testOntSdk.Native.Ont.NewTransferTransaction(0, 0, acc.Address, acc.Address, 10)
		assert.Nil(t, err)
		testOntSdk.SignToTransaction(tx, acc)
		tx2, err := tx.IntoImmutable()
		assert.Nil(t, err)
		res := validation.VerifyTransaction(tx2)
		assert.Equal(t, "not an error", res.Error())
	}
}

func TestGenerateMemory(t *testing.T) {
	expectedPrivateKey := []string{"915f5df65c75afe3293ed613970a1661b0b28d0cb711f21c489d8785977df0cd", "dbf1090889ba8b19aa01fa31c8b1ce29828bd2fa664afd95cc62e6055b74e112",
		"1487a8e53e4f4e2e1991781bcd14b3d334d3b2965cb48c976b234da29d7cf242", "79f85da015f079469c6e04aa0fc23523187d0f72c29450073d858ddeed272617"}
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	mnemonic = "ecology cricket napkin scrap board purpose picnic toe bean heart coast retire"
	testOntSdk := NewOntologySdk()
	for i := 0; i < len(expectedPrivateKey); i++ {
		privk, err := testOntSdk.GetPrivateKeyFromMnemonicCodesStrBip44(mnemonic, uint32(i))
		assert.Nil(t, err)
		assert.Equal(t, expectedPrivateKey[i], common.ToHexString(privk))
	}
}

func TestOntologySdk_CreateWallet(t *testing.T) {
	testOntSdk := NewOntologySdk()
	wal, err := testOntSdk.CreateWallet("./wallet2.dat")
	assert.Nil(t, err)
	_, err = wal.NewDefaultSettingAccount(testPasswd)
	assert.Nil(t, err)
	wal.Save()
}

func TestNewOntologySdk(t *testing.T) {
	testOntSdk = NewOntologySdk()
	testWallet, _ = testOntSdk.OpenWallet("./wallet.json")
	event := &event.NotifyEventInfo{
		ContractAddress: common.ADDRESS_EMPTY,
		States:          []interface{}{"transfer", "Abc3UVbyL1kxd9sK6N9hzAT2u91ftbpoXT", "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK", uint64(10000000)},
	}
	e, err := testOntSdk.ParseNaitveTransferEvent(event)
	assert.Nil(t, err)
	fmt.Println(e)
}

func TestOntologySdk_GetTxData(t *testing.T) {
	testOntSdk = NewOntologySdk()
	testWallet, _ = testOntSdk.OpenWallet("./wallet.json")
	acc, _ := testWallet.GetAccountByAddress("AVBzcUtgdgS94SpBmw4rDMhYA4KDq1YTzy", testPasswd)
	tx, _ := testOntSdk.Native.Ont.NewTransferTransaction(500, 10000, acc.Address, acc.Address, 100)
	testOntSdk.SignToTransaction(tx, acc)
	tx2, _ := tx.IntoImmutable()
	var buffer bytes.Buffer
	tx2.Serialize(&buffer)
	txData := hex.EncodeToString(buffer.Bytes())
	tx3, _ := testOntSdk.GetMutableTx(txData)
	assert.Equal(t, tx, tx3)
}

func Init() {
	testOntSdk = NewOntologySdk()
	testOntSdk.NewRpcClient().SetAddress("http://localhost:20336")

	var err error
	var wallet *Wallet
	if !common.FileExisted("./wallet.json") {
		wallet, err = testOntSdk.CreateWallet("./wallet.json")
		if err != nil {
			fmt.Println("[CreateWallet] error:", err)
			return
		}
	} else {
		wallet, err = testOntSdk.OpenWallet("./wallet.json")
		if err != nil {
			fmt.Println("[CreateWallet] error:", err)
			return
		}
	}
	_, err = wallet.NewDefaultSettingAccount(testPasswd)
	if err != nil {
		fmt.Println("")
		return
	}
	wallet.Save()
	testWallet, err = testOntSdk.OpenWallet("./wallet.json")
	if err != nil {
		fmt.Printf("account.Open error:%s\n", err)
		return
	}
	testDefAcc, err = testWallet.GetDefaultAccount(testPasswd)
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
		return
	}

	ws := testOntSdk.NewWebSocketClient()
	err = ws.Connect("ws://localhost:20335")
	if err != nil {
		fmt.Printf("Connect ws error:%s", err)
		return
	}
}

func TestOnt_Transfer(t *testing.T) {
	testOntSdk = NewOntologySdk()
	testWallet, _ = testOntSdk.OpenWallet("./wallet.json")
	txHash, err := testOntSdk.Native.Ont.Transfer(testGasPrice, testGasLimit, testDefAcc, testDefAcc.Address, 1)
	if err != nil {
		t.Errorf("NewTransferTransaction error:%s", err)
		return
	}
	testOntSdk.WaitForGenerateBlock(30*time.Second, 1)
	evts, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		t.Errorf("GetSmartContractEvent error:%s", err)
		return
	}
	fmt.Printf("TxHash:%s\n", txHash.ToHexString())
	fmt.Printf("State:%d\n", evts.State)
	fmt.Printf("GasConsume:%d\n", evts.GasConsumed)
	for _, notify := range evts.Notify {
		fmt.Printf("ContractAddress:%s\n", notify.ContractAddress)
		fmt.Printf("States:%+v\n", notify.States)
	}
}

func TestOng_WithDrawONG(t *testing.T) {
	Init()
	unboundONG, err := testOntSdk.Native.Ong.UnboundONG(testDefAcc.Address)
	if err != nil {
		t.Errorf("UnboundONG error:%s", err)
		return
	}
	fmt.Printf("Address:%s UnboundONG:%d\n", testDefAcc.Address.ToBase58(), unboundONG)
	_, err = testOntSdk.Native.Ong.WithdrawONG(0, 20000, testDefAcc, unboundONG)
	if err != nil {
		t.Errorf("WithDrawONG error:%s", err)
		return
	}
	fmt.Printf("Address:%s WithDrawONG amount:%d success\n", testDefAcc.Address.ToBase58(), unboundONG)
}

func TestGlobalParam_GetGlobalParams(t *testing.T) {
	Init()
	gasPrice := "gasPrice"
	params := []string{gasPrice}
	results, err := testOntSdk.Native.GlobalParams.GetGlobalParams(params)
	if err != nil {
		t.Errorf("GetGlobalParams:%+v error:%s", params, err)
		return
	}
	fmt.Printf("Params:%s Value:%v\n", gasPrice, results[gasPrice])
}

func TestGlobalParam_SetGlobalParams(t *testing.T) {
	Init()
	gasPrice := "gasPrice"
	globalParams, err := testOntSdk.Native.GlobalParams.GetGlobalParams([]string{gasPrice})
	if err != nil {
		t.Errorf("GetGlobalParams error:%s", err)
		return
	}
	gasPriceValue, err := strconv.Atoi(globalParams[gasPrice])
	if err != nil {
		t.Errorf("Get prama value error:%s", err)
		return
	}
	_, err = testOntSdk.Native.GlobalParams.SetGlobalParams(testGasPrice, testGasLimit, testDefAcc, map[string]string{gasPrice: strconv.Itoa(gasPriceValue + 1)})
	if err != nil {
		t.Errorf("SetGlobalParams error:%s", err)
		return
	}
	testOntSdk.WaitForGenerateBlock(30*time.Second, 1)
	globalParams, err = testOntSdk.Native.GlobalParams.GetGlobalParams([]string{gasPrice})
	if err != nil {
		t.Errorf("GetGlobalParams error:%s", err)
		return
	}
	gasPriceValueAfter, err := strconv.Atoi(globalParams[gasPrice])
	if err != nil {
		t.Errorf("Get prama value error:%s", err)
		return
	}
	fmt.Printf("After set params gasPrice:%d\n", gasPriceValueAfter)
}

func TestWsScribeEvent(t *testing.T) {
	Init()
	wsClient := testOntSdk.ClientMgr.GetWebSocketClient()
	err := wsClient.SubscribeEvent()
	if err != nil {
		t.Errorf("SubscribeTxHash error:%s", err)
		return
	}
	defer wsClient.UnsubscribeTxHash()

	actionCh := wsClient.GetActionCh()
	timer := time.NewTimer(time.Minute * 3)
	for {
		select {
		case <-timer.C:
			return
		case action := <-actionCh:
			fmt.Printf("Action:%s\n", action.Action)
			fmt.Printf("Result:%s\n", action.Result)
		}
	}
}

func TestWsTransfer(t *testing.T) {
	Init()
	wsClient := testOntSdk.ClientMgr.GetWebSocketClient()
	testOntSdk.ClientMgr.SetDefaultClient(wsClient)
	txHash, err := testOntSdk.Native.Ont.Transfer(testGasPrice, testGasLimit, testDefAcc, testDefAcc.Address, 1)
	if err != nil {
		t.Errorf("NewTransferTransaction error:%s", err)
		return
	}
	testOntSdk.WaitForGenerateBlock(30*time.Second, 1)
	evts, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		t.Errorf("GetSmartContractEvent error:%s", err)
		return
	}
	fmt.Printf("TxHash:%s\n", txHash.ToHexString())
	fmt.Printf("State:%d\n", evts.State)
	fmt.Printf("GasConsume:%d\n", evts.GasConsumed)
	for _, notify := range evts.Notify {
		fmt.Printf("ContractAddress:%s\n", notify.ContractAddress)
		fmt.Printf("States:%+v\n", notify.States)
	}
}

func TestTrustNotify(t *testing.T) {
	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://192.168.124.20:20334")
	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.CreateWallet("test1.json")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := wallet.NewAccount(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, []byte("123456"))
	if err != nil {
		t.Fatal(err)
	}
	var identity *Identity

	if wallet.GetIdentityCount() == 0 {
		identity, err = wallet.NewIdentity(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, []byte("123456"))
		if err != nil {
			t.Fatal(err)
		}

		ctrl, _ := identity.GetControllerByIndex(1, []byte("123456"))
		_, err = sdk.Native.OntId.RegIDWithPublicKey(0, 0, acc, identity.ID, ctrl)
		if err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(time.Second * 6)

	id := identity.ID
	raw := []byte("???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????")
	keyIndex := []byte("1")
	// testSignature := []byte("ThisIsAFakeSignauture")

	ctrl, _ := identity.GetControllerByIndex(1, []byte("123456"))
	testSignature, _ := ctrl.Sign(raw)

	tx, err := sdk.Native.TrustNotify.Send(identity.ID, acc, raw, keyIndex, testSignature)
	if err != nil {
		t.Fatalf("TestTrustNotify fail: %s\n", err)
	}

	time.Sleep(time.Second * 6)

	txHash := tx.ToHexString()

	t.Logf("Transaction: %s\n", txHash)

	err = sdk.Native.TrustNotify.Verify(txHash, id)
	if err != nil {
		t.Fatalf("TestTrustNotify fail: %s\n", err)
	}

	t.Log("GetArgs:")
	args, err := sdk.Native.TrustNotify.Get(txHash)
	t.Logf("%#v\n - %s\n", args, err)
}

func TestStorage(t *testing.T) {
	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://127.0.0.1:20334")
	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.OpenWallet("test.json")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := wallet.NewAccount(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, []byte("123456"))
	if err != nil {
		t.Fatal(err)
	}

	var key = []byte("This_Is_A_Key")
	var val = []byte("This is a value")

	_, err = sdk.Native.Storage.Put(acc, key, val)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 6)

	res, err := sdk.Native.Storage.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	s, err := res.Result.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if s != string(val) {
		t.Fatalf("Want %s, got %s", val, s)
	}

	_, err = sdk.Native.Storage.Delete(acc, key)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 6)

	res, err = sdk.Native.Storage.Get(key)
	s, _ = res.Result.ToString()
	if s != "" {
		t.Fatalf("Delete Not Work: want empty, got %s", s)
	}
}

func TestVoucher(t *testing.T) {
	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://127.0.0.1:20334")
	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.OpenWallet("test.json")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := wallet.NewAccount(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, []byte("123456"))
	if err != nil {
		t.Fatal(err)
	}

	var pass = []byte("123456")
	var identity *Identity
	if wallet.GetIdentityCount() < 1 {
		_idt, err := wallet.NewIdentity(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, pass)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		ctrl, _ := _idt.GetControllerByIndex(1, pass)
		_, err = sdk.Native.OntId.RegIDWithPublicKey(0, 0, acc, _idt.ID, ctrl)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		identity = _idt
		time.Sleep(6 * time.Second)
		wallet.Save()
	} else {
		identity, _ = wallet.GetDefaultIdentity()
	}

	// vc, _ := sdk.Native.Voucer.NewVoucher(identity, pass, "??????????????????", "??????????????????")
	// if err != nil {
	// 	t.Fatalf("Voucher.NewVoucher error: %s", err)
	// }

	// tx, err := sdk.Native.Voucer.Put(acc, vc)
	// if err != nil {
	// 	t.Fatalf("Voucher.NewVoucher error: %s", err)
	// }

	// time.Sleep(time.Second * 6)

	// evt, err := sdk.GetSmartContractEvent(tx.ToHexString())
	// if err != nil {
	// 	t.Fatalf("Tx not success: %s", err)
	// }
	// notify := evt.Notify[0].States.([]interface{})
	// t.Logf("Event: %s - %s - %s", notify...)

	// res, err := sdk.Native.Voucer.GetVoucher(identity.ID)
	// if err != nil {
	// 	t.Fatalf("Voucher.NewVoucher error: %s", err)
	// }
	// _, err = res.Result.ToByteArray()
	// if err != nil {
	// 	t.Fatalf("Voucher.Bytes error: %s", err)
	// }

	tx, err := sdk.Native.Voucer.PutEvent(acc, identity, pass, identity.ID, "????????????????????????", "????????????????????????")
	if err != nil {
		t.Fatalf("Voucher.PutEvent error: %s", err)
	}
	time.Sleep(time.Second * 6)
	evt, err := sdk.GetSmartContractEvent(tx.ToHexString())
	if err != nil {
		t.Fatalf("Tx not success: %s", err)
	}
	notify := evt.Notify[0].States.([]interface{})
	t.Logf("Event: %s - %s - %s", notify...)

	_res, err := sdk.Native.Voucer.GetVoucherEvents("sds")
	t.Logf("LastErr: %s", err)
	for _, evt := range _res {
		t.Logf("%#v", evt)
	}
}

func TestAssignCtidPid(t *testing.T) {

	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://192.168.0.20:21334")
	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.OpenWallet("test.json")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		t.Fatal(err)
	}

	var pass = []byte("123456")

	var idt *Identity
	if wallet.GetIdentityCount() < 1 {
		_idt, err := wallet.NewIdentity(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, pass)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		ctrl, _ := _idt.GetControllerByIndex(1, pass)
		tx, err := sdk.Native.OntId.RegIDWithPublicKey(0, 0, acc, _idt.ID, ctrl)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		t.Logf("RegisterTx: %s", tx.ToHexString())
		idt = _idt
		time.Sleep(6 * time.Second)
		wallet.Save()
	} else {
		idt, _ = wallet.GetDefaultIdentity()
	}

	ch := make(chan struct{}, 3)
	w := sync.WaitGroup{}

	for i := 100; i < 500; i++ {
		go func(x int) {
			w.Add(1)
			ch <- struct{}{}
			tx, err := sdk.Native.OntId.AssignCtidPID(idt, acc, []byte(fmt.Sprintf("fecredBas-%d", x)), []byte("ThisIsAFakePid"), []byte("123456"))
			if err != nil {
				t.Errorf("Error: %s", err)
			}

			t.Logf("Tx: %s", tx.ToHexString())
			<-ch
			w.Done()
		}(i)
	}

	w.Wait()
}

func TestGetCtidPid(t *testing.T) {

	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://192.168.0.20:21334")
	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.OpenWallet("test.json")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		t.Fatal(err)
	}

	var pass = []byte("123456")

	var idt *Identity
	if wallet.GetIdentityCount() < 1 {
		_idt, err := wallet.NewIdentity(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, pass)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		ctrl, _ := _idt.GetControllerByIndex(1, pass)
		tx, err := sdk.Native.OntId.RegIDWithPublicKey(0, 0, acc, _idt.ID, ctrl)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		t.Logf("RegisterTx: %s", tx.ToHexString())
		idt = _idt
		time.Sleep(6 * time.Second)
		wallet.Save()
	} else {
		idt, _ = wallet.GetDefaultIdentity()
	}

	ch := make(chan struct{}, 6)
	w := sync.WaitGroup{}

	for i := 100; i < 400; i++ {
		go func(x int) {
			ch <- struct{}{}
			w.Add(1)
			_, err := sdk.Native.OntId.GetCtidPID(idt, acc, []byte(fmt.Sprintf("fecredBas-%d", x)), []byte("123456"))
			if err != nil {
				t.Errorf("Error: %s of %d", err, x)
			}
			<-ch
			w.Done()
		}(i)
	}

	w.Wait()
}

func TestGetCtidEPid(t *testing.T) {

	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://192.168.0.20:21334")
	sdk.ClientMgr.SetDefaultClient(rest)

	bas, epid, signature, err := sdk.Native.OntId.GetCtidEPID([]byte("did:idfor:TD4sWGkFUurqTJvYf4P9Zxfy9TNXNQ2YqK"), []byte("????????????1"))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("\nbas: %s,\nEPid: %s,\nSignature: %s\n", bas, epid, signature)
}

func TestRevokeCtidPid(t *testing.T) {

	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://192.168.0.20:21334")
	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.OpenWallet("test.json")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		t.Fatal(err)
	}

	var pass = []byte("123456")

	var idt *Identity
	if wallet.GetIdentityCount() < 1 {
		_idt, err := wallet.NewIdentity(keypair.PK_SM2, keypair.SM2P256V1, signature.SM3withSM2, pass)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		ctrl, _ := _idt.GetControllerByIndex(1, pass)
		tx, err := sdk.Native.OntId.RegIDWithPublicKey(0, 0, acc, _idt.ID, ctrl)
		if err != nil {
			t.Log("Error: can not create identity")
			t.FailNow()
		}
		t.Logf("RegisterTx: %s", tx.ToHexString())
		idt = _idt
		time.Sleep(6 * time.Second)
		wallet.Save()
	} else {
		idt, _ = wallet.GetDefaultIdentity()
	}

	for i := 0; i < 10000; i++ {
		tx, err := sdk.Native.OntId.RevokeCtidPID(idt, acc, []byte(fmt.Sprintf("fecredBas-%d", i)), []byte("123456"))
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Tx: %s", tx.ToHexString())
	}
}

func TestDidDocument(t *testing.T) {
	sdk := NewOntologySdk()
	rest := sdk.NewRestClient()
	rest.SetAddress("http://127.0.0.1:20334")

	sdk.ClientMgr.SetDefaultClient(rest)

	wallet, err := sdk.OpenWallet("test.json")

	acc, err := wallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		t.Fatal(err)
	}

	raw := `
{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:idfor:aabbcc",
  "authentication": [{
    "id": "did:idfor:aabbcc#keys-1",
    "type": "Ed25519VerificationKey2018",
    "controller": "did:idfor:aabbcc",
    "publicKeyBase58": "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
  }],
  "service": [{
    "id":"did:idfor:aabbcc#vcs",
    "type": "VerifiableCredentialService",
    "serviceEndpoint": "https://example.com/vc/"
  }]
}
`

	var key = []byte("did:idfor:aabbcc")
	var doc = []byte(raw)

	_, err = sdk.Native.DidDoc.Put(acc, key, doc)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 6)

	res, err := sdk.Native.DidDoc.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	s, err := res.Result.ToString()

	t.Logf("Result: %s", s)

	if err != nil {
		t.Fatal(err)
	}
	if s != string(doc) {
		t.Fatalf("Want %s, got %s", doc, s)
	}
}

func Test_A(t *testing.T) {
	u := "/api/v1/transaction"
	r := "/api/v1/idfor/did::scheme::addr/doc"

	l := strings.TrimRight(r, "addr/doc")
	t.Logf("L: %s", l)

	c := strings.Contains(u, l)

	t.Logf("Contains: %#v", c)
}
