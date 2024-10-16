package limit

import (
	currency_lib "github.com/code-payments/code-server/pkg/currency"
)

type SendLimit struct {
	PerTransaction float64
	Daily          float64
}

type MicroPaymentLimit struct {
	Min float64
	Max float64
}

func init() {
	newMicroPaymentLimits := make(map[currency_lib.Code]MicroPaymentLimit)
	for currency, limit := range MicroPaymentLimits {
		newMicroPaymentLimits[currency] = MicroPaymentLimit{
			Min: limit.Min,
			Max: 5.0 * limit.Max,
		}
	}
	MicroPaymentLimits = newMicroPaymentLimits
}

// todo: Better way of managing all of this
var (
	SendLimits = map[currency_lib.Code]SendLimit{
		"aed": {PerTransaction: 1_000.00, Daily: 3_500.00},
		"afn": {PerTransaction: 20_000.00, Daily: 85_000.00},
		"all": {PerTransaction: 25_000.00, Daily: 100_000.00},
		"amd": {PerTransaction: 100_000.00, Daily: 400_000.00},
		"ang": {PerTransaction: 500.00, Daily: 1_750.00},
		"aoa": {PerTransaction: 100_000.00, Daily: 500_000.00},
		"ars": {PerTransaction: 25_000.00, Daily: 125_000.00},
		"aud": {PerTransaction: 250.00, Daily: 1_250.00},
		"awg": {PerTransaction: 500.00, Daily: 1_500.00},
		"azn": {PerTransaction: 250.00, Daily: 1_500.00},
		"bam": {PerTransaction: 500.00, Daily: 1_500.00},
		"bbd": {PerTransaction: 500.00, Daily: 2_000.00},
		"bdt": {PerTransaction: 25_000.00, Daily: 90_000.00},
		"bgn": {PerTransaction: 500.00, Daily: 1_750.00},
		"bhd": {PerTransaction: 100.00, Daily: 350.00},
		"bif": {PerTransaction: 500_000.00, Daily: 2_000_000.00},
		"bmd": {PerTransaction: 250.00, Daily: 1_000.00},
		"bnd": {PerTransaction: 250.00, Daily: 1_250.00},
		"bob": {PerTransaction: 1_500.00, Daily: 6_500.00},
		"brl": {PerTransaction: 1_000.00, Daily: 5_000.00},
		"bsd": {PerTransaction: 250.00, Daily: 1_000.00},
		"btn": {PerTransaction: 20_000.00, Daily: 75_000.00},
		"bwp": {PerTransaction: 2_500.00, Daily: 10_000.00},
		"byn": {PerTransaction: 500.00, Daily: 2_500.00},
		"byr": {PerTransaction: 5_000_000.00, Daily: 17_500_000.00},
		"bzd": {PerTransaction: 500.00, Daily: 2_000.00},
		"cad": {PerTransaction: 250.00, Daily: 1_250.00},
		"cdf": {PerTransaction: 500_000.00, Daily: 2_000_000.00},
		"chf": {PerTransaction: 250.00, Daily: 900.00},
		"clf": {PerTransaction: 8.00, Daily: 30.00},
		"clp": {PerTransaction: 150_000.00, Daily: 800_000.00},
		"cny": {PerTransaction: 1_500.00, Daily: 6_500.00},
		"cop": {PerTransaction: 1_000_000.00, Daily: 4_000_000.00},
		"crc": {PerTransaction: 150_000.00, Daily: 60_0000.00},
		"cuc": {PerTransaction: 250.00, Daily: 1_000.00},
		"cup": {PerTransaction: 5_000.00, Daily: 25_000.00},
		"cve": {PerTransaction: 25_000.00, Daily: 100_000.00},
		"czk": {PerTransaction: 5_000.00, Daily: 22_500.00},
		"djf": {PerTransaction: 40_000.00, Daily: 175_000.00},
		"dkk": {PerTransaction: 1_000.00, Daily: 7_000.00},
		"dop": {PerTransaction: 12_500.00, Daily: 50_000.00},
		"dzd": {PerTransaction: 25_000.00, Daily: 125_000.00},
		"egp": {PerTransaction: 2_500.00, Daily: 17_500.00},
		"ern": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"etb": {PerTransaction: 12_500.00, Daily: 50_000.00},
		"eur": {PerTransaction: 250.00, Daily: 999.99},
		"fjd": {PerTransaction: 500.00, Daily: 2_000.00},
		"fkp": {PerTransaction: 250.00, Daily: 800.00},
		"gbp": {PerTransaction: 250.00, Daily: 800.00},
		"gel": {PerTransaction: 500.00, Daily: 2_500.00},
		"ggp": {PerTransaction: 250.00, Daily: 800.00},
		"ghs": {PerTransaction: 2_250.00, Daily: 12_500.00},
		"gip": {PerTransaction: 250.00, Daily: 800.00},
		"gmd": {PerTransaction: 15_000.00, Daily: 50_000.00},
		"gnf": {PerTransaction: 2_000_000.00, Daily: 8_500_000.00},
		"gtq": {PerTransaction: 1_500.00, Daily: 7_500.00},
		"gyd": {PerTransaction: 50_000.00, Daily: 200_000.00},
		"hkd": {PerTransaction: 1_000.00, Daily: 7_500.00},
		"hnl": {PerTransaction: 5_000.00, Daily: 22_500.00},
		"hrk": {PerTransaction: 1_500.00, Daily: 7_250.00},
		"htg": {PerTransaction: 25_000.00, Daily: 125_000.00},
		"huf": {PerTransaction: 50_000.00, Daily: 375_000.00},
		"idr": {PerTransaction: 1_500_000.00, Daily: 14_000_000.00},
		"ils": {PerTransaction: 750.00, Daily: 3_000.00},
		"imp": {PerTransaction: 250.00, Daily: 800.00},
		"inr": {PerTransaction: 10_000.00, Daily: 80_000.00},
		"iqd": {PerTransaction: 250_000.00, Daily: 1_250_000.00},
		"irr": {PerTransaction: 10_000_000.00, Daily: 40_000_000.00},
		"isk": {PerTransaction: 25_000.00, Daily: 125_000.00},
		"jep": {PerTransaction: 250.00, Daily: 800.00},
		"jmd": {PerTransaction: 25_000.00, Daily: 150_000.00},
		"jod": {PerTransaction: 100.00, Daily: 700.00},
		"jpy": {PerTransaction: 25_000.00, Daily: 125_000.00},
		"kes": {PerTransaction: 25_000.00, Daily: 100_000.00},
		"kgs": {PerTransaction: 10_000.00, Daily: 80_000.00},
		"khr": {PerTransaction: 1_000_000.00, Daily: 4_000_000.00},
		"kmf": {PerTransaction: 100_000.00, Daily: 400_000.00},
		"kpw": {PerTransaction: 250_000.00, Daily: 900_000.00},
		"krw": {PerTransaction: 250_000.00, Daily: 1_250_000.00},
		"kwd": {PerTransaction: 50.00, Daily: 300.00},
		"kyd": {PerTransaction: 250.00, Daily: 800.00},
		"kzt": {PerTransaction: 100_000.00, Daily: 400_000.00},
		"lak": {PerTransaction: 2_500_000.00, Daily: 15_000_000.00},
		"lbp": {PerTransaction: 3_000_000.00, Daily: 15_000_000.00},
		"lkr": {PerTransaction: 50_000.00, Daily: 350_000.00},
		"lrd": {PerTransaction: 25_000.00, Daily: 150_000.00},
		"lsl": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"ltl": {PerTransaction: 1_000.00, Daily: 2_500.00},
		"lvl": {PerTransaction: 150.00, Daily: 600.00},
		"lyd": {PerTransaction: 1_250.00, Daily: 4_500.00},
		"mad": {PerTransaction: 2_500.00, Daily: 10_000.00},
		"mdl": {PerTransaction: 2_500.00, Daily: 17_500.00},
		"mga": {PerTransaction: 1_000_000.00, Daily: 4_000_000.00},
		"mkd": {PerTransaction: 15_000.00, Daily: 50_000.00},
		"mmk": {PerTransaction: 500_000.00, Daily: 2_000_000.00},
		"mnt": {PerTransaction: 750_000.00, Daily: 3_000_000.00},
		"mop": {PerTransaction: 2_000.00, Daily: 8_000.00},
		"mru": {PerTransaction: 10000.00, Daily: 35000.00},
		"mur": {PerTransaction: 10_000.00, Daily: 45_000.00},
		"mvr": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"mwk": {PerTransaction: 250_000.00, Daily: 1_000_000.00},
		"mxn": {PerTransaction: 2_500.00, Daily: 17_500.00},
		"myr": {PerTransaction: 500.00, Daily: 4_000.00},
		"mzn": {PerTransaction: 15_000.00, Daily: 60_000.00},
		"nad": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"ngn": {PerTransaction: 400_000.00, Daily: 1_600_000.00},
		"nio": {PerTransaction: 5_000.00, Daily: 35_000.00},
		"nok": {PerTransaction: 2500.00, Daily: 9000.00},
		"npr": {PerTransaction: 25_000.00, Daily: 125_000.00},
		"nzd": {PerTransaction: 500.00, Daily: 1_500.00},
		"omr": {PerTransaction: 50.00, Daily: 350.00},
		"pab": {PerTransaction: 250.00, Daily: 1_000.00},
		"pen": {PerTransaction: 500.00, Daily: 3_500.00},
		"pgk": {PerTransaction: 1_000.00, Daily: 3_500.00},
		"php": {PerTransaction: 10_000.00, Daily: 50_000.00},
		"pkr": {PerTransaction: 50_000.00, Daily: 250_000.00},
		"pln": {PerTransaction: 750.00, Daily: 4_500.00},
		"pyg": {PerTransaction: 1_500_000.00, Daily: 6_500_000.00},
		"qar": {PerTransaction: 500.00, Daily: 3_500.00},
		"ron": {PerTransaction: 500.00, Daily: 4_500.00},
		"rsd": {PerTransaction: 25_000.00, Daily: 100_000.00},
		"rub": {PerTransaction: 15_000.00, Daily: 60_000.00},
		"rwf": {PerTransaction: 250_000.00, Daily: 1_000_000.00},
		"sar": {PerTransaction: 750.00, Daily: 3_750.00},
		"sbd": {PerTransaction: 2_000.00, Daily: 8_000.00},
		"scr": {PerTransaction: 3_000.00, Daily: 12_500.00},
		"sdg": {PerTransaction: 150_000.00, Daily: 500_000.00},
		"sek": {PerTransaction: 2_500.00, Daily: 10_000.00},
		"sgd": {PerTransaction: 250.00, Daily: 1_250.00},
		"shp": {PerTransaction: 250.00, Daily: 1_250.00},
		"sll": {PerTransaction: 5_000_000.00, Daily: 12_500_000.00},
		"sos": {PerTransaction: 150_000.00, Daily: 500_000.00},
		"srd": {PerTransaction: 5_000.00, Daily: 22_500.00},
		"std": {PerTransaction: 5_000_000.00, Daily: 20_000_000.00},
		"syp": {PerTransaction: 500_000.00, Daily: 2_500_000.00},
		"szl": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"thb": {PerTransaction: 5_000.00, Daily: 35_000.00},
		"tjs": {PerTransaction: 2_500.00, Daily: 10_000.00},
		"tmt": {PerTransaction: 750.00, Daily: 3_500.00},
		"tnd": {PerTransaction: 750.00, Daily: 3_000.00},
		"top": {PerTransaction: 500.00, Daily: 2_250.00},
		"try": {PerTransaction: 2_500.00, Daily: 17_500.00},
		"ttd": {PerTransaction: 1_500.00, Daily: 6_500.00},
		"twd": {PerTransaction: 5_000.00, Daily: 30_000.00},
		"tzs": {PerTransaction: 500_000.00, Daily: 2_250_000.00},
		"uah": {PerTransaction: 10_000.00, Daily: 35_000.00},
		"ugx": {PerTransaction: 1_000_000.00, Daily: 3_500_000.00},
		"usd": {PerTransaction: 250.00, Daily: 1000.00},
		"uyu": {PerTransaction: 10_000.00, Daily: 40_000.00},
		"uzs": {PerTransaction: 2_500_000.00, Daily: 10_000_000.00},
		"vnd": {PerTransaction: 5_000_000.00, Daily: 22_500_000.00},
		"vuv": {PerTransaction: 25_000.00, Daily: 100_000.00},
		"wst": {PerTransaction: 500.00, Daily: 2_500.00},
		"xaf": {PerTransaction: 150_000.00, Daily: 600_000.00},
		"xcd": {PerTransaction: 500.00, Daily: 2_500.00},
		"xdr": {PerTransaction: 150.00, Daily: 750.00},
		"xof": {PerTransaction: 150_000.00, Daily: 600_000.00},
		"xpf": {PerTransaction: 25_000.00, Daily: 100_000.00},
		"yer": {PerTransaction: 50_000.00, Daily: 250_000.00},
		"zar": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"zmk": {PerTransaction: 2_500_000.00, Daily: 9_000_000.00},
		"zmw": {PerTransaction: 2_500.00, Daily: 15_000.00},
		"zwl": {PerTransaction: 50_000.00, Daily: 300_000.00},
	}

	MicroPaymentLimits = map[currency_lib.Code]MicroPaymentLimit{
		"aed": {Min: 0.20, Max: 4.00},
		"afn": {Min: 4.00, Max: 80.00},
		"all": {Min: 5.00, Max: 100.00},
		"amd": {Min: 20.0, Max: 400.00},
		"ang": {Min: 0.10, Max: 2.00},
		"aoa": {Min: 20.00, Max: 400.00},
		"ars": {Min: 5.00, Max: 100.00},
		"aud": {Min: 0.05, Max: 1.00},
		"awg": {Min: 0.10, Max: 2.00},
		"azn": {Min: 0.05, Max: 1.00},
		"bam": {Min: 0.10, Max: 2.00},
		"bbd": {Min: 0.10, Max: 2.00},
		"bdt": {Min: 5.00, Max: 100.00},
		"bgn": {Min: 0.10, Max: 2.00},
		"bhd": {Min: 0.02, Max: 0.40},
		"bif": {Min: 100.00, Max: 2000.00},
		"bmd": {Min: 0.05, Max: 1.00},
		"bnd": {Min: 0.05, Max: 1.00},
		"bob": {Min: 0.30, Max: 6.00},
		"brl": {Min: 0.20, Max: 4.00},
		"bsd": {Min: 0.05, Max: 1.00},
		"btn": {Min: 4.00, Max: 80.00},
		"bwp": {Min: 0.50, Max: 10.00},
		"byn": {Min: 0.10, Max: 2.00},
		"byr": {Min: 1000.0, Max: 20000.00},
		"bzd": {Min: 0.10, Max: 2.00},
		"cad": {Min: 0.05, Max: 1.00},
		"cdf": {Min: 100.00, Max: 2000.00},
		"chf": {Min: 0.05, Max: 1.00},
		"clf": {Min: 0.01, Max: 0.03},
		"clp": {Min: 30.00, Max: 600.00},
		"cny": {Min: 0.30, Max: 6.00},
		"cop": {Min: 200.00, Max: 4000.00},
		"crc": {Min: 30.00, Max: 600.00},
		"cuc": {Min: 0.05, Max: 1.00},
		"cup": {Min: 1.00, Max: 20.00},
		"cve": {Min: 5.00, Max: 100.00},
		"czk": {Min: 1.00, Max: 20.00},
		"djf": {Min: 8.00, Max: 160.00},
		"dkk": {Min: 0.20, Max: 4.00},
		"dop": {Min: 2.50, Max: 50.00},
		"dzd": {Min: 5.00, Max: 100.00},
		"egp": {Min: 0.50, Max: 10.00},
		"ern": {Min: 0.50, Max: 10.00},
		"etb": {Min: 2.50, Max: 50.00},
		"eur": {Min: 0.05, Max: 1.00},
		"fjd": {Min: 0.10, Max: 2.00},
		"fkp": {Min: 0.05, Max: 1.00},
		"gbp": {Min: 0.05, Max: 1.00},
		"gel": {Min: 0.10, Max: 2.00},
		"ggp": {Min: 0.05, Max: 1.00},
		"ghs": {Min: 0.45, Max: 9.00},
		"gip": {Min: 0.05, Max: 1.00},
		"gmd": {Min: 3.00, Max: 60.00},
		"gnf": {Min: 400.00, Max: 8000.00},
		"gtq": {Min: 0.30, Max: 6.00},
		"gyd": {Min: 10.00, Max: 200.00},
		"hkd": {Min: 0.20, Max: 4.00},
		"hnl": {Min: 1.00, Max: 20.00},
		"hrk": {Min: 0.30, Max: 6.00},
		"htg": {Min: 5.00, Max: 100.00},
		"huf": {Min: 10.00, Max: 200.00},
		"idr": {Min: 300.00, Max: 6000.00},
		"ils": {Min: 0.15, Max: 3.00},
		"imp": {Min: 0.05, Max: 1.00},
		"inr": {Min: 2.00, Max: 40.00},
		"iqd": {Min: 50.00, Max: 1000.00},
		"irr": {Min: 2000.00, Max: 40000.00},
		"isk": {Min: 5.00, Max: 100.00},
		"jep": {Min: 0.05, Max: 1.00},
		"jmd": {Min: 5.00, Max: 100.00},
		"jod": {Min: 0.02, Max: 0.40},
		"jpy": {Min: 5.00, Max: 100.00},
		"kes": {Min: 5.00, Max: 100.00},
		"kgs": {Min: 2.00, Max: 40.00},
		"khr": {Min: 200.00, Max: 4000.00},
		"kin": {Min: 2500.00, Max: 50000.00},
		"kmf": {Min: 20.00, Max: 400.00},
		"kpw": {Min: 50.00, Max: 1000.00},
		"krw": {Min: 50.00, Max: 1000.00},
		"kwd": {Min: 0.01, Max: 0.20},
		"kyd": {Min: 0.05, Max: 1.00},
		"kzt": {Min: 20.00, Max: 400.00},
		"lak": {Min: 500.00, Max: 10000.00},
		"lbp": {Min: 600.00, Max: 12000.00},
		"lkr": {Min: 10.00, Max: 200.00},
		"lrd": {Min: 5.00, Max: 100.00},
		"lsl": {Min: 0.50, Max: 10.00},
		"ltl": {Min: 0.20, Max: 4.00},
		"lvl": {Min: 0.03, Max: 0.60},
		"lyd": {Min: 0.25, Max: 5.00},
		"mad": {Min: 0.50, Max: 10.00},
		"mdl": {Min: 0.50, Max: 10.00},
		"mga": {Min: 200.00, Max: 4000.00},
		"mkd": {Min: 3.00, Max: 60.00},
		"mmk": {Min: 100.00, Max: 2000.00},
		"mnt": {Min: 150.00, Max: 3000.00},
		"mop": {Min: 0.40, Max: 8.00},
		"mru": {Min: 2.00, Max: 40.00},
		"mur": {Min: 2.00, Max: 40.00},
		"mvr": {Min: 0.50, Max: 10.00},
		"mwk": {Min: 50.00, Max: 1000.00},
		"mxn": {Min: 0.50, Max: 10.00},
		"myr": {Min: 0.10, Max: 2.00},
		"mzn": {Min: 3.00, Max: 60.00},
		"nad": {Min: 0.50, Max: 10.00},
		"ngn": {Min: 80.00, Max: 1_600.00},
		"nio": {Min: 1.00, Max: 20.00},
		"nok": {Min: 0.50, Max: 10.00},
		"npr": {Min: 5.00, Max: 100.00},
		"nzd": {Min: 0.10, Max: 2.00},
		"omr": {Min: 0.01, Max: 0.20},
		"pab": {Min: 0.05, Max: 1.00},
		"pen": {Min: 0.10, Max: 2.00},
		"pgk": {Min: 0.20, Max: 4.00},
		"php": {Min: 2.00, Max: 40.00},
		"pkr": {Min: 10.00, Max: 200.00},
		"pln": {Min: 0.15, Max: 3.00},
		"pyg": {Min: 300.00, Max: 6000.00},
		"qar": {Min: 0.10, Max: 2.00},
		"ron": {Min: 0.10, Max: 2.00},
		"rsd": {Min: 5.00, Max: 100.00},
		"rub": {Min: 3.00, Max: 60.00},
		"rwf": {Min: 50.00, Max: 1000.00},
		"sar": {Min: 0.15, Max: 3.00},
		"sbd": {Min: 0.40, Max: 8.00},
		"scr": {Min: 0.60, Max: 12.00},
		"sdg": {Min: 30.00, Max: 600.00},
		"sek": {Min: 0.50, Max: 10.00},
		"sgd": {Min: 0.05, Max: 1.00},
		"shp": {Min: 0.05, Max: 1.00},
		"sll": {Min: 1000.00, Max: 20000.00},
		"sos": {Min: 30.00, Max: 600.00},
		"srd": {Min: 1.00, Max: 20.00},
		"std": {Min: 1000.00, Max: 20000.00},
		"syp": {Min: 100.00, Max: 2000.00},
		"szl": {Min: 0.50, Max: 10.00},
		"thb": {Min: 1.00, Max: 20.00},
		"tjs": {Min: 0.50, Max: 10.00},
		"tmt": {Min: 0.15, Max: 3.00},
		"tnd": {Min: 0.15, Max: 3.00},
		"top": {Min: 0.10, Max: 2.00},
		"try": {Min: 0.50, Max: 10.00},
		"ttd": {Min: 0.30, Max: 6.00},
		"twd": {Min: 1.00, Max: 20.00},
		"tzs": {Min: 100.00, Max: 2000.00},
		"uah": {Min: 2.00, Max: 40.00},
		"ugx": {Min: 200.00, Max: 4000.00},
		"usd": {Min: 0.05, Max: 1.00},
		"uyu": {Min: 2.00, Max: 40.00},
		"uzs": {Min: 500.00, Max: 10000.00},
		"vnd": {Min: 1000.00, Max: 20000.00},
		"vuv": {Min: 5.00, Max: 100.00},
		"wst": {Min: 0.10, Max: 2.00},
		"xaf": {Min: 30.00, Max: 600.00},
		"xcd": {Min: 0.10, Max: 2.00},
		"xdr": {Min: 0.03, Max: 0.60},
		"xof": {Min: 30.00, Max: 600.00},
		"xpf": {Min: 5.00, Max: 100.00},
		"yer": {Min: 10.00, Max: 200.00},
		"zar": {Min: 0.50, Max: 10.00},
		"zmk": {Min: 500.00, Max: 10000.00},
		"zmw": {Min: 0.50, Max: 10.00},
		"zwl": {Min: 10.00, Max: 200.00},
	}

	MaxDailyDepositUsdAmount = 1.5 * SendLimits[currency_lib.USD].Daily
	MaxPerDepositUsdAmount   = 1.2 * SendLimits[currency_lib.USD].PerTransaction
)
