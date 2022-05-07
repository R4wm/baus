package baus

const (
	GENESIS             = 1
	EXODUS              = 1534
	LEVITICUS           = 2747
	NUMBERS             = 3606
	DEUTERONOMY         = 4894
	JOSHUA              = 5853
	JUDGES              = 6511
	RUTH                = 7129
	FIRSTSAMUEL         = 7214
	SECONDSAMUEL        = 8024
	FIRSTKINGS          = 8719
	SECONDKINGS         = 9535
	FIRSTCHRONICLES     = 10254
	SECONDCHRONICLES    = 11196
	EZRA                = 12018
	NEHEMIAH            = 12298
	ESTHER              = 12704
	JOB                 = 12871
	PSALMS              = 13941
	PROVERBS            = 16402
	ECCLESIASTES        = 17317
	SONGOFSOLOMON       = 17539
	ISAIAH              = 17656
	JEREMIAH            = 18948
	LAMENTATIONS        = 20312
	EZEKIEL             = 20466
	DANIEL              = 21739
	HOSEA               = 22096
	JOEL                = 22293
	AMOS                = 22366
	OBADIAH             = 22512
	JONAH               = 22533
	MICAH               = 22581
	NAHUM               = 22686
	HABAKKUK            = 22733
	ZEPHANIAH           = 22789
	HAGGAI              = 22842
	ZECHARIAH           = 22880
	MALACHI             = 23091
	MATTHEW             = 23146
	MARK                = 24217
	LUKE                = 24895
	JOHN                = 26046
	ACTS                = 26925
	ROMANS              = 27932
	FIRSTCORINTHIANS    = 28365
	SECONDCORINTHIANS   = 28802
	GALATIANS           = 29059
	EPHESIANS           = 29208
	PHILIPPIANS         = 29363
	COLOSSIANS          = 29467
	FIRSTTHESSALONIANS  = 29562
	SECONDTHESSALONIANS = 29651
	FIRSTTIMOTHY        = 29698
	SECONDTIMOTHY       = 29811
	TITUS               = 29894
	PHILEMON            = 29940
	HEBREWS             = 29965
	JAMES               = 30268
	FIRSTPETER          = 30376
	SECONDPETER         = 30481
	FIRSTJOHN           = 30542
	SECONDJOHN          = 30647
	THIRDJOHN           = 30660
	JUDE                = 30674
	REVELATION          = 30699
)

func getOrdinalBookNumber(book string) int {
	switch book {
	case "GENESIS":
		return GENESIS
	case "EXODUS":
		return EXODUS
	case "LEVITICUS":
		return LEVITICUS
	case "NUMBERS":
		return NUMBERS
	case "DEUTERONOMY":
		return DEUTERONOMY
	case "JOSHUA":
		return JOSHUA
	case "JUDGES":
		return JUDGES
	case "RUTH":
		return RUTH
	case "FIRSTSAMUEL":
		return FIRSTSAMUEL
	case "SECONDSAMUEL":
		return SECONDSAMUEL
	case "FIRSTKINGS":
		return FIRSTKINGS
	case "SECONDKINGS":
		return SECONDKINGS
	case "FIRSTCHRONICLES":
		return FIRSTCHRONICLES
	case "SECONDCHRONICLES":
		return SECONDCHRONICLES
	case "EZRA":
		return EZRA
	case "NEHEMIAH":
		return NEHEMIAH
	case "ESTHER":
		return ESTHER
	case "JOB":
		return JOB
	case "PSALMS":
		return PSALMS
	case "PROVERBS":
		return PROVERBS
	case "ECCLESIASTES":
		return ECCLESIASTES
	case "SONGOFSOLOMON":
		return SONGOFSOLOMON
	case "ISAIAH":
		return ISAIAH
	case "JEREMIAH":
		return JEREMIAH
	case "LAMENTATIONS":
		return LAMENTATIONS
	case "EZEKIEL":
		return EZEKIEL
	case "DANIEL":
		return DANIEL
	case "HOSEA":
		return HOSEA
	case "JOEL":
		return JOEL
	case "AMOS":
		return AMOS
	case "OBADIAH":
		return OBADIAH
	case "JONAH":
		return JONAH
	case "MICAH":
		return MICAH
	case "NAHUM":
		return NAHUM
	case "HABAKKUK":
		return HABAKKUK
	case "ZEPHANIAH":
		return ZEPHANIAH
	case "HAGGAI":
		return HAGGAI
	case "ZECHARIAH":
		return ZECHARIAH
	case "MALACHI":
		return MALACHI
	case "MATTHEW":
		return MATTHEW
	case "MARK":
		return MARK
	case "LUKE":
		return LUKE
	case "JOHN":
		return JOHN
	case "ACTS":
		return ACTS
	case "ROMANS":
		return ROMANS
	case "FIRSTCORINTHIANS":
		return FIRSTCORINTHIANS
	case "SECONDCORINTHIANS":
		return SECONDCORINTHIANS
	case "GALATIANS":
		return GALATIANS
	case "EPHESIANS":
		return EPHESIANS
	case "PHILIPPIANS":
		return PHILIPPIANS
	case "COLOSSIANS":
		return COLOSSIANS
	case "FIRSTTHESSALONIANS":
		return FIRSTTHESSALONIANS
	case "SECONDTHESSALONIANS":
		return SECONDTHESSALONIANS
	case "FIRSTTIMOTHY":
		return FIRSTTIMOTHY
	case "SECONDTIMOTHY":
		return SECONDTIMOTHY
	case "TITUS":
		return TITUS
	case "PHILEMON":
		return PHILEMON
	case "HEBREWS":
		return HEBREWS
	case "JAMES":
		return JAMES
	case "FIRSTPETER":
		return FIRSTPETER
	case "SECONDPETER":
		return SECONDPETER
	case "FIRSTJOHN":
		return FIRSTJOHN
	case "SECONDJOHN":
		return SECONDJOHN
	case "THIRDJOHN":
		return THIRDJOHN
	case "JUDE":
		return JUDE
	case "REVELATION":
		return REVELATION
	default:
		return 0
	}
}