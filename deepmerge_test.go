package deepmerge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		src  interface{}
		dst  interface{}
		want interface{}
		err  bool
	}{
		{
			src:  []interface{}{1, 1, 2, 3, 5, 8},
			dst:  []interface{}{1, 1, 2, 3, 5, 8},
			want: []interface{}{1, 1, 2, 3, 5, 8, 1, 1, 2, 3, 5, 8},
			err:  false,
		},
		{
			src: []interface{}{
				1,
				"foo",
			},
			dst: []interface{}{
				"baz",
				true,
			},
			want: []interface{}{
				1,
				"foo",
				"baz",
				true,
			},
			err: false,
		},
		{
			src: map[string]interface{}{
				"hoge": "hoge",
			},
			dst: map[string]interface{}{
				"hoge": "hoge",
			},
			want: map[string]interface{}{
				"hoge": "hoge",
			},
			err: false,
		},
		{
			src: map[string]interface{}{
				"hoge":  "hoge",
				"array": []string{"hoge", "huga"},
			},
			dst: map[string]interface{}{
				"hoge":  "hoge",
				"array": []string{"hoge", "huga"},
			},
			want: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", "hoge", "huga"},
			},
			err: false,
		},
		{
			src: map[string]interface{}{
				"hoge":  "hoge",
				"array": []string{"hoge", "huga"},
			},
			dst: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", 1},
			},
			want: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", "hoge", "huga", 1},
			},
			err: false,
		},
		{
			src: map[string]interface{}{
				"hoge":  "hoge",
				"array": []string{"hoge", "huga"},
			},
			dst: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", 1},
				"byte":  []byte{0xff, 0x00},
			},
			want: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", "hoge", "huga", 1},
				"byte":  []byte{0xff, 0x00},
			},
			err: false,
		},
		{
			src: map[string]interface{}{
				"hoge":  "hoge",
				"array": []string{"hoge", "huga"},
				"byte":  []byte{0xff, 0x00},
			},
			dst: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", 1},
				"byte":  []byte{0xff, 0x00},
			},
			want: map[string]interface{}{
				"hoge":  "hoge",
				"array": []interface{}{"hoge", "huga", "hoge", "huga", 1},
				"byte":  []interface{}{[]byte{0xff, 0x00}, []byte{0xff, 0x00}},
			},
			err: false,
		},
		{
			src: []interface{}{0.5, 1, "foo"},
			dst: []string{"hoge", "huga"},
			want: []interface{}{
				0.5, 1, "foo", "hoge", "huga",
			},
			err: false,
		},
		{
			src:  "hoge",
			dst:  "hoge",
			want: "hoge",
			err:  false,
		},
		{
			src: map[string]interface{}{
				"hoge": "huga",
				"map": map[string]interface{}{
					"foo":  "bar",
					"john": "doe",
				},
			},
			dst: map[string]interface{}{
				"hoge": "huga",
				"map": map[string]interface{}{
					"fizz": "buzz",
					"john": "doe",
				},
			},
			want: map[string]interface{}{
				"hoge": "huga",
				"map": map[string]interface{}{
					"foo":  "bar",
					"fizz": "buzz",
					"john": "doe",
				},
			},
			err: false,
		},
		{
			src: []string{"hoge", "huga"},
			dst: map[string]interface{}{},
			err: true,
		},
	}

	for _, test := range tests {
		got, err := Merge(test.src, test.dst)
		input := []interface{}{
			test.src,
			test.dst,
		}
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v but: %v", input, err)
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v but not:", input)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want: %v, but: %v", test.want, got)
		}
	}
}
