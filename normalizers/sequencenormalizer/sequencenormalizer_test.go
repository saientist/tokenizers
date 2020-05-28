// Copyright (c) 2020, NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sequencenormalizer

import (
	. "github.com/nlpodyssey/gotokenizers/normalizers"
	. "github.com/nlpodyssey/gotokenizers/normalizers/lowercasenormalizer"
	. "github.com/nlpodyssey/gotokenizers/normalizers/normalizedstring"
	. "github.com/nlpodyssey/gotokenizers/normalizers/stripnormalizer"
	"testing"
)

func TestSequenceNormalizerWithTwoNormalizers(t *testing.T) {
	sn := NewSequenceNormalizer([]Normalizer{
		NewLowerCaseNormalizer(),
		NewStripNormalizer(true, true),
	})
	ns := NewNormalizedString("  Foo Bar SÜẞ  ")
	err := sn.Normalize(ns)
	if err != nil {
		t.Error(err)
	}
	expected := "foo bar süß"
	if actual := ns.Get(); actual != expected {
		t.Errorf("expected %#v, actual %#v", expected, actual)
	}
}

func TestSequenceNormalizerWithOneNormalizer(t *testing.T) {
	sn := NewSequenceNormalizer([]Normalizer{
		NewStripNormalizer(true, true),
	})
	ns := NewNormalizedString("  foo  ")
	err := sn.Normalize(ns)
	if err != nil {
		t.Error(err)
	}
	expected := "foo"
	if actual := ns.Get(); actual != expected {
		t.Errorf("expected %#v, actual %#v", expected, actual)
	}
}

func TestSequenceNormalizerWithEmptySequence(t *testing.T) {
	sn := NewSequenceNormalizer([]Normalizer{})
	ns := NewNormalizedString("  foo  ")
	err := sn.Normalize(ns)
	if err != nil {
		t.Error(err)
	}
	expected := "  foo  "
	if actual := ns.Get(); actual != expected {
		t.Errorf("expected %#v, actual %#v", expected, actual)
	}
}
