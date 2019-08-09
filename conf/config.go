package conf

var CorpID = "ww0c481bb483b34ff2"
var NotifySuiteID = "ww5b2c93da21371621"
var GuardSuiteID = "wwed407427d5408872"
var ScoreSuiteID = "wwa417172c7acf23ec"
var TimetableSuiteID = "wwb5ea7647ad7a9d9b"
var RepairSuiteID = "wwb6ac6cdfe58a944f"
var DEBUG = false
var Detail = "GO_DEV"
var WorkingDirectory  = "./"
var GN_WEB_HOST = "http://xsqywxnotify.xiaosikeji.com"
var API_HOST = "http://xsunifiedapi.xiaosikeji.com"
var REPAIR_WEB_HOST = "http://repairvue.xiaosikeji.com"

var Token = "8w9h86nyei94dcky5gv6uz1y50jj9k309jrva2zcule6pduojlb8t0j229lf81qn"

var PayConfig = map[string]string{
	"appid":"ww0c481bb483b34ff2",
	"mchid":"1526191341",
	"mchkey":"6tqeoi3nxz137uv5p6x94yb5xtyb4rfo",
}


var UnifiedConf = map[string]string{
	"CorpID":                 CorpID,
	"Provider_Secret":        "T5F01OyyB5tdS0zEO8Fr9KCCYfyJJ5_5rB7ZhdgT0d8",
	"ProviderAccessTokenKey": CorpID + "ProviderAccessTokenKey" +Detail,

	"NotifySecret":              "a860ewcfw9AbBup7av761Z51fN90SO_aYH3ErmScfjw",
	"NotifyToken":               "NRBkRt2nrk6z",
	"NotifyEncodingAESKey":      "JFTtPeh9CKsWd8yn2NfcYriTUhR9epQEGCiP1N6lPkN",
	"NotifySuiteTicketKey":      NotifySuiteID + "NotifySuiteTicketKey" +Detail,
	"NotifySuiteAccessTokenKey": CorpID + NotifySuiteID + "NotifySuiteAccessTokenKey" +Detail,
	"NotifyPreAuthCodeKey":      CorpID + NotifySuiteID + "NotifyPreAuthCodeKey" +Detail,

	"GuardSecret":              "plxtseBV1fAzdnO6PFKSagNRHSzUHp1KufCTojit0vM",
	"GuardToken":               "OisB56EgufEtnA",
	"GuardEncodingAESKey":      "aSwbttkSSUqTeWyQtenxEkRw8KwmJpd4o2m4iWPtJsI",
	"GuardSuiteTicketKey":      GuardSuiteID + "GuardSuiteTicketKey" +Detail,
	"GuardSuiteAccessTokenKey": CorpID + GuardSuiteID + "GuardSuiteAccessTokenKey" +Detail,
	"GuardPreAuthCodeKey":      CorpID + GuardSuiteID + "GuardPreAuthCodeKey" +Detail,

	"TimetableSecret":              "rjEZI44PcymhLVoxWS5Z2f5PImc3YFuhwZT5XVWQ2-Y",
	"TimetableToken":               "z8inRzWAbr66DjMAhHdn9D",
	"TimetableEncodingAESKey":      "6BZSNot2LNowmKKfMillApd0ODfRQAOxPMaaejioZek",
	"TimetableSuiteTicketKey":      TimetableSuiteID + "TimetableSuiteTicketKey" +Detail,
	"TimetableSuiteAccessTokenKey": CorpID + TimetableSuiteID + "TimetableSuiteAccessTokenKey" +Detail,
	"TimetablePreAuthCodeKey":      CorpID + TimetableSuiteID + "TimetablePreAuthCodeKey" +Detail,


	"ScoreSecret":              	"AOK_VqkWwbBl7rzZawZXriGMpxtYbCPhygkPHb3yApQ",
	"ScoreToken":              		"fgjHrsLX37C9LYAGwuMilEi",
	"ScoreEncodingAESKey":     	 	"16MLw4sFBmkSlRKAa9W7P9cZKPuwmU69L12QABXPg7p",
	"ScoreSuiteTicketKey":     	    ScoreSuiteID + "ScoreSuiteTicketKey" +Detail,
	"ScoreSuiteAccessTokenKey": 	CorpID + ScoreSuiteID + "ScoreSuiteAccessTokenKey" +Detail,
	"ScorePreAuthCodeKey":      	CorpID + ScoreSuiteID + "ScorePreAuthCodeKey" +Detail,

	"RepairSecret":              	"Fv3q0RFexHm1H-Epo4NkLJ3Mpj3qWtcNNVusGv4Ys6E",
	"RepairToken":              	"M6zWLH44GpnRwPyAXvV9oWb",
	"RepairEncodingAESKey":     	 "14sJ6TFYKgzH543IxSw1jgudqO7CM5M6TtBU9IJRJmX",
	"RepairSuiteTicketKey":     	RepairSuiteID + "RepairSuiteTicketKey" +Detail,
	"RepairSuiteAccessTokenKey": 	CorpID + RepairSuiteID + "RepairSuiteAccessTokenKey" +Detail,
	"RepairPreAuthCodeKey":      	CorpID + RepairSuiteID + "RepairPreAuthCodeKey" +Detail,


}
var SettingCode = map[string]interface{}{
	"IsCafeteriaInstalled": "00001",
	"IsCardVipInstalled": "00002",
}