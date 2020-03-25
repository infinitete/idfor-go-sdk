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

type Voucher struct {
	ontSdk *OntologySdk
	native *NativeContract
}

type t_voucher struct {
	Key       string
	Title     string
	Desc      string
	Creator   string
	Signature string
	Timestamp int64
}

type t_get_voucher struct {
	Key string
}

func (vc *t_voucher) deserialization(source *common.ZeroCopySource) error {
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

	vc.Key = key
	vc.Title = title
	vc.Creator = creator
	vc.Signature = signature
	vc.Desc = content
	vc.Timestamp = timestamp

	return nil
}

func (vc *Voucher) NewVoucher(creator *Identity, pass []byte, title, desc string) (*t_voucher, error) {
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

	return &t_voucher{
		Key:       key,
		Title:     title,
		Desc:      desc,
		Creator:   creator.ID,
		Signature: hex.EncodeToString(bs),
		Timestamp: 0,
	}, nil
}

func (s *Voucher) Put(acc *Account, vc *t_voucher) (common.Uint256, error) {
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

func (s *Voucher) GetVoucher(key string) (*sdkcom.PreExecResult, error) {
	param := t_get_voucher{
		Key: key,
	}
	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), VOUCHER_CONTRACT_ADDRESS, "GetVoucher", []interface{}{param})
	if err != nil {
		return nil, err
	}

	return s.ontSdk.PreExecTransaction(imt)
}
