/*
 * Copyright (c) 2018 VMware, Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is furnished to do
 * so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial
 * portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		s1 := RandStringBytesMaskImpr(10)
		s2 := RandStringBytesMaskImpr(10)
		if s1 == s2 {
			t.Fatalf("failed in generating random string. s1: %s, s2: %s", s1, s2)
		}
		fmt.Println(s1)
		fmt.Println(s2)
	}
}

func TestRandomNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		seed := time.Now().UTC().Second()
		s1 := RandStringBytesMaskImpr(seed)
		s2 := RandStringBytesMaskImpr(seed)
		if s1 == s2 {
			t.Fatalf("failed in generating random string. s1: %s, s2: %s", s1, s2)
		}
		fmt.Println(s1)
		fmt.Println(s2)
	}
}
