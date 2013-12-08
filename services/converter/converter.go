// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"github.com/andreaskoch/allmark2/model"
)

type Converter interface {
	Convert(item *model.Item) (*model.Item, error)
}
