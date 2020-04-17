/*
Package layout represents the strucutre of project files when marshaled.

goals:
  - directly represent desired data in a user-focused manner
    - probably stick with YAML for now
  - be usable with as few fields set as possible
    - construction funcs secondary as helpers OK
  - provide hydration methods like `lo.ParseYAML([]byte) error`
  - provide validation method(s) like `lo.Validate() error`
  - possibly provide normalization method(s) like `lo.Normalize() error`
  - possibly provide convenience method(s) like `lo.ValidateNormalize() error`
    - with better name(s)
  - rely on other, focused, pkgs for string->valid instance->string behavior
    - for example: Namespace
    - types like Namespace should be accessible without dep on this pkg
  - provide type for "resaving" like `[]byte.FromOriginal(d []byte) ([]byte, error)`
    - by requiring the original data, order can be preserved
  - avoid lifetime management - wrap this package for that functionality
  - avoid file handling itself - either wrap or push back to callers
    - note that FromOriginal receives and returns []byte
*/
package layout
