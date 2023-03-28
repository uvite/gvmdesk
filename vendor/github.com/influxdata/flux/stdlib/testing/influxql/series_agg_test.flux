package influxql_test


import "testing"
import "internal/influxql"
import "csv"

inData =
    "
#datatype,string,long,dateTime:RFC3339,string,string,string,double
#group,false,false,false,true,true,true,false
#default,0,,,,,,
,result,table,_time,_measurement,t0,_field,_value
,,0,1970-01-01T00:00:00Z,m,0,f,0.19434194999233168
,,0,1970-01-01T01:00:00Z,m,0,f,0.35586976154169886
,,0,1970-01-01T02:00:00Z,m,0,f,0.9008931119054228
,,0,1970-01-01T03:00:00Z,m,0,f,0.6461505985646413
,,0,1970-01-01T04:00:00Z,m,0,f,0.1340222613556339
,,0,1970-01-01T05:00:00Z,m,0,f,0.3050922896043849
,,0,1970-01-01T06:00:00Z,m,0,f,0.16797790004756785
,,0,1970-01-01T07:00:00Z,m,0,f,0.6859900761088404
,,0,1970-01-01T08:00:00Z,m,0,f,0.3813372334346726
,,0,1970-01-01T09:00:00Z,m,0,f,0.37739800802050527
,,0,1970-01-01T10:00:00Z,m,0,f,0.2670215125945959
,,0,1970-01-01T11:00:00Z,m,0,f,0.19857273235709308
,,0,1970-01-01T12:00:00Z,m,0,f,0.7926413090714327
,,0,1970-01-01T13:00:00Z,m,0,f,0.8488436313118317
,,0,1970-01-01T14:00:00Z,m,0,f,0.1960293435787179
,,0,1970-01-01T15:00:00Z,m,0,f,0.27204741679052236
,,0,1970-01-01T16:00:00Z,m,0,f,0.6045056499409555
,,0,1970-01-01T17:00:00Z,m,0,f,0.21508343480255984
,,0,1970-01-01T18:00:00Z,m,0,f,0.2712545253017199
,,0,1970-01-01T19:00:00Z,m,0,f,0.22728191431845607
,,0,1970-01-01T20:00:00Z,m,0,f,0.8232481787306024
,,0,1970-01-01T21:00:00Z,m,0,f,0.9722054606060748
,,0,1970-01-01T22:00:00Z,m,0,f,0.9332942983017809
,,0,1970-01-01T23:00:00Z,m,0,f,0.009704805042322441
,,0,1970-01-02T00:00:00Z,m,0,f,0.4614776151185129
,,0,1970-01-02T01:00:00Z,m,0,f,0.3972854143424396
,,0,1970-01-02T02:00:00Z,m,0,f,0.024157782439736365
,,0,1970-01-02T03:00:00Z,m,0,f,0.7074351703076142
,,0,1970-01-02T04:00:00Z,m,0,f,0.5819899173941508
,,0,1970-01-02T05:00:00Z,m,0,f,0.2974899730817849
,,0,1970-01-02T06:00:00Z,m,0,f,0.3664899570202347
,,0,1970-01-02T07:00:00Z,m,0,f,0.5666625499409519
,,0,1970-01-02T08:00:00Z,m,0,f,0.2592658730352201
,,0,1970-01-02T09:00:00Z,m,0,f,0.6907206550112025
,,0,1970-01-02T10:00:00Z,m,0,f,0.7184801284027215
,,0,1970-01-02T11:00:00Z,m,0,f,0.363103986952813
,,0,1970-01-02T12:00:00Z,m,0,f,0.938825820840304
,,0,1970-01-02T13:00:00Z,m,0,f,0.7034638846507775
,,0,1970-01-02T14:00:00Z,m,0,f,0.5714903231820487
,,0,1970-01-02T15:00:00Z,m,0,f,0.24449047981396105
,,0,1970-01-02T16:00:00Z,m,0,f,0.14165037565843824
,,0,1970-01-02T17:00:00Z,m,0,f,0.05351135846151062
,,0,1970-01-02T18:00:00Z,m,0,f,0.3450781133356193
,,0,1970-01-02T19:00:00Z,m,0,f,0.23254297482426214
,,0,1970-01-02T20:00:00Z,m,0,f,0.15416851272541165
,,0,1970-01-02T21:00:00Z,m,0,f,0.9287113745228632
,,0,1970-01-02T22:00:00Z,m,0,f,0.8464406026410536
,,0,1970-01-02T23:00:00Z,m,0,f,0.7786237155792206
,,0,1970-01-03T00:00:00Z,m,0,f,0.7222630273842695
,,0,1970-01-03T01:00:00Z,m,0,f,0.5702856518144571
,,0,1970-01-03T02:00:00Z,m,0,f,0.4475020612540418
,,0,1970-01-03T03:00:00Z,m,0,f,0.19482413230523188
,,0,1970-01-03T04:00:00Z,m,0,f,0.14555100659831088
,,0,1970-01-03T05:00:00Z,m,0,f,0.3715313467677773
,,0,1970-01-03T06:00:00Z,m,0,f,0.15710124605981904
,,0,1970-01-03T07:00:00Z,m,0,f,0.05115366925369082
,,0,1970-01-03T08:00:00Z,m,0,f,0.49634673580304356
,,0,1970-01-03T09:00:00Z,m,0,f,0.09850492453963475
,,0,1970-01-03T10:00:00Z,m,0,f,0.07088528667647799
,,0,1970-01-03T11:00:00Z,m,0,f,0.9535958852850828
,,0,1970-01-03T12:00:00Z,m,0,f,0.9473123289831784
,,0,1970-01-03T13:00:00Z,m,0,f,0.6321990998686917
,,0,1970-01-03T14:00:00Z,m,0,f,0.5310985616209651
,,0,1970-01-03T15:00:00Z,m,0,f,0.14010236285353878
,,0,1970-01-03T16:00:00Z,m,0,f,0.5143111322693407
,,0,1970-01-03T17:00:00Z,m,0,f,0.1419555013503121
,,0,1970-01-03T18:00:00Z,m,0,f,0.034988171145264535
,,0,1970-01-03T19:00:00Z,m,0,f,0.4646423361131385
,,0,1970-01-03T20:00:00Z,m,0,f,0.7280775859440926
,,0,1970-01-03T21:00:00Z,m,0,f,0.9605223329866902
,,0,1970-01-03T22:00:00Z,m,0,f,0.6294671473626672
,,0,1970-01-03T23:00:00Z,m,0,f,0.09676486946771183
,,0,1970-01-04T00:00:00Z,m,0,f,0.4846624906255957
,,0,1970-01-04T01:00:00Z,m,0,f,0.9000151629241091
,,0,1970-01-04T02:00:00Z,m,0,f,0.8187520581651648
,,0,1970-01-04T03:00:00Z,m,0,f,0.6356479673253379
,,0,1970-01-04T04:00:00Z,m,0,f,0.9172292568869698
,,0,1970-01-04T05:00:00Z,m,0,f,0.25871413585674596
,,0,1970-01-04T06:00:00Z,m,0,f,0.934030201106989
,,0,1970-01-04T07:00:00Z,m,0,f,0.6300301521545785
,,0,1970-01-04T08:00:00Z,m,0,f,0.9898695895471914
,,0,1970-01-04T09:00:00Z,m,0,f,0.6576532850348832
,,0,1970-01-04T10:00:00Z,m,0,f,0.1095953745610317
,,0,1970-01-04T11:00:00Z,m,0,f,0.20714716664645624
,,0,1970-01-04T12:00:00Z,m,0,f,0.49378319061925324
,,0,1970-01-04T13:00:00Z,m,0,f,0.3244630221410883
,,0,1970-01-04T14:00:00Z,m,0,f,0.1425620337332085
,,0,1970-01-04T15:00:00Z,m,0,f,0.37483772088251627
,,0,1970-01-04T16:00:00Z,m,0,f,0.9386123621523778
,,0,1970-01-04T17:00:00Z,m,0,f,0.2944439301474122
,,0,1970-01-04T18:00:00Z,m,0,f,0.8075592894168399
,,0,1970-01-04T19:00:00Z,m,0,f,0.8131183413273094
,,0,1970-01-04T20:00:00Z,m,0,f,0.6056875144431602
,,0,1970-01-04T21:00:00Z,m,0,f,0.5514021237520469
,,0,1970-01-04T22:00:00Z,m,0,f,0.2904517561416824
,,0,1970-01-04T23:00:00Z,m,0,f,0.7773782053605
,,0,1970-01-05T00:00:00Z,m,0,f,0.1390732850129641
,,0,1970-01-05T01:00:00Z,m,0,f,0.36874812027455345
,,0,1970-01-05T02:00:00Z,m,0,f,0.8497133445947114
,,0,1970-01-05T03:00:00Z,m,0,f,0.2842281672817387
,,0,1970-01-05T04:00:00Z,m,0,f,0.5851186942712497
,,0,1970-01-05T05:00:00Z,m,0,f,0.2754694564842422
,,0,1970-01-05T06:00:00Z,m,0,f,0.03545539694267428
,,0,1970-01-05T07:00:00Z,m,0,f,0.4106208929295988
,,0,1970-01-05T08:00:00Z,m,0,f,0.3680257641839746
,,0,1970-01-05T09:00:00Z,m,0,f,0.7484477843640726
,,0,1970-01-05T10:00:00Z,m,0,f,0.2196945379224781
,,0,1970-01-05T11:00:00Z,m,0,f,0.7377409626382783
,,0,1970-01-05T12:00:00Z,m,0,f,0.4340408821652924
,,0,1970-01-05T13:00:00Z,m,0,f,0.04157784831355819
,,0,1970-01-05T14:00:00Z,m,0,f,0.9005324473445669
,,0,1970-01-05T15:00:00Z,m,0,f,0.6243062492954053
,,0,1970-01-05T16:00:00Z,m,0,f,0.4138274722170456
,,0,1970-01-05T17:00:00Z,m,0,f,0.6559961319794279
,,0,1970-01-05T18:00:00Z,m,0,f,0.09452730201881836
,,0,1970-01-05T19:00:00Z,m,0,f,0.35207875464289057
,,0,1970-01-05T20:00:00Z,m,0,f,0.47000290183266497
,,0,1970-01-05T21:00:00Z,m,0,f,0.13384008497720026
,,0,1970-01-05T22:00:00Z,m,0,f,0.2542495300083506
,,0,1970-01-05T23:00:00Z,m,0,f,0.04357411582677676
,,0,1970-01-06T00:00:00Z,m,0,f,0.2730770850239896
,,0,1970-01-06T01:00:00Z,m,0,f,0.07346719069503016
,,0,1970-01-06T02:00:00Z,m,0,f,0.19296870107837727
,,0,1970-01-06T03:00:00Z,m,0,f,0.8550701670111052
,,0,1970-01-06T04:00:00Z,m,0,f,0.9015279993379257
,,0,1970-01-06T05:00:00Z,m,0,f,0.7681329597853651
,,0,1970-01-06T06:00:00Z,m,0,f,0.13458582961527799
,,0,1970-01-06T07:00:00Z,m,0,f,0.5025964032341974
,,0,1970-01-06T08:00:00Z,m,0,f,0.9660611150198847
,,0,1970-01-06T09:00:00Z,m,0,f,0.7406756350132208
,,0,1970-01-06T10:00:00Z,m,0,f,0.48245323402069856
,,0,1970-01-06T11:00:00Z,m,0,f,0.5396866678590079
,,0,1970-01-06T12:00:00Z,m,0,f,0.24056787192459894
,,0,1970-01-06T13:00:00Z,m,0,f,0.5473495899891297
,,0,1970-01-06T14:00:00Z,m,0,f,0.9939487519980328
,,0,1970-01-06T15:00:00Z,m,0,f,0.7718086454038607
,,0,1970-01-06T16:00:00Z,m,0,f,0.3729231862915519
,,0,1970-01-06T17:00:00Z,m,0,f,0.978216628089757
,,0,1970-01-06T18:00:00Z,m,0,f,0.30410501498270626
,,0,1970-01-06T19:00:00Z,m,0,f,0.36293525766110357
,,0,1970-01-06T20:00:00Z,m,0,f,0.45673893698213724
,,0,1970-01-06T21:00:00Z,m,0,f,0.42887470039944864
,,0,1970-01-06T22:00:00Z,m,0,f,0.42264444401794515
,,0,1970-01-06T23:00:00Z,m,0,f,0.3061909271178175
,,0,1970-01-07T00:00:00Z,m,0,f,0.6681291175687905
,,0,1970-01-07T01:00:00Z,m,0,f,0.5494108420781338
,,0,1970-01-07T02:00:00Z,m,0,f,0.31779594303648045
,,0,1970-01-07T03:00:00Z,m,0,f,0.22502703712265368
,,0,1970-01-07T04:00:00Z,m,0,f,0.03498146847868716
,,0,1970-01-07T05:00:00Z,m,0,f,0.16139395876022747
,,0,1970-01-07T06:00:00Z,m,0,f,0.6335318955521227
,,0,1970-01-07T07:00:00Z,m,0,f,0.5854967453622169
,,0,1970-01-07T08:00:00Z,m,0,f,0.43015814365562627
,,0,1970-01-07T09:00:00Z,m,0,f,0.07215482648098204
,,0,1970-01-07T10:00:00Z,m,0,f,0.09348412983453618
,,0,1970-01-07T11:00:00Z,m,0,f,0.9023793546915768
,,0,1970-01-07T12:00:00Z,m,0,f,0.9055451292861832
,,0,1970-01-07T13:00:00Z,m,0,f,0.3280454144164272
,,0,1970-01-07T14:00:00Z,m,0,f,0.05897468763156862
,,0,1970-01-07T15:00:00Z,m,0,f,0.3686339026679373
,,0,1970-01-07T16:00:00Z,m,0,f,0.7547173975990482
,,0,1970-01-07T17:00:00Z,m,0,f,0.457847526142958
,,0,1970-01-07T18:00:00Z,m,0,f,0.5038320054556072
,,0,1970-01-07T19:00:00Z,m,0,f,0.47058145000588336
,,0,1970-01-07T20:00:00Z,m,0,f,0.5333903317331339
,,0,1970-01-07T21:00:00Z,m,0,f,0.1548508614296064
,,0,1970-01-07T22:00:00Z,m,0,f,0.6837681053869291
,,0,1970-01-07T23:00:00Z,m,0,f,0.9081953381867953
,,1,1970-01-01T00:00:00Z,m,1,f,0.15129694889144107
,,1,1970-01-01T01:00:00Z,m,1,f,0.18038761353721244
,,1,1970-01-01T02:00:00Z,m,1,f,0.23198629938985071
,,1,1970-01-01T03:00:00Z,m,1,f,0.4940776062344333
,,1,1970-01-01T04:00:00Z,m,1,f,0.5654050390735228
,,1,1970-01-01T05:00:00Z,m,1,f,0.3788291715942209
,,1,1970-01-01T06:00:00Z,m,1,f,0.39178743939497507
,,1,1970-01-01T07:00:00Z,m,1,f,0.573740997246541
,,1,1970-01-01T08:00:00Z,m,1,f,0.6171205083791419
,,1,1970-01-01T09:00:00Z,m,1,f,0.2562012267655005
,,1,1970-01-01T10:00:00Z,m,1,f,0.41301351982023743
,,1,1970-01-01T11:00:00Z,m,1,f,0.335808747696944
,,1,1970-01-01T12:00:00Z,m,1,f,0.25034171949067086
,,1,1970-01-01T13:00:00Z,m,1,f,0.9866289864317817
,,1,1970-01-01T14:00:00Z,m,1,f,0.42988399575215924
,,1,1970-01-01T15:00:00Z,m,1,f,0.02602624797587471
,,1,1970-01-01T16:00:00Z,m,1,f,0.9926232260423908
,,1,1970-01-01T17:00:00Z,m,1,f,0.9771153046566231
,,1,1970-01-01T18:00:00Z,m,1,f,0.5680196566957276
,,1,1970-01-01T19:00:00Z,m,1,f,0.01952645919207055
,,1,1970-01-01T20:00:00Z,m,1,f,0.3439692491089684
,,1,1970-01-01T21:00:00Z,m,1,f,0.15596143014601407
,,1,1970-01-01T22:00:00Z,m,1,f,0.7986983212658367
,,1,1970-01-01T23:00:00Z,m,1,f,0.31336565203700295
,,1,1970-01-02T00:00:00Z,m,1,f,0.6398281383647288
,,1,1970-01-02T01:00:00Z,m,1,f,0.14018673322595193
,,1,1970-01-02T02:00:00Z,m,1,f,0.2847409792344233
,,1,1970-01-02T03:00:00Z,m,1,f,0.4295460864480138
,,1,1970-01-02T04:00:00Z,m,1,f,0.9674016258565854
,,1,1970-01-02T05:00:00Z,m,1,f,0.108837862280129
,,1,1970-01-02T06:00:00Z,m,1,f,0.47129460971856907
,,1,1970-01-02T07:00:00Z,m,1,f,0.9175708860682784
,,1,1970-01-02T08:00:00Z,m,1,f,0.3383504562747057
,,1,1970-01-02T09:00:00Z,m,1,f,0.7176237840014899
,,1,1970-01-02T10:00:00Z,m,1,f,0.45631599181081023
,,1,1970-01-02T11:00:00Z,m,1,f,0.58210555704762
,,1,1970-01-02T12:00:00Z,m,1,f,0.44833346180841194
,,1,1970-01-02T13:00:00Z,m,1,f,0.847082665931482
,,1,1970-01-02T14:00:00Z,m,1,f,0.1032050849659337
,,1,1970-01-02T15:00:00Z,m,1,f,0.6342038875836871
,,1,1970-01-02T16:00:00Z,m,1,f,0.47157138392000586
,,1,1970-01-02T17:00:00Z,m,1,f,0.5939195811492147
,,1,1970-01-02T18:00:00Z,m,1,f,0.3907003938279841
,,1,1970-01-02T19:00:00Z,m,1,f,0.3737781066004461
,,1,1970-01-02T20:00:00Z,m,1,f,0.6059179847188622
,,1,1970-01-02T21:00:00Z,m,1,f,0.37459130316766875
,,1,1970-01-02T22:00:00Z,m,1,f,0.529020795101784
,,1,1970-01-02T23:00:00Z,m,1,f,0.5797965259387311
,,1,1970-01-03T00:00:00Z,m,1,f,0.4196060336001739
,,1,1970-01-03T01:00:00Z,m,1,f,0.4423826236661577
,,1,1970-01-03T02:00:00Z,m,1,f,0.7562185239602677
,,1,1970-01-03T03:00:00Z,m,1,f,0.29641000596052747
,,1,1970-01-03T04:00:00Z,m,1,f,0.5511866012217823
,,1,1970-01-03T05:00:00Z,m,1,f,0.477231168882557
,,1,1970-01-03T06:00:00Z,m,1,f,0.5783604476492074
,,1,1970-01-03T07:00:00Z,m,1,f,0.6087147255603924
,,1,1970-01-03T08:00:00Z,m,1,f,0.9779728651411874
,,1,1970-01-03T09:00:00Z,m,1,f,0.8559123961968673
,,1,1970-01-03T10:00:00Z,m,1,f,0.039322803759977897
,,1,1970-01-03T11:00:00Z,m,1,f,0.5107877963474311
,,1,1970-01-03T12:00:00Z,m,1,f,0.36939734036661503
,,1,1970-01-03T13:00:00Z,m,1,f,0.24036834333350818
,,1,1970-01-03T14:00:00Z,m,1,f,0.9041140297145132
,,1,1970-01-03T15:00:00Z,m,1,f,0.3088634061697057
,,1,1970-01-03T16:00:00Z,m,1,f,0.3391757217065211
,,1,1970-01-03T17:00:00Z,m,1,f,0.5709032014080667
,,1,1970-01-03T18:00:00Z,m,1,f,0.023692334151288443
,,1,1970-01-03T19:00:00Z,m,1,f,0.9283397254805887
,,1,1970-01-03T20:00:00Z,m,1,f,0.7897301020744532
,,1,1970-01-03T21:00:00Z,m,1,f,0.5499067643037981
,,1,1970-01-03T22:00:00Z,m,1,f,0.20359811467533634
,,1,1970-01-03T23:00:00Z,m,1,f,0.1946255400705282
,,1,1970-01-04T00:00:00Z,m,1,f,0.44702956746887096
,,1,1970-01-04T01:00:00Z,m,1,f,0.44634342940951505
,,1,1970-01-04T02:00:00Z,m,1,f,0.4462164964469759
,,1,1970-01-04T03:00:00Z,m,1,f,0.5245740015591633
,,1,1970-01-04T04:00:00Z,m,1,f,0.29252555227190247
,,1,1970-01-04T05:00:00Z,m,1,f,0.5137169576742285
,,1,1970-01-04T06:00:00Z,m,1,f,0.1624473579380766
,,1,1970-01-04T07:00:00Z,m,1,f,0.30153697909681254
,,1,1970-01-04T08:00:00Z,m,1,f,0.2324327035115191
,,1,1970-01-04T09:00:00Z,m,1,f,0.034393197916253775
,,1,1970-01-04T10:00:00Z,m,1,f,0.4336629996115634
,,1,1970-01-04T11:00:00Z,m,1,f,0.8790573703532555
,,1,1970-01-04T12:00:00Z,m,1,f,0.9016824143089478
,,1,1970-01-04T13:00:00Z,m,1,f,0.34003737969744235
,,1,1970-01-04T14:00:00Z,m,1,f,0.3848952908759773
,,1,1970-01-04T15:00:00Z,m,1,f,0.9951718603202089
,,1,1970-01-04T16:00:00Z,m,1,f,0.8567450174592717
,,1,1970-01-04T17:00:00Z,m,1,f,0.12389207874832112
,,1,1970-01-04T18:00:00Z,m,1,f,0.6712865769046611
,,1,1970-01-04T19:00:00Z,m,1,f,0.46454363710822305
,,1,1970-01-04T20:00:00Z,m,1,f,0.9625945392247928
,,1,1970-01-04T21:00:00Z,m,1,f,0.7535558804101941
,,1,1970-01-04T22:00:00Z,m,1,f,0.744281664085344
,,1,1970-01-04T23:00:00Z,m,1,f,0.6811372884190415
,,1,1970-01-05T00:00:00Z,m,1,f,0.46171144508557443
,,1,1970-01-05T01:00:00Z,m,1,f,0.7701860606472665
,,1,1970-01-05T02:00:00Z,m,1,f,0.25517367370396854
,,1,1970-01-05T03:00:00Z,m,1,f,0.5564394982112523
,,1,1970-01-05T04:00:00Z,m,1,f,0.18256039263141344
,,1,1970-01-05T05:00:00Z,m,1,f,0.08465044152492789
,,1,1970-01-05T06:00:00Z,m,1,f,0.04682876596739505
,,1,1970-01-05T07:00:00Z,m,1,f,0.5116535677666431
,,1,1970-01-05T08:00:00Z,m,1,f,0.26327513076438025
,,1,1970-01-05T09:00:00Z,m,1,f,0.8551637599549397
,,1,1970-01-05T10:00:00Z,m,1,f,0.04908769638903045
,,1,1970-01-05T11:00:00Z,m,1,f,0.6747954667852788
,,1,1970-01-05T12:00:00Z,m,1,f,0.6701210820394512
,,1,1970-01-05T13:00:00Z,m,1,f,0.6698146693971668
,,1,1970-01-05T14:00:00Z,m,1,f,0.32939712697857165
,,1,1970-01-05T15:00:00Z,m,1,f,0.788384711857412
,,1,1970-01-05T16:00:00Z,m,1,f,0.9435078647906675
,,1,1970-01-05T17:00:00Z,m,1,f,0.05526759807741008
,,1,1970-01-05T18:00:00Z,m,1,f,0.3040576381882256
,,1,1970-01-05T19:00:00Z,m,1,f,0.13057573237533082
,,1,1970-01-05T20:00:00Z,m,1,f,0.438829781443743
,,1,1970-01-05T21:00:00Z,m,1,f,0.16639381298657024
,,1,1970-01-05T22:00:00Z,m,1,f,0.17817868556539768
,,1,1970-01-05T23:00:00Z,m,1,f,0.37006948631938175
,,1,1970-01-06T00:00:00Z,m,1,f,0.7711386953356921
,,1,1970-01-06T01:00:00Z,m,1,f,0.37364593618845465
,,1,1970-01-06T02:00:00Z,m,1,f,0.9285996064937719
,,1,1970-01-06T03:00:00Z,m,1,f,0.8685918613936688
,,1,1970-01-06T04:00:00Z,m,1,f,0.049757835180659744
,,1,1970-01-06T05:00:00Z,m,1,f,0.3562051567466768
,,1,1970-01-06T06:00:00Z,m,1,f,0.9028928456702144
,,1,1970-01-06T07:00:00Z,m,1,f,0.45412719022597203
,,1,1970-01-06T08:00:00Z,m,1,f,0.5210991958721604
,,1,1970-01-06T09:00:00Z,m,1,f,0.5013716125947244
,,1,1970-01-06T10:00:00Z,m,1,f,0.7798859934672562
,,1,1970-01-06T11:00:00Z,m,1,f,0.20777334301449937
,,1,1970-01-06T12:00:00Z,m,1,f,0.12979889080684515
,,1,1970-01-06T13:00:00Z,m,1,f,0.6713165183217583
,,1,1970-01-06T14:00:00Z,m,1,f,0.5267649385791876
,,1,1970-01-06T15:00:00Z,m,1,f,0.2766996970172108
,,1,1970-01-06T16:00:00Z,m,1,f,0.837561303602128
,,1,1970-01-06T17:00:00Z,m,1,f,0.10692091027423688
,,1,1970-01-06T18:00:00Z,m,1,f,0.16161417900026617
,,1,1970-01-06T19:00:00Z,m,1,f,0.7596615857389895
,,1,1970-01-06T20:00:00Z,m,1,f,0.9033476318497203
,,1,1970-01-06T21:00:00Z,m,1,f,0.9281794553091864
,,1,1970-01-06T22:00:00Z,m,1,f,0.7691815845690406
,,1,1970-01-06T23:00:00Z,m,1,f,0.5713941284458292
,,1,1970-01-07T00:00:00Z,m,1,f,0.8319045908167892
,,1,1970-01-07T01:00:00Z,m,1,f,0.5839200214729727
,,1,1970-01-07T02:00:00Z,m,1,f,0.5597883274306116
,,1,1970-01-07T03:00:00Z,m,1,f,0.8448107197504592
,,1,1970-01-07T04:00:00Z,m,1,f,0.39141999130543037
,,1,1970-01-07T05:00:00Z,m,1,f,0.3151057211763145
,,1,1970-01-07T06:00:00Z,m,1,f,0.3812489036241129
,,1,1970-01-07T07:00:00Z,m,1,f,0.03893545284960627
,,1,1970-01-07T08:00:00Z,m,1,f,0.513934438417237
,,1,1970-01-07T09:00:00Z,m,1,f,0.07387412770693513
,,1,1970-01-07T10:00:00Z,m,1,f,0.16131994851623296
,,1,1970-01-07T11:00:00Z,m,1,f,0.8524873225734262
,,1,1970-01-07T12:00:00Z,m,1,f,0.7108229805824855
,,1,1970-01-07T13:00:00Z,m,1,f,0.4087372331379091
,,1,1970-01-07T14:00:00Z,m,1,f,0.5408493060971712
,,1,1970-01-07T15:00:00Z,m,1,f,0.8752116934130074
,,1,1970-01-07T16:00:00Z,m,1,f,0.9569196248412628
,,1,1970-01-07T17:00:00Z,m,1,f,0.5206668595695829
,,1,1970-01-07T18:00:00Z,m,1,f,0.012847952493292788
,,1,1970-01-07T19:00:00Z,m,1,f,0.7155605509853933
,,1,1970-01-07T20:00:00Z,m,1,f,0.8293273149090988
,,1,1970-01-07T21:00:00Z,m,1,f,0.38705272903958904
,,1,1970-01-07T22:00:00Z,m,1,f,0.5459991408731746
,,1,1970-01-07T23:00:00Z,m,1,f,0.7066840478612406
"
outData =
    "
#datatype,string,long,dateTime:RFC3339,string,string,double
#group,false,false,false,true,true,false
#default,0,,,,,
,result,table,_time,_measurement,t0,difference
,,0,1970-01-01T01:00:00Z,m,0,0.16152781154936718
,,0,1970-01-01T02:00:00Z,m,0,0.5450233503637238
,,0,1970-01-01T03:00:00Z,m,0,-0.25474251334078146
,,0,1970-01-01T04:00:00Z,m,0,-0.5121283372090074
,,0,1970-01-01T05:00:00Z,m,0,0.171070028248751
,,0,1970-01-01T06:00:00Z,m,0,-0.13711438955681704
,,0,1970-01-01T07:00:00Z,m,0,0.5180121760612726
,,0,1970-01-01T08:00:00Z,m,0,-0.30465284267416787
,,0,1970-01-01T09:00:00Z,m,0,-0.003939225414167302
,,0,1970-01-01T10:00:00Z,m,0,-0.11037649542590938
,,0,1970-01-01T11:00:00Z,m,0,-0.06844878023750281
,,0,1970-01-01T12:00:00Z,m,0,0.5940685767143397
,,0,1970-01-01T13:00:00Z,m,0,0.056202322240399005
,,0,1970-01-01T14:00:00Z,m,0,-0.6528142877331138
,,0,1970-01-01T15:00:00Z,m,0,0.07601807321180445
,,0,1970-01-01T16:00:00Z,m,0,0.3324582331504331
,,0,1970-01-01T17:00:00Z,m,0,-0.3894222151383956
,,0,1970-01-01T18:00:00Z,m,0,0.056171090499160053
,,0,1970-01-01T19:00:00Z,m,0,-0.04397261098326383
,,0,1970-01-01T20:00:00Z,m,0,0.5959662644121464
,,0,1970-01-01T21:00:00Z,m,0,0.1489572818754724
,,0,1970-01-01T22:00:00Z,m,0,-0.03891116230429392
,,0,1970-01-01T23:00:00Z,m,0,-0.9235894932594585
,,0,1970-01-02T00:00:00Z,m,0,0.4517728100761904
,,0,1970-01-02T01:00:00Z,m,0,-0.0641922007760733
,,0,1970-01-02T02:00:00Z,m,0,-0.3731276319027032
,,0,1970-01-02T03:00:00Z,m,0,0.6832773878678778
,,0,1970-01-02T04:00:00Z,m,0,-0.1254452529134633
,,0,1970-01-02T05:00:00Z,m,0,-0.28449994431236597
,,0,1970-01-02T06:00:00Z,m,0,0.0689999839384498
,,0,1970-01-02T07:00:00Z,m,0,0.2001725929207172
,,0,1970-01-02T08:00:00Z,m,0,-0.3073966769057318
,,0,1970-01-02T09:00:00Z,m,0,0.4314547819759824
,,0,1970-01-02T10:00:00Z,m,0,0.027759473391518963
,,0,1970-01-02T11:00:00Z,m,0,-0.3553761414499085
,,0,1970-01-02T12:00:00Z,m,0,0.575721833887491
,,0,1970-01-02T13:00:00Z,m,0,-0.2353619361895265
,,0,1970-01-02T14:00:00Z,m,0,-0.13197356146872874
,,0,1970-01-02T15:00:00Z,m,0,-0.3269998433680877
,,0,1970-01-02T16:00:00Z,m,0,-0.1028401041555228
,,0,1970-01-02T17:00:00Z,m,0,-0.08813901719692763
,,0,1970-01-02T18:00:00Z,m,0,0.29156675487410866
,,0,1970-01-02T19:00:00Z,m,0,-0.11253513851135713
,,0,1970-01-02T20:00:00Z,m,0,-0.07837446209885049
,,0,1970-01-02T21:00:00Z,m,0,0.7745428617974516
,,0,1970-01-02T22:00:00Z,m,0,-0.08227077188180965
,,0,1970-01-02T23:00:00Z,m,0,-0.06781688706183298
,,0,1970-01-03T00:00:00Z,m,0,-0.0563606881949511
,,0,1970-01-03T01:00:00Z,m,0,-0.1519773755698124
,,0,1970-01-03T02:00:00Z,m,0,-0.12278359056041532
,,0,1970-01-03T03:00:00Z,m,0,-0.2526779289488099
,,0,1970-01-03T04:00:00Z,m,0,-0.049273125706920995
,,0,1970-01-03T05:00:00Z,m,0,0.22598034016946644
,,0,1970-01-03T06:00:00Z,m,0,-0.2144301007079583
,,0,1970-01-03T07:00:00Z,m,0,-0.10594757680612822
,,0,1970-01-03T08:00:00Z,m,0,0.44519306654935276
,,0,1970-01-03T09:00:00Z,m,0,-0.3978418112634088
,,0,1970-01-03T10:00:00Z,m,0,-0.027619637863156757
,,0,1970-01-03T11:00:00Z,m,0,0.8827105986086048
,,0,1970-01-03T12:00:00Z,m,0,-0.006283556301904358
,,0,1970-01-03T13:00:00Z,m,0,-0.3151132291144867
,,0,1970-01-03T14:00:00Z,m,0,-0.10110053824772658
,,0,1970-01-03T15:00:00Z,m,0,-0.39099619876742636
,,0,1970-01-03T16:00:00Z,m,0,0.3742087694158019
,,0,1970-01-03T17:00:00Z,m,0,-0.37235563091902857
,,0,1970-01-03T18:00:00Z,m,0,-0.10696733020504756
,,0,1970-01-03T19:00:00Z,m,0,0.42965416496787395
,,0,1970-01-03T20:00:00Z,m,0,0.26343524983095407
,,0,1970-01-03T21:00:00Z,m,0,0.23244474704259765
,,0,1970-01-03T22:00:00Z,m,0,-0.33105518562402303
,,0,1970-01-03T23:00:00Z,m,0,-0.5327022778949554
,,0,1970-01-04T00:00:00Z,m,0,0.38789762115788384
,,0,1970-01-04T01:00:00Z,m,0,0.4153526722985134
,,0,1970-01-04T02:00:00Z,m,0,-0.08126310475894427
,,0,1970-01-04T03:00:00Z,m,0,-0.18310409083982693
,,0,1970-01-04T04:00:00Z,m,0,0.28158128956163186
,,0,1970-01-04T05:00:00Z,m,0,-0.6585151210302238
,,0,1970-01-04T06:00:00Z,m,0,0.675316065250243
,,0,1970-01-04T07:00:00Z,m,0,-0.3040000489524105
,,0,1970-01-04T08:00:00Z,m,0,0.3598394373926129
,,0,1970-01-04T09:00:00Z,m,0,-0.3322163045123082
,,0,1970-01-04T10:00:00Z,m,0,-0.5480579104738514
,,0,1970-01-04T11:00:00Z,m,0,0.09755179208542454
,,0,1970-01-04T12:00:00Z,m,0,0.28663602397279697
,,0,1970-01-04T13:00:00Z,m,0,-0.16932016847816495
,,0,1970-01-04T14:00:00Z,m,0,-0.1819009884078798
,,0,1970-01-04T15:00:00Z,m,0,0.23227568714930777
,,0,1970-01-04T16:00:00Z,m,0,0.5637746412698614
,,0,1970-01-04T17:00:00Z,m,0,-0.6441684320049655
,,0,1970-01-04T18:00:00Z,m,0,0.5131153592694276
,,0,1970-01-04T19:00:00Z,m,0,0.00555905191046957
,,0,1970-01-04T20:00:00Z,m,0,-0.20743082688414927
,,0,1970-01-04T21:00:00Z,m,0,-0.05428539069111327
,,0,1970-01-04T22:00:00Z,m,0,-0.2609503676103645
,,0,1970-01-04T23:00:00Z,m,0,0.48692644921881756
,,0,1970-01-05T00:00:00Z,m,0,-0.6383049203475358
,,0,1970-01-05T01:00:00Z,m,0,0.22967483526158936
,,0,1970-01-05T02:00:00Z,m,0,0.4809652243201579
,,0,1970-01-05T03:00:00Z,m,0,-0.5654851773129727
,,0,1970-01-05T04:00:00Z,m,0,0.300890526989511
,,0,1970-01-05T05:00:00Z,m,0,-0.30964923778700754
,,0,1970-01-05T06:00:00Z,m,0,-0.24001405954156793
,,0,1970-01-05T07:00:00Z,m,0,0.3751654959869245
,,0,1970-01-05T08:00:00Z,m,0,-0.04259512874562421
,,0,1970-01-05T09:00:00Z,m,0,0.38042202018009796
,,0,1970-01-05T10:00:00Z,m,0,-0.5287532464415945
,,0,1970-01-05T11:00:00Z,m,0,0.5180464247158002
,,0,1970-01-05T12:00:00Z,m,0,-0.30370008047298586
,,0,1970-01-05T13:00:00Z,m,0,-0.3924630338517342
,,0,1970-01-05T14:00:00Z,m,0,0.8589545990310088
,,0,1970-01-05T15:00:00Z,m,0,-0.2762261980491616
,,0,1970-01-05T16:00:00Z,m,0,-0.21047877707835977
,,0,1970-01-05T17:00:00Z,m,0,0.24216865976238228
,,0,1970-01-05T18:00:00Z,m,0,-0.5614688299606095
,,0,1970-01-05T19:00:00Z,m,0,0.2575514526240722
,,0,1970-01-05T20:00:00Z,m,0,0.1179241471897744
,,0,1970-01-05T21:00:00Z,m,0,-0.3361628168554647
,,0,1970-01-05T22:00:00Z,m,0,0.12040944503115034
,,0,1970-01-05T23:00:00Z,m,0,-0.21067541418157384
,,0,1970-01-06T00:00:00Z,m,0,0.22950296919721283
,,0,1970-01-06T01:00:00Z,m,0,-0.19960989432895943
,,0,1970-01-06T02:00:00Z,m,0,0.11950151038334711
,,0,1970-01-06T03:00:00Z,m,0,0.6621014659327279
,,0,1970-01-06T04:00:00Z,m,0,0.04645783232682055
,,0,1970-01-06T05:00:00Z,m,0,-0.1333950395525606
,,0,1970-01-06T06:00:00Z,m,0,-0.6335471301700871
,,0,1970-01-06T07:00:00Z,m,0,0.3680105736189194
,,0,1970-01-06T08:00:00Z,m,0,0.46346471178568727
,,0,1970-01-06T09:00:00Z,m,0,-0.22538548000666392
,,0,1970-01-06T10:00:00Z,m,0,-0.2582224009925222
,,0,1970-01-06T11:00:00Z,m,0,0.05723343383830931
,,0,1970-01-06T12:00:00Z,m,0,-0.29911879593440893
,,0,1970-01-06T13:00:00Z,m,0,0.3067817180645308
,,0,1970-01-06T14:00:00Z,m,0,0.44659916200890304
,,0,1970-01-06T15:00:00Z,m,0,-0.22214010659417205
,,0,1970-01-06T16:00:00Z,m,0,-0.3988854591123088
,,0,1970-01-06T17:00:00Z,m,0,0.605293441798205
,,0,1970-01-06T18:00:00Z,m,0,-0.6741116131070507
,,0,1970-01-06T19:00:00Z,m,0,0.058830242678397315
,,0,1970-01-06T20:00:00Z,m,0,0.09380367932103367
,,0,1970-01-06T21:00:00Z,m,0,-0.027864236582688606
,,0,1970-01-06T22:00:00Z,m,0,-0.006230256381503485
,,0,1970-01-06T23:00:00Z,m,0,-0.11645351690012767
,,0,1970-01-07T00:00:00Z,m,0,0.361938190450973
,,0,1970-01-07T01:00:00Z,m,0,-0.11871827549065661
,,0,1970-01-07T02:00:00Z,m,0,-0.2316148990416534
,,0,1970-01-07T03:00:00Z,m,0,-0.09276890591382678
,,0,1970-01-07T04:00:00Z,m,0,-0.19004556864396652
,,0,1970-01-07T05:00:00Z,m,0,0.1264124902815403
,,0,1970-01-07T06:00:00Z,m,0,0.4721379367918953
,,0,1970-01-07T07:00:00Z,m,0,-0.048035150189905895
,,0,1970-01-07T08:00:00Z,m,0,-0.15533860170659058
,,0,1970-01-07T09:00:00Z,m,0,-0.35800331717464423
,,0,1970-01-07T10:00:00Z,m,0,0.021329303353554138
,,0,1970-01-07T11:00:00Z,m,0,0.8088952248570407
,,0,1970-01-07T12:00:00Z,m,0,0.0031657745946064297
,,0,1970-01-07T13:00:00Z,m,0,-0.577499714869756
,,0,1970-01-07T14:00:00Z,m,0,-0.2690707267848586
,,0,1970-01-07T15:00:00Z,m,0,0.3096592150363687
,,0,1970-01-07T16:00:00Z,m,0,0.38608349493111094
,,0,1970-01-07T17:00:00Z,m,0,-0.29686987145609023
,,0,1970-01-07T18:00:00Z,m,0,0.04598447931264926
,,0,1970-01-07T19:00:00Z,m,0,-0.033250555449723884
,,0,1970-01-07T20:00:00Z,m,0,0.06280888172725052
,,0,1970-01-07T21:00:00Z,m,0,-0.37853947030352747
,,0,1970-01-07T22:00:00Z,m,0,0.5289172439573226
,,0,1970-01-07T23:00:00Z,m,0,0.2244272327998662
,,1,1970-01-01T01:00:00Z,m,1,0.02909066464577137
,,1,1970-01-01T02:00:00Z,m,1,0.05159868585263827
,,1,1970-01-01T03:00:00Z,m,1,0.2620913068445826
,,1,1970-01-01T04:00:00Z,m,1,0.07132743283908949
,,1,1970-01-01T05:00:00Z,m,1,-0.18657586747930188
,,1,1970-01-01T06:00:00Z,m,1,0.012958267800754153
,,1,1970-01-01T07:00:00Z,m,1,0.18195355785156597
,,1,1970-01-01T08:00:00Z,m,1,0.043379511132600856
,,1,1970-01-01T09:00:00Z,m,1,-0.3609192816136414
,,1,1970-01-01T10:00:00Z,m,1,0.15681229305473693
,,1,1970-01-01T11:00:00Z,m,1,-0.07720477212329341
,,1,1970-01-01T12:00:00Z,m,1,-0.08546702820627317
,,1,1970-01-01T13:00:00Z,m,1,0.7362872669411108
,,1,1970-01-01T14:00:00Z,m,1,-0.5567449906796225
,,1,1970-01-01T15:00:00Z,m,1,-0.4038577477762845
,,1,1970-01-01T16:00:00Z,m,1,0.9665969780665161
,,1,1970-01-01T17:00:00Z,m,1,-0.015507921385767731
,,1,1970-01-01T18:00:00Z,m,1,-0.4090956479608955
,,1,1970-01-01T19:00:00Z,m,1,-0.5484931975036571
,,1,1970-01-01T20:00:00Z,m,1,0.3244427899168979
,,1,1970-01-01T21:00:00Z,m,1,-0.18800781896295435
,,1,1970-01-01T22:00:00Z,m,1,0.6427368911198227
,,1,1970-01-01T23:00:00Z,m,1,-0.4853326692288338
,,1,1970-01-02T00:00:00Z,m,1,0.3264624863277259
,,1,1970-01-02T01:00:00Z,m,1,-0.4996414051387769
,,1,1970-01-02T02:00:00Z,m,1,0.1445542460084714
,,1,1970-01-02T03:00:00Z,m,1,0.14480510721359047
,,1,1970-01-02T04:00:00Z,m,1,0.5378555394085716
,,1,1970-01-02T05:00:00Z,m,1,-0.8585637635764564
,,1,1970-01-02T06:00:00Z,m,1,0.3624567474384401
,,1,1970-01-02T07:00:00Z,m,1,0.44627627634970934
,,1,1970-01-02T08:00:00Z,m,1,-0.5792204297935727
,,1,1970-01-02T09:00:00Z,m,1,0.3792733277267841
,,1,1970-01-02T10:00:00Z,m,1,-0.2613077921906796
,,1,1970-01-02T11:00:00Z,m,1,0.12578956523680979
,,1,1970-01-02T12:00:00Z,m,1,-0.13377209523920808
,,1,1970-01-02T13:00:00Z,m,1,0.39874920412307
,,1,1970-01-02T14:00:00Z,m,1,-0.7438775809655482
,,1,1970-01-02T15:00:00Z,m,1,0.5309988026177533
,,1,1970-01-02T16:00:00Z,m,1,-0.1626325036636812
,,1,1970-01-02T17:00:00Z,m,1,0.12234819722920887
,,1,1970-01-02T18:00:00Z,m,1,-0.20321918732123062
,,1,1970-01-02T19:00:00Z,m,1,-0.016922287227538024
,,1,1970-01-02T20:00:00Z,m,1,0.23213987811841608
,,1,1970-01-02T21:00:00Z,m,1,-0.2313266815511934
,,1,1970-01-02T22:00:00Z,m,1,0.15442949193411526
,,1,1970-01-02T23:00:00Z,m,1,0.050775730836947086
,,1,1970-01-03T00:00:00Z,m,1,-0.1601904923385572
,,1,1970-01-03T01:00:00Z,m,1,0.022776590065983815
,,1,1970-01-03T02:00:00Z,m,1,0.31383590029410996
,,1,1970-01-03T03:00:00Z,m,1,-0.4598085179997402
,,1,1970-01-03T04:00:00Z,m,1,0.25477659526125485
,,1,1970-01-03T05:00:00Z,m,1,-0.0739554323392253
,,1,1970-01-03T06:00:00Z,m,1,0.10112927876665034
,,1,1970-01-03T07:00:00Z,m,1,0.030354277911185057
,,1,1970-01-03T08:00:00Z,m,1,0.369258139580795
,,1,1970-01-03T09:00:00Z,m,1,-0.1220604689443201
,,1,1970-01-03T10:00:00Z,m,1,-0.8165895924368894
,,1,1970-01-03T11:00:00Z,m,1,0.4714649925874532
,,1,1970-01-03T12:00:00Z,m,1,-0.14139045598081607
,,1,1970-01-03T13:00:00Z,m,1,-0.12902899703310686
,,1,1970-01-03T14:00:00Z,m,1,0.663745686381005
,,1,1970-01-03T15:00:00Z,m,1,-0.5952506235448074
,,1,1970-01-03T16:00:00Z,m,1,0.030312315536815404
,,1,1970-01-03T17:00:00Z,m,1,0.2317274797015456
,,1,1970-01-03T18:00:00Z,m,1,-0.5472108672567783
,,1,1970-01-03T19:00:00Z,m,1,0.9046473913293003
,,1,1970-01-03T20:00:00Z,m,1,-0.13860962340613558
,,1,1970-01-03T21:00:00Z,m,1,-0.239823337770655
,,1,1970-01-03T22:00:00Z,m,1,-0.3463086496284618
,,1,1970-01-03T23:00:00Z,m,1,-0.008972574604808131
,,1,1970-01-04T00:00:00Z,m,1,0.2524040273983428
,,1,1970-01-04T01:00:00Z,m,1,-0.0006861380593559119
,,1,1970-01-04T02:00:00Z,m,1,-0.000126932962539128
,,1,1970-01-04T03:00:00Z,m,1,0.07835750511218742
,,1,1970-01-04T04:00:00Z,m,1,-0.23204844928726087
,,1,1970-01-04T05:00:00Z,m,1,0.22119140540232607
,,1,1970-01-04T06:00:00Z,m,1,-0.3512695997361519
,,1,1970-01-04T07:00:00Z,m,1,0.13908962115873594
,,1,1970-01-04T08:00:00Z,m,1,-0.06910427558529345
,,1,1970-01-04T09:00:00Z,m,1,-0.1980395055952653
,,1,1970-01-04T10:00:00Z,m,1,0.3992698016953096
,,1,1970-01-04T11:00:00Z,m,1,0.4453943707416921
,,1,1970-01-04T12:00:00Z,m,1,0.022625043955692314
,,1,1970-01-04T13:00:00Z,m,1,-0.5616450346115054
,,1,1970-01-04T14:00:00Z,m,1,0.04485791117853494
,,1,1970-01-04T15:00:00Z,m,1,0.6102765694442316
,,1,1970-01-04T16:00:00Z,m,1,-0.1384268428609372
,,1,1970-01-04T17:00:00Z,m,1,-0.7328529387109506
,,1,1970-01-04T18:00:00Z,m,1,0.54739449815634
,,1,1970-01-04T19:00:00Z,m,1,-0.20674293979643804
,,1,1970-01-04T20:00:00Z,m,1,0.4980509021165697
,,1,1970-01-04T21:00:00Z,m,1,-0.2090386588145987
,,1,1970-01-04T22:00:00Z,m,1,-0.009274216324850038
,,1,1970-01-04T23:00:00Z,m,1,-0.06314437566630249
,,1,1970-01-05T00:00:00Z,m,1,-0.2194258433334671
,,1,1970-01-05T01:00:00Z,m,1,0.30847461556169203
,,1,1970-01-05T02:00:00Z,m,1,-0.5150123869432979
,,1,1970-01-05T03:00:00Z,m,1,0.30126582450728373
,,1,1970-01-05T04:00:00Z,m,1,-0.37387910557983883
,,1,1970-01-05T05:00:00Z,m,1,-0.09790995110648555
,,1,1970-01-05T06:00:00Z,m,1,-0.03782167555753284
,,1,1970-01-05T07:00:00Z,m,1,0.464824801799248
,,1,1970-01-05T08:00:00Z,m,1,-0.2483784370022628
,,1,1970-01-05T09:00:00Z,m,1,0.5918886291905594
,,1,1970-01-05T10:00:00Z,m,1,-0.8060760635659092
,,1,1970-01-05T11:00:00Z,m,1,0.6257077703962484
,,1,1970-01-05T12:00:00Z,m,1,-0.004674384745827598
,,1,1970-01-05T13:00:00Z,m,1,-0.0003064126422843705
,,1,1970-01-05T14:00:00Z,m,1,-0.3404175424185952
,,1,1970-01-05T15:00:00Z,m,1,0.45898758487884034
,,1,1970-01-05T16:00:00Z,m,1,0.15512315293325551
,,1,1970-01-05T17:00:00Z,m,1,-0.8882402667132574
,,1,1970-01-05T18:00:00Z,m,1,0.24879004011081554
,,1,1970-01-05T19:00:00Z,m,1,-0.1734819058128948
,,1,1970-01-05T20:00:00Z,m,1,0.3082540490684122
,,1,1970-01-05T21:00:00Z,m,1,-0.2724359684571728
,,1,1970-01-05T22:00:00Z,m,1,0.011784872578827432
,,1,1970-01-05T23:00:00Z,m,1,0.19189080075398407
,,1,1970-01-06T00:00:00Z,m,1,0.4010692090163104
,,1,1970-01-06T01:00:00Z,m,1,-0.3974927591472375
,,1,1970-01-06T02:00:00Z,m,1,0.5549536703053173
,,1,1970-01-06T03:00:00Z,m,1,-0.060007745100103094
,,1,1970-01-06T04:00:00Z,m,1,-0.818834026213009
,,1,1970-01-06T05:00:00Z,m,1,0.30644732156601706
,,1,1970-01-06T06:00:00Z,m,1,0.5466876889235377
,,1,1970-01-06T07:00:00Z,m,1,-0.4487656554442424
,,1,1970-01-06T08:00:00Z,m,1,0.06697200564618833
,,1,1970-01-06T09:00:00Z,m,1,-0.019727583277435956
,,1,1970-01-06T10:00:00Z,m,1,0.2785143808725318
,,1,1970-01-06T11:00:00Z,m,1,-0.5721126504527568
,,1,1970-01-06T12:00:00Z,m,1,-0.07797445220765423
,,1,1970-01-06T13:00:00Z,m,1,0.5415176275149132
,,1,1970-01-06T14:00:00Z,m,1,-0.1445515797425707
,,1,1970-01-06T15:00:00Z,m,1,-0.2500652415619768
,,1,1970-01-06T16:00:00Z,m,1,0.5608616065849172
,,1,1970-01-06T17:00:00Z,m,1,-0.7306403933278911
,,1,1970-01-06T18:00:00Z,m,1,0.054693268726029295
,,1,1970-01-06T19:00:00Z,m,1,0.5980474067387233
,,1,1970-01-06T20:00:00Z,m,1,0.14368604611073088
,,1,1970-01-06T21:00:00Z,m,1,0.024831823459466107
,,1,1970-01-06T22:00:00Z,m,1,-0.15899787074014582
,,1,1970-01-06T23:00:00Z,m,1,-0.19778745612321147
,,1,1970-01-07T00:00:00Z,m,1,0.2605104623709601
,,1,1970-01-07T01:00:00Z,m,1,-0.24798456934381652
,,1,1970-01-07T02:00:00Z,m,1,-0.02413169404236115
,,1,1970-01-07T03:00:00Z,m,1,0.2850223923198476
,,1,1970-01-07T04:00:00Z,m,1,-0.4533907284450288
,,1,1970-01-07T05:00:00Z,m,1,-0.07631427012911585
,,1,1970-01-07T06:00:00Z,m,1,0.06614318244779838
,,1,1970-01-07T07:00:00Z,m,1,-0.34231345077450664
,,1,1970-01-07T08:00:00Z,m,1,0.4749989855676307
,,1,1970-01-07T09:00:00Z,m,1,-0.4400603107103018
,,1,1970-01-07T10:00:00Z,m,1,0.08744582080929783
,,1,1970-01-07T11:00:00Z,m,1,0.6911673740571932
,,1,1970-01-07T12:00:00Z,m,1,-0.1416643419909407
,,1,1970-01-07T13:00:00Z,m,1,-0.3020857474445764
,,1,1970-01-07T14:00:00Z,m,1,0.13211207295926208
,,1,1970-01-07T15:00:00Z,m,1,0.33436238731583623
,,1,1970-01-07T16:00:00Z,m,1,0.08170793142825539
,,1,1970-01-07T17:00:00Z,m,1,-0.4362527652716799
,,1,1970-01-07T18:00:00Z,m,1,-0.5078189070762902
,,1,1970-01-07T19:00:00Z,m,1,0.7027125984921005
,,1,1970-01-07T20:00:00Z,m,1,0.11376676392370555
,,1,1970-01-07T21:00:00Z,m,1,-0.4422745858695098
,,1,1970-01-07T22:00:00Z,m,1,0.15894641183358554
,,1,1970-01-07T23:00:00Z,m,1,0.16068490698806603
"

testcase difference {
    got =
        csv.from(csv: inData)
            |> testing.load()
            |> range(start: influxql.minTime, stop: influxql.maxTime)
            |> filter(fn: (r) => r._measurement == "m")
            |> filter(fn: (r) => r._field == "f")
            |> difference()
            |> drop(columns: ["_start", "_stop", "_field"])
            |> rename(columns: {_value: "difference"})
    want = csv.from(csv: outData)

    testing.diff(got, want)
}
