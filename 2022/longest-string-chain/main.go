package main

import "fmt"

// Time: O(n^2*k) where n is len(words) & k is the largest word length
// Space: O(n)
//
// The first intuition is that predecessors have one less character, so a chain of strings
// must increase length one by one.
//
// There doesn't seem to be a good way to speed up checking if a string is a predecessor of another
// one, so for every necessary predecessor check, it's gonna be linear time, constant space.
//
// 0) bucket words by length, so we can do step 1.
// 1) compare all words of each length to all words of length+1, forming a graph of
//    successors.
// 2) this graph is really an array of trees (because it's directed with no cycles!), so, for each
//    tree, calculate the max depth, and the largest one is the answer!
func longestStrChain(words []string) int {
	var (
		// (0) bucket words by length (linear time, linear space)
		lengthToWords = bucketByLength(words)
		graph         = make([][]int, len(words))
	)
	// (1) Build graph of successors (time is O(len(words)^2*k) where k is len of longest word, linear space)
	for length, idxsOfWords := range lengthToWords {
		for _, i1 := range idxsOfWords {
			for _, i2 := range lengthToWords[length+1] {
				if isPredecessor(words[i1], words[i2]) {
					graph[i1] = append(graph[i1], i2)
				}
			}
		}
	}

	// (2) calculate max length of graph which is really an array of trees (linear time with memo)
	return largestLengthPathInGraph(graph)
}

func bucketByLength(words []string) map[int][]int {
	buckets := map[int][]int{}
	for i, word := range words {
		buckets[len(word)] = append(buckets[len(word)], i)
	}
	return buckets
}

func isPredecessor(w1, w2 string) bool {
	// Compare with "two pointers" allowing up to one difference
	var diff int
	for i := 0; i < len(w1); i++ {
		if w1[i] == w2[i+diff] {
			continue
		}
		if diff == 1 {
			return false
		}
		diff++
		i--
	}
	// Extra edge case not contemplated by loop above: difference is in the last character
	if diff == 1 && w1[len(w1)-1] != w2[len(w2)-1] {
		return false
	}

	return true
}

func largestLengthPathInGraph(graph [][]int) int {
	mx := 0
	for idx := range graph {
		mx = max(mx, maxDepthOfTree(idx, graph, map[int]int{}))
	}
	return mx
}

// Added memoisation, because in an unfortunate case of a word being successor of a million predecessors,
// all continuations of those trees will be recomputed multiple times. This memoisation ensures
// each cell in the entire graph will be visited once, making time linear.
func maxDepthOfTree(idx int, graph [][]int, memo map[int]int) int {
	if len(graph[idx]) == 0 {
		return 1
	}
	if md, ok := memo[idx]; ok {
		return md
	}
	var maxDepth int
	for _, successorIdx := range graph[idx] {
		maxDepth = max(maxDepth, 1+maxDepthOfTree(successorIdx, graph, memo))
	}
	memo[idx] = maxDepth
	return memo[idx]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    []string
		expected int
	}{
		// {
		// 	input:    []string{"a", "b", "ba", "bca", "bda", "bdca"},
		// 	expected: 4,
		// },
		// {
		// 	input:    []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
		// 	expected: 5,
		// },
		// {
		// 	input:    []string{"abcd", "dbqca"},
		// 	expected: 1,
		// },
		// {
		// 	input:    []string{"qyssedya", "pabouk", "mjwdrbqwp", "vylodpmwp", "nfyqeowa", "pu", "paboukc", "qssedya", "lopmw", "nfyqowa", "vlodpmw", "mwdrqwp", "opmw", "qsda", "neo", "qyssedhyac", "pmw", "lodpmw", "mjwdrqwp", "eo", "nfqwa", "pabuk", "nfyqwa", "qssdya", "qsdya", "qyssedhya", "pabu", "nqwa", "pabqoukc", "pbu", "mw", "vlodpmwp", "x", "xr"},
		// 	expected: 8,
		// },
		{
			input:    []string{"uiykgmcc", "jrgbss", "mhkqodcpy", "lkj", "bwqktun", "s", "nrctyzifwytjblwy", "wrp", "scqlcwmxw", "irqvnxdcxoejuu", "gmlckvofwyifmrw", "wbzbyrcppaljigvo", "lk", "kfeouqyyrer", "efzzpvi", "ubkcitcmwxk", "txihn", "mdwdmbtx", "vuzvcoaif", "jwmboqvhpqodsj", "wscfvrfl", "pzye", "waxyoxftvrgqmkg", "wwdidopozinxxn", "dclpg", "xjsvlxktxs", "ajj", "pvsdastm", "tatjxhygidhn", "feafycxdxagn", "irqvnxxoeuu", "kwjo", "tztoovsyfwz", "prllrw", "sclmx", "bbmjnwaxcwaml", "gl", "wiax", "uzvcoaif", "ztovyfwz", "qxy", "zuexoxyp", "qxyyrl", "pvsdasvtm", "femafycxdxaagn", "rspvccjcm", "wvyiax", "vst", "efzi", "fjmdcc", "icsinrbpql", "ctybiizlcr", "ntyzfwytjblw", "tatjxhygidhpn", "e", "kykizdandafusu", "pnepuwcsxl", "kfeuqyyrer", "afplzhbqguu", "hvajtj", "prll", "ildzdimea", "zueoxp", "ezi", "lqr", "jkaagljikwamaqvf", "mlzwhkxsn", "rspvccbcjjtcm", "wscfvrl", "m", "msygukwlkrqboc", "pifojogoveub", "bkcmwx", "jercgybhss", "wrpi", "aicsinkgrbpqli", "aplzbuu", "sclcmxw", "atpepgsz", "govrcuuglaer", "bdxjpsvlxkytxs", "uikgm", "bm", "wvyhiqax", "znvaasgfvqi", "hatpepgsz", "hrzebpa", "bnfz", "lybtqrfzw", "taxhygihn", "bjnfzk", "mhqp", "ide", "znvcaasgfvqi", "ftv", "afplzhbqsguuu", "thn", "pdbccbe", "mxevopfoimgjww", "fjmdrcce", "rspvccjjcm", "jv", "motnfwohwule", "xjsvlxtxs", "bqeb", "eug", "jftavwgl", "rzebpa", "lybtqrfazw", "zuexoxp", "jercgybhsys", "hajtj", "bkcitcmwxk", "mbpvxsdastvtm", "mowlznwhkxsn", "dvenn", "rsacxe", "tatjxhygihn", "cotybiizlcr", "bbmnaxaml", "pkwrpsi", "nqpdbccbkxens", "mbpbovxsdastvtm", "mj", "pxpsvikwekuq", "qeug", "dmelddga", "aicsinkgrbpxqli", "bdxjpsvlxktytxs", "pkrllrxw", "jkgljikwmaqf", "iddie", "ctybiizcr", "nyzfwytjblw", "yvuhmiuehspi", "keuqre", "wzbypaigvo", "sck", "uzcoaf", "dlpg", "ubkcpitlscmwxk", "molzwhkxsn", "pepuwcsxl", "laplm", "dclpgc", "mahkxqodcpy", "sclcmx", "hvrzebpaz", "bgovrcuuglaer", "clazpulmw", "yvuyhmiuehspiq", "wzbycpaljigvo", "sceqalciwmxw", "hjytflmvsgv", "u", "hjyvxytfflhmvsgv", "jkgjikwmaqf", "fefycxdxagn", "ftvw", "ofncgxrkqvcr", "spvcjc", "pvsdastvtm", "kykzdandaus", "wbzbycppaljigvo", "haytpepgsz", "jmowlznwhkxsn", "aplzhbguu", "zvyz", "nfvqi", "jfvtavwsgl", "xejnllhfulns", "zhhvbiqiw", "jkgljikwmaqvf", "tyizc", "irqvnxcxoejuu", "clvazzpulmw", "oncgxrqvcr", "qlupvpdkhrm", "mtnfwohwule", "wwdidopzozinxxn", "auiykgmcc", "wscfvrfyl", "pfksmrullrxw", "jwmoqvhpqods", "ftavwg", "iddiea", "kcmw", "ykkwjwo", "pe", "aplzbguu", "eu", "bbmnaxal", "ntyswtnlab", "zhhhvbhbiqiw", "jwmoqvpqods", "kykzdndaus", "bbmjnaxcwaml", "zunvcaasgfvqi", "icsingrbpql", "sceqalciwmsxyw", "yvuhmiuehsp", "bxjsvlxktxs", "waxoxftvrgqmkg", "cogxxpaknks", "scllvazzpulmw", "tatjxhygeidhpn", "ftvwg", "tyz", "nafvqi", "oby", "pgzpkhqog", "irqvnxxoejuu", "oxwpkxlakcp", "bnf", "oxwnpkxlakcp", "bwqktu", "ufybbaozoqk", "ntydswtnlab", "zvyfz", "znaafvqi", "npdbccbke", "mhkqocpy", "kuq", "bjnfz", "taxhyihn", "kwrpsi", "qifepmcatbdjlf", "lzwhks", "kfeuqre", "mxevopfoimgww", "spvcjcm", "oncgxrkqvcr", "jftavwsgl", "soifcbya", "jpzyeg", "jwmboqvhpqods", "lapulm", "jrgbhss", "xejfnllhfulns", "zhhhvbbiqiw", "km", "kuqre", "scxlzlvazzpulmw", "ztvyfwz", "wbzbycpaljigvo", "rzbpa", "vsastm", "uybaooqk", "dn", "ykwjwo", "ufybmvbaozoqk", "nknm", "mbpvsdastvtm", "dpgzpxykhqog", "wzbypajigvo", "bnjnfzk", "eollbigtftpdrd", "zhbiqiw", "yvuhiuehp", "zhhhvbhbiqiwg", "pfksrullrxw", "pzyeg", "aplzhbqguu", "z", "hvrzecbpazw", "clvazpulmw", "tajxhygihn", "pgzpxykhqog", "fefyxdxagn", "wimomuvhn", "lqrzw", "xejnlhfulns", "jhrc", "xsxxs", "slmx", "jrgss", "uikgmc", "ncgqvcr", "womuhn", "aryouvtnmme", "uzco", "zhhhvbiqiw", "hjytflhmvsgv", "znvaasfvqi", "kuqr", "ojrpp", "ztoovyfwz", "zvz", "pxpsviweuq", "ufybaooqk", "xy", "jfvvtavwksvgl", "raiachv", "bmnaxl", "rspvccjjtcm", "pgzpxkhqog", "xhbtfnqebaj", "sceqalciwmsxw", "jssctk", "uzvcoaf", "fefydxagn", "jhrvc", "mbj", "raiahv", "nrtyzifwytjblwy", "mhqcp", "jkgjkwmaqf", "wscfvrfylhi", "lqrz", "ahabucermswyrvl", "wxoxftvrgqmkg", "ku", "uyaoq", "mhqocp", "ykwjo", "vstm", "ofncgxrkqvcwr", "dqvh", "taxyihn", "idie", "bwqtu", "tztoovyfwz", "rspvcccjjtcm", "uojrpp", "wmomuhn", "cotycbiizlxcr", "nrtyzfwytjblw", "ocbya", "sceqlciwmxw", "ajtj", "rspvccbcjjthcm", "kfeuqyyre", "dmelddg", "txyihn", "ubkcitlscmwxk", "ntyswtnla", "bdxjpstvlxktytxs", "odqdvbh", "pxpsvikeewekuq", "mdwdmbdtux", "vs", "bma", "wzbypigvo", "qxyy", "vsstm", "hbtnqeba", "hrzebpaz", "xhbtfnjsqebbaj", "ahaucermswyrv", "ddmbtx", "zhhbiqiw", "pxpsvikewekuq", "odqdvgbh", "bxjpsvlxktxs", "jsck", "fjmdc", "mdwdmbdtx", "jqxyyrl", "pxpsvikweuq", "ctybizcr", "dqvbh", "lpl", "lqrfzw", "ufybaozoqk", "znvaafvqi", "yvuhmiuehp", "hvrzebpazw", "pfksrllrxw", "alzuu", "xjsvxtxs", "afplzhbqguuu", "icsingrbpqli", "hjxytflhmvsgv", "femafycxdxagn", "uyaoqk", "gmlckvofwyifrw", "cinrbpql", "jrcgbhss", "oxwpkxlkcp", "jkagljikwamaqvf", "eollbigtftpdrdy", "rspvcjcm", "socbya", "clapulm", "qeb", "kwrpi", "efzpi", "hbtfnqebaj", "kykizdnandafusu", "sclvazzpulmw", "efzzpvvi", "jfvvtavwsvgl", "mhqocpy", "v", "mbpbvxsdastvtm", "irqvnxouu", "hvaajtj", "ofnlcgxrkqvcwr", "hbtqeba", "hbtqeb", "jwmqpds", "ntrnlhujdslco", "zv", "npdbccbken", "mhp", "ddb", "prllw", "mddmbtx", "clazpulm", "cogxxpaknkse", "bkitcmwxk", "oxwpklkcp", "tyiz", "jwmqvpqods", "waxyoxftvrgqmkgb", "afplzhbbqsgujuu", "bwtu", "jercgbhss", "rsacx", "mahkqodcpy", "cotycbiizlcr", "ahabucermswyrv", "lupvpkhr", "dvnn", "b", "atpepsz", "ncgxqvcr", "qe", "ubkcitlcmwxk", "lyqrfzw", "wimomuhn", "bbmnaxl", "motnfwohrwule", "yvuyhmiuehspi", "jfvvtavwsgl", "rac", "fefdxagn", "bwqkctun", "uotjrpp", "ddbtx", "afplzhbbqsguuu", "xss", "xsxs", "wvyiqax", "kykizdandaus", "npdbccbkens", "r", "oxwnpkxjlakcp", "tzmteoovsyfwz", "kykizdnandafuspu", "ahabulcermswyrvl", "xjsxxs", "qxyyr", "ck", "xhbtfnqebbaj", "nqpdbccbkens", "mpvsdastvtm", "zuexqoxyp", "gmlkvofwyifrw", "kmw", "txhn", "kykizdandausu", "molznwhkxsn", "lupvpdkhr", "jwmqvpds", "bktcmwx", "wyiax", "hzvaajtj", "ddbx", "pifojogveub", "naafvqi", "motnfwjohrwule", "odqvbh", "aicsingrbpqli", "jopzyeg", "lybtqrfazrw", "pijogveub", "xzejfnllhfulns", "scxllvazzpulmw", "irqyvnxdcxfoejuu", "cogxpaknks", "pdkwrpsi", "wzbycpajigvo", "xjsxtxs", "irqvnxdcxfoejuu", "xhbtfnjqebbaj", "uybaoqk", "oncgxqvcr", "aj", "pepuwsxl", "lytqrfzw", "nkm", "jrgs", "pkrllrw", "wscfvrfyli", "bbmjnaxcaml", "jftavwg", "vuzvcozaif", "pifjogveub", "cmogxxpaknkse", "cinrbql", "scqlciwmxw", "ztvyfz", "mxyevopfoimgjpww", "soicbya", "lupvpdkhrm", "ahaucermsyrv", "ufybmvbaouzoqk", "bdxjpsvlxktxs", "hjxytfflhmvsgv", "hjvxytfflhmvsgv", "nqpdbccbzkxens", "wr", "kykzdndus", "iddimea", "fjmdrcc", "efzzpi", "vsdastm", "btqeb", "pfkrllrxw", "ocby", "irqvnxxouu", "ildzpdimea", "lzwhkxsn", "ilddimea", "ufybvbaozoqk", "mxyevopfoimgjww", "jhr", "kcmwx", "dvn", "uzcof", "glw", "hbtnqebaj", "riahv", "w", "qeugv", "kfeuqyre", "ilrdzpdimea", "lplm", "icinrbpql", "scqlcmxw", "bbmjnaxaml", "e", "rsac", "bf", "jwmqvpqds", "tzteoovsyfwz", "rc", "lzwhkxs", "jkgljikwamaqvf", "tybizc", "aplzuu", "nrtyzifwytjblw", "pze", "bktcmwxk", "uiykgmc", "jsctk", "npdbccbe", "tybizcr"},
			expected: 15,
		},
	}
	for _, tc := range ts {
		actual := longestStrChain(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
