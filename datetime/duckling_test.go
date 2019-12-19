// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package datetime

import "testing"

import "time"

func TestDuckling(t *testing.T) {
	d, err := NewDuckling()
	if err != nil {
		t.Error("Error while initalizing the duckling service", err)
		return
	}
	ch := d.Query([]rune("this month"))
	res := <-ch
	if len(res.Res) == 0 {
		t.Error("Expected today to resolved. Got empty result")
		return
	}

	if !res.Res[0].IsValid() {
		t.Error("Expected today to resolved. Didn't a valid response")
		return
	}

	if res.Res[0].Value.Time.Month() != time.Now().Month() {
		t.Error("Expected today to be resolved today's date", time.Now().Month(), "got", res.Res[0].Value.Time.Month())
		return
	}
}
