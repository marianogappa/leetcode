package main

import "fmt"

// Time: O(n^2)
// Space: O(n^2)
func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
	}
	// This exercise requires some mental work to figure out why
	// a letter match makes the value at memo[i1][i2] +1 over the
	// previous letter on both strings, and why a non-match
	// carries the largest of each of the previous letters.
	for i1 := 1; i1 <= len(text1); i1++ {
		for i2 := 1; i2 <= len(text2); i2++ {
			if text1[i1-1] == text2[i2-1] {
				memo[i1][i2] = 1 + memo[i1-1][i2-1]
			} else {
				memo[i1][i2] = max(memo[i1][i2-1], memo[i1-1][i2])
			}
		}
	}
	return memo[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		text1    string
		text2    string
		expected int
	}{
		{
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			text1:    "fcvafurqjylclorwfoladwfqzkbebslwnmpmlkbezkxoncvwhstwzwpqxqtyxozkpgtgtsjobujezgrkvevklmludgtyrmjaxyputqbyxqvupojutsjwlwluzsbmvyxifqtglwvcnkfsfglwjwrmtyxmdgjifyjwrsnenuvsdedsbqdovwzsdghclcdexmtsbexwrszihcpibwpidixmpmxshwzmjgtadmtkxqfkrsdqjcrmxkbkfoncrcvoxuvcdytajgfwrcxivixanuzerebuzklyhezevonqdsrkzetsrgfgxibqpmfuxcrinetyzkvudghgrytsvwzkjulmhanankxqfihenuhmfsfkfepibkjmzybmlkzozmluvybyzsleludsxkpinizoraxonmhwtkfkhudizepyzijafqlepcbihofepmjqtgrsxorunshgpazovuhktatmlcfklafivivefyfubunszyvarcrkpsnglkduzaxqrerkvcnmrurkhkpargvcxefovwtapedaluhclmzynebczodwropwdenqxmrutuhehadyfspcpuxyzodifqdqzgbwhodcjonypyjwbwxepcpujerkrelunstebopkncdazexsbezmhynizsvarafwfmnclerafejgnizcbsrcvcnwrolofyzulcxaxqjqzunedidulspslebifinqrchyvapkzmzwbwjgbyrqhqpolwjijmzyduzerqnadapudmrazmzadstozytonuzarizszubkzkhenaxivytmjqjgvgzwpgxefatetoncjgjsdilmvgtgpgbibexwnexstipkjylalqnupexytkradwxmlmhsnmzuxcdkfkxyfgrmfqtajatgjctenqhkvyrgvapctqtyrufcdobibizihuhsrsterozotytubefutaxcjarknynetipehoduxyjstufwvkvwvwnuletybmrczgtmxctuny",
			text2:    "nohgdazargvalupetizezqpklktojqtqdivcpsfgjopaxwbkvujilqbclehulatshehmjqhyfkpcfwxovajkvankjkvevgdovazmbgtqfwvejczsnmbchkdibstklkxarwjqbqxwvixavkhylqvghqpifijohudenozotejoxavkfkzcdqnoxydynavwdylwhatslyrwlejwdwrmpevmtwpahatwlaxmjmdgrebmfyngdcbmbgjcvqpcbadujkxaxujudmbejcrevuvcdobolcbstifedcvmngnqhudixgzktcdqngxmruhcxqxypwhahobudelivgvynefkjqdyvalmvudcdivmhghqrelurodwdsvuzmjixgdexonwjczghalsjopixsrwjixuzmjgxydqnipelgrivkzkxgjchibgnqbknstspujwdydszohqjsfuzstyjgnwhsrebmlwzkzijgnmnczmrehspihspyfedabotwvwxwpspypctizyhcxypqzctwlspszonsrmnyvmhsvqtkbyhmhwjmvazaviruzqxmbczaxmtqjexmdudypovkjklynktahupanujylylgrajozobsbwpwtohkfsxeverqxylwdwtojoxydepybavwhgdehafurqtcxqhuhkdwxkdojipolctcvcrsvczcxedglgrejerqdgrsvsxgjodajatsnixutihwpivihadqdotsvyrkxehodybapwlsjexixgponcxifijchejoxgxebmbclczqvkfuzgxsbshqvgfcraxytaxeviryhexmvqjybizivyjanwxmpojgxgbyhcruvqpafwjslkbohqlknkdqjixsfsdurgbsvclmrcrcnulinqvcdqhcvwdaxgvafwravunurqvizqtozuxinytafopmhchmxsxgfanetmdcjalmrolejidylkjktunqhkxchyjmpkvsfgnybsjedmzkrkhwryzan",
			expected: 323,
		},
	}
	for _, tc := range ts {
		actual := longestCommonSubsequence(tc.text1, tc.text2)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.text1, tc.text2, tc.expected, actual)
		}
	}
}
