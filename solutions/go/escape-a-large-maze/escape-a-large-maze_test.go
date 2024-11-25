package escapealargemaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEscapePossible(t *testing.T) {
	cases := []struct {
		blocked [][]int
		source  []int
		target  []int
		result  bool
	}{
		{
			blocked: [][]int{{0, 1}, {1, 0}}, source: []int{0, 0}, target: []int{0, 2}, result: false,
		},
		{
			blocked: [][]int{}, source: []int{0, 0}, target: []int{999999, 999999}, result: true,
		},
		{
			blocked: [][]int{{100005, 100073}, {100006, 100074}, {100007, 100075}, {100008, 100076}, {100009, 100077}, {100010, 100078}, {100011, 100079}, {100012, 100080}, {100013, 100081}, {100014, 100082}, {100015, 100083}, {100016, 100084}, {100017, 100085}, {100018, 100086}, {100019, 100087}, {100020, 100088}, {100021, 100089}, {100022, 100090}, {100023, 100091}, {100024, 100092}, {100025, 100091}, {100026, 100090}, {100027, 100089}, {100028, 100088}, {100029, 100087}, {100030, 100086}, {100031, 100085}, {100032, 100084}, {100033, 100083}, {100034, 100082}, {100035, 100081}, {100036, 100080}, {100037, 100079}, {100038, 100078}, {100039, 100077}, {100040, 100076}, {100041, 100075}, {100042, 100074}, {100043, 100073}, {100044, 100072}, {100043, 100071}, {100042, 100070}, {100041, 100069}, {100040, 100068}, {100039, 100067}, {100038, 100066}, {100037, 100065}, {100036, 100064}, {100035, 100063}, {100034, 100062}, {100033, 100061}, {100032, 100060}, {100031, 100059}, {100030, 100058}, {100029, 100057}, {100028, 100056}, {100027, 100055}, {100026, 100054}, {100025, 100053}, {100024, 100052}, {100023, 100053}, {100022, 100054}, {100021, 100055}, {100020, 100056}, {100019, 100057}, {100018, 100058}, {100017, 100059}, {100016, 100060}, {100015, 100061}, {100014, 100062}, {100013, 100063}, {100012, 100064}, {100011, 100065}, {100010, 100066}, {100009, 100067}, {100008, 100068}, {100007, 100069}, {100006, 100070}, {100005, 100071}},
			source:  []int{100024, 100072},
			target:  []int{999994, 999990},
			result:  true,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, isEscapePossible(tc.blocked, tc.source, tc.target))
		})
	}
}
