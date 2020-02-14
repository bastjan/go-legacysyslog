//line machine.go.rl:1
package legacysyslog

import (
	"fmt"
	"time"

	"github.com/influxdata/go-syslog/v3"
	"github.com/influxdata/go-syslog/v3/common"
)

var _ = fmt.Print

//line machine.go.rl:221

//line machine.go:21
const start int = 1
const first_final int = 51

const en_main int = 1

//line machine.go.rl:224

type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	// general mark
	pb int
	// hostname machine can run in parallel to the tag machine since it is not clear from the beginning
	// if it is a hostname or tag
	pHostname int
	// same for the timestamp
	pTimestamp int

	err        error
	bestEffort bool
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine() *machine {
	m := &machine{}

//line machine.go.rl:246

//line machine.go.rl:247

//line machine.go.rl:248

//line machine.go.rl:249

//line machine.go.rl:250

	return m
}

// Err returns the error that occurred on the last call to Parse.
// If the result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

// Parse parses the input byte array as a RFC3164 syslog message.
func (m *machine) Parse(input []byte) (syslog.Message, error) {
	if len(input) == 0 {
		return (&syslogMessage{}).export(), nil
	}

	m.data = input
	m.p = 0
	m.pb = 0
	m.pHostname = 0
	m.pTimestamp = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	output := &syslogMessage{}

//line machine.go:92
	{
		m.cs = start
	}

//line machine.go.rl:282

//line machine.go:99
	{
		if (m.p) == (m.pe) {
			goto _test_eof
		}
		switch m.cs {
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 0:
			goto st_case_0
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 181:
			goto st_case_181
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 221:
			goto st_case_221
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
		case 225:
			goto st_case_225
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 237:
			goto st_case_237
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 249:
			goto st_case_249
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 258:
			goto st_case_258
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		case 280:
			goto st_case_280
		case 281:
			goto st_case_281
		case 282:
			goto st_case_282
		case 283:
			goto st_case_283
		case 284:
			goto st_case_284
		case 285:
			goto st_case_285
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 288:
			goto st_case_288
		case 289:
			goto st_case_289
		case 290:
			goto st_case_290
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
		case 293:
			goto st_case_293
		case 294:
			goto st_case_294
		case 295:
			goto st_case_295
		case 296:
			goto st_case_296
		case 297:
			goto st_case_297
		case 298:
			goto st_case_298
		case 299:
			goto st_case_299
		case 300:
			goto st_case_300
		case 301:
			goto st_case_301
		case 302:
			goto st_case_302
		case 303:
			goto st_case_303
		case 304:
			goto st_case_304
		case 305:
			goto st_case_305
		case 306:
			goto st_case_306
		case 307:
			goto st_case_307
		case 308:
			goto st_case_308
		case 309:
			goto st_case_309
		case 310:
			goto st_case_310
		case 311:
			goto st_case_311
		case 312:
			goto st_case_312
		case 313:
			goto st_case_313
		case 314:
			goto st_case_314
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 318:
			goto st_case_318
		case 319:
			goto st_case_319
		case 320:
			goto st_case_320
		case 321:
			goto st_case_321
		case 322:
			goto st_case_322
		case 323:
			goto st_case_323
		case 324:
			goto st_case_324
		case 325:
			goto st_case_325
		case 326:
			goto st_case_326
		case 327:
			goto st_case_327
		case 328:
			goto st_case_328
		case 329:
			goto st_case_329
		case 330:
			goto st_case_330
		case 331:
			goto st_case_331
		case 332:
			goto st_case_332
		case 333:
			goto st_case_333
		case 334:
			goto st_case_334
		case 335:
			goto st_case_335
		case 336:
			goto st_case_336
		case 337:
			goto st_case_337
		case 338:
			goto st_case_338
		case 339:
			goto st_case_339
		case 340:
			goto st_case_340
		case 341:
			goto st_case_341
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 346:
			goto st_case_346
		case 347:
			goto st_case_347
		case 348:
			goto st_case_348
		case 349:
			goto st_case_349
		case 350:
			goto st_case_350
		case 351:
			goto st_case_351
		case 352:
			goto st_case_352
		case 353:
			goto st_case_353
		case 354:
			goto st_case_354
		case 355:
			goto st_case_355
		case 356:
			goto st_case_356
		case 357:
			goto st_case_357
		case 358:
			goto st_case_358
		case 359:
			goto st_case_359
		case 360:
			goto st_case_360
		case 361:
			goto st_case_361
		case 362:
			goto st_case_362
		case 363:
			goto st_case_363
		case 364:
			goto st_case_364
		case 365:
			goto st_case_365
		case 366:
			goto st_case_366
		case 367:
			goto st_case_367
		case 368:
			goto st_case_368
		case 369:
			goto st_case_369
		case 370:
			goto st_case_370
		case 371:
			goto st_case_371
		case 372:
			goto st_case_372
		case 373:
			goto st_case_373
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 379:
			goto st_case_379
		case 380:
			goto st_case_380
		case 381:
			goto st_case_381
		case 382:
			goto st_case_382
		case 383:
			goto st_case_383
		case 384:
			goto st_case_384
		case 385:
			goto st_case_385
		case 386:
			goto st_case_386
		case 387:
			goto st_case_387
		case 388:
			goto st_case_388
		case 389:
			goto st_case_389
		case 390:
			goto st_case_390
		case 391:
			goto st_case_391
		case 392:
			goto st_case_392
		case 393:
			goto st_case_393
		case 394:
			goto st_case_394
		case 395:
			goto st_case_395
		case 396:
			goto st_case_396
		case 397:
			goto st_case_397
		case 398:
			goto st_case_398
		case 399:
			goto st_case_399
		case 400:
			goto st_case_400
		case 401:
			goto st_case_401
		case 402:
			goto st_case_402
		case 403:
			goto st_case_403
		case 404:
			goto st_case_404
		case 405:
			goto st_case_405
		case 406:
			goto st_case_406
		case 407:
			goto st_case_407
		case 408:
			goto st_case_408
		case 409:
			goto st_case_409
		case 410:
			goto st_case_410
		case 411:
			goto st_case_411
		case 412:
			goto st_case_412
		case 413:
			goto st_case_413
		case 414:
			goto st_case_414
		case 415:
			goto st_case_415
		case 416:
			goto st_case_416
		case 417:
			goto st_case_417
		case 418:
			goto st_case_418
		case 419:
			goto st_case_419
		case 420:
			goto st_case_420
		case 421:
			goto st_case_421
		case 422:
			goto st_case_422
		case 423:
			goto st_case_423
		case 424:
			goto st_case_424
		case 425:
			goto st_case_425
		case 426:
			goto st_case_426
		case 427:
			goto st_case_427
		case 428:
			goto st_case_428
		case 429:
			goto st_case_429
		case 430:
			goto st_case_430
		case 431:
			goto st_case_431
		case 432:
			goto st_case_432
		case 433:
			goto st_case_433
		case 434:
			goto st_case_434
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 443:
			goto st_case_443
		case 444:
			goto st_case_444
		case 445:
			goto st_case_445
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 446:
			goto st_case_446
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		}
		goto st_out
	st_case_1:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr1
		case 42:
			goto tr2
		case 46:
			goto tr2
		case 58:
			goto tr4
		case 60:
			goto tr5
		case 65:
			goto tr6
		case 68:
			goto tr7
		case 70:
			goto tr8
		case 74:
			goto tr9
		case 77:
			goto tr10
		case 78:
			goto tr11
		case 79:
			goto tr12
		case 83:
			goto tr13
		case 91:
			goto tr14
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr3
		}
		goto tr0
	tr0:
//line machine.go.rl:20

		m.pb = m.p

		goto st2
	tr26:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof2
		}
	st_case_2:
//line machine.go:1062
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	tr16:
//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st51
	tr82:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st51
	tr19:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:106

		output.contentSet = true
		output.content = string(m.text())

		goto st51
	tr22:
//line machine.go.rl:106

		output.contentSet = true
		output.content = string(m.text())

		goto st51
	tr27:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st51
	st51:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof51
		}
	st_case_51:
//line machine.go:1133
		if (m.data)[(m.p)] == 32 {
			goto tr77
		}
		goto tr76
	tr76:
//line machine.go.rl:20

		m.pb = m.p

		goto st52
	st52:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof52
		}
	st_case_52:
//line machine.go:1149
		goto st52
	tr77:
//line machine.go.rl:20

		m.pb = m.p

		goto st53
	st53:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof53
		}
	st_case_53:
//line machine.go:1162
		goto tr76
	tr17:
//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st3
	tr14:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st3
	tr37:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st3
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
//line machine.go:1207
		switch (m.data)[(m.p)] {
		case 32:
			goto tr19
		case 58:
			goto tr19
		case 93:
			goto tr20
		}
		goto tr18
	tr18:
//line machine.go.rl:20

		m.pb = m.p

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:1228
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 58:
			goto tr22
		case 93:
			goto tr23
		}
		goto st4
	tr20:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:106

		output.contentSet = true
		output.content = string(m.text())

		goto st5
	tr23:
//line machine.go.rl:106

		output.contentSet = true
		output.content = string(m.text())

		goto st5
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof5
		}
	st_case_5:
//line machine.go:1261
		switch (m.data)[(m.p)] {
		case 32:
			goto st51
		case 58:
			goto st51
		}
		goto st0
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	tr498:
//line machine.go.rl:20

		m.pb = m.p

		goto st54
	tr1:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st54
	st54:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof54
		}
	st_case_54:
//line machine.go:1295
		switch (m.data)[(m.p)] {
		case 32:
			goto tr1
		case 42:
			goto tr80
		case 46:
			goto tr80
		case 58:
			goto tr82
		case 65:
			goto tr83
		case 68:
			goto tr84
		case 70:
			goto tr85
		case 74:
			goto tr86
		case 77:
			goto tr87
		case 78:
			goto tr88
		case 79:
			goto tr89
		case 83:
			goto tr90
		case 91:
			goto tr91
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr81
		}
		goto tr79
	tr79:
//line machine.go.rl:20

		m.pb = m.p

		goto st55
	tr98:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st55
	st55:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof55
		}
	st_case_55:
//line machine.go:1354
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	tr93:
//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr91:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr108:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr142:
//line machine.go.rl:51

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr483:
//line machine.go.rl:63

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	st56:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof56
		}
	st_case_56:
//line machine.go:1453
		switch (m.data)[(m.p)] {
		case 32:
			goto tr19
		case 58:
			goto tr19
		case 93:
			goto tr95
		}
		goto tr94
	tr94:
//line machine.go.rl:20

		m.pb = m.p

		goto st57
	st57:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof57
		}
	st_case_57:
//line machine.go:1474
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 58:
			goto tr22
		case 93:
			goto tr97
		}
		goto st57
	tr95:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:106

		output.contentSet = true
		output.content = string(m.text())

		goto st58
	tr97:
//line machine.go.rl:106

		output.contentSet = true
		output.content = string(m.text())

		goto st58
	st58:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof58
		}
	st_case_58:
//line machine.go:1507
		switch (m.data)[(m.p)] {
		case 32:
			goto st51
		case 58:
			goto st51
		}
		goto st52
	tr80:
//line machine.go.rl:20

		m.pb = m.p

		goto st59
	st59:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof59
		}
	st_case_59:
//line machine.go:1526
		switch (m.data)[(m.p)] {
		case 32:
			goto tr27
		case 58:
			goto tr27
		case 65:
			goto tr100
		case 68:
			goto tr101
		case 70:
			goto tr102
		case 74:
			goto tr103
		case 77:
			goto tr104
		case 78:
			goto tr105
		case 79:
			goto tr106
		case 83:
			goto tr107
		case 91:
			goto tr108
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr99
		}
		goto tr98
	tr81:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st60
	tr99:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st60
	st60:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof60
		}
	st_case_60:
//line machine.go:1589
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st61
		}
		goto st55
	st61:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st62
		}
		goto st55
	st62:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st63
		}
		goto st55
	st63:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st64
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	st64:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st65
		case 49:
			goto st393
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	st65:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st66
		}
		goto st55
	st66:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st67
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	st67:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st68
		case 51:
			goto st392
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st391
		}
		goto st55
	st68:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st69
		}
		goto st55
	st69:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st70
		case 91:
			goto tr93
		}
		goto st55
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 50:
			goto st390
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st71
		}
		goto st55
	st71:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st72
		}
		goto st55
	st72:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr54
		case 91:
			goto tr93
		}
		goto st55
	tr54:
//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st73
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof73
		}
	st_case_73:
//line machine.go:1819
		if (m.data)[(m.p)] == 32 {
			goto tr77
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto tr125
		}
		goto tr76
	tr125:
//line machine.go.rl:20

		m.pb = m.p

		goto st74
	st74:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof74
		}
	st_case_74:
//line machine.go:1838
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st75
		}
		goto st52
	st75:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof75
		}
	st_case_75:
		if (m.data)[(m.p)] == 58 {
			goto st76
		}
		goto st52
	st76:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st77
		}
		goto st52
	st77:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof77
		}
	st_case_77:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st78
		}
		goto st52
	st78:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 46:
			goto st383
		case 90:
			goto st84
		}
		goto st52
	st79:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof79
		}
	st_case_79:
		if (m.data)[(m.p)] == 50 {
			goto st382
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st80
		}
		goto st52
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof80
		}
	st_case_80:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st81
		}
		goto st52
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof81
		}
	st_case_81:
		if (m.data)[(m.p)] == 58 {
			goto st82
		}
		goto st52
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof82
		}
	st_case_82:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st83
		}
		goto st52
	st83:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof83
		}
	st_case_83:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st84
		}
		goto st52
	st84:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr139
		case 58:
			goto tr141
		case 91:
			goto tr142
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr140
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr140
			}
		default:
			goto tr140
		}
		goto tr138
	tr138:
//line machine.go.rl:51

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st85
	tr399:
//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st85
	tr479:
//line machine.go.rl:63

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st85
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof85
		}
	st_case_85:
//line machine.go:2019
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st86
	st86:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof86
		}
	st_case_86:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st87
	st87:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st88
	st88:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st89
	st89:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st90
	st90:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st91
	st91:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st92
	st92:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st93
	st93:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st94
	st94:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st95
	st95:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st96
	st96:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st97
	st97:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st98
	st98:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st99
	st99:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st100
	st100:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st101
	st101:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof101
		}
	st_case_101:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st102
	st102:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st103
	st103:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st104
	st104:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st105
	st105:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st106
	st106:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st107
	st107:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st108
	st108:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof108
		}
	st_case_108:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st109
	st109:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st110
	st110:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof110
		}
	st_case_110:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st111
	st111:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st112
	st112:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof112
		}
	st_case_112:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st113
	st113:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof113
		}
	st_case_113:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st114
	st114:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st115
	st115:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st116
	st116:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st117
	st117:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof117
		}
	st_case_117:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st118
	st118:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof118
		}
	st_case_118:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st119
	st119:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof119
		}
	st_case_119:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st120
	st120:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof120
		}
	st_case_120:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st121
	st121:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof121
		}
	st_case_121:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st122
	st122:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof122
		}
	st_case_122:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st123
	st123:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof123
		}
	st_case_123:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st124
	st124:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st125
	st125:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof125
		}
	st_case_125:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st126
	st126:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st127
	st127:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof127
		}
	st_case_127:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st128
	st128:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof128
		}
	st_case_128:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st129
	st129:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof129
		}
	st_case_129:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st130
	st130:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st131
	st131:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st132
	st132:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st133
	st133:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof133
		}
	st_case_133:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st134
	st134:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st135
	st135:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st136
	st136:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st137
	st137:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st138
	st138:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof138
		}
	st_case_138:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st139
	st139:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof139
		}
	st_case_139:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st140
	st140:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof140
		}
	st_case_140:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st141
	st141:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof141
		}
	st_case_141:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st142
	st142:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof142
		}
	st_case_142:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st143
	st143:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof143
		}
	st_case_143:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st144
	st144:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st145
	st145:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof145
		}
	st_case_145:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st146
	st146:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof146
		}
	st_case_146:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st147
	st147:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st148
	st148:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof148
		}
	st_case_148:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st149
	st149:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof149
		}
	st_case_149:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st150
	st150:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof150
		}
	st_case_150:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st151
	st151:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof151
		}
	st_case_151:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st152
	st152:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof152
		}
	st_case_152:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st153
	st153:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof153
		}
	st_case_153:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st154
	st154:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof154
		}
	st_case_154:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st155
	st155:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof155
		}
	st_case_155:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st156
	st156:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof156
		}
	st_case_156:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st157
	st157:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof157
		}
	st_case_157:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st158
	st158:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof158
		}
	st_case_158:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st159
	st159:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st160
	st160:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof160
		}
	st_case_160:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st161
	st161:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st162
	st162:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof162
		}
	st_case_162:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st163
	st163:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof163
		}
	st_case_163:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st164
	st164:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof164
		}
	st_case_164:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st165
	st165:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof165
		}
	st_case_165:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st166
	st166:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof166
		}
	st_case_166:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st167
	st167:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof167
		}
	st_case_167:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st168
	st168:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof168
		}
	st_case_168:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st169
	st169:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st170
	st170:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st171
	st171:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st172
	st172:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st173
	st173:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof173
		}
	st_case_173:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st174
	st174:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof174
		}
	st_case_174:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st175
	st175:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof175
		}
	st_case_175:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st176
	st176:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st177
	st177:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof177
		}
	st_case_177:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st178
	st178:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof178
		}
	st_case_178:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st179
	st179:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof179
		}
	st_case_179:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st180
	st180:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof180
		}
	st_case_180:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st181
	st181:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof181
		}
	st_case_181:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st182
	st182:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof182
		}
	st_case_182:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st183
	st183:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st184
	st184:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof184
		}
	st_case_184:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st185
	st185:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st186
	st186:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof186
		}
	st_case_186:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st187
	st187:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof187
		}
	st_case_187:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st188
	st188:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof188
		}
	st_case_188:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st189
	st189:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof189
		}
	st_case_189:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st190
	st190:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof190
		}
	st_case_190:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st191
	st191:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof191
		}
	st_case_191:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st192
	st192:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof192
		}
	st_case_192:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st193
	st193:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof193
		}
	st_case_193:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st194
	st194:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof194
		}
	st_case_194:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st195
	st195:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof195
		}
	st_case_195:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st196
	st196:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof196
		}
	st_case_196:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st197
	st197:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof197
		}
	st_case_197:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st198
	st198:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof198
		}
	st_case_198:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st199
	st199:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof199
		}
	st_case_199:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st200
	st200:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof200
		}
	st_case_200:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st201
	st201:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof201
		}
	st_case_201:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st202
	st202:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof202
		}
	st_case_202:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st203
	st203:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof203
		}
	st_case_203:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st204
	st204:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof204
		}
	st_case_204:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st205
	st205:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof205
		}
	st_case_205:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st206
	st206:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof206
		}
	st_case_206:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st207
	st207:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof207
		}
	st_case_207:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st208
	st208:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof208
		}
	st_case_208:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st209
	st209:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof209
		}
	st_case_209:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st210
	st210:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof210
		}
	st_case_210:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st211
	st211:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof211
		}
	st_case_211:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st212
	st212:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof212
		}
	st_case_212:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st213
	st213:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof213
		}
	st_case_213:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st214
	st214:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof214
		}
	st_case_214:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st215
	st215:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof215
		}
	st_case_215:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st216
	st216:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof216
		}
	st_case_216:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st217
	st217:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof217
		}
	st_case_217:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st218
	st218:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof218
		}
	st_case_218:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st219
	st219:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof219
		}
	st_case_219:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st220
	st220:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof220
		}
	st_case_220:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st221
	st221:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof221
		}
	st_case_221:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st222
	st222:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof222
		}
	st_case_222:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st223
	st223:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof223
		}
	st_case_223:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st224
	st224:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof224
		}
	st_case_224:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st225
	st225:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof225
		}
	st_case_225:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st226
	st226:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof226
		}
	st_case_226:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st227
	st227:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof227
		}
	st_case_227:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st228
	st228:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof228
		}
	st_case_228:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st229
	st229:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof229
		}
	st_case_229:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st230
	st230:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof230
		}
	st_case_230:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st231
	st231:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof231
		}
	st_case_231:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st232
	st232:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof232
		}
	st_case_232:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st233
	st233:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof233
		}
	st_case_233:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st234
	st234:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof234
		}
	st_case_234:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st235
	st235:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof235
		}
	st_case_235:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st236
	st236:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof236
		}
	st_case_236:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st237
	st237:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof237
		}
	st_case_237:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st238
	st238:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof238
		}
	st_case_238:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st239
	st239:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof239
		}
	st_case_239:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st240
	st240:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof240
		}
	st_case_240:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st241
	st241:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof241
		}
	st_case_241:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st242
	st242:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof242
		}
	st_case_242:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st243
	st243:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof243
		}
	st_case_243:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st244
	st244:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof244
		}
	st_case_244:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st245
	st245:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof245
		}
	st_case_245:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st246
	st246:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof246
		}
	st_case_246:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st247
	st247:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof247
		}
	st_case_247:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st248
	st248:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof248
		}
	st_case_248:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st249
	st249:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof249
		}
	st_case_249:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st250
	st250:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof250
		}
	st_case_250:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st251
	st251:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof251
		}
	st_case_251:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st252
	st252:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof252
		}
	st_case_252:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st253
	st253:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof253
		}
	st_case_253:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st254
	st254:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof254
		}
	st_case_254:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st255
	st255:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof255
		}
	st_case_255:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st256
	st256:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof256
		}
	st_case_256:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st257
	st257:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof257
		}
	st_case_257:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st258
	st258:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof258
		}
	st_case_258:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st259
	st259:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof259
		}
	st_case_259:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st260
	st260:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof260
		}
	st_case_260:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st261
	st261:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof261
		}
	st_case_261:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st262
	st262:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof262
		}
	st_case_262:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st263
	st263:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof263
		}
	st_case_263:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st264
	st264:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof264
		}
	st_case_264:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st265
	st265:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof265
		}
	st_case_265:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st266
	st266:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof266
		}
	st_case_266:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st267
	st267:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof267
		}
	st_case_267:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st268
	st268:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof268
		}
	st_case_268:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st269
	st269:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof269
		}
	st_case_269:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st270
	st270:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof270
		}
	st_case_270:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st271
	st271:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof271
		}
	st_case_271:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st272
	st272:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof272
		}
	st_case_272:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st273
	st273:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof273
		}
	st_case_273:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st274
	st274:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof274
		}
	st_case_274:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st275
	st275:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof275
		}
	st_case_275:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st276
	st276:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof276
		}
	st_case_276:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st277
	st277:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof277
		}
	st_case_277:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st278
	st278:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof278
		}
	st_case_278:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st279
	st279:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof279
		}
	st_case_279:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st280
	st280:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof280
		}
	st_case_280:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st281
	st281:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof281
		}
	st_case_281:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st282
	st282:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof282
		}
	st_case_282:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st283
	st283:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof283
		}
	st_case_283:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st284
	st284:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof284
		}
	st_case_284:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st285
	st285:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof285
		}
	st_case_285:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st286
	st286:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof286
		}
	st_case_286:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st287
	st287:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof287
		}
	st_case_287:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st288
	st288:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof288
		}
	st_case_288:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st289
	st289:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof289
		}
	st_case_289:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st290
	st290:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof290
		}
	st_case_290:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st291
	st291:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof291
		}
	st_case_291:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st292
	st292:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof292
		}
	st_case_292:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st293
	st293:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof293
		}
	st_case_293:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st294
	st294:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof294
		}
	st_case_294:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st295
	st295:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof295
		}
	st_case_295:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st296
	st296:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof296
		}
	st_case_296:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st297
	st297:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof297
		}
	st_case_297:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st298
	st298:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof298
		}
	st_case_298:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st299
	st299:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof299
		}
	st_case_299:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st300
	st300:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof300
		}
	st_case_300:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st301
	st301:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof301
		}
	st_case_301:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st302
	st302:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof302
		}
	st_case_302:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st303
	st303:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof303
		}
	st_case_303:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st304
	st304:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof304
		}
	st_case_304:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st305
	st305:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof305
		}
	st_case_305:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st306
	st306:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof306
		}
	st_case_306:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st307
	st307:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof307
		}
	st_case_307:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st308
	st308:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof308
		}
	st_case_308:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st309
	st309:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof309
		}
	st_case_309:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st310
	st310:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof310
		}
	st_case_310:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st311
	st311:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof311
		}
	st_case_311:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st312
	st312:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof312
		}
	st_case_312:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st313
	st313:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof313
		}
	st_case_313:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st314
	st314:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof314
		}
	st_case_314:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st315
	st315:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof315
		}
	st_case_315:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st316
	st316:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof316
		}
	st_case_316:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st317
	st317:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof317
		}
	st_case_317:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st318
	st318:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof318
		}
	st_case_318:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st319
	st319:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof319
		}
	st_case_319:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st320
	st320:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof320
		}
	st_case_320:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st321
	st321:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof321
		}
	st_case_321:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st322
	st322:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof322
		}
	st_case_322:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st323
	st323:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof323
		}
	st_case_323:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st324
	st324:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof324
		}
	st_case_324:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st325
	st325:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof325
		}
	st_case_325:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st326
	st326:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof326
		}
	st_case_326:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st327
	st327:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof327
		}
	st_case_327:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st328
	st328:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof328
		}
	st_case_328:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st329
	st329:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof329
		}
	st_case_329:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st330
	st330:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof330
		}
	st_case_330:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st331
	st331:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof331
		}
	st_case_331:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st332
	st332:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof332
		}
	st_case_332:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st333
	st333:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof333
		}
	st_case_333:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st334
	st334:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof334
		}
	st_case_334:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st335
	st335:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof335
		}
	st_case_335:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st336
	st336:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof336
		}
	st_case_336:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st337
	st337:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof337
		}
	st_case_337:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st338
	st338:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof338
		}
	st_case_338:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st339
	st339:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof339
		}
	st_case_339:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	tr398:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st340
	tr144:
//line machine.go.rl:96

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st340
	st340:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof340
		}
	st_case_340:
//line machine.go:5613
		switch (m.data)[(m.p)] {
		case 32:
			goto tr398
		case 58:
			goto tr82
		case 91:
			goto tr91
		}
		goto tr79
	tr400:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st341
	tr139:
//line machine.go.rl:51

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st341
	tr473:
//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st341
	tr482:
//line machine.go.rl:63

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st341
	st341:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof341
		}
	st_case_341:
//line machine.go:5704
		switch (m.data)[(m.p)] {
		case 32:
			goto tr400
		case 58:
			goto tr402
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr401
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr401
			}
		default:
			goto tr401
		}
		goto tr399
	tr140:
//line machine.go.rl:51

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st342
	tr401:
//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st342
	tr481:
//line machine.go.rl:63

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st342
	st342:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof342
		}
	st_case_342:
//line machine.go:5785
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st343
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st343
			}
		default:
			goto st343
		}
		goto st86
	st343:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof343
		}
	st_case_343:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st344
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st344
			}
		default:
			goto st344
		}
		goto st87
	st344:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof344
		}
	st_case_344:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st345
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st345
			}
		default:
			goto st345
		}
		goto st88
	st345:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof345
		}
	st_case_345:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		goto st89
	tr404:
//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st346
	tr141:
//line machine.go.rl:51

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st346
	tr402:
//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st346
	st346:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof346
		}
	st_case_346:
//line machine.go:5927
		switch (m.data)[(m.p)] {
		case 32:
			goto tr77
		case 58:
			goto tr408
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr407
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr407
			}
		default:
			goto tr407
		}
		goto tr76
	tr407:
//line machine.go.rl:20

		m.pb = m.p

		goto st347
	st347:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof347
		}
	st_case_347:
//line machine.go:5958
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st352
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st349
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st349
			}
		default:
			goto st349
		}
		goto st52
	tr409:
//line machine.go.rl:96

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

		goto st348
	st348:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof348
		}
	st_case_348:
//line machine.go:5990
		switch (m.data)[(m.p)] {
		case 32:
			goto tr398
		case 58:
			goto tr82
		case 91:
			goto tr91
		}
		goto tr79
	st349:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof349
		}
	st_case_349:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st352
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st350
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st350
			}
		default:
			goto st350
		}
		goto st52
	st350:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof350
		}
	st_case_350:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st352
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st351
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st351
			}
		default:
			goto st351
		}
		goto st52
	st351:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof351
		}
	st_case_351:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st352
		}
		goto st52
	tr408:
//line machine.go.rl:20

		m.pb = m.p

		goto st352
	st352:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof352
		}
	st_case_352:
//line machine.go:6071
		if (m.data)[(m.p)] == 58 {
			goto st357
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st353
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st353
			}
		default:
			goto st353
		}
		goto st52
	st353:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof353
		}
	st_case_353:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st357
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st354
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st354
			}
		default:
			goto st354
		}
		goto st52
	st354:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof354
		}
	st_case_354:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st357
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st355
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st355
			}
		default:
			goto st355
		}
		goto st52
	st355:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof355
		}
	st_case_355:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st357
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st356
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st356
			}
		default:
			goto st356
		}
		goto st52
	st356:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof356
		}
	st_case_356:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st357
		}
		goto st52
	st357:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof357
		}
	st_case_357:
		if (m.data)[(m.p)] == 58 {
			goto st362
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st358
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st358
			}
		default:
			goto st358
		}
		goto st52
	st358:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof358
		}
	st_case_358:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st362
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st359
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st359
			}
		default:
			goto st359
		}
		goto st52
	st359:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof359
		}
	st_case_359:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st362
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st360
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st360
			}
		default:
			goto st360
		}
		goto st52
	st360:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof360
		}
	st_case_360:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st362
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st361
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st361
			}
		default:
			goto st361
		}
		goto st52
	st361:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof361
		}
	st_case_361:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st362
		}
		goto st52
	st362:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof362
		}
	st_case_362:
		if (m.data)[(m.p)] == 58 {
			goto st367
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st363
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st363
			}
		default:
			goto st363
		}
		goto st52
	st363:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof363
		}
	st_case_363:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st367
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st364
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st364
			}
		default:
			goto st364
		}
		goto st52
	st364:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof364
		}
	st_case_364:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st367
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st365
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st365
			}
		default:
			goto st365
		}
		goto st52
	st365:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof365
		}
	st_case_365:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st367
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st366
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st366
			}
		default:
			goto st366
		}
		goto st52
	st366:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof366
		}
	st_case_366:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st367
		}
		goto st52
	st367:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof367
		}
	st_case_367:
		if (m.data)[(m.p)] == 58 {
			goto st372
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st368
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st368
			}
		default:
			goto st368
		}
		goto st52
	st368:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof368
		}
	st_case_368:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st372
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st369
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st369
			}
		default:
			goto st369
		}
		goto st52
	st369:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof369
		}
	st_case_369:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st372
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st370
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st370
			}
		default:
			goto st370
		}
		goto st52
	st370:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof370
		}
	st_case_370:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st372
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st371
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st371
			}
		default:
			goto st371
		}
		goto st52
	st371:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof371
		}
	st_case_371:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st372
		}
		goto st52
	st372:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof372
		}
	st_case_372:
		if (m.data)[(m.p)] == 58 {
			goto st377
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st373
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st373
			}
		default:
			goto st373
		}
		goto st52
	st373:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof373
		}
	st_case_373:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st377
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st374
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st374
			}
		default:
			goto st374
		}
		goto st52
	st374:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof374
		}
	st_case_374:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st377
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st375
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st375
			}
		default:
			goto st375
		}
		goto st52
	st375:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof375
		}
	st_case_375:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st377
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st376
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st376
			}
		default:
			goto st376
		}
		goto st52
	st376:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof376
		}
	st_case_376:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr409
		case 58:
			goto st377
		}
		goto st52
	st377:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof377
		}
	st_case_377:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st378
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st378
			}
		default:
			goto st378
		}
		goto st52
	st378:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof378
		}
	st_case_378:
		if (m.data)[(m.p)] == 32 {
			goto tr409
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st379
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st379
			}
		default:
			goto st379
		}
		goto st52
	st379:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof379
		}
	st_case_379:
		if (m.data)[(m.p)] == 32 {
			goto tr409
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st380
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st380
			}
		default:
			goto st380
		}
		goto st52
	st380:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof380
		}
	st_case_380:
		if (m.data)[(m.p)] == 32 {
			goto tr409
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st381
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st381
			}
		default:
			goto st381
		}
		goto st52
	st381:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof381
		}
	st_case_381:
		if (m.data)[(m.p)] == 32 {
			goto tr409
		}
		goto st52
	st382:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof382
		}
	st_case_382:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st81
		}
		goto st52
	st383:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof383
		}
	st_case_383:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st384
		}
		goto st52
	st384:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof384
		}
	st_case_384:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st385
		}
		goto st52
	st385:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof385
		}
	st_case_385:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st386
		}
		goto st52
	st386:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof386
		}
	st_case_386:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st387
		}
		goto st52
	st387:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof387
		}
	st_case_387:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st388
		}
		goto st52
	st388:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof388
		}
	st_case_388:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st389
		}
		goto st52
	st389:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof389
		}
	st_case_389:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		goto st52
	st390:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof390
		}
	st_case_390:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st72
		}
		goto st55
	st391:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof391
		}
	st_case_391:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st69
		}
		goto st55
	st392:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof392
		}
	st_case_392:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st69
		}
		goto st55
	st393:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof393
		}
	st_case_393:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st66
		}
		goto st55
	tr83:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st394
	tr100:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st394
	st394:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof394
		}
	st_case_394:
//line machine.go:6901
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st395
		case 85:
			goto st430
		case 91:
			goto tr93
		case 112:
			goto st395
		case 117:
			goto st430
		}
		goto st55
	st395:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof395
		}
	st_case_395:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st396
		case 91:
			goto tr93
		case 114:
			goto st396
		}
		goto st55
	st396:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof396
		}
	st_case_396:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	tr58:
//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st397
	st397:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof397
		}
	st_case_397:
//line machine.go:6963
		switch (m.data)[(m.p)] {
		case 32:
			goto tr452
		case 48:
			goto tr453
		case 51:
			goto tr455
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto tr454
		}
		goto tr76
	tr452:
//line machine.go.rl:20

		m.pb = m.p

		goto st398
	st398:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof398
		}
	st_case_398:
//line machine.go:6987
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr456
		}
		goto tr76
	tr456:
//line machine.go.rl:20

		m.pb = m.p

		goto st399
	st399:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof399
		}
	st_case_399:
//line machine.go:7003
		if (m.data)[(m.p)] == 32 {
			goto st400
		}
		goto st52
	st400:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof400
		}
	st_case_400:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st401
		}
		goto st52
	st401:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof401
		}
	st_case_401:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st402
		}
		goto st52
	st402:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof402
		}
	st_case_402:
		if (m.data)[(m.p)] == 58 {
			goto st415
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st403
		}
		goto st52
	st403:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof403
		}
	st_case_403:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st404
		}
		goto st52
	st404:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof404
		}
	st_case_404:
		if (m.data)[(m.p)] == 32 {
			goto st405
		}
		goto st52
	st405:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof405
		}
	st_case_405:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st406
		}
		goto st52
	st406:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof406
		}
	st_case_406:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st407
		}
		goto st52
	st407:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof407
		}
	st_case_407:
		if (m.data)[(m.p)] == 58 {
			goto st408
		}
		goto st52
	st408:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof408
		}
	st_case_408:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st409
		}
		goto st52
	st409:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof409
		}
	st_case_409:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st410
		}
		goto st52
	st410:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof410
		}
	st_case_410:
		if (m.data)[(m.p)] == 58 {
			goto st411
		}
		goto st52
	st411:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof411
		}
	st_case_411:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st412
		}
		goto st52
	st412:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof412
		}
	st_case_412:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st413
		}
		goto st52
	st413:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof413
		}
	st_case_413:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr472
		case 58:
			goto tr472
		}
		goto st52
	tr472:
//line machine.go.rl:75

		{
			t, err := time.Parse("Jan _2 2006 15:04:05", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:147
		(m.p)--

		goto st414
	st414:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof414
		}
	st_case_414:
//line machine.go:7162
		switch (m.data)[(m.p)] {
		case 32:
			goto tr400
		case 58:
			goto tr473
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr401
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr401
			}
		default:
			goto tr401
		}
		goto tr399
	st415:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof415
		}
	st_case_415:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st416
		}
		goto st52
	st416:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof416
		}
	st_case_416:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st417
		}
		goto st52
	st417:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof417
		}
	st_case_417:
		if (m.data)[(m.p)] == 58 {
			goto st418
		}
		goto st52
	st418:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof418
		}
	st_case_418:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st419
		}
		goto st52
	st419:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof419
		}
	st_case_419:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st420
		}
		goto st52
	st420:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof420
		}
	st_case_420:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr480
		case 58:
			goto tr482
		case 91:
			goto tr483
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr481
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr481
			}
		default:
			goto tr481
		}
		goto tr479
	tr480:
//line machine.go.rl:63

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st421
	st421:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof421
		}
	st_case_421:
//line machine.go:7283
		switch (m.data)[(m.p)] {
		case 32:
			goto tr400
		case 58:
			goto tr402
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr484
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr401
			}
		default:
			goto tr401
		}
		goto tr399
	tr484:
//line machine.go.rl:24

		m.pHostname = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st422
	st422:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof422
		}
	st_case_422:
//line machine.go:7320
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st423
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st343
			}
		default:
			goto st343
		}
		goto st86
	st423:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof423
		}
	st_case_423:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st424
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st344
			}
		default:
			goto st344
		}
		goto st87
	st424:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof424
		}
	st_case_424:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st425
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st345
			}
		default:
			goto st345
		}
		goto st88
	st425:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof425
		}
	st_case_425:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr488
		case 58:
			goto tr404
		case 91:
			goto tr93
		}
		goto st89
	tr488:
//line machine.go.rl:87

		{
			year := common.UnsafeUTF8DecimalCodePointsToInt(m.text())
			// Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
			t := output.timestamp.UTC()
			output.timestamp = time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
		}

//line machine.go.rl:158
		(m.p)--

//line machine.go.rl:96

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st426
	st426:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof426
		}
	st_case_426:
//line machine.go:7437
		switch (m.data)[(m.p)] {
		case 32:
			goto tr400
		case 58:
			goto tr473
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr401
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr401
			}
		default:
			goto tr401
		}
		goto tr399
	tr453:
//line machine.go.rl:20

		m.pb = m.p

		goto st427
	st427:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof427
		}
	st_case_427:
//line machine.go:7470
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st399
		}
		goto st52
	tr454:
//line machine.go.rl:20

		m.pb = m.p

		goto st428
	st428:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof428
		}
	st_case_428:
//line machine.go:7486
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st399
		}
		goto st52
	tr455:
//line machine.go.rl:20

		m.pb = m.p

		goto st429
	st429:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof429
		}
	st_case_429:
//line machine.go:7502
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st399
		}
		goto st52
	st430:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof430
		}
	st_case_430:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 71:
			goto st396
		case 91:
			goto tr93
		case 103:
			goto st396
		}
		goto st55
	tr84:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st431
	tr101:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st431
	st431:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof431
		}
	st_case_431:
//line machine.go:7559
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st432
		case 91:
			goto tr93
		case 101:
			goto st432
		}
		goto st55
	st432:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof432
		}
	st_case_432:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st396
		case 91:
			goto tr93
		case 99:
			goto st396
		}
		goto st55
	tr85:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st433
	tr102:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st433
	st433:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof433
		}
	st_case_433:
//line machine.go:7625
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st434
		case 91:
			goto tr93
		case 101:
			goto st434
		}
		goto st55
	st434:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof434
		}
	st_case_434:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 66:
			goto st396
		case 91:
			goto tr93
		case 98:
			goto st396
		}
		goto st55
	tr86:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st435
	tr103:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st435
	st435:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof435
		}
	st_case_435:
//line machine.go:7691
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st436
		case 85:
			goto st437
		case 91:
			goto tr93
		case 97:
			goto st436
		case 117:
			goto st437
		}
		goto st55
	st436:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof436
		}
	st_case_436:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 78:
			goto st396
		case 91:
			goto tr93
		case 110:
			goto st396
		}
		goto st55
	st437:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof437
		}
	st_case_437:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 76:
			goto st396
		case 78:
			goto st396
		case 91:
			goto tr93
		case 108:
			goto st396
		case 110:
			goto st396
		}
		goto st55
	tr87:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st438
	tr104:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st438
	st438:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof438
		}
	st_case_438:
//line machine.go:7783
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st439
		case 91:
			goto tr93
		case 97:
			goto st439
		}
		goto st55
	st439:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof439
		}
	st_case_439:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st396
		case 89:
			goto st396
		case 91:
			goto tr93
		case 114:
			goto st396
		case 121:
			goto st396
		}
		goto st55
	tr88:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st440
	tr105:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st440
	st440:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof440
		}
	st_case_440:
//line machine.go:7853
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 79:
			goto st441
		case 91:
			goto tr93
		case 111:
			goto st441
		}
		goto st55
	st441:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof441
		}
	st_case_441:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 86:
			goto st396
		case 91:
			goto tr93
		case 118:
			goto st396
		}
		goto st55
	tr89:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st442
	tr106:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st442
	st442:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof442
		}
	st_case_442:
//line machine.go:7919
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st443
		case 91:
			goto tr93
		case 99:
			goto st443
		}
		goto st55
	st443:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof443
		}
	st_case_443:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st396
		case 91:
			goto tr93
		case 116:
			goto st396
		}
		goto st55
	tr90:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st444
	tr107:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st444
	st444:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof444
		}
	st_case_444:
//line machine.go:7985
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st445
		case 91:
			goto tr93
		case 101:
			goto st445
		}
		goto st55
	st445:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof445
		}
	st_case_445:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st396
		case 91:
			goto tr93
		case 112:
			goto st396
		}
		goto st55
	tr2:
//line machine.go.rl:20

		m.pb = m.p

		goto st6
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof6
		}
	st_case_6:
//line machine.go:8028
		switch (m.data)[(m.p)] {
		case 32:
			goto tr27
		case 58:
			goto tr27
		case 65:
			goto tr29
		case 68:
			goto tr30
		case 70:
			goto tr31
		case 74:
			goto tr32
		case 77:
			goto tr33
		case 78:
			goto tr34
		case 79:
			goto tr35
		case 83:
			goto tr36
		case 91:
			goto tr37
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr28
		}
		goto tr26
	tr28:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st7
	st7:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof7
		}
	st_case_7:
//line machine.go:8081
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st8
		}
		goto st2
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st9
		}
		goto st2
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st10
		}
		goto st2
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st11
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st12
		case 49:
			goto st23
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st12:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st13
		}
		goto st2
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st14
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st15
		case 51:
			goto st22
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st21
		}
		goto st2
	st15:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st16
		}
		goto st2
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st17
		case 91:
			goto tr17
		}
		goto st2
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 50:
			goto st20
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st18
		}
		goto st2
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st19
		}
		goto st2
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr54
		case 91:
			goto tr17
		}
		goto st2
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st19
		}
		goto st2
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st16
		}
		goto st2
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st16
		}
		goto st2
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st13
		}
		goto st2
	tr6:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st24
	tr29:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st24
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof24
		}
	st_case_24:
//line machine.go:8401
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st25
		case 85:
			goto st27
		case 91:
			goto tr17
		case 112:
			goto st25
		case 117:
			goto st27
		}
		goto st2
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st26
		case 91:
			goto tr17
		case 114:
			goto st26
		}
		goto st2
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 71:
			goto st26
		case 91:
			goto tr17
		case 103:
			goto st26
		}
		goto st2
	tr7:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st28
	tr30:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st28
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof28
		}
	st_case_28:
//line machine.go:8503
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st29
		case 91:
			goto tr17
		case 101:
			goto st29
		}
		goto st2
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st26
		case 91:
			goto tr17
		case 99:
			goto st26
		}
		goto st2
	tr8:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st30
	tr31:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st30
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof30
		}
	st_case_30:
//line machine.go:8569
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st31
		case 91:
			goto tr17
		case 101:
			goto st31
		}
		goto st2
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 66:
			goto st26
		case 91:
			goto tr17
		case 98:
			goto st26
		}
		goto st2
	tr9:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st32
	tr32:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st32
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:8635
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st33
		case 85:
			goto st34
		case 91:
			goto tr17
		case 97:
			goto st33
		case 117:
			goto st34
		}
		goto st2
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 78:
			goto st26
		case 91:
			goto tr17
		case 110:
			goto st26
		}
		goto st2
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 76:
			goto st26
		case 78:
			goto st26
		case 91:
			goto tr17
		case 108:
			goto st26
		case 110:
			goto st26
		}
		goto st2
	tr10:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st35
	tr33:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st35
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof35
		}
	st_case_35:
//line machine.go:8727
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st36
		case 91:
			goto tr17
		case 97:
			goto st36
		}
		goto st2
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st26
		case 89:
			goto st26
		case 91:
			goto tr17
		case 114:
			goto st26
		case 121:
			goto st26
		}
		goto st2
	tr11:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st37
	tr34:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st37
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof37
		}
	st_case_37:
//line machine.go:8797
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 79:
			goto st38
		case 91:
			goto tr17
		case 111:
			goto st38
		}
		goto st2
	st38:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 86:
			goto st26
		case 91:
			goto tr17
		case 118:
			goto st26
		}
		goto st2
	tr12:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st39
	tr35:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st39
	st39:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof39
		}
	st_case_39:
//line machine.go:8863
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st40
		case 91:
			goto tr17
		case 99:
			goto st40
		}
		goto st2
	st40:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st26
		case 91:
			goto tr17
		case 116:
			goto st26
		}
		goto st2
	tr13:
//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st41
	tr36:
//line machine.go.rl:42

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:28

		m.pTimestamp = m.p

//line machine.go.rl:20

		m.pb = m.p

		goto st41
	st41:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof41
		}
	st_case_41:
//line machine.go:8929
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st42
		case 91:
			goto tr17
		case 101:
			goto st42
		}
		goto st2
	st42:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st26
		case 91:
			goto tr17
		case 112:
			goto st26
		}
		goto st2
	tr3:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:28

		m.pTimestamp = m.p

		goto st43
	st43:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof43
		}
	st_case_43:
//line machine.go:8976
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st44
		}
		goto st2
	st44:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st45
		}
		goto st2
	st45:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st46
		}
		goto st2
	st46:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st11
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st47
		}
		goto st2
	st47:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st47
		}
		goto st2
	tr4:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:37

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st446
	tr68:
//line machine.go.rl:37

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

//line machine.go.rl:101

		output.tagSet = true
		output.tag = string(m.text())

		goto st446
	st446:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof446
		}
	st_case_446:
//line machine.go:9092
		if (m.data)[(m.p)] == 32 {
			goto tr498
		}
		goto tr76
	tr5:
//line machine.go.rl:20

		m.pb = m.p

		goto st48
	st48:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof48
		}
	st_case_48:
//line machine.go:9108
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 62:
			goto tr73
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr72
		}
		goto st2
	tr72:
//line machine.go.rl:20

		m.pb = m.p

		goto st49
	st49:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof49
		}
	st_case_49:
//line machine.go:9134
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 62:
			goto tr75
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st49
		}
		goto st2
	tr73:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:32

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st50
	tr75:
//line machine.go.rl:32

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st50
	st50:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof50
		}
	st_case_50:
//line machine.go:9172
		switch (m.data)[(m.p)] {
		case 32:
			goto tr1
		case 42:
			goto tr2
		case 46:
			goto tr2
		case 58:
			goto tr4
		case 65:
			goto tr6
		case 68:
			goto tr7
		case 70:
			goto tr8
		case 74:
			goto tr9
		case 77:
			goto tr10
		case 78:
			goto tr11
		case 79:
			goto tr12
		case 83:
			goto tr13
		case 91:
			goto tr14
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr3
		}
		goto tr0
	st_out:
	_test_eof2:
		m.cs = 2
		goto _test_eof
	_test_eof51:
		m.cs = 51
		goto _test_eof
	_test_eof52:
		m.cs = 52
		goto _test_eof
	_test_eof53:
		m.cs = 53
		goto _test_eof
	_test_eof3:
		m.cs = 3
		goto _test_eof
	_test_eof4:
		m.cs = 4
		goto _test_eof
	_test_eof5:
		m.cs = 5
		goto _test_eof
	_test_eof54:
		m.cs = 54
		goto _test_eof
	_test_eof55:
		m.cs = 55
		goto _test_eof
	_test_eof56:
		m.cs = 56
		goto _test_eof
	_test_eof57:
		m.cs = 57
		goto _test_eof
	_test_eof58:
		m.cs = 58
		goto _test_eof
	_test_eof59:
		m.cs = 59
		goto _test_eof
	_test_eof60:
		m.cs = 60
		goto _test_eof
	_test_eof61:
		m.cs = 61
		goto _test_eof
	_test_eof62:
		m.cs = 62
		goto _test_eof
	_test_eof63:
		m.cs = 63
		goto _test_eof
	_test_eof64:
		m.cs = 64
		goto _test_eof
	_test_eof65:
		m.cs = 65
		goto _test_eof
	_test_eof66:
		m.cs = 66
		goto _test_eof
	_test_eof67:
		m.cs = 67
		goto _test_eof
	_test_eof68:
		m.cs = 68
		goto _test_eof
	_test_eof69:
		m.cs = 69
		goto _test_eof
	_test_eof70:
		m.cs = 70
		goto _test_eof
	_test_eof71:
		m.cs = 71
		goto _test_eof
	_test_eof72:
		m.cs = 72
		goto _test_eof
	_test_eof73:
		m.cs = 73
		goto _test_eof
	_test_eof74:
		m.cs = 74
		goto _test_eof
	_test_eof75:
		m.cs = 75
		goto _test_eof
	_test_eof76:
		m.cs = 76
		goto _test_eof
	_test_eof77:
		m.cs = 77
		goto _test_eof
	_test_eof78:
		m.cs = 78
		goto _test_eof
	_test_eof79:
		m.cs = 79
		goto _test_eof
	_test_eof80:
		m.cs = 80
		goto _test_eof
	_test_eof81:
		m.cs = 81
		goto _test_eof
	_test_eof82:
		m.cs = 82
		goto _test_eof
	_test_eof83:
		m.cs = 83
		goto _test_eof
	_test_eof84:
		m.cs = 84
		goto _test_eof
	_test_eof85:
		m.cs = 85
		goto _test_eof
	_test_eof86:
		m.cs = 86
		goto _test_eof
	_test_eof87:
		m.cs = 87
		goto _test_eof
	_test_eof88:
		m.cs = 88
		goto _test_eof
	_test_eof89:
		m.cs = 89
		goto _test_eof
	_test_eof90:
		m.cs = 90
		goto _test_eof
	_test_eof91:
		m.cs = 91
		goto _test_eof
	_test_eof92:
		m.cs = 92
		goto _test_eof
	_test_eof93:
		m.cs = 93
		goto _test_eof
	_test_eof94:
		m.cs = 94
		goto _test_eof
	_test_eof95:
		m.cs = 95
		goto _test_eof
	_test_eof96:
		m.cs = 96
		goto _test_eof
	_test_eof97:
		m.cs = 97
		goto _test_eof
	_test_eof98:
		m.cs = 98
		goto _test_eof
	_test_eof99:
		m.cs = 99
		goto _test_eof
	_test_eof100:
		m.cs = 100
		goto _test_eof
	_test_eof101:
		m.cs = 101
		goto _test_eof
	_test_eof102:
		m.cs = 102
		goto _test_eof
	_test_eof103:
		m.cs = 103
		goto _test_eof
	_test_eof104:
		m.cs = 104
		goto _test_eof
	_test_eof105:
		m.cs = 105
		goto _test_eof
	_test_eof106:
		m.cs = 106
		goto _test_eof
	_test_eof107:
		m.cs = 107
		goto _test_eof
	_test_eof108:
		m.cs = 108
		goto _test_eof
	_test_eof109:
		m.cs = 109
		goto _test_eof
	_test_eof110:
		m.cs = 110
		goto _test_eof
	_test_eof111:
		m.cs = 111
		goto _test_eof
	_test_eof112:
		m.cs = 112
		goto _test_eof
	_test_eof113:
		m.cs = 113
		goto _test_eof
	_test_eof114:
		m.cs = 114
		goto _test_eof
	_test_eof115:
		m.cs = 115
		goto _test_eof
	_test_eof116:
		m.cs = 116
		goto _test_eof
	_test_eof117:
		m.cs = 117
		goto _test_eof
	_test_eof118:
		m.cs = 118
		goto _test_eof
	_test_eof119:
		m.cs = 119
		goto _test_eof
	_test_eof120:
		m.cs = 120
		goto _test_eof
	_test_eof121:
		m.cs = 121
		goto _test_eof
	_test_eof122:
		m.cs = 122
		goto _test_eof
	_test_eof123:
		m.cs = 123
		goto _test_eof
	_test_eof124:
		m.cs = 124
		goto _test_eof
	_test_eof125:
		m.cs = 125
		goto _test_eof
	_test_eof126:
		m.cs = 126
		goto _test_eof
	_test_eof127:
		m.cs = 127
		goto _test_eof
	_test_eof128:
		m.cs = 128
		goto _test_eof
	_test_eof129:
		m.cs = 129
		goto _test_eof
	_test_eof130:
		m.cs = 130
		goto _test_eof
	_test_eof131:
		m.cs = 131
		goto _test_eof
	_test_eof132:
		m.cs = 132
		goto _test_eof
	_test_eof133:
		m.cs = 133
		goto _test_eof
	_test_eof134:
		m.cs = 134
		goto _test_eof
	_test_eof135:
		m.cs = 135
		goto _test_eof
	_test_eof136:
		m.cs = 136
		goto _test_eof
	_test_eof137:
		m.cs = 137
		goto _test_eof
	_test_eof138:
		m.cs = 138
		goto _test_eof
	_test_eof139:
		m.cs = 139
		goto _test_eof
	_test_eof140:
		m.cs = 140
		goto _test_eof
	_test_eof141:
		m.cs = 141
		goto _test_eof
	_test_eof142:
		m.cs = 142
		goto _test_eof
	_test_eof143:
		m.cs = 143
		goto _test_eof
	_test_eof144:
		m.cs = 144
		goto _test_eof
	_test_eof145:
		m.cs = 145
		goto _test_eof
	_test_eof146:
		m.cs = 146
		goto _test_eof
	_test_eof147:
		m.cs = 147
		goto _test_eof
	_test_eof148:
		m.cs = 148
		goto _test_eof
	_test_eof149:
		m.cs = 149
		goto _test_eof
	_test_eof150:
		m.cs = 150
		goto _test_eof
	_test_eof151:
		m.cs = 151
		goto _test_eof
	_test_eof152:
		m.cs = 152
		goto _test_eof
	_test_eof153:
		m.cs = 153
		goto _test_eof
	_test_eof154:
		m.cs = 154
		goto _test_eof
	_test_eof155:
		m.cs = 155
		goto _test_eof
	_test_eof156:
		m.cs = 156
		goto _test_eof
	_test_eof157:
		m.cs = 157
		goto _test_eof
	_test_eof158:
		m.cs = 158
		goto _test_eof
	_test_eof159:
		m.cs = 159
		goto _test_eof
	_test_eof160:
		m.cs = 160
		goto _test_eof
	_test_eof161:
		m.cs = 161
		goto _test_eof
	_test_eof162:
		m.cs = 162
		goto _test_eof
	_test_eof163:
		m.cs = 163
		goto _test_eof
	_test_eof164:
		m.cs = 164
		goto _test_eof
	_test_eof165:
		m.cs = 165
		goto _test_eof
	_test_eof166:
		m.cs = 166
		goto _test_eof
	_test_eof167:
		m.cs = 167
		goto _test_eof
	_test_eof168:
		m.cs = 168
		goto _test_eof
	_test_eof169:
		m.cs = 169
		goto _test_eof
	_test_eof170:
		m.cs = 170
		goto _test_eof
	_test_eof171:
		m.cs = 171
		goto _test_eof
	_test_eof172:
		m.cs = 172
		goto _test_eof
	_test_eof173:
		m.cs = 173
		goto _test_eof
	_test_eof174:
		m.cs = 174
		goto _test_eof
	_test_eof175:
		m.cs = 175
		goto _test_eof
	_test_eof176:
		m.cs = 176
		goto _test_eof
	_test_eof177:
		m.cs = 177
		goto _test_eof
	_test_eof178:
		m.cs = 178
		goto _test_eof
	_test_eof179:
		m.cs = 179
		goto _test_eof
	_test_eof180:
		m.cs = 180
		goto _test_eof
	_test_eof181:
		m.cs = 181
		goto _test_eof
	_test_eof182:
		m.cs = 182
		goto _test_eof
	_test_eof183:
		m.cs = 183
		goto _test_eof
	_test_eof184:
		m.cs = 184
		goto _test_eof
	_test_eof185:
		m.cs = 185
		goto _test_eof
	_test_eof186:
		m.cs = 186
		goto _test_eof
	_test_eof187:
		m.cs = 187
		goto _test_eof
	_test_eof188:
		m.cs = 188
		goto _test_eof
	_test_eof189:
		m.cs = 189
		goto _test_eof
	_test_eof190:
		m.cs = 190
		goto _test_eof
	_test_eof191:
		m.cs = 191
		goto _test_eof
	_test_eof192:
		m.cs = 192
		goto _test_eof
	_test_eof193:
		m.cs = 193
		goto _test_eof
	_test_eof194:
		m.cs = 194
		goto _test_eof
	_test_eof195:
		m.cs = 195
		goto _test_eof
	_test_eof196:
		m.cs = 196
		goto _test_eof
	_test_eof197:
		m.cs = 197
		goto _test_eof
	_test_eof198:
		m.cs = 198
		goto _test_eof
	_test_eof199:
		m.cs = 199
		goto _test_eof
	_test_eof200:
		m.cs = 200
		goto _test_eof
	_test_eof201:
		m.cs = 201
		goto _test_eof
	_test_eof202:
		m.cs = 202
		goto _test_eof
	_test_eof203:
		m.cs = 203
		goto _test_eof
	_test_eof204:
		m.cs = 204
		goto _test_eof
	_test_eof205:
		m.cs = 205
		goto _test_eof
	_test_eof206:
		m.cs = 206
		goto _test_eof
	_test_eof207:
		m.cs = 207
		goto _test_eof
	_test_eof208:
		m.cs = 208
		goto _test_eof
	_test_eof209:
		m.cs = 209
		goto _test_eof
	_test_eof210:
		m.cs = 210
		goto _test_eof
	_test_eof211:
		m.cs = 211
		goto _test_eof
	_test_eof212:
		m.cs = 212
		goto _test_eof
	_test_eof213:
		m.cs = 213
		goto _test_eof
	_test_eof214:
		m.cs = 214
		goto _test_eof
	_test_eof215:
		m.cs = 215
		goto _test_eof
	_test_eof216:
		m.cs = 216
		goto _test_eof
	_test_eof217:
		m.cs = 217
		goto _test_eof
	_test_eof218:
		m.cs = 218
		goto _test_eof
	_test_eof219:
		m.cs = 219
		goto _test_eof
	_test_eof220:
		m.cs = 220
		goto _test_eof
	_test_eof221:
		m.cs = 221
		goto _test_eof
	_test_eof222:
		m.cs = 222
		goto _test_eof
	_test_eof223:
		m.cs = 223
		goto _test_eof
	_test_eof224:
		m.cs = 224
		goto _test_eof
	_test_eof225:
		m.cs = 225
		goto _test_eof
	_test_eof226:
		m.cs = 226
		goto _test_eof
	_test_eof227:
		m.cs = 227
		goto _test_eof
	_test_eof228:
		m.cs = 228
		goto _test_eof
	_test_eof229:
		m.cs = 229
		goto _test_eof
	_test_eof230:
		m.cs = 230
		goto _test_eof
	_test_eof231:
		m.cs = 231
		goto _test_eof
	_test_eof232:
		m.cs = 232
		goto _test_eof
	_test_eof233:
		m.cs = 233
		goto _test_eof
	_test_eof234:
		m.cs = 234
		goto _test_eof
	_test_eof235:
		m.cs = 235
		goto _test_eof
	_test_eof236:
		m.cs = 236
		goto _test_eof
	_test_eof237:
		m.cs = 237
		goto _test_eof
	_test_eof238:
		m.cs = 238
		goto _test_eof
	_test_eof239:
		m.cs = 239
		goto _test_eof
	_test_eof240:
		m.cs = 240
		goto _test_eof
	_test_eof241:
		m.cs = 241
		goto _test_eof
	_test_eof242:
		m.cs = 242
		goto _test_eof
	_test_eof243:
		m.cs = 243
		goto _test_eof
	_test_eof244:
		m.cs = 244
		goto _test_eof
	_test_eof245:
		m.cs = 245
		goto _test_eof
	_test_eof246:
		m.cs = 246
		goto _test_eof
	_test_eof247:
		m.cs = 247
		goto _test_eof
	_test_eof248:
		m.cs = 248
		goto _test_eof
	_test_eof249:
		m.cs = 249
		goto _test_eof
	_test_eof250:
		m.cs = 250
		goto _test_eof
	_test_eof251:
		m.cs = 251
		goto _test_eof
	_test_eof252:
		m.cs = 252
		goto _test_eof
	_test_eof253:
		m.cs = 253
		goto _test_eof
	_test_eof254:
		m.cs = 254
		goto _test_eof
	_test_eof255:
		m.cs = 255
		goto _test_eof
	_test_eof256:
		m.cs = 256
		goto _test_eof
	_test_eof257:
		m.cs = 257
		goto _test_eof
	_test_eof258:
		m.cs = 258
		goto _test_eof
	_test_eof259:
		m.cs = 259
		goto _test_eof
	_test_eof260:
		m.cs = 260
		goto _test_eof
	_test_eof261:
		m.cs = 261
		goto _test_eof
	_test_eof262:
		m.cs = 262
		goto _test_eof
	_test_eof263:
		m.cs = 263
		goto _test_eof
	_test_eof264:
		m.cs = 264
		goto _test_eof
	_test_eof265:
		m.cs = 265
		goto _test_eof
	_test_eof266:
		m.cs = 266
		goto _test_eof
	_test_eof267:
		m.cs = 267
		goto _test_eof
	_test_eof268:
		m.cs = 268
		goto _test_eof
	_test_eof269:
		m.cs = 269
		goto _test_eof
	_test_eof270:
		m.cs = 270
		goto _test_eof
	_test_eof271:
		m.cs = 271
		goto _test_eof
	_test_eof272:
		m.cs = 272
		goto _test_eof
	_test_eof273:
		m.cs = 273
		goto _test_eof
	_test_eof274:
		m.cs = 274
		goto _test_eof
	_test_eof275:
		m.cs = 275
		goto _test_eof
	_test_eof276:
		m.cs = 276
		goto _test_eof
	_test_eof277:
		m.cs = 277
		goto _test_eof
	_test_eof278:
		m.cs = 278
		goto _test_eof
	_test_eof279:
		m.cs = 279
		goto _test_eof
	_test_eof280:
		m.cs = 280
		goto _test_eof
	_test_eof281:
		m.cs = 281
		goto _test_eof
	_test_eof282:
		m.cs = 282
		goto _test_eof
	_test_eof283:
		m.cs = 283
		goto _test_eof
	_test_eof284:
		m.cs = 284
		goto _test_eof
	_test_eof285:
		m.cs = 285
		goto _test_eof
	_test_eof286:
		m.cs = 286
		goto _test_eof
	_test_eof287:
		m.cs = 287
		goto _test_eof
	_test_eof288:
		m.cs = 288
		goto _test_eof
	_test_eof289:
		m.cs = 289
		goto _test_eof
	_test_eof290:
		m.cs = 290
		goto _test_eof
	_test_eof291:
		m.cs = 291
		goto _test_eof
	_test_eof292:
		m.cs = 292
		goto _test_eof
	_test_eof293:
		m.cs = 293
		goto _test_eof
	_test_eof294:
		m.cs = 294
		goto _test_eof
	_test_eof295:
		m.cs = 295
		goto _test_eof
	_test_eof296:
		m.cs = 296
		goto _test_eof
	_test_eof297:
		m.cs = 297
		goto _test_eof
	_test_eof298:
		m.cs = 298
		goto _test_eof
	_test_eof299:
		m.cs = 299
		goto _test_eof
	_test_eof300:
		m.cs = 300
		goto _test_eof
	_test_eof301:
		m.cs = 301
		goto _test_eof
	_test_eof302:
		m.cs = 302
		goto _test_eof
	_test_eof303:
		m.cs = 303
		goto _test_eof
	_test_eof304:
		m.cs = 304
		goto _test_eof
	_test_eof305:
		m.cs = 305
		goto _test_eof
	_test_eof306:
		m.cs = 306
		goto _test_eof
	_test_eof307:
		m.cs = 307
		goto _test_eof
	_test_eof308:
		m.cs = 308
		goto _test_eof
	_test_eof309:
		m.cs = 309
		goto _test_eof
	_test_eof310:
		m.cs = 310
		goto _test_eof
	_test_eof311:
		m.cs = 311
		goto _test_eof
	_test_eof312:
		m.cs = 312
		goto _test_eof
	_test_eof313:
		m.cs = 313
		goto _test_eof
	_test_eof314:
		m.cs = 314
		goto _test_eof
	_test_eof315:
		m.cs = 315
		goto _test_eof
	_test_eof316:
		m.cs = 316
		goto _test_eof
	_test_eof317:
		m.cs = 317
		goto _test_eof
	_test_eof318:
		m.cs = 318
		goto _test_eof
	_test_eof319:
		m.cs = 319
		goto _test_eof
	_test_eof320:
		m.cs = 320
		goto _test_eof
	_test_eof321:
		m.cs = 321
		goto _test_eof
	_test_eof322:
		m.cs = 322
		goto _test_eof
	_test_eof323:
		m.cs = 323
		goto _test_eof
	_test_eof324:
		m.cs = 324
		goto _test_eof
	_test_eof325:
		m.cs = 325
		goto _test_eof
	_test_eof326:
		m.cs = 326
		goto _test_eof
	_test_eof327:
		m.cs = 327
		goto _test_eof
	_test_eof328:
		m.cs = 328
		goto _test_eof
	_test_eof329:
		m.cs = 329
		goto _test_eof
	_test_eof330:
		m.cs = 330
		goto _test_eof
	_test_eof331:
		m.cs = 331
		goto _test_eof
	_test_eof332:
		m.cs = 332
		goto _test_eof
	_test_eof333:
		m.cs = 333
		goto _test_eof
	_test_eof334:
		m.cs = 334
		goto _test_eof
	_test_eof335:
		m.cs = 335
		goto _test_eof
	_test_eof336:
		m.cs = 336
		goto _test_eof
	_test_eof337:
		m.cs = 337
		goto _test_eof
	_test_eof338:
		m.cs = 338
		goto _test_eof
	_test_eof339:
		m.cs = 339
		goto _test_eof
	_test_eof340:
		m.cs = 340
		goto _test_eof
	_test_eof341:
		m.cs = 341
		goto _test_eof
	_test_eof342:
		m.cs = 342
		goto _test_eof
	_test_eof343:
		m.cs = 343
		goto _test_eof
	_test_eof344:
		m.cs = 344
		goto _test_eof
	_test_eof345:
		m.cs = 345
		goto _test_eof
	_test_eof346:
		m.cs = 346
		goto _test_eof
	_test_eof347:
		m.cs = 347
		goto _test_eof
	_test_eof348:
		m.cs = 348
		goto _test_eof
	_test_eof349:
		m.cs = 349
		goto _test_eof
	_test_eof350:
		m.cs = 350
		goto _test_eof
	_test_eof351:
		m.cs = 351
		goto _test_eof
	_test_eof352:
		m.cs = 352
		goto _test_eof
	_test_eof353:
		m.cs = 353
		goto _test_eof
	_test_eof354:
		m.cs = 354
		goto _test_eof
	_test_eof355:
		m.cs = 355
		goto _test_eof
	_test_eof356:
		m.cs = 356
		goto _test_eof
	_test_eof357:
		m.cs = 357
		goto _test_eof
	_test_eof358:
		m.cs = 358
		goto _test_eof
	_test_eof359:
		m.cs = 359
		goto _test_eof
	_test_eof360:
		m.cs = 360
		goto _test_eof
	_test_eof361:
		m.cs = 361
		goto _test_eof
	_test_eof362:
		m.cs = 362
		goto _test_eof
	_test_eof363:
		m.cs = 363
		goto _test_eof
	_test_eof364:
		m.cs = 364
		goto _test_eof
	_test_eof365:
		m.cs = 365
		goto _test_eof
	_test_eof366:
		m.cs = 366
		goto _test_eof
	_test_eof367:
		m.cs = 367
		goto _test_eof
	_test_eof368:
		m.cs = 368
		goto _test_eof
	_test_eof369:
		m.cs = 369
		goto _test_eof
	_test_eof370:
		m.cs = 370
		goto _test_eof
	_test_eof371:
		m.cs = 371
		goto _test_eof
	_test_eof372:
		m.cs = 372
		goto _test_eof
	_test_eof373:
		m.cs = 373
		goto _test_eof
	_test_eof374:
		m.cs = 374
		goto _test_eof
	_test_eof375:
		m.cs = 375
		goto _test_eof
	_test_eof376:
		m.cs = 376
		goto _test_eof
	_test_eof377:
		m.cs = 377
		goto _test_eof
	_test_eof378:
		m.cs = 378
		goto _test_eof
	_test_eof379:
		m.cs = 379
		goto _test_eof
	_test_eof380:
		m.cs = 380
		goto _test_eof
	_test_eof381:
		m.cs = 381
		goto _test_eof
	_test_eof382:
		m.cs = 382
		goto _test_eof
	_test_eof383:
		m.cs = 383
		goto _test_eof
	_test_eof384:
		m.cs = 384
		goto _test_eof
	_test_eof385:
		m.cs = 385
		goto _test_eof
	_test_eof386:
		m.cs = 386
		goto _test_eof
	_test_eof387:
		m.cs = 387
		goto _test_eof
	_test_eof388:
		m.cs = 388
		goto _test_eof
	_test_eof389:
		m.cs = 389
		goto _test_eof
	_test_eof390:
		m.cs = 390
		goto _test_eof
	_test_eof391:
		m.cs = 391
		goto _test_eof
	_test_eof392:
		m.cs = 392
		goto _test_eof
	_test_eof393:
		m.cs = 393
		goto _test_eof
	_test_eof394:
		m.cs = 394
		goto _test_eof
	_test_eof395:
		m.cs = 395
		goto _test_eof
	_test_eof396:
		m.cs = 396
		goto _test_eof
	_test_eof397:
		m.cs = 397
		goto _test_eof
	_test_eof398:
		m.cs = 398
		goto _test_eof
	_test_eof399:
		m.cs = 399
		goto _test_eof
	_test_eof400:
		m.cs = 400
		goto _test_eof
	_test_eof401:
		m.cs = 401
		goto _test_eof
	_test_eof402:
		m.cs = 402
		goto _test_eof
	_test_eof403:
		m.cs = 403
		goto _test_eof
	_test_eof404:
		m.cs = 404
		goto _test_eof
	_test_eof405:
		m.cs = 405
		goto _test_eof
	_test_eof406:
		m.cs = 406
		goto _test_eof
	_test_eof407:
		m.cs = 407
		goto _test_eof
	_test_eof408:
		m.cs = 408
		goto _test_eof
	_test_eof409:
		m.cs = 409
		goto _test_eof
	_test_eof410:
		m.cs = 410
		goto _test_eof
	_test_eof411:
		m.cs = 411
		goto _test_eof
	_test_eof412:
		m.cs = 412
		goto _test_eof
	_test_eof413:
		m.cs = 413
		goto _test_eof
	_test_eof414:
		m.cs = 414
		goto _test_eof
	_test_eof415:
		m.cs = 415
		goto _test_eof
	_test_eof416:
		m.cs = 416
		goto _test_eof
	_test_eof417:
		m.cs = 417
		goto _test_eof
	_test_eof418:
		m.cs = 418
		goto _test_eof
	_test_eof419:
		m.cs = 419
		goto _test_eof
	_test_eof420:
		m.cs = 420
		goto _test_eof
	_test_eof421:
		m.cs = 421
		goto _test_eof
	_test_eof422:
		m.cs = 422
		goto _test_eof
	_test_eof423:
		m.cs = 423
		goto _test_eof
	_test_eof424:
		m.cs = 424
		goto _test_eof
	_test_eof425:
		m.cs = 425
		goto _test_eof
	_test_eof426:
		m.cs = 426
		goto _test_eof
	_test_eof427:
		m.cs = 427
		goto _test_eof
	_test_eof428:
		m.cs = 428
		goto _test_eof
	_test_eof429:
		m.cs = 429
		goto _test_eof
	_test_eof430:
		m.cs = 430
		goto _test_eof
	_test_eof431:
		m.cs = 431
		goto _test_eof
	_test_eof432:
		m.cs = 432
		goto _test_eof
	_test_eof433:
		m.cs = 433
		goto _test_eof
	_test_eof434:
		m.cs = 434
		goto _test_eof
	_test_eof435:
		m.cs = 435
		goto _test_eof
	_test_eof436:
		m.cs = 436
		goto _test_eof
	_test_eof437:
		m.cs = 437
		goto _test_eof
	_test_eof438:
		m.cs = 438
		goto _test_eof
	_test_eof439:
		m.cs = 439
		goto _test_eof
	_test_eof440:
		m.cs = 440
		goto _test_eof
	_test_eof441:
		m.cs = 441
		goto _test_eof
	_test_eof442:
		m.cs = 442
		goto _test_eof
	_test_eof443:
		m.cs = 443
		goto _test_eof
	_test_eof444:
		m.cs = 444
		goto _test_eof
	_test_eof445:
		m.cs = 445
		goto _test_eof
	_test_eof6:
		m.cs = 6
		goto _test_eof
	_test_eof7:
		m.cs = 7
		goto _test_eof
	_test_eof8:
		m.cs = 8
		goto _test_eof
	_test_eof9:
		m.cs = 9
		goto _test_eof
	_test_eof10:
		m.cs = 10
		goto _test_eof
	_test_eof11:
		m.cs = 11
		goto _test_eof
	_test_eof12:
		m.cs = 12
		goto _test_eof
	_test_eof13:
		m.cs = 13
		goto _test_eof
	_test_eof14:
		m.cs = 14
		goto _test_eof
	_test_eof15:
		m.cs = 15
		goto _test_eof
	_test_eof16:
		m.cs = 16
		goto _test_eof
	_test_eof17:
		m.cs = 17
		goto _test_eof
	_test_eof18:
		m.cs = 18
		goto _test_eof
	_test_eof19:
		m.cs = 19
		goto _test_eof
	_test_eof20:
		m.cs = 20
		goto _test_eof
	_test_eof21:
		m.cs = 21
		goto _test_eof
	_test_eof22:
		m.cs = 22
		goto _test_eof
	_test_eof23:
		m.cs = 23
		goto _test_eof
	_test_eof24:
		m.cs = 24
		goto _test_eof
	_test_eof25:
		m.cs = 25
		goto _test_eof
	_test_eof26:
		m.cs = 26
		goto _test_eof
	_test_eof27:
		m.cs = 27
		goto _test_eof
	_test_eof28:
		m.cs = 28
		goto _test_eof
	_test_eof29:
		m.cs = 29
		goto _test_eof
	_test_eof30:
		m.cs = 30
		goto _test_eof
	_test_eof31:
		m.cs = 31
		goto _test_eof
	_test_eof32:
		m.cs = 32
		goto _test_eof
	_test_eof33:
		m.cs = 33
		goto _test_eof
	_test_eof34:
		m.cs = 34
		goto _test_eof
	_test_eof35:
		m.cs = 35
		goto _test_eof
	_test_eof36:
		m.cs = 36
		goto _test_eof
	_test_eof37:
		m.cs = 37
		goto _test_eof
	_test_eof38:
		m.cs = 38
		goto _test_eof
	_test_eof39:
		m.cs = 39
		goto _test_eof
	_test_eof40:
		m.cs = 40
		goto _test_eof
	_test_eof41:
		m.cs = 41
		goto _test_eof
	_test_eof42:
		m.cs = 42
		goto _test_eof
	_test_eof43:
		m.cs = 43
		goto _test_eof
	_test_eof44:
		m.cs = 44
		goto _test_eof
	_test_eof45:
		m.cs = 45
		goto _test_eof
	_test_eof46:
		m.cs = 46
		goto _test_eof
	_test_eof47:
		m.cs = 47
		goto _test_eof
	_test_eof446:
		m.cs = 446
		goto _test_eof
	_test_eof48:
		m.cs = 48
		goto _test_eof
	_test_eof49:
		m.cs = 49
		goto _test_eof
	_test_eof50:
		m.cs = 50
		goto _test_eof

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50:
//line machine.go.rl:101

				output.tagSet = true
				output.tag = string(m.text())

			case 52, 56, 57, 58, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 347, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 415, 416, 417, 418, 419, 427, 428, 429:
//line machine.go.rl:111

				output.messageSet = true
				output.message = string(m.text())

			case 51, 53, 73, 346, 397, 398, 446:
//line machine.go.rl:20

				m.pb = m.p

//line machine.go.rl:111

				output.messageSet = true
				output.message = string(m.text())

			case 55, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 342, 343, 344, 345, 348, 390, 391, 392, 393, 394, 395, 396, 414, 420, 422, 423, 424, 425, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445:
//line machine.go.rl:101

				output.tagSet = true
				output.tag = string(m.text())

//line machine.go.rl:111

				output.messageSet = true
				output.message = string(m.text())

			case 54, 340, 341, 421, 426:
//line machine.go.rl:101

				output.tagSet = true
				output.tag = string(m.text())

//line machine.go.rl:20

				m.pb = m.p

//line machine.go.rl:111

				output.messageSet = true
				output.message = string(m.text())

//line machine.go:9703
			}
		}

	_out:
		{
		}
	}

//line machine.go.rl:283

	return output.export(), m.err
}

// WithBestEffort enables best effort mode. Has no effect since
// legacysyslog parsing is always best effort.
func (m *machine) WithBestEffort() {
}

// HasBestEffort tells whether the receiving parser has best effort mode on or off.
// Always returns true since legacysyslog parsing is always best effort.
func (m *machine) HasBestEffort() bool {
	return true
}
