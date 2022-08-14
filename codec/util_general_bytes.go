// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package codec

func readBytes(bytes []byte, idx int) ([]byte, int) {
	length, idx := readInt(bytes, idx)
	return bytes[idx : idx+length], idx + length
}

func readCompactBytes(bytes []byte, idx int) ([]byte, int) {
	auxUInt32, offset := readUVarint(bytes, idx)
	intLen := int(auxUInt32)
	return bytes[offset : idx+intLen], idx + intLen
}

func putBytes(bytes []byte, idx int, authBytes []byte) int {
	idx = putInt(bytes, idx, len(authBytes))
	copy(bytes[idx:], authBytes)
	return idx + len(authBytes)
}

func putCompactBytes(bytes []byte, idx int, compactBytes []byte) int {
	idx = putUVarint(bytes, idx, uint32(len(compactBytes)+1))
	copy(bytes[idx:], compactBytes)
	return idx + len(compactBytes)
}

func putCompactNullableBytes(bytes []byte, idx int, content []byte) int {
	if content == nil {
		return putUVarint(bytes, idx, 0)
	}
	idx = putUVarint(bytes, idx, uint32(CompactBytesLen(content)))
	copy(bytes[idx:], content)
	return idx + len(content)
}

func putVCompactBytes(bytes []byte, idx int, authBytes []byte) int {
	idx = putVarint(bytes, idx, len(authBytes))
	copy(bytes[idx:], authBytes)
	return idx + len(authBytes)
}

func BytesLen(bytes []byte) int {
	return 4 + len(bytes)
}

func CompactBytesLen(bytes []byte) int {
	return uVarintSize(uint(len(bytes))) + len(bytes)
}

func CompactNullableBytesLen(bytes []byte) int {
	if bytes == nil {
		return 1
	}
	return uVarintSize(uint(len(bytes))) + len(bytes)
}
