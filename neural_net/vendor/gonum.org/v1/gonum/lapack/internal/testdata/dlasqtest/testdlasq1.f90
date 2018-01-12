program randomsys1
implicit none
integer, parameter :: nmax=1000
real(kind=8), dimension(nmax) :: b, X
real(kind=8), dimension(nmax,nmax) :: a
real(kind=8) :: err
integer :: i, info, lda, ldb, nrhs, n,iter
integer, dimension(nmax) :: ipiv

real(kind=8), dimension(100) :: d
real(kind=8), dimension(99) :: e
real(kind=8), dimension(400) :: work

d(1:100) = (/1.8334043365537367D+00, &
1.4451749896846686D+00, &
1.0018566447551758D-01, &
-7.2143260438744417D-01, &
-3.7864653015502087D-01, &
-9.0270111568850808D-01, &
1.2204305489831029D+00, &
-9.7628177811136485D-01, &
8.4199233511256721D-01, &
-2.7938817329922050D-01, &
3.6157779599908046D-01, &
-1.8563411313998144D+00, &
-5.7930081140992240D-01, &
7.4080550463379169D-01, &
1.7021409147402005D+00, &
-5.7992035328196923D-01, &
4.0877426434579855D-01, &
-7.1297236049446144D-01, &
-1.2214095798914903D+00, &
2.9037983248746674D-01, &
7.4685018821608473D-01, &
3.0213735481316539D-01, &
-1.5207207136907624D-01, &
-2.1332671668411556D+00, &
6.8744661834930676D-01, &
-2.0946670404018297D-01, &
-1.5221059713957628D+00, &
1.1117190383859539D+00, &
-6.1175948159744020D-01, &
-4.4149212620857964D-01, &
-5.5702632609947533D-01, &
1.4071858950692646D+00, &
-3.2329881667362437D-01, &
-3.1958092104323499D-01, &
9.0934520529412111D-01, &
9.7881421364746712D-01, &
-5.1202970940327841D-01, &
1.5040024724520102D+00, &
-7.1993831181468571D-01, &
-7.1819661000094503D-01, &
-1.3481185445933910D+00, &
-1.4984426192966893D+00, &
1.1356626654278745D+00, &
1.6427417967661164D+00, &
-1.4184643787388000D+00, &
2.9811560271518989D-01, &
7.8630022575860559D-01, &
-1.8262830018047089D+00, &
6.3058255632564841D-01, &
-2.0692752940382309D-02, &
-7.2726648905906033D-01, &
-1.0461446937034022D+00, &
1.2530345094987356D+00, &
-2.3583665341168443D+00, &
1.9177654334479410D-01, &
-1.3563410975095058D+00, &
-1.0669710425337906D+00, &
1.4840245472507219D+00, &
-6.9185935518981789D-01, &
1.6813910559942205D-01, &
-7.1255209442204559D-01, &
-1.0112797453604008D+00, &
2.8591746998403011D-01, &
-1.9403039239509563D+00, &
-8.1434141084858885D-02, &
1.3873918713367210D+00, &
-8.8212258376548647D-01, &
-1.2253510598547583D+00, &
-3.8677386127356073D-01, &
-1.0262656526479850D+00, &
2.9468734022014376D-01, &
2.3442965677966704D-01, &
1.2219251054024911D+00, &
2.6066505150099868D+00, &
-7.8543147636303856D-01, &
-9.8126277325503253D-01, &
1.1722358680271947D+00, &
-8.1477749181289072D-01, &
1.3437777060446568D-01, &
3.4626341297821356D-01, &
-4.5672026157532375D-01, &
3.0496975959999184D-01, &
3.4636683737604146D-01, &
1.5114807806635011D-01, &
-1.1376177393945328D+00, &
9.3419670621891793D-01, &
7.9186416310239138D-01, &
6.7230421440462595D-01, &
-2.3236847424852280D-01, &
-1.0927137499922757D+00, &
9.8562766620822340D-01, &
-1.1382935431007701D-01, &
-9.2072237463768225D-01, &
6.1142850054965170D-01, &
8.2752057022739134D-03, &
6.7197122515126417D-01, &
-1.1007816668204429D+00, &
-1.2196829073430047D+00, &
-6.1428585523321222D-01, &
6.4599803465517280D-01/)
e(1:99) = (/-9.6393084631802151D-01, &
2.5532567563781607D+00, &
8.2492664240014357D-01, &
8.2628261757058474D-01, &
7.3021658879350859D-01, &
3.4014431785419519D-02, &
3.2121571974900542D-01, &
2.5987166374572213D-02, &
-3.1150206355945814D-01, &
1.3429094629249927D+00, &
8.6246434952180806D-01, &
-8.3756967113851388D-01, &
9.5277237959592009D-01, &
1.1717152179539618D+00, &
2.5980977671460709D-01, &
-6.4468162556415265D-01, &
-1.3751364204170078D+00, &
2.9677172586579936D-01, &
-3.7071376979215720D-01, &
8.2912690407496381D-01, &
-8.6820437618589197D-01, &
5.2500961173269689D-01, &
1.0657701704030644D+00, &
-4.6621502244820201D-03, &
-1.9013997092621748D-01, &
1.5098985741543924D-01, &
1.0102557493909003D+00, &
8.8830298507891103D-01, &
2.0464938169302065D+00, &
4.7910192662606277D-01, &
1.4155288808120892D+00, &
-5.8169388172737679D-01, &
-9.8007278321065916D-01, &
2.4369633027015425D-01, &
1.6173163491335715D-01, &
6.6887624704464499D-01, &
-1.6500999383869115D+00, &
1.4380895281962367D+00, &
4.7508565250807777D-01, &
-3.1332991280327299D-01, &
3.1402552392574451D-01, &
5.6246373170551534D-01, &
2.5944662334710866D-01, &
4.8101648688789655D-01, &
1.7823376751423265D+00, &
3.0160656491545923D-01, &
-8.8915960863742050D-01, &
-4.4783548340444157D-01, &
8.9985836172311440D-01, &
-1.5626460660617920D+00, &
8.9972644535054036D-01, &
2.4456452268563592D-01, &
-3.1377944726557985D+00, &
1.6874136691232020D+00, &
2.4791290942030142D-01, &
1.7055713617986679D+00, &
1.7027580566127303D+00, &
-5.2969836953828042D-01, &
-8.6858804294195124D-01, &
7.6588136514601834D-01, &
8.6161822555855139D-01, &
6.5387844189250555D-01, &
7.0164941351276944D-01, &
4.1171318512873312D-01, &
7.6075070364872455D-01, &
8.5708035578209718D-02, &
-4.3558500874018535D-01, &
-6.2302104134015979D-01, &
8.4912051051824700D-01, &
-1.7120108380813925D-01, &
-9.7880552224113848D-01, &
1.1904436348486702D+00, &
7.0273864977367972D-01, &
-1.0213785672492079D+00, &
4.8392839864322634D-02, &
1.2611184618297511D-01, &
5.3330169134056482D-01, &
1.8070298106837654D+00, &
-2.8022831541922144D-01, &
8.0235047640662738D-01, &
-1.2615220404695868D+00, &
1.1878769364434660D+00, &
-2.1059219864297674D-01, &
3.2897539618854971D-01, &
-5.8928028913554642D-01, &
1.9164347352074701D-02, &
2.8035162764822374D-01, &
-9.6622429734784299D-02, &
3.4216241143907045D-01, &
-2.2358052317750254D+00, &
6.6284070879481805D-01, &
7.4316074303777269D-01, &
1.0280848437626724D+00, &
-2.0939898252763922D-01, &
-1.0268515265064981D+00, &
-1.2648527910628871D-01, &
4.8663846308033204D-01, &
1.2270171407392749D+00, &
-1.6189022502021406D+00/)
work(1:400) = (/6.0466028797961957D-01, &
9.4050908804501243D-01, &
6.6456005321849043D-01, &
4.3771418718698019D-01, &
4.2463749707126569D-01, &
6.8682307286710942D-01, &
6.5637019217476222D-02, &
1.5651925473279124D-01, &
9.6969518914484562D-02, &
3.0091186058528707D-01, &
5.1521262850206540D-01, &
8.1363996099009683D-01, &
2.1426387258237492D-01, &
3.8065718929968601D-01, &
3.1805817433032985D-01, &
4.6888984490242319D-01, &
2.8303415118044517D-01, &
2.9310185733681576D-01, &
6.7908467592021626D-01, &
2.1855305259276428D-01, &
2.0318687664732285D-01, &
3.6087141685690599D-01, &
5.7067327607102258D-01, &
8.6249143744788637D-01, &
2.9311424455385804D-01, &
2.9708256355629153D-01, &
7.5257303555161192D-01, &
2.0658266191369859D-01, &
8.6533501300156102D-01, &
6.9671916574663473D-01, &
5.2382030605000085D-01, &
2.8303083325889995D-02, &
1.5832827774512764D-01, &
6.0725343954551536D-01, &
9.7524161886057836D-01, &
7.9453623373871976D-02, &
5.9480859768306260D-01, &
5.9120651313875290D-02, &
6.9202458735311201D-01, &
3.0152268100655999D-01, &
1.7326623818270528D-01, &
5.4109985500873525D-01, &
5.4415557300088502D-01, &
2.7850762181610883D-01, &
4.2315220157182809D-01, &
5.3058571535070520D-01, &
2.5354050051506050D-01, &
2.8208099496492467D-01, &
7.8860491501934493D-01, &
3.6180548048031691D-01, &
8.8054312274161706D-01, &
2.9711226063977081D-01, &
8.9436172933045366D-01, &
9.7454618399116566D-02, &
9.7691686858626237D-01, &
7.4290998949843021D-02, &
2.2228941700678773D-01, &
6.8107831239257088D-01, &
2.4151508854715265D-01, &
3.1152244431052484D-01, &
9.3284642851843402D-01, &
7.4184895999182299D-01, &
8.0105504265266125D-01, &
7.3023147729480831D-01, &
1.8292491645390843D-01, &
4.2835708180680782D-01, &
8.9699195756187267D-01, &
6.8265348801324377D-01, &
9.7892935557668759D-01, &
9.2221225892172687D-01, &
9.0837275353887081D-02, &
4.9314199770488037D-01, &
9.2698680357441421D-01, &
9.5494544041678175D-01, &
3.4795396362822290D-01, &
6.9083883150567893D-01, &
7.1090719529999513D-01, &
5.6377959581526438D-01, &
6.4948946059294044D-01, &
5.5176504901277490D-01, &
7.5582350749159777D-01, &
4.0380328579570035D-01, &
1.3065111702897217D-01, &
9.8596472934024670D-01, &
8.9634174539621614D-01, &
3.2208397052088172D-01, &
7.2114776519267410D-01, &
6.4453978250932942D-01, &
8.5520507541911234D-02, &
6.6957529769977453D-01, &
6.2272831736370449D-01, &
3.6969284363982191D-01, &
2.3682254680548520D-01, &
5.3528189063440612D-01, &
1.8724610140105305D-01, &
2.3884070280531861D-01, &
6.2809817121836331D-01, &
1.2675292937260130D-01, &
2.8133029380535923D-01, &
4.1032284435628247D-01, &
4.3491247389145765D-01, &
6.2509502830053043D-01, &
5.5014692050772329D-01, &
6.2360882645293014D-01, &
7.2918072673429812D-01, &
8.3053391899480622D-01, &
5.1381551612136129D-04, &
7.3606860149543141D-01, &
3.9998376285699544D-01, &
4.9786811334270198D-01, &
6.0397810228292748D-01, &
4.0961827788499267D-01, &
2.9671281274886468D-02, &
1.9038945142366389D-03, &
2.8430411748625642D-03, &
9.1582131461295702D-01, &
5.8983418500491935D-01, &
5.5939244907101404D-01, &
8.1540517093336062D-01, &
8.7801175865240000D-01, &
4.5844247857565062D-01, &
6.0016559532333080D-01, &
2.6265150609689439D-02, &
8.4583278724804167D-01, &
2.4969320116349378D-01, &
6.4178429079958299D-01, &
2.4746660783662855D-01, &
1.7365584472313275D-01, &
5.9262375321244554D-01, &
8.1439455096702107D-01, &
6.9383813651720949D-01, &
3.0322547833006870D-02, &
5.3921010589094598D-01, &
9.7567481498731645D-01, &
7.5076305647959851D-01, &
2.9400631279501488D-01, &
7.5316127773675856D-01, &
1.5096404497960700D-01, &
3.5576726540923664D-01, &
8.3193085296981628D-01, &
2.3183004193767690D-01, &
6.2783460500002275D-01, &
4.9839430127597562D-01, &
8.9836089260366833D-02, &
2.5193959794895041D-02, &
3.9221618315402479D-01, &
5.8938308640079917D-01, &
9.2961163544903025D-01, &
5.7208680144308399D-01, &
5.8857634514348212D-01, &
4.1176268834501623D-01, &
5.5258038981424384D-01, &
4.9160739613162047D-01, &
9.5795391353751358D-01, &
7.9720854091080284D-01, &
1.0738111282075208D-01, &
7.8303497339600214D-01, &
3.9325099922888668D-01, &
1.3041384617379179D-01, &
1.9003276633920804D-01, &
7.3982578101583363D-01, &
6.5404140923127974D-01, &
9.8383788985732593D-02, &
5.2038028571222783D-01, &
9.9729663719935122D-02, &
1.5184340208190175D-01, &
7.6190262303755044D-02, &
3.1520808532012451D-01, &
1.5965092146489504D-01, &
1.3780406161952607D-01, &
3.2261068286779754D-01, &
5.3907451703947940D-01, &
5.7085162734549566D-01, &
5.1278175811108151D-01, &
6.8417513009745512D-01, &
6.5304020513536076D-01, &
5.2449975954986505D-01, &
6.5427013442414605D-01, &
7.1636837490167116D-01, &
6.3664421403817983D-01, &
1.2825909106361078D-02, &
3.0682195787138565D-02, &
9.8030874806304999D-02, &
3.6911170916434483D-01, &
8.2645412563474197D-01, &
3.4768170859156955D-01, &
3.4431501772636058D-01, &
2.5299982364784412D-01, &
2.1647114665497036D-01, &
5.5500213563479417D-01, &
4.0207084527183062D-01, &
5.0649706367641834D-01, &
1.6867966833433606D-01, &
3.3136826030698385D-01, &
8.2792809615055885D-01, &
7.0028787314581509D-01, &
5.7926259664335768D-02, &
9.9915949022033324D-01, &
4.1154036322047599D-01, &
1.1167463676480495D-01, &
7.8075408455849260D-01, &
9.2117624440742188D-02, &
5.3494624494407637D-02, &
7.1469581589162956D-01, &
2.5076227542918023D-01, &
8.4863292090315690D-01, &
9.7388187407067284D-01, &
2.1256094905031958D-01, &
2.1533783325605065D-02, &
9.4519476038882588D-01, &
9.2970155499924934D-02, &
6.4583337452397671D-01, &
3.1188554282705405D-01, &
4.4846436394045647D-01, &
4.8723924858036949D-01, &
8.2479676511350006D-02, &
6.7182910623463954D-01, &
4.0018828942364343D-01, &
9.0027514726431157D-01, &
9.4988320610125321D-01, &
3.1933126760711733D-01, &
4.9938549375241320D-01, &
4.0043231714181288D-01, &
1.9808670325451940D-02, &
6.4503886601944815D-01, &
4.2868843006993296D-01, &
3.3959675138730994D-01, &
8.8744750085050050D-01, &
2.3632747430436052D-01, &
7.6500821493327975D-01, &
3.5754647436084384D-02, &
7.2757725604152290D-01, &
6.2583662695812525D-01, &
5.1308750608785669D-01, &
7.2448356792351315D-02, &
7.2422905845916841D-01, &
8.7984484630570914D-01, &
9.7776347735771851D-01, &
8.4750026226468134D-01, &
8.3219793814993315D-01, &
2.4784452318699535D-01, &
9.1339906293647088D-01, &
7.5037210134653420D-02, &
8.3510380115435290D-01, &
6.2933169164530067D-01, &
7.5174057889673473D-01, &
6.3200343378879975D-01, &
9.6934213238731665D-02, &
1.4827369494876504D-02, &
5.8383474186253115D-01, &
6.8756195202154743D-02, &
9.9827381100849455D-01, &
6.4918841659842363D-01, &
9.8546557863324791D-01, &
8.3480576021921249D-01, &
3.3205608571906026D-01, &
6.6139318058334262D-01, &
9.5602062659660969D-01, &
3.1051027622482125D-01, &
1.8439069400202679D-01, &
9.6709434137177297D-01, &
8.3324181552815457D-01, &
3.0954845052732810D-01, &
8.0587176753764456D-01, &
4.1732584219038238D-01, &
7.1853044935277477D-01, &
4.0673677545039083D-01, &
8.9580326774414576D-01, &
9.5817636260259365D-01, &
1.8713221139656417D-02, &
7.9167230908208319D-01, &
4.2355315388584103D-01, &
1.5181277223073395D-02, &
4.3269824007906393D-01, &
9.0477623706573340D-01, &
8.5570441457488644D-01, &
4.2921642176334200D-02, &
6.5903053300775438D-01, &
3.4785904313005395D-01, &
5.0348679004869112D-01, &
8.3994742117055976D-01, &
2.3109568410543832D-02, &
1.2436351859954159D-01, &
2.6117561918821841D-01, &
8.3494750649349414D-01, &
3.1480479595597533D-01, &
7.6812064740880894D-03, &
8.9975012571752733D-01, &
3.7026753645051064D-01, &
1.0019940926941497D-01, &
6.4320402657020315D-01, &
7.6988908998308336D-01, &
7.9112533566198451D-01, &
2.6238190747072776D-01, &
3.4686388037925503D-01, &
2.1465371537694145D-01, &
8.2209289717657175D-01, &
3.5113429966521320D-01, &
5.9919425250588099D-01, &
5.7835125693111211D-01, &
4.1358098797631293D-01, &
1.1985050890286310D-01, &
9.1161370679159903D-01, &
5.3785580105748208D-02, &
2.2891758676059876D-01, &
3.2417396306138829D-01, &
3.5076512764716117D-01, &
3.4928874777426255D-01, &
3.0380212985436572D-01, &
9.6874615996581170D-01, &
6.7152655046083776D-01, &
2.0794312837315651D-01, &
9.6313940120247044D-01, &
3.0220237504213365D-01, &
8.0794108095480799D-01, &
1.3408416275024179D-01, &
9.4776028919455635D-01, &
6.4086482116825383D-01, &
9.5325875425035178D-01, &
8.0987422593395209D-01, &
1.8159084675756379D-01, &
9.4275737153737327D-01, &
8.3124103554376771D-01, &
4.9468043578205978D-01, &
8.5531034647693982D-01, &
7.1074391181909824D-01, &
2.7349475629159786D-01, &
4.0763287189198161D-01, &
9.0976128251911847D-01, &
9.4439713870030451D-01, &
4.9863245185560190D-01, &
2.8863831012730923D-01, &
9.7589525649963815D-01, &
4.5258447627808124D-01, &
4.4990698677957075D-02, &
3.1536198151820755D-01, &
9.5190614812037189D-01, &
7.5156308247423609D-01, &
5.3579099898961424D-01, &
6.6971458883510748D-01, &
8.6517499748328641D-01, &
4.5888445390388938D-01, &
5.7855090249582031D-01, &
4.8152982184966137D-01, &
5.5061576198318274D-01, &
9.5062324380815433D-01, &
5.0986542047295536D-01, &
7.4251472966182985D-01, &
4.9079401441435533D-01, &
6.6151414870689360D-02, &
2.6249066264989940D-01, &
9.2546794407799982D-01, &
3.7148665165822231D-01, &
4.0941940003107308D-01, &
4.1575196973399631D-01, &
9.7261599736539445D-02, &
9.0162762447969347D-01, &
4.4446597981328932D-03, &
2.7392454335102678D-01, &
1.0930666111680035D-01, &
8.5544841289295426D-01, &
2.5705535663902546D-01, &
9.8913209203202213D-01, &
9.2641142236812712D-01, &
1.7094603208839290D-01, &
3.0388712489325242D-01, &
5.3345144978115477D-01, &
1.7648961347647024D-01, &
8.1359077477652830D-01, &
7.0513712380125892D-01, &
2.5720755742139950D-01, &
2.5036892046498466D-01, &
3.3509436689927874D-01, &
7.5124063162526056D-01, &
4.8797826077860845D-03, &
8.4099320643626019D-01, &
2.2957358869665739D-01, &
1.3285547727582237D-02, &
9.4993740716879371D-01, &
8.9937146465701423D-01, &
9.6262420114388625D-01, &
4.3000361954927006D-02, &
7.1266261216467264D-01, &
5.1094098258212241D-02, &
4.0753210485857738D-01, &
4.7569737399615403D-01, &
3.4746838606940983D-01, &
4.0719938711096422D-02, &
5.9756620514440806D-01, &
2.6012467360309705D-01, &
8.3285585557738717D-01, &
9.6049750529821787D-01, &
9.3670756890653750D-01, &
2.2932023844733959D-01, &
7.2031310018914962D-01, &
7.5648232426876405D-01, &
4.5015392507594826D-01, &
3.3897738839543617D-01, &
4.7249205225111501D-01, &
9.8599436000817042D-01/)
n =  100
info = 0

open(unit = 4, STATUS='REPLACE', file = "gen4tests.txt")
open(unit = 3, STATUS='REPLACE', file = "gen3tests.txt")
open(unit = 5, STATUS='REPLACE', file = "gen5tests.txt")

call dlasq1(n, d, e, work, info)

close(5)
close(3)
close(4)

end