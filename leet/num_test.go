package leet

import (
	"fmt"
	"reflect"
	"testing"
)
//输入：n = 3
//输出：3
//示例 2：
//
//输入：n = 11
//输出：0
//解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。

func TestNum400(t *testing.T)  {
	datas := []struct{
		n, ans int
	}{
		{3, 3},
		{11, 0},
	}
	for _, x:= range datas {
		fmt.Println(reflect.DeepEqual(findNthDigit(x.n),x.ans))
	}
}
func TestNum34(t *testing.T)  {
	datas := []struct{
		nums, ans []int
		target int
	}{
		{[]int{5,7,7,8,8,10}, []int{3, 4}, 8},
		{[]int{5,7,7,8,8,10}, []int{-1,-1}, 6},
	}
	for _, x:= range datas {
	   fmt.Println(reflect.DeepEqual(searchRange(x.nums, x.target),x.ans))
	}
}
func TestNum33(t *testing.T)  {
	//[4,5,6,7,0,1,2]
	//0
	datas := []struct {
		nums   []int
		target int
		ans int
	} {
		{[]int{4,5,6,7,0,1,2}, 0, 4},
		{[]int{1}, 0, -1},
		{[]int{3,5,1}, 1, 2},
	}
	for _, tt :=range datas {
		fmt.Println(search(tt.nums,tt.target) == tt.ans)
	}
}
func TestNum32(t *testing.T)  {
   //fmt.Println(longestValidParentheses(")("))
   //fmt.Println(longestValidParentheses("())())(()"))
   //fmt.Println(longestValidParentheses("()(()"))
   //fmt.Println(longestValidParentheses("(())(()"))
   //fmt.Println(longestValidParentheses("(())(()"))
   fmt.Println(longestValidParentheses(")(()())"))//6
   fmt.Println(longestValidParentheses("(())((()))"))//10
   fmt.Println(longestValidParentheses(")()())"))//4
}
//[1,3,2]
//预期结果：
//[2,1,3]
func TestNum31(t *testing.T)  {
	ns := []int{}
	reverse(ns)
	fmt.Println(ns)
	nums:= []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
func TestNum30(t *testing.T)  {
	//stat := 0b1100
	//findAFreeWord([]int{3, 0, 4, 0}, []int{2,1,1,1}, 2, &stat)
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"foo","bar"}))//[0,9]
	s:= "ejwwmybnorgshugzmoxopwuvshlcwasclobxmckcvtxfndeztdqiakfusswqsovdfwatanwxgtctyjvsmlcoxijrahivwfybbbudosawnfpmomgczirzscqvlaqhfqkithlhbodptvdhjljltckogcjsdbbktotnxgwyuapnxuwgfirbmdrvgapldsvwgqjfxggtixjhshnzphcemtzsvodygbxpriwqockyavfscvtsewyqpxlnnqnvrkmjtjbjllilinflkbfoxdhocsbpirmcbznuioevcojkdqvoraeqdlhffkwqbjsdkfxstdpxryixrdligpzldgtiqryuasxmxwgtcwsvwasngdwovxzafuixmjrobqbbnhwpdokcpfpxinlfmkfrfqrtzkhabidqszhxorzfypcjcnopzwigmbznmjnpttflsmjifknezrneedvgzfmnhoavxqksjreddpmibbodtbhzfehgluuukupjmbbvshzxyniaowdjamlfssndojyyephstlonsplrettspwepipwcjmfyvfybxiuqtkdlzqedjxxbvdsfurhedneauccrkyjfiptjfxmpxlssrkyldfriuvjranikluqtjjcoiqffdxaukagphzycvjtvwdhhxzagkevvuccxccuoccdkbboymjtimdrmerspxpktsmrwrlkvpnhqrvpdekmtpdfuxzjwpvqjjhfaupylefbvbsbhdncsshmrhxoyuejenqgjheulkxjnqkwvzznriclrbzryfaeuqkfxrbldyusoeoldpbwadhrgijeplijcvqbormrqglgmzsprtmryvkeevlthvflsvognbxfjilwkdndyzwwxgdbeqwlldyezmkopktzugxgkklimhhjqkmuaifnodtpredhqygmedtqpezboimeuyyujfjxkdmbjpizpqltvgknnlodtbhnbhjkmuhwxvzgmkhbcvvadhnssbvneecglnqxhavhvxpkjxlluilzpysjcnwguyofnhfvhaceztoiscumkhociglkvispihvyoatxcxbeqsmluixgsliatukrecgoldmzfhwkgaqzsckonjuhxdhqztjfxstjvikdrhpyjfxbjjryslfpqoiphrwfjqqhaamrjbrsiovrxmqsyxhqmritjeauwqbwtpqcqhvyyssvfknfhxvtodpzipueixdbntdfcaeatyyainfpkclbgaaqrwwzwbcjwiqzkwzfuxfclmsxpdyvfbnwxjytnaejivivriamhgqsskqhnqeurttrfrmstrbeokzhuzvbfmwywohmgogyhzpmsdemugqkspsmoppwbnwabdmiruibwznqcuczculujfiavzwynsyqxmarjkshjhxobandwyzggjibjgzyaaqxorqxbkenscbveqbaociwmqxxyzvyblypeongzrttvwqzmrccwkzidyfbxcaypyquodcpwxkstbthuvjqgialhfmgjohzoxvdaxuywfqrgmyahhtpqtazbphmfoluliznftodyguesshcacrsvutylalqrykehjuofisdookjhrljvedsywrlyccpaowjaqyfaqioesxnlkwgpbznzszyudpwrlgrdgwdyhucztsneqttsuirmjriohhgunzatyfrfzvgvptbgpwajgtysligupoqeoqxoyqtzozufvvlktnvahvsseymtpeyfvxttqosgpplkmxwgmsgtpantazppgnubmpwcdqkvhwfuvcahwibniohiqyywnuzzmxeppokxksrfwrpuzqhjgqryorwboxdauhrkxehiwaputeouwxdfoudcoagcxjcuqvenznxxnprgvhasffxtzaxpcfrcovwgrcwqptoekhmgpoywtxruxokcubekzcrqengviwbtgnzvdzrwwkqvacxwgdhffyvjldgvchoiwnfzoyvkiogisdfyjmfomcazigukqlumyzmnzjzhzfpslwsukykwckvktswjdqxdrlsqvsxwxpqkljeyjpulbswwmuhplfueqnvnhukgjarxlxvwmriqjgmxawmndhsvwnjdjvjtxcsjapfogpesxtpypenunfpjuyoevzztctecilqqbxkaqcyhiobvtqgqruumvvhxolbyzsqcrdchhdqprtkkjsccowrjtyjjmkhleanvfpemuublnnyzfabtxsestncfalqenfcswgerbfcqsapzdtscnzugmwlmidtxkvqhbuaecevwhmwkfqmvpgbefpqpsjmdecmixmmbsjxzwvjdmxydechlraajjmoqpcyoqmrjwoiumuzatydzcnktnkeyztoqvogodxxznhvzduzxudwwqhpftwdspuimioanlzobhjakgajafgzxpqckmhdbbnqmcszpuoqbztnftzgahhxwxbgkilnmzfydyxusnnvngksbjabqjaohdvrniezhmxmkxhemwbbclwdxwgngicplzgajmaryzfkyoqlkrmmfirchzrphveuwmvgaxzbwenvteifxuuefnimnadwxhruvoavlzyhfmeasmgrjawongccgfbgoualiaivbhcgvjjnxpggrewglalthmzvgziobrjeanlvyukwlscexbkibvdjhdgnepdiimmkcxhattwglbkicvsfswocbvphmtpwhcgjbnmxgidtlqcnnwtfujhvgzdussqbwynylzvtjapvqtidpdjkpshvrmqlhindhabubyokzdfrwqvnvgzkyhistydagsgnujiviyijdnabfxqbdqnexvwsvzvcsbrmkbkuzsdehghndyqjodnnblfwmaygdstotfkvxozgwhtbhlkvrzismnozqpfthajafuxekzlgigjpsukjvsdihrjzgovnreqwapdkoqswyclqyvbvpedzyoyedvuuamscbxnqnfmmjyehvidnoimmxmtcinwkbqmcobubjjpshucechrqrffqsyscnqoohcsxenypyqhfklloudgmklcejvgynwouzhtfwuuukdbwpmkjrqxeeaipxrokncholathupdetgaktmvmftqjvzyssocftjwemroghrncynmtchhhcaqxbqpthuaafwgrouaxonzocljeuslzsdwvuoodipdpnlhdihaywzmymxdjrqikughquwtenyucjdgrmipiidiwclhuepgyynoslhzahtdqwliktzsddaahohbszhqxxgripqlwlomjbwtuynydoakejmwkvojuwbfltqjfgxqhwkduzbxpdhtpvrzrfjndmsqfizmqxdxtpbpoemekvxzrrakwjxcxqsdasptruqmjtbaapgmkfnbwnlvzlxwdpzfjryanrmzmpzoefapmnsjdgecrdywsabctaegttffigupnwgakylngrrxurtotxqmzxvsqazajvrwsxyeyjteakeudzjxwbjvagnsjntskmocmpgkybqbnwvrwgoskzqkgffpsyhfmxhymqinrbohxlytsmoeleqrjvievpjipsgdkrqeuglrsjnmvdsihicsgkybcjltcswolpsfxdypmlbjotuxewskisnmczfgreuevnjssjifvlqlhkllifxrxkdbjlhcpegmtrelbosyajljvwwedtxbdccpnmreqaqjrxwulpunagwxesbilalrdniqbzxrbpcvmzpyqklsskpwctgqtrjwhrpisocwderqfiqxsdpkphjsapkvhvsqojyixaechvuoemmyqdlfkuzmlliugckuljfkljoshjhlvvlnywvjswvekfyqhjnsusefdtakejxbejrchoncklguqgnyrcslwztbstmycjziuskegagtlonducdogwbevugppsptdqbajmepmmizaycwcgmjeopbivsyphtvxvvgjbyxpgwpganjiaumojpyhhywosrmnouwpstgbrvhtlqcnmqbygbfnabesvshjmdbhyhirfrkqkmfwdgujhzyjdcbyuijjnkqluaczrnrbbwaeeupnwqzbsazplkyaxqorqsshhlljjlpphhedxdepgfgrqerpuhgmaawhnhqwsgnznrfmxjbdrkwjopylxezxgvetcvrwdewsxdeumhzfrvoilmvksuhyqltuimrnsphqslmgvmmojawwptghonigbdclqtbikiacwpjrbxhmzejozpypfixglatdvuogdoizdtsgsztsfcihtgwyqugeuahpuvvzmgarbsyuutmbxuisdfrvbxzxzhmuektssuktoknkfbmcwwubbnwenybmfqglaceuyqnoadzfenjcjfdlvcpiatuhjdujhaffqsvqvuxchgerokejovrqonxxstibunikiedfyahijobxyhimebctobsjudkqstbcxgixgrhpfiofpwruzvpqyjzvollheoldutddnksutjakhtghpxxnjykxjwgqmsvhnykclexepxqxqzghwfxfdhfmflesfabvanxlrurjtigkjotftqnwyskffpxlragrnfffawqtgyfpmzxfpkdpenxlewyxxgrkmwrmshhzfnorolyfxbvdrspxqnxnuoygkruczddgssygfymdcjgvdxutlrhffhnpyjuxmxefrelxezcgikdliyhvpocvvpkvagvmezrxffujeysplvavtjqjxsgujqsjznxforctwzecxyrkwufpdxadrgzczrnyelfschnagucguuqqqwitviynrypsrdswqxqsegulcwrwsjnihxedfcqychqumiscfkwmqqxunqrfbgqjdwmkyelbldxympctbzfupeocwhkypchuyvhybsbmvymjppfrqmlfrbkpjwpyyytytawuuyjrwxboogfessmltwdcssdqtwomymjskujjtmxiueopwacrwfuqazitvyhvlspvoaeipdsjhgyfjbxhityisidnhlksfznubucqxwaheamndjxmcxwufajmnveuwuoyosqnoqwvtjkwuhkzghvmjhawcfszbhzrbpgsidnbmxxihihnrfbamcyojqpkzodbejtmmipahojoysepzhpljpaugrghgjimtdahnpivdtlcnptnxjyiaafislqavamqgmxtdfoiaakorebqpbbpegawrqymqkewycsdjglkiwaacdqterkixkgraedtqirqmjtvsfhadhafktyrmkzmvidxmisfskvevpcnujqxrqedleuyowkjgphsxzzqlvujkwwgiodbfjesnbsbzcnftuzrvzjjudsgcqmmfpnmyrenuxotbbyvxyovzxgtcyzgqnsvcfhczoptnfnojnlinbfmylhdlijcvcxzjhdixuckaralemvsnbgooorayceuedtomzyjtctvtwgyiesxhynvogxnjdjphcftbefxgasawzagfugmuthjahylkhatlgpnkuksuesrduxkodwjzgubpsmzzmvkskzeglxaqrrvmrgcwcnvkhwzbibaxwnriowoavosminabvfxastkcrkdclgzjvqrjofjjvbyfragofeoazzeqljuypthkmywaffmcjkickqqsuhsviyovhitxeajqahshpejaqtcdkuvgdpclnsguabtgbfwdmrmbvydorfrbcokfdmtsgboidkpgpnmdeyhawkqqshtwxdbarwuxykgduxjlkxppwyruihkcqgynjcpbylayvgdqfpbqmshksyfbhrfxxemhgbkgmkhjtkzyzdqmxxwqvdtevyducpdksntgyaqtkrrkwiyuhukfadjvdnrievszilfinxbyrvknfihmetreydbcstkwoexwsfhfekfvfplmxszcosgovisnbemrjlndqwkvhqsofdbdychmupcsxvhazvrihhnxfyumonbvqeyoghccxfuwacxzxqkezxefxarnnujgyjugrzjoefmghjfhcrnbrtgouaehwnnxwkdplodpuqxdbemfwahptpfppjzowoltyqijfoabgzejerpatwponuefgdtcrgxswiddygeeflpjeelzccnsztxfyqhqyhkuppapvgvdtkmxraytcolbhkiiasaazkvqzvfxbaaxkoudovxrjkusxdazxaawmvoostlvvnsfbpjqkijvudpriqrfsrdfortimgdhtypunakzituezjyhbrpuksbamuiycngvlvpyvczfxvlwhjgicvempfobbwadkiavdswyuxdttoqaaykctprkwfmyeodowglzyjzuhencufcwdobydslazxadnftllhmjslfbrtdlahkgwlebdpdeofidldoymakfnpgekmsltcrrnxvspywfggjrmxryybdltmsfykstmlnzjitaipfoyohkmzimcozxardydxtpjgquoluzbznzqvlewtqyhryjldjoadgjlyfckzbnbootlzxhupieggntjxilcqxnocpyesnhjbauaxcvmkzusmodlyonoldequfunsbwudquaurogsiyhydswsimflrvfwruouskxjfzfynmrymyyqsvkajpnanvyepnzixyteyafnmwnbwmtojdpsucthxtopgpxgnsmnsrdhpskledapiricvdmtwaifrhnebzuttzckroywranbrvgmashxurelyrrbslxnmzyeowchwpjplrdnjlkfcoqdhheavbnhdlltjpahflwscafnnsspikuqszqpcdyfrkaabdigogatgiitadlinfyhgowjuvqlhrniuvrketfmboibttkgakohbmsvhigqztbvrsgxlnjndrqwmcdnntwofojpyrhamivfcdcotodwhvtuyyjlthbaxmrvfzxrhvzkydartfqbalxyjilepmemawjfxhzecyqcdswxxmaaxxyifmouauibstgpcfwgfmjlfhketkeshfcorqirmssfnbuqiqwqfhbmol"
	//s := "qbbnhwpdokcpfpxinlfmkfrfqrtzkhabidqszhxorzfypcjcnopzwigmbznmjnpttflsmjifknezrneedvgzfmnhoavxqksjreddpmibbodtbhzfehgluuukupjmbbvshzxyniaowdjamlfssndojyyephstlonsplrettspwepipwcjmfyvfybxiuqtkdlzqedjxxbvdsfurhedneauccrkyjfiptjfxmpxlssrkyldfriuvjranikluqtjjcoiqffdxaukagphzycvjtvwdhhxzagkevvuccxccuoccdkbboymjtimdrmerspxpktsmrwrlkvpnhqrvpdekmtpdfuxzjwpvqjjhfaupylefbvbsbhdncsshmrhxoyuejenqgjheulkxjnqkwvzznriclrbzryfaeuqkfxrbldyusoeoldpbwadhrgijeplijcvqbormrqglgmzsprtmryvkeevlthvflsvognbxfjilwkdndyzwwxgdbeqwlldyezmkopktzugxgkklimhhjqkmuaifnodtpredhqygmedtqpezboimeuyyujfjxkdmbjpizpqltvgknnlodtbhnbhjkmuhwxvzgmkhbcvvadhnssbvneecglnqxhavhvxpkjxlluilzpysjcnwguyofnhfvhaceztoiscumkhociglkvispihvyoatxcxbeqsmluixgsliatukrecgoldmzfhwkgaqzsckonjuhxdhqztjfxstjvikdrhpyjfxbjjryslfpqoiphrwfjqqhaamrjbrsiovrxmqsyxhqmritjeauwqbwtpqcqhvyyssvfknfhxvtodpzipueixdbntdfcaeatyyainfpkclbgaaqrwwzwbcjwiqzkwzfuxfclmsxpdyvfbnwxjytnaejivivriamhgqsskqhnqeurttrfrmstrbeokzhuzvbfmwywohmgogyhzpmsdemugqkspsmoppwbnwabdmiruibwznqcuczculujfiavzwynsyqxmarjkshjhxobandwyzggjibjgzyaaqxorqxbkenscbveqbaociwmqxxyzvyblypeongzrttvwqzmrccwkzidyfbxcaypyquodcpwxkstbthuvjqgialhfmgjohzoxvdaxuywfqrgmyahhtpqtazbphmfoluliznftodyguesshcacrsvutylalqrykehjuofisdookjhrljvedsywrlyccpaowjaqyfaqioesxnlkwgpbznzszyudpwrlgrdgwdyhucztsneqttsuirmjriohhgunzatyfrfzvgvptbgpwajgtysligupoqeoqxoyqtzozufvvlktnvahvsseymtpeyfvxttqosgpplkmxwgmsgtpantazppgnubmpwcdqkvhwfuvcahwibniohiqyywnuzzmxeppokxksrfwrpuzqhjgqryorwboxdauhrkxehiwaputeouwxdfoudcoagcxjcuqvenznxxnprgvhasffxtzaxpcfrcovwgrcwqptoekhmgpoywtxruxokcubekzcrqengviwbtgnzvdzrwwkqvacxwgdhffyvjldgvchoiwnfzoyvkiogisdfyjmfomcazigukqlumyzmnzjzhzfpslwsukykwckvktswjdqxdrlsqvsxwxpqkljeyjpulbswwmuhplfueqnvnhukgjarxlxvwmriqjgmxawmndhsvwnjdjvjtxcsjapfogpesxtpypenunfpjuyoevzztctecilqqbxkaqcyhiobvtqgqruumvvhxolbyzsqcrdchhdqprtkkjsccowrjtyjjmkhleanvfpemuublnnyzfabtxsestncfalqenfcswgerbfcqsapzdtscnzugmwlmidtxkvqhbuaecevwhmwkfqmvpgbefpqpsjmdecmixmmbsjxzwvjdmxydechlraajjmoqpcyoqmrjwoiumuzatydzcnktnkeyztoqvogodxxznhvzduzxudwwqhpftwdspuimioanlzobhjakgajafgzxpqckmhdbbnqmcszpuoqbztnftzgahhxwxbgkilnmzfydyxusnnvngksbjabqjaohdvrniezhmxmkxhemwbbclwdxwgngicplzgajmaryzfkyoqlkrmmfirchzrphveuwmvgaxzbwenvteifxuuefnimnadwxhruvoavlzyhfmeasmgrjawongccgfbgoualiaivbhcgvjjnxpggrewglalthmzvgziobrjeanlvyukwlscexbkibvdjhdgnepdiimmkcxhattwglbkicvsfswocbvphmtpwhcgjbnmxgidtlqcnnwtfujhvgzdussqbwynylzvtjapvqtidpdjkpshvrmqlhindhabubyokzdfrwqvnvgzkyhistydagsgnujiviyijdnabfxqbdqnexvwsvzvcsbrmkbkuzsdehghndyqjodnnblfwmaygdstotfkvxozgwhtbhlkvrzismnozqpfthajafuxekzlgigjpsukjvsdihrjzgovnreqwapdkoqswyclqyvbvpedzyoyedvuuamscbxnqnfmmjyehvidnoimmxmtcinwkbqmcobubjjpshucechrqrffqsyscnqoohcsxenypyqhfklloudgmklcejvgynwouzhtfwuuukdbwpmkjrqxeeaipxrokncholathupdetgaktmvmftqjvzyssocftjwemroghrncynmtchhhcaqxbqpthuaafwgrouaxonzocljeuslzsdwvuoodipdpnlhdihaywzmymxdjrqikughquwtenyucjdgrmipiidiwclhuepgyynoslhzahtdqwliktzsddaahohbszhqxxgripqlwlomjbwtuynydoakejmwkvojuwbfltqjfgxqhwkduzbxpdhtpvrzrfjndmsqfizmqxdxtpbpoemekvxzrrakwjxcxqsdasptruqmjtbaapgmkfnbwnlvzlxwdpzfjryanrmzmpzoefapmnsjdgecrdywsabctaegttffigupnwgakylngrrxurtotxqmzxvsqazajvrwsxyeyjteakeudzjxwbjvagnsjntskmocmpgkybqbnwvrwgoskzqkgffpsyhfmxhymqinrbohxlytsmoeleqrjvievpjipsgdkrqeuglrsjnmvdsihicsgkybcjltcswolpsfxdypmlbjotuxewskisnmczfgreuevnjssjifvlqlhkllifxrxkdbjlhcpegmtrelbosyajljvwwedtxbdccpnmreqaqjrxwulpunagwxesbilalrdniqbzxrbpcvmzpyqklsskpwctgqtrjwhrpisocwderqfiqxsdpkphjsapkvhvsqojyixaechvuoemmyqdlfkuzmlliugckuljfkljoshjhlvvlnywvjswvekfyqhjnsusefdtakejxbejrchoncklguqgnyrcslwztbstmycjziuskegagtlonducdogwbevugppsptdqbajmepmmizaycwcgmjeopbivsyphtvxvvgjbyxpgwpganjiaumojpyhhywosrmnouwpstgbrvhtlqcnmqbygbfnabesvshjmdbhyhirfrkqkmfwdgujhzyjdcbyuijjnkqluaczrnrbbwaeeupnwqzbsazplkyaxqorqsshhlljjlpphhedxdepgfgrqerpuhgmaawhnhqwsgnznrfmxjbdrkwjopylxezxgvetcvrwdewsxdeumhzfrvoilmvksuhyqltuimrnsphqslmgvmmojawwptghonigbdclqtbikiacwpjrbxhmzejozpypfixglatdvuogdoizdtsgsztsfcihtgwyqugeuahpuvvzmgarbsyuutmbxuisdfrvbxzxzhmuektssuktoknkfbmcwwubbnwenybmfqglaceuyqnoadzfenjcjfdlvcpiatuhjdujhaffqsvqvuxchgerokejovrqonxxstibunikiedfyahijobxyhimebctobsjudkqstbcxgixgrhpfiofpwruzvpqyjzvollheoldutddnksutjakhtghpxxnjykxjwgqmsvhnykclexepxqxqzghwfxfdhfmflesfabvanxlrurjtigkjotftqnwyskffpxlragrnfffawqtgyfpmzxfpkdpenxlewyxxgrkmwrmshhzfnorolyfxbvdrspxqnxnuoygkruczddgssygfymdcjgvdxutlrhffhnpyjuxmxefrelxezcgikdliyhvpocvvpkvagvmezrxffujeysplvavtjqjxsgujqsjznxforctwzecxyrkwufpdxadrgzczrnyelfschnagucguuqqqwitviynrypsrdswqxqsegulcwrwsjnihxedfcqychqumiscfkwmqqxunqrfbgqjdwmkyelbldxympctbzfupeocwhkypchuyvhybsbmvymjppfrqmlfrbkpjwpyyytytawuuyjrwxboogfessmltwdcssdqtwomymjskujjtmxiueopwacrwfuqazitvyhvlspvoaeipdsjhgyfjbxhityisidnhlksfznubucqxwaheamndjxmcxwufajmnveuwuoyosqnoqwvtjkwuhkzghvmjhawcfszbhzrbpgsidnbmxxihihnrfbamcyojqpkzodbejtmmipahojoysepzhpljpaugrghgjimtdahnpivdtlcnptnxjyiaafislqavamqgmxtdfoiaakorebqpbbpegawrqymqkewycsdjglkiwaacdqterkixkgraedtqirqmjtvsfhadhafktyrmkzmvidxmisfskvevpcnujqxrqedleuyowkjgphsxzzqlvujkwwgiodbfjesnbsbzcnftuzrvzjjudsgcqmmfpnmyrenuxotbbyvxyovzxgtcyzgqnsvcfhczoptnfnojnlinbfmylhdlijcvcxzjhdixuckaralemvsnbgooorayceuedtomzyjtctvtwgyiesxhynvogxnjdjphcftbefxgasawzagfugmuthjahylkhatlgpnkuksuesrduxkodwjzgubpsmzzmvkskzeglxaqrrvmrgcwcnvkhwzbibaxwnriowoavosminabvfxastkcrkdclgzjvqrjofjjvbyfragofeoazzeqljuypthkmywaffmcjkickqqsuhsviyovhitxeajqahshpejaqtcdkuvgdpclnsguabtgbfwdmrmbvydorfrbcokfdmtsgboidkpgpnmdeyhawkqqshtwxdbarwuxykgduxjlkxppwyruihkcqgynjcpbylayvgdqfpbqmshksyfbhrfxxemhgbkgmkhjtkzyzdqmxxwqvdtevyducpdksntgyaqtkrrkwiyuhukfadjvdnrievszilfinxbyrvknfihmetreydbcstkwoexwsfhfekfvfplmxszcosgovisnbemrjlndqwkvhqsofdbdychmupcsxvhazvrihhnxfyumonbvqeyoghccxfuwacxzxqkezxefxarnnujgyjugrzjoefmghjfhcrnbrtgouaehwnnxwkdplodpuqxdbemfwahptpfppjzowoltyqijfoabgzejerpatwponuefgdtcrgxswiddygeeflpjeelzccnsztxfyqhqyhkuppapvgvdtkmxraytcolbhkiiasaazkvqzvfxbaaxkoudovxrjkusxdazxaawmvoostlvvnsfbpjqkijvudpriqrfsrdfortimgdhtypunakzituezjyhbrpuksbamuiycngvlvpyvczfxvlwhjgicvempfobbwadkiavdswyuxdttoqaaykctprkwfmyeodowglzyjzuhencufcwdobydslazxadnftllhmjslfbrtdlahkgwlebdpdeofidldoymakfnpgekmsltcrrnxvspywfggjrmxryybdltmsfykstmlnzjitaipfoyohkmzimcozxardydxtpjgquoluzbznzqvlewtqyhryjldjoadgjlyfckzbnbootlzxhupieggntjxilcqxnocpyesnhjbauaxcvmkzusmodlyonoldequfunsbwudquaurogsiyhydswsimflrvfwruouskxjfzfynmrymyyqsvkajpnanvyepnzixyteyafnmwnbwmtojdpsucthxtopgpxgnsmnsrdhpskledapiricvdmtwaifrhnebzuttzckroywranbrvgmashxurelyrrbslxnmzyeowchwpjplrdnjlkfcoqdhheavbnhdlltjpahflwscafnnsspikuqszqpcdyfrkaabdigogatgiitadlinfyhgowjuvqlhrniuvrketfmboibttkgakohbmsvhigqztbvrsgxlnjndrqwmcdnntwofojpyrhamivfcdcotodwhvtuyyjlthbaxmrvfzxrhvzkydartfqbalxyjilepmemawjfxhzecyqcdswxxmaaxxyifmouauibstgpcfwgfmjlfhketkeshfcorqirmssfnbuqiqwqfhbmol"
	words := []string{
		"toiscumkhociglkvispihvyoatxcx",
		"ndojyyephstlonsplrettspwepipw",
		"yzfkyoqlkrmmfirchzrphveuwmvga",
		"mxxihihnrfbamcyojqpkzodbejtmm",
		"fenjcjfdlvcpiatuhjdujhaffqsvq",
		"ehghndyqjodnnblfwmaygdstotfkv",
		"heoldutddnksutjakhtghpxxnjykx",
		"cvrwdewsxdeumhzfrvoilmvksuhyq",
		"ftqjvzyssocftjwemroghrncynmtc",
		"idiwclhuepgyynoslhzahtdqwlikt",
		"eurttrfrmstrbeokzhuzvbfmwywoh",
		"jxlluilzpysjcnwguyofnhfvhacez",
		"uskegagtlonducdogwbevugppsptd",
		"xmcxwufajmnveuwuoyosqnoqwvtjk",
		"wolpsfxdypmlbjotuxewskisnmczf",
		"fjryanrmzmpzoefapmnsjdgecrdyw",
		"jgmxawmndhsvwnjdjvjtxcsjapfog",
		"wuhkzghvmjhawcfszbhzrbpgsidnb",
		"yelbldxympctbzfupeocwhkypchuy",
		"vzduzxudwwqhpftwdspuimioanlzo",
		"bdpdeofidldoymakfnpgekmsltcrr",
		"fmyeodowglzyjzuhencufcwdobyds",
		"dhtypunakzituezjyhbrpuksbamui",
		"bdmiruibwznqcuczculujfiavzwyn",
		"eudzjxwbjvagnsjntskmocmpgkybq",
		"tuynydoakejmwkvojuwbfltqjfgxq",
		"psrdswqxqsegulcwrwsjnihxedfcq",
		"cokfdmtsgboidkpgpnmdeyhawkqqs",
		"fujhvgzdussqbwynylzvtjapvqtid",
		"rqeuglrsjnmvdsihicsgkybcjltcs",
		"vhybsbmvymjppfrqmlfrbkpjwpyyy",
		"aukagphzycvjtvwdhhxzagkevvucc",
		"hwkduzbxpdhtpvrzrfjndmsqfizmq",
		"ywnuzzmxeppokxksrfwrpuzqhjgqr",
		"qbajmepmmizaycwcgmjeopbivsyph",
		"uamscbxnqnfmmjyehvidnoimmxmtc",
		"nxvspywfggjrmxryybdltmsfykstm",
		"amrjbrsiovrxmqsyxhqmritjeauwq",
		"yorwboxdauhrkxehiwaputeouwxdf",
		"qkewycsdjglkiwaacdqterkixkgra",
		"ycngvlvpyvczfxvlwhjgicvempfob",
		"jgphsxzzqlvujkwwgiodbfjesnbsb",
		"mkxhemwbbclwdxwgngicplzgajmar",
		"mryvkeevlthvflsvognbxfjilwkdn",
		"mezrxffujeysplvavtjqjxsgujqsj",
		"rtotxqmzxvsqazajvrwsxyeyjteak",
		"sabctaegttffigupnwgakylngrrxu",
		"xccuoccdkbboymjtimdrmerspxpkt",
		"xusnnvngksbjabqjaohdvrniezhmx",
		"oyuejenqgjheulkxjnqkwvzznricl",
		"mxszcosgovisnbemrjlndqwkvhqso",
		"wsgnznrfmxjbdrkwjopylxezxgvet",
		"dxmisfskvevpcnujqxrqedleuyowk",
		"dhrgijeplijcvqbormrqglgmzsprt",
		"vuxchgerokejovrqonxxstibuniki",
		"lumyzmnzjzhzfpslwsukykwckvkts",
		"inwkbqmcobubjjpshucechrqrffqs",
		"ywtxruxokcubekzcrqengviwbtgnz",
		"ccpnmreqaqjrxwulpunagwxesbila",
		"pesxtpypenunfpjuyoevzztctecil",
		"sygfymdcjgvdxutlrhffhnpyjuxmx",
		"uisdfrvbxzxzhmuektssuktoknkfb",
		"cejvgynwouzhtfwuuukdbwpmkjrqx",
		"oudcoagcxjcuqvenznxxnprgvhasf",
		"sxnlkwgpbznzszyudpwrlgrdgwdyh",
		"qqbxkaqcyhiobvtqgqruumvvhxolb",
		"mkhleanvfpemuublnnyzfabtxsest",
		"bibaxwnriowoavosminabvfxastkc",
		"bcxgixgrhpfiofpwruzvpqyjzvoll",
		"lzccnsztxfyqhqyhkuppapvgvdtkm",
		"pdjkpshvrmqlhindhabubyokzdfrw",
		"qbbnhwpdokcpfpxinlfmkfrfqrtzk",
		"rnyelfschnagucguuqqqwitviynry",
		"qtrjwhrpisocwderqfiqxsdpkphjs",
		"vxttqosgpplkmxwgmsgtpantazppg",
		"tyisidnhlksfznubucqxwaheamndj",
		"kgaqzsckonjuhxdhqztjfxstjvikd",
		"jeuslzsdwvuoodipdpnlhdihaywzm",
		"vdzrwwkqvacxwgdhffyvjldgvchoi",
		"cftbefxgasawzagfugmuthjahylkh",
		"xraytcolbhkiiasaazkvqzvfxbaax",
		"oyqtzozufvvlktnvahvsseymtpeyf",
		"rnnujgyjugrzjoefmghjfhcrnbrtg",
		"rfzvgvptbgpwajgtysligupoqeoqx",
		"igbdclqtbikiacwpjrbxhmzejozpy",
		"dyzwwxgdbeqwlldyezmkopktzugxg",
		"hmetreydbcstkwoexwsfhfekfvfpl",
		"zcnftuzrvzjjudsgcqmmfpnmyrenu",
		"zzmvkskzeglxaqrrvmrgcwcnvkhwz",
		"vjswvekfyqhjnsusefdtakejxbejr",
		"rwwzwbcjwiqzkwzfuxfclmsxpdyvf",
		"fdbdychmupcsxvhazvrihhnxfyumo",
		"vdtevyducpdksntgyaqtkrrkwiyuh",
		"nbvqeyoghccxfuwacxzxqkezxefxa",
		"vpgbefpqpsjmdecmixmmbsjxzwvjd",
		"jwgqmsvhnykclexepxqxqzghwfxfd",
		"olyfxbvdrspxqnxnuoygkruczddgs",
		"qgmxtdfoiaakorebqpbbpegawrqym",
		"liaivbhcgvjjnxpggrewglalthmzv",
		"choncklguqgnyrcslwztbstmycjzi",
		"fpkdpenxlewyxxgrkmwrmshhzfnor",
		"hhhcaqxbqpthuaafwgrouaxonzocl",
		"ipahojoysepzhpljpaugrghgjimtd",
		"wosrmnouwpstgbrvhtlqcnmqbygbf",
		"nwyskffpxlragrnfffawqtgyfpmzx",
		"bcvvadhnssbvneecglnqxhavhvxpk",
		"hoavxqksjreddpmibbodtbhzfehgl",
		"lazxadnftllhmjslfbrtdlahkgwle",
		"uuukupjmbbvshzxyniaowdjamlfss",
		"tpqtazbphmfoluliznftodyguessh",
		"ychqumiscfkwmqqxunqrfbgqjdwmk",
		"rkdclgzjvqrjofjjvbyfragofeoaz",
		"pphhedxdepgfgrqerpuhgmaawhnhq",
		"cacrsvutylalqrykehjuofisdookj",
		"kyldfriuvjranikluqtjjcoiqffdx",
		"bnwvrwgoskzqkgffpsyhfmxhymqin",
		"uzmlliugckuljfkljoshjhlvvlnyw",
		"abfxqbdqnexvwsvzvcsbrmkbkuzsd",
		"xotbbyvxyovzxgtcyzgqnsvcfhczo",
		"bwtpqcqhvyyssvfknfhxvtodpzipu",
		"nsfbpjqkijvudpriqrfsrdfortimg",
		"tgwyqugeuahpuvvzmgarbsyuutmbx",
		"upnwqzbsazplkyaxqorqsshhlljjl",
		"edfyahijobxyhimebctobsjudkqst",
		"ialhfmgjohzoxvdaxuywfqrgmyahh",
		"jlhcpegmtrelbosyajljvwwedtxbd",
		"tpfppjzowoltyqijfoabgzejerpat",
		"mgogyhzpmsdemugqkspsmoppwbnwa",
		"nubmpwcdqkvhwfuvcahwibniohiqy",
		"ukfadjvdnrievszilfinxbyrvknfi",
		"dgnepdiimmkcxhattwglbkicvsfsw",
		"syqxmarjkshjhxobandwyzggjibjg",
		"bnwxjytnaejivivriamhgqsskqhnq",
		"hzyjdcbyuijjnkqluaczrnrbbwaee",
		"yscnqoohcsxenypyqhfklloudgmkl",
		"habidqszhxorzfypcjcnopzwigmbz",
		"wjdqxdrlsqvsxwxpqkljeyjpulbsw",
		"tytawuuyjrwxboogfessmltwdcssd",
		"pfixglatdvuogdoizdtsgsztsfcih",
		"apkvhvsqojyixaechvuoemmyqdlfk",
		"ouaehwnnxwkdplodpuqxdbemfwahp",
		"ixuckaralemvsnbgooorayceuedto",
		"ymxdjrqikughquwtenyucjdgrmipi",
		"smrwrlkvpnhqrvpdekmtpdfuxzjwp",
		"bhjakgajafgzxpqckmhdbbnqmcszp",
		"beqsmluixgsliatukrecgoldmzfhw",
		"greuevnjssjifvlqlhkllifxrxkdb",
		"yzsqcrdchhdqprtkkjsccowrjtyjj",
		"sviyovhitxeajqahshpejaqtcdkuv",
		"qtwomymjskujjtmxiueopwacrwfuq",
		"mzyjtctvtwgyiesxhynvogxnjdjph",
		"dyfbxcaypyquodcpwxkstbthuvjqg",
		"hfmflesfabvanxlrurjtigkjotftq",
		"mxydechlraajjmoqpcyoqmrjwoium",
		"nabesvshjmdbhyhirfrkqkmfwdguj",
		"bhrfxxemhgbkgmkhjtkzyzdqmxxwq",
		"gziobrjeanlvyukwlscexbkibvdjh",
		"mcwwubbnwenybmfqglaceuyqnoadz",
		"xyzvyblypeongzrttvwqzmrccwkzi",
		"ncfalqenfcswgerbfcqsapzdtscnz",
		"dtqpezboimeuyyujfjxkdmbjpizpq",
		"wmuhplfueqnvnhukgjarxlxvwmriq",
		"qwapdkoqswyclqyvbvpedzyoyedvu",
		"uoqbztnftzgahhxwxbgkilnmzfydy",
		"zsddaahohbszhqxxgripqlwlomjbw",
		"bwadkiavdswyuxdttoqaaykctprkw",
		"eixdbntdfcaeatyyainfpkclbgaaq",
		"nmjnpttflsmjifknezrneedvgzfmn",
		"avlzyhfmeasmgrjawongccgfbgoua",
		"kklimhhjqkmuaifnodtpredhqygme",
		"xzbwenvteifxuuefnimnadwxhruvo",
		"ugmwlmidtxkvqhbuaecevwhmwkfqm",
		"rhpyjfxbjjryslfpqoiphrwfjqqha",
		"eeaipxrokncholathupdetgaktmvm",
		"ltuimrnsphqslmgvmmojawwptghon",
		"azitvyhvlspvoaeipdsjhgyfjbxhi",
		"efrelxezcgikdliyhvpocvvpkvagv",
		"znxforctwzecxyrkwufpdxadrgzcz",
		"kcqgynjcpbylayvgdqfpbqmshksyf",
		"hrljvedsywrlyccpaowjaqyfaqioe",
		"cjmfyvfybxiuqtkdlzqedjxxbvdsf",
		"zeqljuypthkmywaffmcjkickqqsuh",
		"wnfzoyvkiogisdfyjmfomcazigukq",
		"zyaaqxorqxbkenscbveqbaociwmqx",
		"ahnpivdtlcnptnxjyiaafislqavam",
		"edtqirqmjtvsfhadhafktyrmkzmvi",
		"wponuefgdtcrgxswiddygeeflpjee",
		"xozgwhtbhlkvrzismnozqpfthajaf",
		"ptnfnojnlinbfmylhdlijcvcxzjhd",
		"uxekzlgigjpsukjvsdihrjzgovnre",
		"rbohxlytsmoeleqrjvievpjipsgdk",
		"fxtzaxpcfrcovwgrcwqptoekhmgpo",
		"tvxvvgjbyxpgwpganjiaumojpyhhy",
		"vqjjhfaupylefbvbsbhdncsshmrhx",
		"urhedneauccrkyjfiptjfxmpxlssr",
		"ltvgknnlodtbhnbhjkmuhwxvzgmkh",
		"ucztsneqttsuirmjriohhgunzatyf",
		"rbzryfaeuqkfxrbldyusoeoldpbwa",
		"atlgpnkuksuesrduxkodwjzgubpsm",
		"lrdniqbzxrbpcvmzpyqklsskpwctg",
		"qvnvgzkyhistydagsgnujiviyijdn",
		"uzatydzcnktnkeyztoqvogodxxznh",
		"ocbvphmtpwhcgjbnmxgidtlqcnnwt",
		"koudovxrjkusxdazxaawmvoostlvv",
		"ptruqmjtbaapgmkfnbwnlvzlxwdpz",
		"xdxtpbpoemekvxzrrakwjxcxqsdas",
		"gdpclnsguabtgbfwdmrmbvydorfrb",
		"htwxdbarwuxykgduxjlkxppwyruih",
	}
	//fmt.Println(findSubstring1("ab", []string{"a", "b"}))//[]
	fmt.Println(findSubstring1(s, words))//[]
}
func TestNum29(t *testing.T)  {
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(-10, 3))
}
func TestNum28(t *testing.T) {
	fmt.Println(strStr("mississippi", "issip"))//issip
	fmt.Println("abc"[1:2])
}

func TestNum27(t *testing.T) {
	fmt.Println(removeElement([]int{3,2,2,3}, 3))
}

func TestNum26(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func TestNum16(t *testing.T)  {
	fmt.Println(threeSumClosest([]int{1,1,-1,-1,3}, -1))
}

func TestNum15(t *testing.T)  {
	//输入：nums = [-1,0,1,-1,2,-1, -1,-4]
	//输出：[[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0,1, 2,-1, -1,-4}))
	fmt.Println(threeSum([]int{0,0,0,0}))
	fmt.Println(Sort([]int{0,1, 2,-1, -1,-4}))
}
func TestNum14(t *testing.T)  {
	fmt.Println(longestCommonPrefix([]string{"flabbbb", "flakdfjkd"}))
}
func TestNum318(t *testing.T)  {
	fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"}) )// 16
}
func TestNum22(t *testing.T)  {
	//fmt.Println(AddString([]string{"1", "11"}, "(", []string{"2", "22"}))
	fmt.Println(generateParenthesis(3))
	fmt.Println(countParenthesisNum(3))
	fmt.Println(generateParenthesis2(3))
}
func TestNum12(t *testing.T) {
	//1994 "MCMXCIV"
	fmt.Println(intToRoman(1994))
	fmt.Println(romanToInt("MCMXCIV"))
}
func TestNum10(t *testing.T)  {
	s:="mississippi"
	p:="mis*is*p*."
	isMatch(s, p)

}
func TestQuick(t *testing.T) {
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//Quick([]int{-2,1,3,4,7,8,100,22}, 2, 2, &wg)
	tt:= []struct {
		Nums []int
	}{{[]int{1, 4, 8, 22, 7, 3 , 100, -2}}}
	for _, x:= range tt {
		Quick(x.Nums, 0, len(x.Nums) - 1, nil)
		fmt.Println(x.Nums)
	}
}

func TestNum58(t *testing.T)  {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon"))
}
var i = 0
var j = 2
func TestNum8(t *testing.T)  {
	arr := []int{1, 2, 3}
	fmt.Println(arr[3:])
}

func TestNum5(t *testing.T)  {
	tests := []struct{s , exp string} {
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"aacabdkacaa", "aca"},
		{"a", "a"},
		{"ac", "a"},
	}
	for _, row := range tests {
		if res:= longestPalindrome(row.s); res != row.exp {
			t.Errorf("%v != %v", res, row.exp)
		} else {
			t.Logf("right!!!")
		}
	}
}
func TestNum4(t *testing.T)  {
	tests:= []struct{
		nums1, nums2 []int
		exp float64
	} {
		{[]int{1, 4},  []int{2, 3, 5, 6}, 3.5},
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
	}
	for _, row:= range tests {
		if res:= findMedianSortedArrays2(row.nums1, row.nums2); res != row.exp {
			t.Errorf("%v != %v", res, row.exp)
		} else {
			t.Logf("right!!!")
		}
	}
}
func TestNum3(t *testing.T) {
	tests := []struct{
		s string
		exp int
	} {
		//{"abcabcbb", 3},
		//{"bbbbb", 1},
		//{"pwwkew", 3},
		//{"", 0},
		{"abba", 2},
	}
	for _, row := range tests {
		res := lengthOfLongestSubstring(row.s)
		if res != row.exp {
			t.Errorf("%v != %v", res, row.exp)
		} else {
			t.Logf("right")
		}

	}
}

func TestNum2(t *testing.T)  {
	tests := []struct {
		a, b, exp []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7,0,8}},
		{[]int{9,9,9,9,9,9,9},[]int{9,9,9,9},[]int{8,9,9,9,0,0,0,1}},
	}
	for _, row:= range tests {
		res := addTwoNumbers(initListNode(row.a), initListNode(row.b))
		exp := initListNode(row.exp)
		if ok, _:= exp.Equal(res); !ok {
			t.Errorf("%v != %v", res, exp)
		} else {
			t.Logf("%v == %v, right !!!", res, exp)
		}

	}

}
