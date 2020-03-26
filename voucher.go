/*
 * Copyright (C) 2020 The Idfor Team
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
	"encoding/hex"
	"io"

	sdkcom "git.fe-cred.com/idfor/idfor-go-sdk/common"
	"git.fe-cred.com/idfor/idfor/common"
)

type voucher struct {
	ontSdk *OntologySdk
	native *NativeContract
}

type Voucher struct {
	Key       string
	Title     string
	Desc      string
	Creator   string
	Signature string
	Timestamp int64
}

type VoucherEvent struct {
	Key       string // Voucher.Key
	Index     uint8  // 排序
	Title     string
	Content   string
	Operator  string
	Signature string
	Timestamp int64
}

type t_get_voucher struct {
	Key string
}

func (vt *Voucher) deserialization(source *common.ZeroCopySource) error {
	key, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	title, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	content, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	creator, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	if eof {
		return io.ErrUnexpectedEOF
	}
	signature, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	timestamp, eof := source.NextInt64()
	if eof {
		return io.ErrUnexpectedEOF
	}

	vt.Key = key
	vt.Title = title
	vt.Creator = creator
	vt.Signature = signature
	vt.Desc = content
	vt.Timestamp = timestamp

	return nil
}

func (vc *VoucherEvent) deserialization(source *common.ZeroCopySource) error {

	key, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	index, eof := source.NextUint8()
	if eof {
		return io.ErrUnexpectedEOF
	}
	title, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	content, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	operator, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	signature, _, _, eof := source.NextString()
	if eof {
		return io.ErrUnexpectedEOF
	}
	timestamp, eof := source.NextInt64()
	if eof {
		return io.ErrUnexpectedEOF
	}

	vc.Key = key
	vc.Index = index
	vc.Title = title
	vc.Content = content
	vc.Operator = operator
	vc.Signature = signature
	vc.Timestamp = timestamp

	return nil
}

func (vc *voucher) NewVoucher(creator *Identity, pass []byte, title, desc string) (*Voucher, error) {
	key, err := GenerateID()
	if err != nil {
		return nil, err
	}

	ctrl, err := creator.GetControllerByIndex(1, pass)
	if err != nil {
		return nil, err
	}
	bs, err := ctrl.Sign([]byte(desc))
	if err != nil {
		return nil, err
	}

	return &Voucher{
		Key:       key,
		Title:     title,
		Desc:      desc,
		Creator:   creator.ID,
		Signature: hex.EncodeToString(bs),
		Timestamp: 0,
	}, nil
}

func (s *voucher) Put(acc *Account, vc *Voucher) (common.Uint256, error) {
	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), VOUCHER_CONTRACT_ADDRESS, "PutVoucher", []interface{}{vc})
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	err = s.ontSdk.SignToTransaction(imt, acc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return s.ontSdk.SendTransaction(imt)
}

func (s *voucher) GetVoucher(key string) (*sdkcom.PreExecResult, error) {
	param := t_get_voucher{
		Key: key,
	}
	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), VOUCHER_CONTRACT_ADDRESS, "GetVoucher", []interface{}{param})
	if err != nil {
		return nil, err
	}

	return s.ontSdk.PreExecTransaction(imt)
}

func (s *voucher) PutEvent(acc *Account, operator *Identity, pass []byte, voucher, title, content string) (common.Uint256, error) {
	type t_args struct {
		Key       string
		Title     string
		Content   string
		Operator  string
		Signature string
	}

	ctrl, err := operator.GetControllerByIndex(1, pass)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	bs, err := ctrl.Sign([]byte(content))
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	signature := hex.EncodeToString(bs)

	args := t_args{
		Key:       voucher,
		Title:     title,
		Content:   content,
		Operator:  operator.ID,
		Signature: signature,
	}

	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), VOUCHER_CONTRACT_ADDRESS, "AddVoucherEvent", []interface{}{args})
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	err = s.ontSdk.SignToTransaction(imt, acc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return s.ontSdk.SendTransaction(imt)
}

func (vc *voucher) GetVoucherEvents(key string) ([]*VoucherEvent, error) {
	param := t_get_voucher{
		Key: key,
	}
	imt, err := vc.native.NewNativeInvokeTransaction(0, 0, byte(0), VOUCHER_CONTRACT_ADDRESS, "GetVoucherEvents", []interface{}{param})
	if err != nil {
		return nil, err
	}

	res, err := vc.ontSdk.PreExecTransaction(imt)
	if err != nil {
		return nil, err
	}
	bs, err := res.Result.ToByteArray()
	if err != nil {
		return nil, err
	}

	source := common.NewZeroCopySource(bs)
	size, _ := source.NextByte()
	ret := []*VoucherEvent{}
	for i := 0; i < int(size); i++ {
		_ve := VoucherEvent{}
		if err := _ve.deserialization(source); err != nil {
			return ret, err
		}
		ret = append(ret, &_ve)
	}

	return ret, err
}
