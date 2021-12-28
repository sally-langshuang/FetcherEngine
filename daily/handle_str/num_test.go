package handle_str

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum472(t *testing.T) {
	datas := []struct {
		words []string
		ans   []string
	}{
		//{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},
		//{[]string{"cat", "dog", "catdog"}, []string{"catdog"}},
		{[]string{"ab", "bc", "abc"}, []string{}},
		{[]string{"rfkqyuqfjkx", "", "vnrtysfrzrmzl", "gfve", "qfpd", "lqdqrrcrwdnxeuo", "q", "klaitgdphcspij", "hbsfyfv", "adzpbfudkklrw", "aozmixr", "ife", "feclhbvfuk", "yeqfqojwtw", "sileeztxwjl", "ngbqqmbxqcqp", "khhqr", "dwfcayssyoqc", "omwufbdfxu", "zhift", "kczvhsybloet", "crfhpxprbsshsjxd", "ilebxwbcto", "yaxzfbjbkrxi", "imqpzwmshlpj", "ta", "hbuxhwadlpto", "eziwkmg", "ovqzgdixrpddzp", "c", "wnqwqecyjyib", "jy", "mjfqwltvzk", "tpvo", "phckcyufdqml", "lim", "lfz", "tgygdt", "nhcvpf", "fbrpzlk", "shwywshtdgmb", "bkkxcvg", "monmwvytby", "nuqhmfj", "qtg", "cwkuzyamnerp", "fmwevhwlezo", "ye", "hbrcewjxvcezi", "tiq", "tfsrptug", "iznorvonzjfea", "gama", "apwlmbzit", "s", "hzkosvn", "nberblt", "kggdgpljfisylt", "mf", "h", "bljvkypcflsaqe", "cijcyrgmqirz", "iaxakholawoydvch", "e", "gttxwpuk", "jf", "xbrtspfttota", "sngqvoijxuv", "bztvaal", "zxbshnrvbykjql", "zz", "mlvyoshiktodnsjj", "qplci", "lzqrxl", "qxru", "ygjtyzleizme", "inx", "lwhhjwsl", "endjvxjyghrveu", "phknqtsdtwxcktmw", "wsdthzmlmbhjkm", "u", "pbqurqfxgqlojmws", "mowsjvpvhznbsi", "hdkbdxqg", "ge", "pzchrgef", "ukmcowoe", "nwhpiid", "xdnnl", "n", "yjyssbsoc", "cdzcuunkrf", "uvouaghhcyvmlk", "aajpfpyljt", "jpyntsefxi", "wjute", "y", "pbcnmhf", "qmmidmvkn", "xmywegmtuno", "vuzygv", "uxtrdsdfzfssmel", "odjgdgzfmrazvnd", "a", "rdkugsbdpawxi", "ivd", "bbqeonycaegxfj", "lrfkraoheucsvpi", "eqrswgkaaaohxx", "hqjtkqaqh", "berbpmglbjipnuj", "wogwczlkyrde", "aqufowbig", "snjniegvdvotu", "ocedkt", "bbufnxorixibbd", "rzuqsyr", "qghoy", "evcuanuujszitaoa", "wsx", "glafbwzdd", "znrvjqeyqi", "npitruijvyllsi", "objltu", "ryp", "nvybsfrxtlfmp", "id", "zoolzslgd", "owijatklvjzscizr", "upmsoxftumyxifyu", "xucubv", "fctkqlroq", "zjv", "wzi", "ppvs", "mflvioemycnphfjt", "nwedtubynsb", "repgcx", "gsfomhvpmy", "kdohe", "tyycsibbeaxn", "wjkfvabn", "llkmagl", "thkglauzgkeuly", "paeurdvexqlw", "akdt", "ihmfrj", "janxk", "rqdll", "cyhbsuxnlftmjc", "yybwsjmajbwtuhkk", "ovytgaufpjl", "iwbnzhybsx", "mumbh", "jqmdabmyu", "br", "lwstjkoxbczkj", "vhsgzvwiixxaob", "fso", "qnebmfl", "ooetjiz", "lq", "msxphqdgz", "mqhoggvrvjqrp", "xbhkkfg", "zxjegsyovdrmw", "jav", "mshoj", "ax", "biztkfomz", "hujdmcyxdqteqja", "gqgsomonv", "reqqzzpw", "lihdnvud", "lznfhbaokxvce", "fhxbldylqqewdnj", "rlbskqgfvn", "lfvobeyolyy", "v", "iwh", "fpbuiujlolnjl", "gvwxljbo", "ypaotdzjxxrsc", "mwrvel", "umzpnoiei", "ogwilaswn", "yw", "egdgye", "hsrznlzrf", "mwdgxaigmxpy", "yaqgault", "dtlg", "cyvfiykmkllf", "zxqyhvizqmamj", "lvvgoifltzywueyp", "abinmy", "ppzaecvmx", "qsmzc", "iddymnl", "uskihek", "evxtehxtbthq", "jvtfzddlgch", "czohpyewf", "ufzazyxtqxcu", "brxpfymuvfvs", "xrrcfuusicc", "aqhlswbzievij", "rv", "udvmara", "upityz", "fecd", "suxteeitxtg", "dfuydrtbfypbn", "cypqodxr", "wikfuxwjht", "jrliuaifpp", "vkmxys", "wvpfyfpkvgthq", "rmajxis", "jncxgviyu", "av", "nmhskodmidaj", "lkfrimprrhen", "uip", "hstyopbvuiqc", "p", "vwduwmjpblqo", "fnxwgqtvwztje", "xwnbcuggl", "iehimvoymyjasin", "spsqiu", "flhyfac", "mqrbq", "pstsxhplrrmbeddv", "hnegtuxx", "alsyxezjwtlwmxv", "jtxytykkcku", "bhhlovgcx", "xhhivxnutkx", "had", "aysulvk", "m", "anhsyxli", "jdkgfc", "potn", "lcibpxkidmwexp", "gwoxjicdkv", "tltienw", "ngiutnuqbzi", "o", "tzlyb", "vumnwehj", "os", "np", "lhv", "uzvgyeette", "ipfvr", "lpprjjalchhhcmh", "k", "pciulccqssaqgd", "tp", "dmzdzveslyjad", "wtsbhgkd", "eouxbldsxzm", "vhtonlampljgzyve", "xhnlcrldtfthul", "xhflc", "upgei", "rlaks", "yfqvnvtnqspyjbxr", "phouoyhvls", "voibuvbhhjcdflvl", "rgorfbjrofokggaf", "dqhqats", "zchpicyuawpovm", "yzwfor", "koat", "pybf", "fhdzsbiyjld", "gznfnqydisn", "xz", "po", "tcjup", "wygsnxk", "kqlima", "fgxnuohrnhg", "publurhztntgmimc", "zuufzphd", "iucrmmmjqtcey", "wnnbq", "rghzyz", "ukjqsjbmp", "mdtrgv", "vyeikgjdnml", "kxwldnmi", "apzuhsbssaxj", "tkbkoljyodlipof", "nkq", "ktwtj", "vgmkgjwle", "t", "agylw", "vomtuy", "jbtvitkqn", "vtdxwrclpspcn", "rdrls", "yxfeoh", "upj", "myctacn", "fdnor", "ahqghzhoqprgkym", "phiuvdv", "jp", "fdgpouzjwbq", "hqoyefmugjvewhxu", "qfzwuwe", "fnsbijkeepyxry", "oja", "qthkcij", "zpmqfbmnr", "ybaibmzonzqlnmd", "svo", "gjftyfehik", "jfrfgznuaytvaegm", "aljhrx", "odjq", "ogwaxrssjxgvnka", "zaqswwofedxj", "lugpktauixp", "dc", "odknlbvxrs", "jeobu", "vqythyvzxbcgrlbg", "hwc", "erpbaxq", "ujxcxck", "rrklkb", "wlrwyuy", "zmg", "yyhga", "xwdbycdu", "htedgvsrhchox", "wr", "suhesetv", "jonqwhkwezjvjgg", "sqqyrxtjkcalswq", "hvyimhe", "pjzdkmoue", "zbphmgoxq", "lbdlcumdgixjbcq", "ztzdjqmadthtdmv", "qcagsyqggcf", "if", "jpjxcjyi", "chyicqibxdgkqtg", "iwpdklhum", "wljmg", "micmun", "npdbamofynykqv", "ijsnfkpfy", "lmq", "oyjmeqvhcrvgm", "mqopusqktdthpvz", "fz", "r", "qbsqtipq", "nxtsnason", "xbpipyhh", "topsuqomfjrd", "islif", "gbndakaq", "bwnkxnwpzeoohlx", "hrtbfnq", "fguvomeepxoffg", "mat", "dzfpfnwbfuj", "onlvy", "cwcchvsasdylb", "rxfcztzqopdi", "ybrhodjn", "oqkijy", "ncvrjo", "dphbfaal", "xgtpdtkz", "sebevsopjvciwljf", "rcumyacqdapwczen", "mabkapuoud", "pbozezeygljfftvy", "bvazmzbndl", "vl", "qiaixdtbhqvlzd", "ffjfb", "svthrfmkoxbho", "cvet", "ucgqyvopafyttrh", "lbgihet", "naiqyufxffdw", "vruh", "uz", "ukffmudygjavem", "dccamymhp", "wofwgjkykm", "fbuujzxhln", "kmm", "lzandlltowjpwsal", "fapfvrmezbsjxs", "wiw", "sc", "soqlh", "hzaplclkwl", "gcdqbcdwbwa", "gadgt", "pgowefka", "juffuguqepwnfh", "nbuinl", "cpdxf", "sox", "fq", "lfnrhgsxkhx", "xrcorfygjxpi", "mwtqjwbhgh", "loc", "fkglorkkvx", "nlzdhucvayrz", "azefobxutitrf", "rlrstkcbtikklmh", "ggk", "sbphcejuylh", "nraoenhd", "zngyodiqlchxyycx", "rrbhfwohfv", "krzolrglgn", "cpjesdzy", "yoifoyg", "hqqevqjugi", "ahmv", "xgaujnyclcjq", "evhyfnlohavrj", "byyvhgh", "hyw", "kedhvwy", "ysljsqminajfipds", "rglnpxfqwu", "cibpynkxg", "su", "mbntqrlwyampdg", "nig", "ldhlhqdyjcfhu", "jfymrbafmyoc", "tyjmnhlfnrtz", "dlazixtlxyvm", "fbiguhsfuqo", "rhymsno", "rkbdlchs", "ocbbwwd", "astaiamnepwkya", "mplirup", "edkxjq", "g", "exlwulswtvot", "tlnc", "vnrrzerz", "ygeraoozbtt", "yyifkin", "eo", "ua", "qgztvqdolf", "rlzddjzcshvd", "khxkdxflwxme", "kk", "zylbhoaac", "cw", "iizic", "gcdxstpz", "kjwdqeg", "earjrncmmkdel", "kbesuhquepj", "nrzbllldgdmyrpgl", "hllwnqozf", "djpchowhwevbqvjj", "zsmhylnjpktb", "pxnktxkm", "fxwiaqqb", "qjwufmwresfsfaok", "aa", "d", "iobioqm", "svjgzk", "khbzp", "euexyudhrioi", "yqsj", "ngrwqpoh", "rwuvd", "eruffmlg", "bxzovyew", "faz", "pmvfvyguqdi", "jlxnoixsy", "hyfrdngjf", "ly", "eibcapetpmeaid", "tpnwwiif", "pfgsp", "kvhhwkzvtvlhhb", "pjxurgqbtldims", "rncplkeweoirje", "akyprzzphew", "wyvfpjyglzrmhfqp", "ubheeqt", "rmbxlcmn", "taqakgim", "apsbu", "khwnykughmwrlk", "vtdlzwpbhcsbvjno", "tffmjggrmyil", "schgwrrzt", "mvndmua", "nlwpw", "glvbtkegzjs", "piwllpgnlpcnezqs", "xkelind", "urtxsezrwz", "zechoc", "vfaimxrqnyiq", "ybugjsblhzfravzn", "btgcpqwovwp", "zgxgodlhmix", "sfzdknoxzassc", "wgzvqkxuqrsqxs", "dwneyqisozq", "fg", "vhfsf", "uspujvqhydw", "eadosqafyxbmzgr", "tyff", "blolplosqnfcwx", "uwkl", "puenodlvotb", "iizudxqjvfnky", "cjcywjkfvukvveq", "jrxd", "igwb", "dftdyelydzyummmt", "uvfmaicednym", "oai", "higfkfavgeemcgo", "naefganqo", "iqebfibigljbc", "ulicojzjfrc", "igxprunj", "cymbrl", "fqmwciqtynca", "zjyagi", "mzuejrttefhdwqc", "zyiurxvf", "wrjxffzbjexsh", "wrxw", "mhrbdxjwi", "htknfa", "wfrvxqdkhbwwef", "vqsghhhutdget", "cwupzrts", "hbjnb", "wpccoa", "nx", "howbzhaoscgyk", "bilt", "wqqatye", "zceuuwg", "jxzon", "kkfj", "bwsezd", "ifdegsyjtswselk", "xweimxlnzoh", "tqthlftjblnpht", "ww", "ss", "b", "jmruuqscwjp", "nxbk", "wd", "cqkrtbxgzg", "xhppcjnq", "cfq", "tkkolzcfi", "wblxki", "ijeglxsvc", "kcqjjwcwuhvzydm", "gubqavlqffhrzz", "hiwxrgftittd", "caybc", "ncsyjlzlxyyklc", "poxcgnexmaajzuha", "dhaccuualacyl", "mtkewbprs", "oncggqvr", "sqqoffmwkplsgbrp", "ioajuppvqluhbdet", "dzwwzaelmo", "afumtqugec", "wglucmugwqi", "zveswrjevfz", "nxlbkak", "pzcejvxzeoybb", "fd", "vewj", "ivws", "zjhudtpqsfc", "zcmukotirrxx", "zksmx", "umofzhhowyftz", "zbotrokaxaryxlk", "ueolqk", "dxmzhoq", "zvu", "cjl", "esfmqgvxwfy", "npbep", "vbgjtbv", "poeugoqynkbfiv", "fewjjscjrei", "yqssxzsydgllfzmo", "urxkwcypctjkabi", "wqtldwhjouas", "tovdtkr", "onzgeyddkqwuhnim", "ffxviyvsktqrfa", "qujhd", "pvcz", "hiyjlkxmeplnrvxg", "hdykehkefp", "vepcxhozpjxtreyn", "liguhuxudbnh", "f", "ordxzm", "klgohcmmbukz", "yrmooliaobbnlap", "dutnbetocxylcey", "ywdsjegd", "cr", "blbxhjsgcuoxmqft", "ngzdc", "srfyjjumcbxole", "dazwzwtdjoyuqeqj", "xazjarqgfm", "fxyfqbeoktcc", "qrsjchxp", "iltaqzawhgu", "sgenjcfxr", "yfikp", "dvwhbyumthkiktb", "walsx", "jyajrkcvysicisab", "brdeumb", "tviihjwxdcz", "dnrrgmem", "ydgxlrjzucxyid", "cdvdpvjlagwmg", "ngnpxjkxims", "gvyhnchlimsxc", "w", "jtizpezjl", "qe", "rjzv", "vhnqvi", "qm", "iedzqswrsnfmnn", "lt", "utqfcqyrrwm", "wtelvsqrru", "fjwrhjcrtbcytn", "qmqxceuohpiffaq", "rmoybqjjgdyo", "pmxttqftypfexlv", "tg", "qa", "iqbqjlnpbf", "kgaynkddbzllecd", "tccvslp", "curkxfoimnw", "fvnyqkzlheruxr", "iiygnzfov", "coqs", "oa", "eiu", "vzemmxtklis", "lxu", "nrwsjaxzwmh", "tdayz", "oxbbemejgosgcynf", "ykbcn", "hesvnctfvdsp", "ku", "rjhykpadahbhj", "at", "sxlngbtxmqr", "wqrom", "qzyabzrco", "rbbyklndcqdj", "cnsmgmwmpbgjq", "krvnaf", "qrwfajnfahyqocdb", "fnlaozmff", "vmoymbmytjvfcgt", "cijyy", "jdgwjbztl", "swmalgbgpaplqgz", "hfl", "typttkrpfvx", "tkzpzrscwbx", "bwfqqvjcukjbsg", "nxqmxr", "x", "eyavnz", "il", "dhthp", "eyelg", "npsoqsw", "reogbmveofvusdsx", "jvdrjkhxkq", "qzjbrpljwuzpl", "czqeevvbvcwh", "vzuszqvhlmapty", "yu", "yldwwgezlqur", "vorxwgdtgjilgydq", "pknt", "bgihl", "ckorgrm", "ixylxjmlfv", "bpoaboylced", "zea", "igfagitkrext", "ipvqq", "dmoerc", "oqxbypihdv", "dtjrrkxro", "rexuhucxpi", "bvmuyarjwqpcoywa", "qwdmfpwvamisns", "bhopoqdsref", "tmnm", "cre", "ktrniqwoofoeenbz", "vlrfcsftapyujmw", "updqikocrdyex", "bcxw", "eaum", "oklsqebuzeziisw", "fzgyhvnwjcns", "dybjywyaodsyw", "lmu", "eocfru", "ztlbggsuzctoc", "ilfzpszgrgj", "imqypqo", "fump", "sjvmsbrcfwretbie", "oxpmplpcg", "wmqigymr", "qevdyd", "gmuyytguexnyc", "hwialkbjgzc", "lmg", "gijjy", "lplrsxznfkoklxlv", "xrbasbznvxas", "twn", "bhqultkyfq", "saeq", "xbuw", "zd", "kng", "uoay", "kfykd", "armuwp", "gtghfxf", "gpucqwbihemixqmy", "jedyedimaa", "pbdrx", "toxmxzimgfao", "zlteob", "adoshnx", "ufgmypupei", "rqyex", "ljhqsaneicvaerqx", "ng", "sid", "zagpiuiia", "re", "oadojxmvgqgdodw", "jszyeruwnupqgmy", "jxigaskpj", "zpsbhgokwtfcisj", "vep", "ebwrcpafxzhb", "gjykhz", "mfomgxjphcscuxj", "iwbdvusywqlsc", "opvrnx", "mkgiwfvqfkotpdz", "inpobubzbvstk", "vubuucilxyh", "bci", "dibmye", "rlcnvnuuqfvhw", "oorbyyiigppuft", "swpksfdxicemjbf", "goabwrqdoudf", "yjutkeqakoarru", "wuznnlyd", "vfelxvtggkkk", "mxlwbkbklbwfsvr", "advraqovan", "smkln", "jxxvzdjlpyurxpj", "ssebtpznwoytjefo", "dynaiukctgrzjx", "irzosjuncvh", "hcnhhrajahitn", "vwtifcoqepqyzwya", "kddxywvgqxo", "syxngevs", "batvzmziq", "mjewiyo", "pzsupxoflq", "byzhtvvpj", "cqnlvlzr", "akvmxzbaei", "mwo", "vg", "ekfkuajjogbxhjii", "isdbplotyak", "jvkmxhtmyznha", "lqjnqzrwrmgt", "mbbhfli", "bpeohsufree", "ajrcsfogh", "lucidbnlysamvy", "tutjdfnvhahxy", "urbrmmadea", "hghv", "acnjx", "athltizloasimp", "gu", "rjfozvgmdakdhao", "iephs", "uztnpqhdl", "rfuyp", "crcszmgplszwfn", "zihegt", "xbspa", "cjbmsamjyqqrasz", "ghzlgnfoas", "ljxl", "cnumquohlcgt", "jm", "mfccj", "hfedli", "vtpieworwhyiucs", "tdtuquartspkotm", "pnkeluekvelj", "ugrloq", "zljmwt", "fkyvqguqq", "tpjidglpxqfxv", "l", "tvvimvroz", "yy", "opwyfovdz", "pwlumocnyuoume", "vjqpzkcfc", "ihicd", "dtttiixlhpikbv", "goblttgvmndkqgg", "gwsibcqahmyyeagk", "prtvoju", "lcblwidhjpu", "kbu", "pey", "gkzrpc", "bqajopjjlfthe", "bc", "lqs", "zkndgojnjnxqsoqi", "zyesldujjlp", "drswybwlfyzph", "xzluwbtmoxokk", "bedrqfui", "opajzeahv", "lehdfnr", "mnlpimduzgmwszc", "velbhj", "miwdn", "wruqc", "kscfodjxg", "wcbm"}, []string{"give"}},
	}

	for _, d := range datas {
		expected := findAllConcatenatedWordsInADict0(d.words)
		if ans := findAllConcatenatedWordsInADict(d.words); !reflect.DeepEqual(toMap(ans), toMap(expected)) {
			fmt.Printf("wrong == > \n%v \n%v\n", ans, expected)
		} else {
			fmt.Printf("ok ! %v == %v\n", ans, expected)
		}
	}
}

func toMap(arr []string) map[string]struct{} {
	ans := map[string]struct{}{}
	for _, s := range arr {
		ans[s] = struct{}{}
	}
	return ans
}

func TestNum1044(t *testing.T) {
	datas := []struct {
		s, ans string
	}{
		//{"banana", "ana"},
		//{"abcd", ""},
		//{"maddmauhrsy", "ma"},
		//{"hhslphslbf", "hsl"},
		{"njsihcfynghcfpynjsi", "njsi"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
	}
	for _, d := range datas {
		if v := longestDupSubstring(d.s); v != d.ans {
			fmt.Printf("%s != %s\n", v, d.ans)
		} else {
			fmt.Println("yes")
		}
	}
}
func TestNum686(t *testing.T) {
	datas := []struct {
		a, b string
		ans  int
	}{
		{"abcd", "cdabcdab", 3},
		{"a", "aa", 2},
		{"a", "a", 1},
		{"abc", "wxyz", -1},
	}
	for _, x := range datas {
		fmt.Println(repeatedStringMatch(x.a, x.b) == x.ans)
	}
}
func TestNum1154(t *testing.T) {
	dayOfYear("2004-03-01")
}
func TestNum748(t *testing.T) {
	datas := []struct {
		licensePlate, ans string
		words             []string
	}{
		//{"iMSlpe4", "simple", []string{"claim","consumer","student","camera","public","never","wonder","simple","thought","use"}},
		//{"OgEu755", "enough", []string{"enough","these","play","wide","wonder","box","arrive","money","tax","thus"}},
		//{"Ah71752", "husband", []string{"suggest","letter","of","husband","easy","education","drug","prevent","writer","old"}},
		//{"1s3 456", "pest", []string{"looks", "pest", "stew", "show"}},
		//{"1s3 PSt", "steps", []string{"step", "steps", "stripe", "stepple"}},
		{"AN87005", "important", []string{"participant", "individual", "start", "exist", "above", "already", "easy", "attack", "player", "important"}},
	}
	for _, d := range datas {
		fmt.Println(shortestCompletingWord(d.licensePlate, d.words) == d.ans)
	}
}
func TestNum44(t *testing.T) {
	datas := []struct {
		s, p string
		ans  bool
	}{
		{"abcabczzzde", "*abc???de*", true},
		{"zacabz", "*a?b*", false},
		{"acdcb", "a*c?b", false},
		{"adceb", "*a*b", true},
		{"cb", "?a", false},
		{"aa", "*", true},
		{"aa", "a", false},
		{"", "*", true},
		{"", "", true},
	}
	for _, x := range datas {
		fmt.Println(isMatch(x.s, x.p) == x.ans)
	}
}
func TestNum43(t *testing.T) {
	datas := []struct {
		num1, num2, ans string
	}{
		{"12", "3", "36"},
		{"9", "9", "81"},
		{"123", "456", "56088"},
		{"123", "0", "0"},
	}
	for _, x := range datas {
		fmt.Println(multiply(x.num1, x.num2) == x.ans)
	}
}
func TestNum1446(t *testing.T) {
	datas := []struct {
		s   string
		ans int
	}{
		{"j", 1},
		{"cc", 2},
		{"tourist", 1},
		{"hooraaaaaaaaaaay", 11},
		{"triplepillooooow", 5},
		{"abbcccddddeeeeedcba", 5},
		{"leetcode", 2},
	}
	for _, x := range datas {
		fmt.Println(reflect.DeepEqual(maxPower(x.s), x.ans))
	}
}
