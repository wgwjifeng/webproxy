//
// Copyright 2017 by Progtramder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package webproxy

type Sniffer interface {
	BeforeRequest(*Session)
	BeforeResponse(*Session)
}

