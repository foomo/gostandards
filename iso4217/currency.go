package iso4217

// Currency https://en.wikipedia.org/wiki/ISO_4217
type Currency string

const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BOV Currency = "BOV"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHE Currency = "CHE"
	CHF Currency = "CHF"
	CHW Currency = "CHW"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	COU Currency = "COU"
	CRC Currency = "CRC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MRU Currency = "MRU"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MXV Currency = "MXV"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SLE Currency = "SLE"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	SSP Currency = "SSP"
	STN Currency = "STN"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	USN Currency = "USN"
	UYI Currency = "UYI"
	UYU Currency = "UYU"
	UYW Currency = "UYW"
	UZS Currency = "UZS"
	VED Currency = "VED"
	VES Currency = "VES"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XBA Currency = "XBA"
	XBB Currency = "XBB"
	XBC Currency = "XBC"
	XBD Currency = "XBD"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPD Currency = "XPD"
	XPF Currency = "XPF"
	XPT Currency = "XPT"
	XSU Currency = "XSU"
	XTS Currency = "XTS"
	XUA Currency = "XUA"
	XXX Currency = "XXX"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMW Currency = "ZMW"
	ZWL Currency = "ZWL"
)

func (c Currency) String() string {
	return string(c)
}

func (c Currency) Valid() bool {
	switch c {
	case AED,
		AFN,
		ALL,
		AMD,
		ANG,
		AOA,
		ARS,
		AUD,
		AWG,
		AZN,
		BAM,
		BBD,
		BDT,
		BGN,
		BHD,
		BIF,
		BMD,
		BND,
		BOB,
		BOV,
		BRL,
		BSD,
		BTN,
		BWP,
		BYN,
		BZD,
		CAD,
		CDF,
		CHE,
		CHF,
		CHW,
		CLF,
		CLP,
		CNY,
		COP,
		COU,
		CRC,
		CUP,
		CVE,
		CZK,
		DJF,
		DKK,
		DOP,
		DZD,
		EGP,
		ERN,
		ETB,
		EUR,
		FJD,
		FKP,
		GBP,
		GEL,
		GHS,
		GIP,
		GMD,
		GNF,
		GTQ,
		GYD,
		HKD,
		HNL,
		HTG,
		HUF,
		IDR,
		ILS,
		INR,
		IQD,
		IRR,
		ISK,
		JMD,
		JOD,
		JPY,
		KES,
		KGS,
		KHR,
		KMF,
		KPW,
		KRW,
		KWD,
		KYD,
		KZT,
		LAK,
		LBP,
		LKR,
		LRD,
		LSL,
		LYD,
		MAD,
		MDL,
		MGA,
		MKD,
		MMK,
		MNT,
		MOP,
		MRU,
		MUR,
		MVR,
		MWK,
		MXN,
		MXV,
		MYR,
		MZN,
		NAD,
		NGN,
		NIO,
		NOK,
		NPR,
		NZD,
		OMR,
		PAB,
		PEN,
		PGK,
		PHP,
		PKR,
		PLN,
		PYG,
		QAR,
		RON,
		RSD,
		RUB,
		RWF,
		SAR,
		SBD,
		SCR,
		SDG,
		SEK,
		SGD,
		SHP,
		SLE,
		SLL,
		SOS,
		SRD,
		SSP,
		STN,
		SVC,
		SYP,
		SZL,
		THB,
		TJS,
		TMT,
		TND,
		TOP,
		TRY,
		TTD,
		TWD,
		TZS,
		UAH,
		UGX,
		USD,
		USN,
		UYI,
		UYU,
		UYW,
		UZS,
		VED,
		VES,
		VND,
		VUV,
		WST,
		XAF,
		XAG,
		XAU,
		XBA,
		XBB,
		XBC,
		XBD,
		XCD,
		XDR,
		XOF,
		XPD,
		XPF,
		XPT,
		XSU,
		XTS,
		XUA,
		XXX,
		YER,
		ZAR,
		ZMW,
		ZWL:
		return true
	default:
		return false
	}
}
