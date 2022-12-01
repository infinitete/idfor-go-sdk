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
	sdkcom "github.com/infinitete/idfor-go-sdk/common"
	"github.com/infinitete/idfor/common"
)

type Storage struct {
	ontSdk *OntologySdk
	native *NativeContract
}

type storageKey struct {
	Key []byte
}

type storageKeyValue struct {
	Key   []byte
	Value []byte
}

func (s *Storage) Put(acc *Account, key []byte, value []byte) (common.Uint256, error) {
	param := storageKeyValue{
		Key:   key,
		Value: value,
	}
	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), STORAGE_CONTRACT_ADDRESS, "Put", []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	err = s.ontSdk.SignToTransaction(imt, acc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return s.ontSdk.SendTransaction(imt)
}

func (s *Storage) Get(key []byte) (*sdkcom.PreExecResult, error) {
	param := storageKey{
		Key: key,
	}
	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), STORAGE_CONTRACT_ADDRESS, "Get", []interface{}{param})
	if err != nil {
		return nil, err
	}

	return s.ontSdk.PreExecTransaction(imt)
}

func (s *Storage) Delete(acc *Account, key []byte) (common.Uint256, error) {
	param := storageKeyValue{
		Key: key,
	}
	imt, err := s.native.NewNativeInvokeTransaction(0, 0, byte(0), STORAGE_CONTRACT_ADDRESS, "Delete", []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	err = s.ontSdk.SignToTransaction(imt, acc)
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return s.ontSdk.SendTransaction(imt)
}
