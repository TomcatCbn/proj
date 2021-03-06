// Copyright (C) 2018, Michael P. Gerlek (Flaxen Consulting)
//
// Portions of this code were derived from the PROJ.4 software
// In keeping with the terms of the PROJ.4 project, this software
// is provided under the MIT-style license in `LICENSE.md` and may
// additionally be subject to the copyrights of the PROJ.4 authors.

package merror_test

import (
	"testing"

	"github.com/go-spatial/proj/merror"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	showSource := merror.ShowSource
	merror.ShowSource = true
	defer func() { merror.ShowSource = showSource }()

	assert := assert.New(t)

	err1 := merror.New("errtest-%d", 1)
	assert.Error(err1)
	exp1 := "errtest-1 (from merror_test.TestError at Error_test.go:24)"
	assert.Equal(exp1, err1.Error())

	err2 := merror.Wrap(err1, "errtest-%d", 2)
	assert.Error(err2)
	exp2 := "errtest-2 (from merror_test.TestError at Error_test.go:29)"
	exp2 += " // Inner: " + exp1
	assert.Equal(exp2, err2.Error())

	err3 := merror.Wrap(err2)
	assert.Error(err3)
	exp3 := "wrapped error (from merror_test.TestError at Error_test.go:35)"
	exp3 += " // Inner: " + exp2
	assert.Equal(exp3, err3.Error())

	err4 := merror.Pass(err2)
	assert.Error(err4)
	assert.Equal(exp2, err4.Error())
}
