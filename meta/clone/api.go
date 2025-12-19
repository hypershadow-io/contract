package clone

import (
	"github.com/hypershadow-io/contract/json"
	"github.com/hypershadow-io/contract/meta"
)

// Clone - make a copy of Meta.
func Clone(m meta.Meta) (meta.Meta, error) {
	if !m.IsValid() {
		return nil, nil
	}
	if m.IsZero() {
		return make(meta.Meta), nil
	}
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var mCopy meta.Meta
	if err = json.Unmarshal(data, &mCopy); err != nil {
		return nil, err
	}
	return mCopy, nil
}

// CloneSilent attempt to make a copy and return the original in case of an error.
func CloneSilent(m meta.Meta) meta.Meta {
	if mCopy, err := Clone(m); err == nil {
		return mCopy
	}
	return m
}
