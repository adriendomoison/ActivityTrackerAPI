package language

// This file is generated by i18n/language/codegen/generate.sh

func init() {
	
	registerPluralSpec([]string{"bm", "bo", "dz", "id", "ig", "ii", "in", "ja", "jbo", "jv", "jw", "kde", "kea", "km", "ko", "lkt", "lo", "ms", "my", "nqo", "root", "sah", "ses", "sg", "th", "to", "vi", "wo", "yo", "zh"}, &PluralSpec{
		Plurals: newPluralSet(Other),
		PluralFunc: func(ops *operands) Plural {
			return Other
		},
	})
	registerPluralSpec([]string{"am", "as", "bn", "fa", "gu", "hi", "kn", "mr", "zu"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 0 or n = 1
			if intEqualsAny(ops.I, 0) ||
				ops.NequalsAny(1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ff", "fr", "hy", "kab"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 0,1
			if intEqualsAny(ops.I, 0, 1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ast", "ca", "de", "en", "et", "fi", "fy", "gl", "it", "ji", "nl", "sv", "sw", "ur", "yi"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"si"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0,1 or i = 0 and f = 1
			if ops.NequalsAny(0, 1) ||
				intEqualsAny(ops.I, 0) && intEqualsAny(ops.F, 1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ak", "bh", "guw", "ln", "mg", "nso", "pa", "ti", "wa"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0..1
			if ops.NinRange(0, 1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"tzm"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0..1 or n = 11..99
			if ops.NinRange(0, 1) ||
				ops.NinRange(11, 99) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"pt"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0..2 and n != 2
			if ops.NinRange(0, 2) && !ops.NequalsAny(2) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"af", "asa", "az", "bem", "bez", "bg", "brx", "ce", "cgg", "chr", "ckb", "dv", "ee", "el", "eo", "es", "eu", "fo", "fur", "gsw", "ha", "haw", "hu", "jgo", "jmc", "ka", "kaj", "kcg", "kk", "kkj", "kl", "ks", "ksb", "ku", "ky", "lb", "lg", "mas", "mgo", "ml", "mn", "nah", "nb", "nd", "ne", "nn", "nnh", "no", "nr", "ny", "nyn", "om", "or", "os", "pap", "ps", "rm", "rof", "rwk", "saq", "sdh", "seh", "sn", "so", "sq", "ss", "ssy", "st", "syr", "ta", "te", "teo", "tig", "tk", "tn", "tr", "ts", "ug", "uz", "ve", "vo", "vun", "wae", "xh", "xog"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"pt_PT"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1 and v = 0
			if ops.NequalsAny(1) && intEqualsAny(ops.V, 0) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"da"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1 or t != 0 and i = 0,1
			if ops.NequalsAny(1) ||
				!intEqualsAny(ops.T, 0) && intEqualsAny(ops.I, 0, 1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"is"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// t = 0 and i % 10 = 1 and i % 100 != 11 or t != 0
			if intEqualsAny(ops.T, 0) && intEqualsAny(ops.I % 10, 1) && !intEqualsAny(ops.I % 100, 11) ||
				!intEqualsAny(ops.T, 0) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"mk"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i % 10 = 1 or f % 10 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 10, 1) ||
				intEqualsAny(ops.F % 10, 1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"fil", "tl"}, &PluralSpec{
		Plurals: newPluralSet(One, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i = 1,2,3 or v = 0 and i % 10 != 4,6,9 or v != 0 and f % 10 != 4,6,9
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I, 1, 2, 3) ||
				intEqualsAny(ops.V, 0) && !intEqualsAny(ops.I % 10, 4, 6, 9) ||
				!intEqualsAny(ops.V, 0) && !intEqualsAny(ops.F % 10, 4, 6, 9) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"lv", "prg"}, &PluralSpec{
		Plurals: newPluralSet(Zero, One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n % 10 = 0 or n % 100 = 11..19 or v = 2 and f % 100 = 11..19
			if ops.NmodEqualsAny(10, 0) ||
				ops.NmodInRange(100, 11, 19) ||
				intEqualsAny(ops.V, 2) && intInRange(ops.F % 100, 11, 19) {
				return Zero
			}
			// n % 10 = 1 and n % 100 != 11 or v = 2 and f % 10 = 1 and f % 100 != 11 or v != 2 and f % 10 = 1
			if ops.NmodEqualsAny(10, 1) && !ops.NmodEqualsAny(100, 11) ||
				intEqualsAny(ops.V, 2) && intEqualsAny(ops.F % 10, 1) && !intEqualsAny(ops.F % 100, 11) ||
				!intEqualsAny(ops.V, 2) && intEqualsAny(ops.F % 10, 1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"lag"}, &PluralSpec{
		Plurals: newPluralSet(Zero, One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0
			if ops.NequalsAny(0) {
				return Zero
			}
			// i = 0,1 and n != 0
			if intEqualsAny(ops.I, 0, 1) && !ops.NequalsAny(0) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ksh"}, &PluralSpec{
		Plurals: newPluralSet(Zero, One, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0
			if ops.NequalsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			return Other
		},
	})
	registerPluralSpec([]string{"iu", "kw", "naq", "se", "sma", "smi", "smj", "smn", "sms"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			// n = 2
			if ops.NequalsAny(2) {
				return Two
			}
			return Other
		},
	})
	registerPluralSpec([]string{"shi"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 0 or n = 1
			if intEqualsAny(ops.I, 0) ||
				ops.NequalsAny(1) {
				return One
			}
			// n = 2..10
			if ops.NinRange(2, 10) {
				return Few
			}
			return Other
		},
	})
	registerPluralSpec([]string{"mo", "ro"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// v != 0 or n = 0 or n != 1 and n % 100 = 1..19
			if !intEqualsAny(ops.V, 0) ||
				ops.NequalsAny(0) ||
				!ops.NequalsAny(1) && ops.NmodInRange(100, 1, 19) {
				return Few
			}
			return Other
		},
	})
	registerPluralSpec([]string{"bs", "hr", "sh", "sr"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i % 10 = 1 and i % 100 != 11 or f % 10 = 1 and f % 100 != 11
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 10, 1) && !intEqualsAny(ops.I % 100, 11) ||
				intEqualsAny(ops.F % 10, 1) && !intEqualsAny(ops.F % 100, 11) {
				return One
			}
			// v = 0 and i % 10 = 2..4 and i % 100 != 12..14 or f % 10 = 2..4 and f % 100 != 12..14
			if intEqualsAny(ops.V, 0) && intInRange(ops.I % 10, 2, 4) && !intInRange(ops.I % 100, 12, 14) ||
				intInRange(ops.F % 10, 2, 4) && !intInRange(ops.F % 100, 12, 14) {
				return Few
			}
			return Other
		},
	})
	registerPluralSpec([]string{"gd"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Few, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1,11
			if ops.NequalsAny(1, 11) {
				return One
			}
			// n = 2,12
			if ops.NequalsAny(2, 12) {
				return Two
			}
			// n = 3..10,13..19
			if ops.NinRange(3, 10) || ops.NinRange(13, 19) {
				return Few
			}
			return Other
		},
	})
	registerPluralSpec([]string{"sl"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Few, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i % 100 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 100, 1) {
				return One
			}
			// v = 0 and i % 100 = 2
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 100, 2) {
				return Two
			}
			// v = 0 and i % 100 = 3..4 or v != 0
			if intEqualsAny(ops.V, 0) && intInRange(ops.I % 100, 3, 4) ||
				!intEqualsAny(ops.V, 0) {
				return Few
			}
			return Other
		},
	})
	registerPluralSpec([]string{"dsb", "hsb"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Few, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i % 100 = 1 or f % 100 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 100, 1) ||
				intEqualsAny(ops.F % 100, 1) {
				return One
			}
			// v = 0 and i % 100 = 2 or f % 100 = 2
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 100, 2) ||
				intEqualsAny(ops.F % 100, 2) {
				return Two
			}
			// v = 0 and i % 100 = 3..4 or f % 100 = 3..4
			if intEqualsAny(ops.V, 0) && intInRange(ops.I % 100, 3, 4) ||
				intInRange(ops.F % 100, 3, 4) {
				return Few
			}
			return Other
		},
	})
	registerPluralSpec([]string{"he", "iw"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// i = 2 and v = 0
			if intEqualsAny(ops.I, 2) && intEqualsAny(ops.V, 0) {
				return Two
			}
			// v = 0 and n != 0..10 and n % 10 = 0
			if intEqualsAny(ops.V, 0) && !ops.NinRange(0, 10) && ops.NmodEqualsAny(10, 0) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"cs", "sk"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// i = 2..4 and v = 0
			if intInRange(ops.I, 2, 4) && intEqualsAny(ops.V, 0) {
				return Few
			}
			// v != 0
			if !intEqualsAny(ops.V, 0) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"pl"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// v = 0 and i % 10 = 2..4 and i % 100 != 12..14
			if intEqualsAny(ops.V, 0) && intInRange(ops.I % 10, 2, 4) && !intInRange(ops.I % 100, 12, 14) {
				return Few
			}
			// v = 0 and i != 1 and i % 10 = 0..1 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 12..14
			if intEqualsAny(ops.V, 0) && !intEqualsAny(ops.I, 1) && intInRange(ops.I % 10, 0, 1) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I % 10, 5, 9) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I % 100, 12, 14) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"be"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n % 10 = 1 and n % 100 != 11
			if ops.NmodEqualsAny(10, 1) && !ops.NmodEqualsAny(100, 11) {
				return One
			}
			// n % 10 = 2..4 and n % 100 != 12..14
			if ops.NmodInRange(10, 2, 4) && !ops.NmodInRange(100, 12, 14) {
				return Few
			}
			// n % 10 = 0 or n % 10 = 5..9 or n % 100 = 11..14
			if ops.NmodEqualsAny(10, 0) ||
				ops.NmodInRange(10, 5, 9) ||
				ops.NmodInRange(100, 11, 14) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"lt"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n % 10 = 1 and n % 100 != 11..19
			if ops.NmodEqualsAny(10, 1) && !ops.NmodInRange(100, 11, 19) {
				return One
			}
			// n % 10 = 2..9 and n % 100 != 11..19
			if ops.NmodInRange(10, 2, 9) && !ops.NmodInRange(100, 11, 19) {
				return Few
			}
			// f != 0
			if !intEqualsAny(ops.F, 0) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"mt"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			// n = 0 or n % 100 = 2..10
			if ops.NequalsAny(0) ||
				ops.NmodInRange(100, 2, 10) {
				return Few
			}
			// n % 100 = 11..19
			if ops.NmodInRange(100, 11, 19) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ru", "uk"}, &PluralSpec{
		Plurals: newPluralSet(One, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i % 10 = 1 and i % 100 != 11
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 10, 1) && !intEqualsAny(ops.I % 100, 11) {
				return One
			}
			// v = 0 and i % 10 = 2..4 and i % 100 != 12..14
			if intEqualsAny(ops.V, 0) && intInRange(ops.I % 10, 2, 4) && !intInRange(ops.I % 100, 12, 14) {
				return Few
			}
			// v = 0 and i % 10 = 0 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 11..14
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 10, 0) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I % 10, 5, 9) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I % 100, 11, 14) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"br"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n % 10 = 1 and n % 100 != 11,71,91
			if ops.NmodEqualsAny(10, 1) && !ops.NmodEqualsAny(100, 11, 71, 91) {
				return One
			}
			// n % 10 = 2 and n % 100 != 12,72,92
			if ops.NmodEqualsAny(10, 2) && !ops.NmodEqualsAny(100, 12, 72, 92) {
				return Two
			}
			// n % 10 = 3..4,9 and n % 100 != 10..19,70..79,90..99
			if (ops.NmodInRange(10, 3, 4) || ops.NmodEqualsAny(10, 9)) && !(ops.NmodInRange(100, 10, 19) || ops.NmodInRange(100, 70, 79) || ops.NmodInRange(100, 90, 99)) {
				return Few
			}
			// n != 0 and n % 1000000 = 0
			if !ops.NequalsAny(0) && ops.NmodEqualsAny(1000000, 0) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ga"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			// n = 2
			if ops.NequalsAny(2) {
				return Two
			}
			// n = 3..6
			if ops.NinRange(3, 6) {
				return Few
			}
			// n = 7..10
			if ops.NinRange(7, 10) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"gv"}, &PluralSpec{
		Plurals: newPluralSet(One, Two, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// v = 0 and i % 10 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 10, 1) {
				return One
			}
			// v = 0 and i % 10 = 2
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 10, 2) {
				return Two
			}
			// v = 0 and i % 100 = 0,20,40,60,80
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I % 100, 0, 20, 40, 60, 80) {
				return Few
			}
			// v != 0
			if !intEqualsAny(ops.V, 0) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"ar"}, &PluralSpec{
		Plurals: newPluralSet(Zero, One, Two, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0
			if ops.NequalsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			// n = 2
			if ops.NequalsAny(2) {
				return Two
			}
			// n % 100 = 3..10
			if ops.NmodInRange(100, 3, 10) {
				return Few
			}
			// n % 100 = 11..99
			if ops.NmodInRange(100, 11, 99) {
				return Many
			}
			return Other
		},
	})
	registerPluralSpec([]string{"cy"}, &PluralSpec{
		Plurals: newPluralSet(Zero, One, Two, Few, Many, Other),
		PluralFunc: func(ops *operands) Plural {
			// n = 0
			if ops.NequalsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NequalsAny(1) {
				return One
			}
			// n = 2
			if ops.NequalsAny(2) {
				return Two
			}
			// n = 3
			if ops.NequalsAny(3) {
				return Few
			}
			// n = 6
			if ops.NequalsAny(6) {
				return Many
			}
			return Other
		},
	})
}
