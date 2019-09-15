package maxnumberofballoons

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input string
	want  int
}{
	{input: "nlaebolko", want: 1},
	{input: "loonbalxballpoon", want: 2},
	{input: "leetcode", want: 0},
	{input: "krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuyznragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvglhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazxdbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfgghogqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjeycjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxddgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglhpudovigwvpimttiketopkvqw", want: 10},
}

func TestHelloWorld(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			actual := maxNumberOfBalloons(tt.input)
			if actual != tt.want {
				t.Errorf("%d, want: %d", actual, tt.want)
			}
		})
	}
}
