/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package boltdb

import (
	"fmt"
	"testing"
)

func TestL5Store_GenNextL5Sid(t *testing.T) {
	handler, err := NewBoltHandler(&BoltConfig{FileName: "./table.bolt"})
	if nil != err {
		t.Fatal(err)
	}
	defer handler.Close()

	l5store := &l5Store{handler: handler}

	if err = l5store.InitL5Data(); nil != err {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		sid, err := l5store.GenNextL5Sid(uint32(i+1) % 6)
		if nil != err {
			t.Fatal(err)
		}
		fmt.Printf("sid %d is %s\n", i, sid)
	}
}
