package maps

import "testing"

func TestRange(t *testing.T) {
	var (
		example1 = map[string]interface{}{
			"asdasd": "12asdjha",
			"a2":     1223,
			"a3":     true,
			"a4":     1.234,
		}

		keys   = []string{}
		values = []interface{}{}
	)

	Range(example1, &keys, func(k, v interface{}) interface{} {
		return k
	})

	t.Error(keys)

	Range(example1, &values, func(k, v interface{}) interface{} {
		return v
	})

	t.Error(values)
}

func TestKeys(t *testing.T) {
	var (
		example1 = map[string]interface{}{
			"asdasd": "12asdjha",
			"a2":     1223,
			"a3":     true,
			"a4":     1.234,
		}

		example2 = map[interface{}]interface{}{
			"asdasd": "12asdjha",
			123:      1223,
			false:    true,
			1.123:    1.234,
		}

		example3 = map[string]interface{}{}

		keys1 = []string{}
		keys2 = []interface{}{}
		keys3 = []string{}
	)

	Keys(example1, &keys1)
	t.Error(keys1)

	Keys(example2, &keys2)
	t.Error(keys2)

	Keys(example3, &keys3)
	t.Error(keys3)
}

func TestKeysString(t *testing.T) {
	var (
		example1 = map[string]interface{}{
			"asdasd": "12asdjha",
			"azd":    1223,
			"and":    true,
			"aef":    1.234,
		}

		example2 = map[string]interface{}{}

		keys1 []string
		keys2 []string
	)

	keys1 = KeysString(example1)
	t.Error(keys1)

	keys2 = KeysString(example2)
	t.Error(keys2)
}
