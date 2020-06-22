package GoatLatin

import "testing"

func TestToGoatLatin(t *testing.T) {
	input := "I speak Goat Latin"
	output := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
	res := toGoatLatin(input)
	t.Log(res)
	if res != output {
		t.Fatalf("test case %s error ", input)
	}
	input = "The quick brown fox jumped over the lazy dog"
	output = "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
	res = toGoatLatin(input)
	t.Log(res)
	if res != output {
		t.Fatalf("test case %s error ", input)
	}
	input = "   "
	t.Log(toGoatLatin(input))
}
