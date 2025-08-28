package impl

// limitNil provides a no-op implementation of the schema.Limit interface.
// It represents the absence of any limit constraints and always returns zero values.
type limitNil struct{}

func (a limitNil) IsValid() bool         { return false }
func (a limitNil) GetEnum() []string     { return nil }
func (a limitNil) GetMin() *float64      { return nil }
func (a limitNil) IsExclusiveMin() bool  { return false }
func (a limitNil) GetMax() *float64      { return nil }
func (a limitNil) IsExclusiveMax() bool  { return false }
func (a limitNil) GetMultiple() *float64 { return nil }
func (a limitNil) IsUnique() bool        { return false }
func (a limitNil) GetPattern() string    { return "" }
func (a limitNil) GetFormat() string     { return "" }
