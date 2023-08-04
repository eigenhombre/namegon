package namegon

// Adapted from
// https://www.math.ubc.ca/~cass/frivs/latin/latin-dict-full.html
var LatinWords = []string{
	"abbas", "abbatia", "abbatis", "abduco", "abeo", "abscido", "absconditus",
	"absens", "absentis", "absorbeo", "absque", "abstergo", "absum",
	"abundans", "abundantia", "abutor", "ac", "accedo", "accendo", "acceptus",
	"accerso", "accipio", "accommodo", "accusator", "accuso", "acer",
	"acerbitas", "acerbus", "acervus", "acidus", "acies", "acquiro", "acsi",
	"actum", "adamo", "adaugeo", "addo", "adduco", "ademptio", "adeo",
	"adeptio", "adepto", "adfectus", "adfero", "adficio", "adflicto",
	"adhaero", "adhuc", "adicio", "adimpleo", "adinventitias", "adipiscor",
	"adiuvo", "administratio", "admiratio", "admiror", "admitto", "admoneo",
	"admonitio", "admoveo", "adnuo", "adopto", "adsidue", "adstringo",
	"adsuesco", "adsum", "adsumo", "adulatio", "adulescens", "adulescentia",
	"adultus", "aduro", "adustum", "advenio", "adversus", "adverto", "advoco",
	"ae", "aedificium", "aeger", "aegre", "aegresco", "aegretudo",
	"aegrotatio", "aegrus", "aeneus", "aequitas", "aequus", "aer", "aeris",
	"aestas", "aestivus", "aestus", "aetas", "aeternus", "affectusadfero",
	"affero", "affligo", "ager", "agere", "aggero", "aggredior", "agnitio",
	"agnosco", "ago", "agri", "ait", "aiunt", "album", "alia", "alienus", "alii",
	"alioqui", "alioquin", "aliqua", "aliquando", "aliquanta", "aliquanto",
	"aliquantum", "aliquantus", "aliqui", "aliquid", "aliquis", "aliquo",
	"aliquot", "aliquotiens", "aliud", "alius", "allatus", "alo", "alter",
	"altera", "alterum", "altum", "altus", "alui", "alveus", "amaritudo",
	"ambitus", "ambulo", "amicitia", "amiculum", "amicus", "amissio",
	"amissus", "amita", "amitto", "amo", "amor", "amoveo", "amplexusamplitudo",
	"amplus", "an", "ancilla", "angelus", "angulus", "angustus", "animadverto",
	"animi", "animus", "annus", "anser", "ante", "antea", "antepono",
	"antiquus", "aperio", "aperte", "apostolus", "apparatus", "appareo",
	"appello", "appono", "appositus", "approbo", "appropinquo", "apto",
	"aptus", "apud", "aqua", "ara", "aranea", "arbitro", "arbitror", "arbor",
	"arbustum", "arbustus", "arca", "arceo", "arcesso", "archa", "arcis",
	"arcus", "argentum", "argumentum", "arguo", "arma", "armarium", "armo",
	"aro", "ars", "articulus", "artificiose", "artificiosus", "artis", "arto",
	"arx", "ascisco", "ascitasper", "asperitas", "aspicio", "asporto",
	"assentator", "assidue", "assuesco", "assumo", "astringo", "astrum", "at",
	"atavus", "ater", "atque", "atqui", "atra", "atrocitas", "atrox", "atrum",
	"attero", "attollo", "attonbitus", "auctor", "auctoritas", "auctus",
	"audacia", "audaciter", "audacter", "audax", "audentia", "audeoaudio",
	"auditor", "aufero", "aureus", "auris", "aurum", "aut", "autem", "autus",
	"auxilium", "avaritia", "avarus", "aveho", "averto", "avoco", "baiulus",
	"balbus", "barba", "bardus", "basium", "beatus", "bellicus", "bellum",
	"bellus", "benebenevolentia", "benigne", "bestia", "bibo", "bis",
	"blandior", "blanditia", "bonus", "bos", "bovis", "brevis", "brevitas",
	"breviter", "cado", "caecus", "caelestis", "caelum", "calamitas",
	"calamus", "calcar", "calco", "calculus", "callide", "callidus", "campana",
	"candidus", "canis", "canonicus", "canonus", "canto", "capillus", "capio",
	"capitis", "capitulus", "capto", "caput", "carbo", "carbonis", "carcer",
	"careo", "caries", "cariosus", "carisma", "caritas", "carmen", "carpo",
	"carus", "casso", "caste", "casus", "catena", "caterva", "catervatim",
	"cattus", "cauda", "causa", "caute", "cautela", "cautim", "cautum", "caveo",
	"cavi", "cavus", "cedo", "celebrer", "celebrus", "celer", "celeritas",
	"celeriter", "celo", "cena", "cenaculum", "ceno", "censura", "centum",
	"cerno", "cernuus", "certe", "certo", "certus", "cervus", "cessi", "cessum",
	"cetera", "ceteri", "ceterum", "ceterus", "charisma", "chirographum",
	"cibo", "cibus", "cicuta", "cilicium", "cimentarius", "ciminatio",
	"ciminosus", "cineris", "cinis", "circumvenio", "cito", "civilis", "civis",
	"civitas", "clam", "clamo", "clamor", "claro", "clarus", "claudeo",
	"claudo", "claudus", "claustrum", "clausus", "clementia", "clibanus",
	"coactum", "coadunatio", "coaegresco", "coegi", "coepicoerceo", "cogito",
	"cognatus", "cognomen", "cognosco", "cogo", "cohaero", "cohero", "cohesi",
	"cohesum", "cohibeo", "cohors", "cohortor", "colligo", "colloco", "collum",
	"colo", "color", "coloratus", "coloro", "colui", "coma", "combibo",
	"comburo", "comedi", "comedo", "comes", "comesum", "cometes", "cometissa",
	"comis", "comitatus", "comiter", "comitis", "comitissa", "comitto",
	"commemoro", "commeo", "commessatio", "comminor", "comminuo", "comminus",
	"commisceo", "commissum", "commodo", "commodum", "commoneo", "commoveo",
	"communis", "comparo", "compater", "compatior", "compedis", "compello",
	"comperio", "comperte", "compes", "competo", "complectus", "compleo",
	"compono", "compositio", "compositus", "comprehendo", "comprobo",
	"comprovincialis", "comptus", "conatus", "concedo", "concepta", "concero",
	"concido", "concilium", "concipio", "concisus", "concito", "conculco",
	"concupiscentia", "concupisco", "concutio", "condico", "conduco",
	"confero", "confessus", "confestim", "confido", "confiteor", "conforto",
	"confugo", "congregatio", "congrego", "congruus", "conicio", "coniecto",
	"conitor", "coniuratio", "coniuratus", "coniuro", "conor", "conqueror",
	"conscendo", "conscientia", "conscindo", "conscius", "conservo",
	"considero", "consido", "consilio", "consilium", "consisto", "consitor",
	"conspergo", "conspicio", "constans", "constanter", "constituo", "consto",
	"constrictum", "constringo", "constrixi", "constructum", "construo",
	"construxi", "constupator", "constupro", "consuasor", "consuefacio",
	"consuesco", "consueta", "consuetudo", "consui", "consulatio", "consulo",
	"consulto", "consultum", "consummatio", "consummo", "consumo", "consuo",
	"consurgo", "consutum", "contabesco", "contactus", "contages", "contagio",
	"contagium", "contamino", "contectum", "contego", "contemno",
	"contemplatio", "contemplor", "contemptim", "contemptio", "contemptus",
	"contendo", "contente", "contentus", "contigi", "contigo", "contineo",
	"contingo", "continuo", "continuus", "contra", "contradictio", "contrado",
	"contraho", "contristo", "conturbo", "conventus", "conversatio",
	"converto", "convoco", "copia", "copiae", "copie", "copiose", "cornu",
	"corona", "corporis", "corpus", "correptius", "corrigocorripio",
	"corroboro", "corrumpo", "corruo", "coruscus", "cotidie", "crapula",
	"cras", "crastinus", "creator", "creatura", "creber", "crebro", "credo",
	"creo", "creptio", "crepusculum", "crescocreta", "cribro", "cribrum",
	"crinis", "crinitus", "cruciamentum", "crucio", "crucis", "crudelis",
	"cruentus", "crur", "cruris", "crustulum", "crux", "cubicularis",
	"cubicularius", "cubiculum", "cubitum", "cubo", "cui", "cuius",
	"cuiusmodi", "culpa", "culpo", "cultellus", "cultum", "cultura", "cum",
	"cunabula", "cunae", "cunctatio", "cunctator", "cunctor", "cunctus",
	"cupiditas", "cupido", "cupio", "cuppedia", "cupressus", "cur", "cura",
	"curatio", "curator", "curia", "curiositas", "curiosus", "curis", "curo",
	"curriculum", "currus", "cursim", "cursito", "curso", "cursor", "cursus",
	"curto", "curtus", "curvo", "curvus", "custodia", "custodiae", "custodie",
	"custos", "damnatio", "damno", "dapifer", "dare", "datum", "de", "debeo",
	"debilito", "decens", "decenter", "decerno", "decerto", "decet", "decimus",
	"decipio", "decor", "decoro", "decorus", "decretum", "decumbo", "dedecor",
	"dedecus", "dedi", "dedico", "deduco", "defaeco", "defendo", "defero",
	"defessus", "defetiscor", "deficiodefigo", "defleo", "defluo", "defungo",
	"degenero", "degero", "degusto", "deinde", "delectatio", "delecto",
	"delego", "deleo", "delibero", "delicate", "deliciae", "delinquo",
	"deludo", "demensdemergo", "demitto", "demo", "demonstro", "demoror",
	"demulceo", "demum", "denego", "denique", "densdenuncio", "denuntio",
	"denuo", "deorsum", "depereo", "depono", "depopulo", "depopulor",
	"deporto", "depraedor", "deprecator", "deprecor", "depredor",
	"deprehensio", "depressus", "deprimo", "depromo", "depulso", "deputo",
	"derelinquo", "derideo", "deripio", "desidero", "desidiosus", "desiit",
	"desino", "desipio", "desolo", "desparatus", "despecto", "despero",
	"despiciens", "despicio", "desposco", "destinatus", "destituo",
	"detectum", "detego", "determino", "detineo", "detrimentum", "deus",
	"devenio", "devito", "devoco", "devotio", "devoveo", "dexter", "dextera",
	"diabolus", "dico", "dictata", "dictator", "dictito", "dicto", "didicerat",
	"didico", "dididi", "dido", "didtum", "diei", "dies", "diffama", "differo",
	"differtus", "difficilis", "difficultas", "digestor", "dignitas",
	"dignosco", "dignus", "digredi", "digredior", "digressio", "digressus",
	"dilabor", "dilato", "dilgenter", "diligens", "diligentia", "diligo",
	"diluculo", "diluo", "dimidium", "dimitto", "directus", "diripio",
	"dirunitas", "diruo", "discedodiscidium", "discipulus", "disco",
	"discrepo", "dispono", "disputatio", "disputo", "dissero", "dissimilis",
	"dissimulo", "dissolutus", "distinguo", "distribuo", "districtus",
	"distringo", "distulo", "dito", "diu", "diutinus", "diutius", "diuturnus",
	"diversus", "dives", "divinitus", "divinus", "divitiae", "divitie", "do",
	"doceo", "doctor", "doctrina", "doctum", "doctus", "docui", "dolens",
	"doleo", "dolor", "dolose", "dolosus", "dolus", "domesticus", "domina",
	"dominatus", "domino", "dominus", "domito", "domus", "donec", "donum",
	"dormio", "dubito", "dubium", "ducis", "duco", "dudum", "dulcedo",
	"dulcidine", "dulcis", "dulcitudinisdummodo", "dulcitudo", "dum",
	"dumtaxat", "duo", "duro", "durus", "dux", "e", "ea", "eadem", "eatenus",
	"ebullio", "ecclesia", "econtra", "ecquando", "edi", "edico", "edificium",
	"editio", "edo", "edoceo", "educo", "effectus", "effero", "effervo",
	"efficio", "effrego", "effringo", "effugioeffundo", "egenus", "egeo",
	"eger", "egi", "ego", "egre", "egredioreicio", "egresco", "egretudo",
	"egrotatio", "eiectum", "elatum", "elatus", "electus", "elementum",
	"elemosina", "elemosinarius", "eligo", "eloquens", "eloquentia", "eluo",
	"eluvies", "eluvio", "emanio", "emendo", "emercor", "emerio", "emi",
	"emineo", "eminor", "eminus", "emiror", "emo", "emoveo", "emptio", "emptum",
	"eneus", "enim", "enumero", "eo", "episcopalis", "episcopus", "epistula",
	"epulo", "epulor", "eques", "equidem", "equina", "equitas", "equitatus",
	"equitis", "equus", "erepo", "erga", "ergo", "eripioerogo", "erogo", "erro",
	"error", "erubesco", "erudio", "eruo", "esse", "essum", "est", "estas",
	"estivus", "estus", "esurio", "et", "etc", "etenim", "eternus", "etiam",
	"etsi", "eum", "evenio", "eventus", "everto", "evito", "evoco", "ex",
	"excedo", "excellentia", "excessum", "excipioexcito", "exclamo",
	"excludo", "excolo", "excrucio", "excusatio", "excuso", "exemplar",
	"exemplum", "exequo", "exequor", "exerceo", "exercitus", "exertus",
	"exesto", "exexaequo", "exheredo", "exheres", "exhibeo", "exhilaro",
	"exhorresco", "exibeo", "exigo", "exilis", "eximietate", "eximius",
	"eximo", "exinde", "exitiabilis", "exitialis", "exitium", "exitosus",
	"exitus", "exordium", "exorior", "exoro", "exorsus", "expedio", "expello",
	"experior", "experir", "expers", "expertus", "expetens", "expeto",
	"expilatio", "expiscor", "expleo", "expletio", "expletum", "expletus",
	"explevi", "explicatus", "explico", "expono", "expositum", "expositus",
	"expostulo", "exposui", "expugno", "expuli", "expulsum", "exquisitus",
	"exsequor", "exsertus", "exsilium", "exspecto", "exstinctum", "exstingui",
	"exstinguo", "exsto", "externus", "extollo", "extorqueo", "extra",
	"extremus", "extuli", "exturbo", "exulto", "exuro", "exussum", "exustio",
	"fabula", "facile", "facilis", "facillimus", "facina", "facinoris",
	"facio", "factum", "factus", "facultas", "facundiafacunditas", "faenum",
	"falcis", "falsus", "fama", "familia", "familiaris", "famulatus",
	"famulus", "fas", "fateor", "fatigo", "fatum", "fautor", "faveo",
	"feculentia", "fefello", "felicis", "feliciter", "felix", "femina",
	"fenestra", "fenum", "fere", "feretrum", "feritas", "ferme", "fero",
	"ferre", "ferrum", "ferus", "festinatio", "festino", "festinus", "feteo",
	"ficus", "fidelis", "fidelitas", "fidens", "fides", "fiducia", "fieri",
	"filia", "filius", "fimus", "fines", "finis", "finitimus", "finitumus",
	"finium", "fio", "firmo", "firmus", "flamma", "flatus", "flax", "fleo",
	"fluctus", "flumen", "fluo", "fodio", "fore", "forem", "forensis", "forma",
	"formica", "formo", "fors", "forsan", "forsit", "forsitan", "fortasse",
	"forte", "fortis", "fortiter", "fortitudo", "fortuna", "fortunate",
	"fortunatus", "forum", "fotum", "foveo", "fovi", "fracta", "frango",
	"frater", "frendo", "frequentatio", "frequentia", "frequento", "frigus",
	"frofui", "from", "frons", "fructuarius", "fructus", "frugalitas", "frugi",
	"frumentum", "fruor", "frustra", "frux", "fuga", "fugio", "fugitivusfugo",
	"fui", "fulcio", "fulgeo", "fulsi", "fultus", "functus", "fundo", "fungi",
	"fungor", "funis", "furibundus", "furor", "furs", "furtificus", "furtim",
	"furtumfususG", "futurus", "galea", "gaudium", "gemo", "gens", "genuit",
	"genus", "gero", "gesto", "gestum", "gigno", "glacialis", "gladius",
	"gloria", "glorior", "grando", "grassor", "gratia", "gratulor", "gratus",
	"gravatus", "gravis", "gravitas", "graviter", "gravo", "gregatim", "gusto",
	"habeo", "habere", "habitum", "habitus", "habui", "hac", "hactenus", "hae",
	"haec", "haesito", "hanc", "harum", "has", "haud", "he", "hec", "hereditas",
	"hesito", "hi", "hic", "hilaris", "hinc", "his", "ho", "hoc", "hodie",
	"hodiernus", "hominis", "homo", "honor", "honorabilis", "hora", "hordeum",
	"horrendus", "hortor", "hortus", "horum", "hos", "hospes", "hostes",
	"hostis", "hostium", "huic", "huius", "humanitas", "humanus", "humerus",
	"humilis", "humo", "humus", "hunc", "hypocrita", "iaceo", "iacio",
	"iaculator", "iaculum", "iam", "ianua", "ibi", "ictus", "id", "idcirco",
	"idem", "identidem", "ideo", "idoneus", "igitur", "ignarus", "ignavus",
	"ignis", "ignoro", "ignosco", "ignotus", "ilico", "illa", "illacrimo",
	"illae", "illarum", "illas", "illataillaturos", "ille", "illi", "illiam",
	"illic", "illis", "illius", "illo", "illorum", "illos", "illuc", "illud",
	"illudo", "illum", "illusi", "illusum", "imago", "imber", "imbrium",
	"imcompositeimitabilis", "imitor", "immanitas", "immerito", "immineo",
	"immo", "immodicus", "immortalis", "immotus", "immunda", "immundus",
	"impedimentum", "impedio", "impedito", "impedo", "impello", "impendeo",
	"impendium", "impendo", "impenetrabiilis", "impensa", "imperator",
	"imperceptus", "imperiosus", "imperium", "impero", "impetro", "impetus",
	"impleo", "importo", "importunus", "impraesentiarum", "imprimis",
	"improbus", "improvidus", "improviso", "impudens", "impudenter", "impuli",
	"impulsum", "impunitus", "imputo", "in", "inanis", "incassum", "inceptor",
	"inceptum", "incertus", "incido", "incipio", "incito", "inclino",
	"inclitus", "includo", "inclutus", "incola", "incompositus", "inconsulte",
	"incontinencia", "incorruptus", "incredibilis", "increpare", "increpo",
	"incubo", "incurro", "indagatio", "inde", "indebitus", "indico", "indidi",
	"indigeo", "indignatio", "indignus", "indo", "indomitus", "induco",
	"inductum", "industria", "industrius", "indutiae", "indux",
	"inedicabilis", "ineptio", "inexpugnabilis", "infamo", "infantia",
	"infeci", "infector", "infectum", "infectus", "infecunditas",
	"infecundus", "infelicitas", "infeliciter", "infelixinfelix", "infenso",
	"infensus", "inferi", "inferius", "inferne", "infero", "inferoinferus",
	"infervesco", "infeste", "infesto", "infestus", "inficio", "infidelis",
	"infidelitas", "infideliter", "infidi", "infidus", "infigo", "infimus",
	"infindo", "infinitas", "infinitio", "infinitus", "infirmatio", "infirme",
	"infirmitas", "infirmo", "infirmus", "infissum", "infit", "infitialis",
	"infitias", "infitior", "inflammatio", "inflammo", "inflatio",
	"inflatius", "inflatus", "inflecto", "inflectum", "infletus", "inflexi",
	"inflexio", "inflexus", "inflictum", "infligo", "inflixi", "inflo",
	"influi", "influo", "influxum", "infodi", "infodio", "informatio",
	"informis", "infortunatus", "infortunium", "infossum", "infra", "infula",
	"ingemuo", "ingenium", "ingens", "ingero", "ingratus", "ingravesco",
	"inicio", "inieci", "iniectum", "inimicus", "iniquus", "initium",
	"iniuria", "iniustus", "innotesco", "innotui", "innuo", "inolesco",
	"inops", "inquam", "inquis", "inquit", "inrideo", "inritus", "inruo",
	"insania", "insciens", "inscribo", "insensatus", "insequor", "inservio",
	"insideo", "insidiae", "insinuo", "insisto", "insolita", "insolitus",
	"insons", "insontis", "insperatus", "instanter", "instar", "instigo",
	"instituo", "insto", "instructus", "instruo", "insurgi", "insurgo",
	"insurrectum", "integer", "intellectum", "intellego", "intellexi",
	"intempestivus", "intendo", "intentio", "intentus", "inter", "intercepi",
	"interceptum", "intercipio", "interdico", "interdum", "intereo",
	"interfeci", "interfectum", "interficio", "interrogatio", "intro",
	"introduco", "intueor", "intumesco", "intus", "inultus", "invado",
	"invalesco", "invenio", "inventor", "investigo", "inveteratus",
	"invetero", "invicem", "invictus", "invideo", "invidia", "inviso",
	"invisus", "invito", "invitus", "ioco", "iocor", "iocus", "ipsa", "ipse",
	"ipsemet", "ipsum", "ira", "irascor", "iratus", "ire", "irrito", "irritum",
	"irritus", "irruo", "is", "ista", "iste", "istud", "ita", "itaque", "iter",
	"itero", "iterum", "itineris", "itis", "itum", "iubeo", "iucunditas",
	"iucundus", "iudex", "iudicium", "iudico", "iuge", "iugis", "iumentum",
	"iunctum", "iungo", "iunxi", "iurandi", "iurandum", "iuris", "iuro", "ius",
	"iussi", "iussu", "iussum", "iustus", "iuvo", "iuxta", "jaculum",
	"judicium", "jugis", "jugiter", "jumentum", "juvenis", "juventus",
	"labefacio", "labefacto", "labefactum", "labefeci", "labellum", "labes",
	"labia", "labiae", "labiosus", "labis", "labium", "labo", "labor",
	"labores", "laboriose", "laboriosus", "laboris", "laboro", "labrum",
	"labrusca", "labruscum", "lac", "lacer", "laceratio", "lacero", "lacerta",
	"lacertosus", "lacertus", "lacesso", "lacrima", "lacrimabilis", "lacrimo",
	"lacrimosus", "lactans", "lactatio", "lacteus", "lactis", "lacto",
	"lactuca", "lacuna", "lacunar", "lacus", "laedo", "laesio", "laetabilis",
	"laetans", "laetatio", "laetifico", "laetificus", "laetitia", "laeto",
	"laetor", "laeve", "laevus", "laganum", "lama", "lambo", "lamenta",
	"lamentabilis", "lamentatio", "lamentor", "lamia", "lammina", "lamna",
	"lamnia", "lapsus", "laqueum", "laqueus", "largior", "lascivio", "lasesco",
	"lateque", "latum", "latus", "latuseris", "laudator", "laudo", "laus",
	"lebes", "lectica", "lector", "lectus", "ledo", "legatarius", "legatus",
	"legens", "legentis", "legio", "legis", "lego", "lemiscus", "lemma",
	"lemures", "lenimentus", "lenio", "lenis", "lenitas", "lenitudo", "leno",
	"lenocinium", "lenocinor", "lens", "lente", "lentesco", "lentis",
	"lentitudo", "lento", "lentulus", "lentus", "lepide", "lepidus", "lepor",
	"lepos", "lepus", "lesio", "letabilis", "letalis", "letaliter", "letanie",
	"letatio", "lethargus", "letifer", "letifico", "letificus", "letitia",
	"leto", "letor", "letum", "levamen", "levamentum", "levatio", "leve",
	"leviculus", "levidensis", "levis", "levitas", "leviter", "levo", "levus",
	"lex", "libatio", "libellus", "libenter", "liber", "libera", "liberalis",
	"liberalitas", "liberaliter", "liberatio", "libere", "libero", "libertas",
	"liberum", "libido", "libri", "licet", "ligo", "lima", "limen", "limina",
	"lingua", "lino", "linteum", "liquidus", "litigo", "littera", "litterae",
	"loci", "loco", "locupleto", "locus", "locutus", "loginquitas", "longe",
	"longus", "loquacis", "loquax", "loquor", "loricatus", "lubricus",
	"lucerna", "lucis", "lucror", "lucrosus", "lucrum", "luctisonus", "luctus",
	"ludio", "ludius", "ludo", "ludus", "lues", "lugeo", "luna", "lupus", "lusi",
	"lusum", "lux", "luxuria", "macellarius", "macer", "macero", "macies",
	"macresco", "mactabilis", "macto", "macula", "maculo", "maculosus",
	"madesco", "madide", "madidus", "mador", "maero", "maeror", "magis",
	"magister", "magnopere", "magnus", "magus", "maiestas", "maior", "maiores",
	"mairia", "male", "malens", "mallui", "malo", "malum", "malus", "mancepo",
	"mancipo", "mandatum", "mando", "mane", "manentia", "maneo", "mansuetus",
	"manus", "mare", "maris", "maritus", "mater", "matera", "materia",
	"matertera", "matrimonium", "matris", "maxime", "maximus", "me", "medicus",
	"mediocris", "meditatus", "meditor", "medius", "mei", "mel", "melior",
	"mellis", "mellitus", "membrana", "memini", "meminisse", "memor",
	"memoratus", "memoria", "mens", "mensa", "mensis", "mentionem", "mentis",
	"mercedis", "merces", "mereo", "mereor", "meretricis", "meretrix",
	"meridianus", "meror", "mestitia", "metuo", "metus", "meus", "mica", "mihi",
	"miles", "milia", "milies", "militaris", "militis", "mille", "millies",
	"minime", "minimus", "ministro", "minor", "minui", "minuo", "minus",
	"minutum", "mire", "miro", "miror", "mirus", "misceo", "miscui", "miser",
	"misere", "misereo", "misereor", "misericordia", "misi", "missa", "missum",
	"mitesco", "mitigo", "mitis", "mitto", "mixtum", "modestus", "modica",
	"modicus", "modio", "modo", "modus", "moleste", "molestia", "molestus",
	"molior", "mollio", "mollis", "monachus", "monasterium", "monastery",
	"moneo", "monitio", "mons", "monstro", "monstrum", "montis", "mora",
	"moratlis", "morbus", "mores", "morior", "moris", "morium", "mors",
	"morsus", "mortifera", "mortis", "mortuus", "mos", "moti", "motum", "moveo",
	"mox", "mucro", "mugio", "mulier", "multi", "multo", "multum", "multus",
	"multusmultus", "mundo", "mundus", "munerior", "muneris", "munero",
	"munimentum", "munio", "munitio", "munus", "muris", "mus", "mussito",
	"mutatio", "muto", "mutuo", "mutuus", "mutuusnascor", "nasci", "natalis",
	"natio", "natura", "natus", "nauta", "navigatio", "navigo", "navis", "ne",
	"nec", "necdum", "necessarius", "necesse", "necne", "neco", "nefas", "nego",
	"negotium", "nemo", "neo", "nepos", "nepotis", "nequam", "nequaquam",
	"neque", "nequeo", "nequitia", "nequities", "nescio", "nichilominus",
	"nidor", "niger", "nihil", "nihilum", "nimirum", "nimis", "nimium", "nisi",
	"niteo", "nitesco", "nitor", "niveus", "no", "nobis", "nocens", "noceo",
	"noctis", "nolle", "nolo", "nolui", "nomen", "nominatim", "nomine",
	"nominetenus", "non", "nondum", "nonnisi", "nonnullus", "nonnumquam",
	"nonus", "nos", "nosco", "noster", "nostra", "nostri", "nostrum", "nota",
	"notarius", "novem", "novitas", "novo", "novus", "nox", "nullus", "numerus",
	"numquam", "nunc", "nunquam", "nuntio", "nuntius", "nuper", "nusquam",
	"nutrimens", "nutrimentus", "nutrio", "nutus", "ob", "obdormio", "obduro",
	"obicio", "obieci", "obiectum", "obligatus", "obliquo", "oblittero",
	"oblivio", "obruo", "obsequium", "obstinatus", "obtestor", "obtineo",
	"obviam", "obvius", "occasio", "occasum", "occidi", "occido", "occulto",
	"occupo", "occurro", "occursus", "ocius", "oculus", "odio", "odium", "of",
	"offensio", "offero", "officina", "officium", "olim", "omitto",
	"omnigenus", "omnino", "omnipotens", "omnis", "onero", "onis", "onus",
	"opera", "opere", "operis", "operor", "opes", "opinio", "opisthotonos",
	"oporotheca", "oportet", "oportunitas", "oppono", "opportune",
	"opportunitas", "opportunitatus", "opportunus", "oppositum", "opposui",
	"oppressi", "oppressum", "opprimo", "opprobrium", "oppugno", "ops",
	"optimates", "optimus", "opto", "opus", "or", "oratio", "orator", "orbis",
	"ordeum", "ordinatio", "ordine", "ordinem", "ordo", "orior", "oriri",
	"oris", "ornatus", "orno", "oro", "ortus", "os", "ostendo", "ostium",
	"otium", "ovis", "pacis", "paciscor", "pactum", "pactus", "paene",
	"paganus", "pala", "palam", "palea", "pallium", "palma", "pando", "panis",
	"par", "paratus", "parco", "parens", "parentis", "pareo", "paries",
	"parietis", "parilis", "pario", "pariter", "paro", "pars", "partim",
	"partis", "parum", "parvus", "pasco", "passer", "passim", "past",
	"patefacio", "pateo", "pater", "paternus", "patiens", "patientia",
	"patior", "patria", "patris", "patrocinor", "patronus", "patruus", "pauci",
	"paulatim", "pauper", "paupertas", "pax", "peccatus", "pecco", "pecto",
	"pectoris", "pectus", "pecunia", "pecus", "pedis", "peior", "peius",
	"pello", "pendeo", "pendo", "pene", "penitus", "penus", "pepulli", "pepulo",
	"per", "peracto", "peragro", "percepi", "perceptum", "percipio",
	"percontor", "perculsus", "percunctor", "percussum", "percutio",
	"perdignus", "perdo", "perduco", "peregrinus", "peremi", "peremptum",
	"pereo", "perfeci", "perfectum", "perfectus", "perfero", "perficio",
	"perfruor", "perfusus", "pergo", "periclitatus", "periclitor",
	"periculosus", "periculum", "perimo", "peritus", "periurium", "perlustro",
	"permissi", "permissum", "permitto", "permoveo", "perniciosus",
	"perperam", "perpetro", "perpetuus", "perscitus", "perscribo", "perseco",
	"persecutus", "persequi", "persequor", "perseverantia", "persevero",
	"persisto", "persolvo", "personam", "perspicuus", "persuadeo", "persuasi",
	"persuasum", "perterreo", "pertimesco", "pertinacia", "pertinaciter",
	"pertinax", "pertineo", "pertingo", "pertorqueo", "pertraho", "perturbo",
	"perturpis", "peruro", "perussi", "perustum", "pervalidus", "pervenio",
	"perversum", "perverti", "perverto", "pervideo", "pervidi", "pervisum",
	"pes", "pessime", "pessimus", "pessum", "pestifer", "pestifere", "pestis",
	"petitus", "peto", "petram", "pharetra", "phasma", "phasmatis",
	"phitonicum", "pia", "pica", "picea", "pictor", "pictoratus", "piger",
	"pignus", "pigra", "pigrum", "piper", "piperis", "pipio", "pirum", "pirus",
	"piscator", "piscis", "pium", "pius", "placeo", "placet", "placide",
	"placidus", "placitum", "placo", "plactum", "plaga", "plagiarius", "plane",
	"plango", "planxi", "platea", "plaustrum", "plebis", "plebs", "plecto",
	"plector", "plene", "plenus", "plerumque", "plerusque", "plico",
	"plorabilis", "plorator", "ploratus", "ploro", "pluit", "pluma",
	"plumbeus", "plumbum", "pluo", "plura", "plures", "plurimi", "plurimum",
	"plurimus", "pluris", "plus", "plusculus", "pluvia", "pluvialis", "pluvit",
	"pocius", "poema", "poematis", "poenapoeta", "polenta", "pollen", "polleo",
	"pollex", "polliceor", "pollicitus", "pomum", "pono", "pons", "pontis",
	"poposcopopulus", "porro", "porta", "posco", "positum", "positus", "posse",
	"possessio", "possum", "post", "postea", "posteri", "posterus",
	"posthabeopostpono", "postquam", "postremo", "postulo", "posui", "potens",
	"potestas", "potior", "potissimum", "potissimus", "potius", "potum",
	"prae", "praebeo", "praecedo", "praecelsus", "praecepio", "praeceptum",
	"praecido", "praecipio", "praecipuus", "praeclarus", "praeconor",
	"praecox", "praeda", "praedico", "praeeo", "praefero", "praeficio",
	"praefinio", "praefoco", "praegravo", "praemitto", "praemium", "praemo",
	"praenuntio", "praenuntius", "praepono", "praepositus", "praeproperus",
	"praesentia", "praesidium", "praestans", "praestantia", "praesto",
	"praesul", "praesum", "praesumo", "praeter", "praeterea", "praetereo",
	"praeteritus", "praetermissio", "praetorgredior", "praevenio",
	"pravitas", "pre", "preastolatio", "prebeo", "precedo", "precelsus",
	"precepio", "preceptum", "precipio", "precipue", "precipuus", "precis",
	"preclarus", "preconor", "precor", "precox", "preda", "predico", "preeo",
	"prefeci", "prefectum", "prefero", "preficio", "prefinio", "prefoco",
	"pregravo", "prehendo", "premium", "premo", "prenda", "prenuncius",
	"prenuntio", "prepono", "prepositus", "preproperus", "presencia",
	"presentia", "presidium", "pressi", "pressum", "prestans", "prestantia",
	"presto", "prestolatio", "presul", "presum", "presumo", "preterea",
	"pretereo", "pretergredior", "preteritus", "pretermissio", "pretium",
	"prevenio", "prevenire", "prex", "primitus", "primo", "primoris",
	"primumprimum", "princeps", "principatus", "principium", "prior",
	"priores", "priscus", "pristinus", "prius", "priusquam", "privatus",
	"privigna", "privo", "privus", "pro", "probitas", "probo", "procedo",
	"procella", "procer", "procinctu", "procul", "procurator", "prodigiosus",
	"proditor", "proelium", "profari", "profatus", "profecto", "profero",
	"proficio", "proficiscor", "proficuusprofiteor", "profiteor", "profor",
	"profugus", "profundo", "profundum", "profundus", "profusum",
	"profuturus", "progener", "progenero", "progenies", "progenitor",
	"progenitum", "progenui", "progigno", "prognatus", "progredior",
	"progressio", "progressus", "prohibeo", "prohibitio", "proicio",
	"proinde", "prolabor", "prolapsio", "prolatio", "prolato", "prolecto",
	"prolesproletarius", "prolicio", "prolix", "prolixi", "prolixus",
	"proloquor", "prolu", "proluo", "prolusio", "prolutum", "proluvier",
	"promereo", "promereor", "promeritumprominens", "promineo", "promisce",
	"promiscue", "promiscus", "promiscuus", "promissio", "promissor",
	"promitto", "promo", "promontorium", "promotum", "promoveo", "promovi",
	"prompsi", "prompte", "promptum", "promptupromptupromptupromptus",
	"promptus", "promulgatio", "promus", "promutuus", "pronepos", "proneptos",
	"pronuntio", "prope", "propello", "propero", "propinquo", "propinquus",
	"propono", "propositum", "proprie", "proprius", "propter",
	"propugnaculum", "prorsus", "prosecutus", "prosequor", "prosper",
	"prosperitas", "prosum", "protesto", "protestor", "protinus",
	"protractus", "protraho", "prout", "provectus", "proveho", "proventus",
	"provideo", "provisor", "provolvere", "proximus", "prudens", "prudenter",
	"prudentia", "publica", "publicus", "puchre", "pudendus", "pudeo",
	"pudicus", "pudor", "puella", "puer", "puerilis", "pueriliter", "puga",
	"pugna", "pugnacitas", "pugnaculum", "pugnax", "pugno", "pugnus",
	"pulchellus", "pulcher", "pulchra", "pulchritudinis", "pulchritudo",
	"pulchrum", "pulex", "pullulo", "pullus", "pulmentum", "pulmo", "pulpa",
	"pulpitum", "pulso", "pulsum", "pulsus", "pulvis", "pumilio", "pumilius",
	"punctum", "pungo", "puniceus", "punio", "punitor", "pupa", "pupilla",
	"pupillus", "puppis", "pupugi", "pupula", "purgamentum", "purgatio",
	"purgo", "purpura", "purus", "pusillus", "putator", "puteo", "puter",
	"putesco", "puteus", "puto", "putus", "pyga", "pyropus", "pyus", "pyxidis",
	"qua", "quadraginta", "quadratus", "quadrigae", "quadrivium", "quadrum",
	"quadruplator", "quadruplor", "quae", "quaedam", "quaenam", "quaero",
	"quaesitio", "quaeso", "quaestio", "quaestuosus", "quaestus", "qualis",
	"qualiscumque", "qualislibet", "qualitas", "qualiter", "quam", "quamdiu",
	"quamobrem", "quamquam", "quamtotius", "quamvis", "quandoquandoquidem",
	"quantacumque", "quanti", "quanto", "quantocius", "quantotius", "quantum",
	"quantumcumque", "quantus", "quantuscumque", "quantuslibet", "quantuvis",
	"quapropter", "quare", "quartus", "quarum", "quas", "quasi", "quassatio",
	"quasso", "quatenus", "quater", "quatinus", "quattuor", "que", "quedam",
	"quem", "quemadmodum", "quenam", "queo", "quercetum", "quercus", "quereia",
	"querella", "queribundus", "querimonia", "queritor", "quernus", "quero",
	"queror", "querulus", "quesitio", "queso", "questio", "questuosus",
	"questus", "qui", "quia", "quibus", "quicquid", "quidam", "quidem",
	"quidemquies", "quidnam", "quidquidam", "quietis", "quilibet", "quin",
	"quinam", "quingenti", "quinquennis", "quippe", "quisnam", "quisquam",
	"quisque", "quisquis", "quo", "quod", "quodammodo", "quodnam", "quomodo",
	"quondam", "quoniam", "quoque", "quorum", "quos", "quot", "quotiens",
	"quotienscumque", "quovis", "quris", "radicitus", "rapio", "rapui",
	"rarus", "ratio", "ratum", "re", "recedo", "recepi", "receptum", "recipio",
	"recito", "recognosco", "recolo", "reconcilio", "recondo", "recordatio",
	"recordor", "recro", "rectum", "rectus", "recuperatio", "recupero",
	"recuso", "redactum", "redarguo", "reddere", "reddo", "redemptio",
	"redemptor", "redeo", "redigo", "redono", "reduco", "redundantia",
	"redundo", "refectorium", "refero", "reformo", "regina", "regius",
	"regnum", "rego", "regula", "rei", "relaxo", "relectum", "relegi", "relego",
	"relevo", "relictus", "relinquo", "reliquum", "relucesco", "reluctor",
	"rem", "remando", "remaneo", "rememdium", "removeo", "remuneror",
	"renuntio", "renuo", "rependo", "repens", "repente", "repere", "reperio",
	"repetitio", "repeto", "repleo", "repletus", "repo", "repono",
	"reprehendo", "repsi", "reptum", "repugno", "requiesco", "requietum",
	"requievi", "requiro", "res", "resisto", "respicio", "respondeo",
	"respondi", "responsum", "restituo", "resumo", "resumps", "resumptum",
	"retineo", "retractum", "retraho", "retraxi", "retribuo", "reus",
	"revenio", "reversus", "reverti", "reverto", "revertor", "revoco",
	"revolvo", "rex", "rexi", "rgis", "rhetor", "rhetoricus", "rideo", "rigor",
	"risi", "risum", "ritus", "rogo", "rostrum", "rota", "rotundus", "rubor",
	"rudimentum", "rui", "rumor", "ruo", "ruris", "rursus", "rus", "rusticus",
	"rutum", "sabbatum", "sacculus", "sacrificum", "sacrilegus", "saepe",
	"saepenumero", "saepius", "saeta", "saevio", "sal", "salis", "salsus",
	"saltem", "salus", "saluto", "salutor", "salveo", "salvus", "sanctifico",
	"sanctimonia", "sanctimonialis", "sanctus", "sane", "sanitas", "sano",
	"sanus", "sapiens", "sapienter", "sapientia", "sarcina", "satago", "satio",
	"satis", "sato", "satura", "saturo", "scaber", "scabies", "scamnum",
	"scaphium", "sceleratus", "sceleris", "scelero", "scelestus", "scelus",
	"schola", "scientia", "scilicet", "scindo", "scio", "sciphu", "scisco",
	"scitum", "scivi", "scribo", "scrinium", "scripsi", "scriptor", "scriptum",
	"se", "secedo", "secerno", "secretum", "secrevi", "secundum", "secundus",
	"securus", "secus", "secutus", "secuutus", "sed", "sedeo", "sedi",
	"seditio", "sedo", "seductor", "semel", "semper", "senectus", "senex",
	"senis", "sensus", "sententia", "sentio", "seorsum", "sepe", "sepelio",
	"sepius", "septem", "sepulchrum", "seputus", "sequax", "sequi", "sequor",
	"serio", "serius", "sermo", "sero", "servio", "servitus", "servo", "servus",
	"seseseveritas", "sessum", "seu", "si", "sibi", "sibimet", "sic", "siccus",
	"sicut", "sidus", "signum", "silens", "silenti", "silentium", "sileo",
	"siliginis", "siligo", "silva", "similis", "similitudo", "simplex",
	"simul", "simulatio", "sine", "singularis", "singuli", "singultim",
	"singultus", "singulus", "sino", "siquidem", "sitio", "sitis", "sive",
	"socer", "socius", "sodalitas", "sol", "soleo", "solio", "solis",
	"solitudinis", "solitudo", "solium", "sollers", "sollicito",
	"sollicitudo", "sollicitus", "solum", "solumsolum", "solus", "solutio",
	"solvo", "somniculosus", "somniculouse", "somnio", "somnium", "somnus",
	"sonitus", "sono", "sophismata", "sopor", "sordeo", "sordes", "sordesco",
	"sortitus", "spargo", "sparsi", "sparsum", "speciosus", "spectaculum",
	"specto", "speculum", "specus", "sperno", "spero", "spes", "spiculum",
	"spiritus", "spoliatio", "spolio", "spolium", "sponte", "spretum",
	"sprevi", "stabilis", "stabilitas", "statim", "statua", "statum", "statuo",
	"stella", "steti", "stillicidium", "stipes", "stipis", "stips", "sto",
	"strenuus", "strues", "studio", "studiose", "studium", "stultus", "sua",
	"suadeo", "suasoria", "sub", "subito", "subitus", "subiungo", "sublatum",
	"sublime", "subnecto", "subpono", "subseco", "subsequor", "substantia",
	"subvenio", "succedo", "succendo", "successio", "succurro", "sufficio",
	"suffoco", "suffragium", "suggero", "sui", "sulum", "sum", "summa",
	"summisse", "summissus", "summitto", "summopere", "summus", "sumo",
	"sumptus", "supellectilis", "supellex", "super", "superbia", "superbus",
	"superficies", "superfluo", "superior", "superna", "superne", "supernus",
	"supero", "supersum", "superus", "supervacuussuppellex",
	"suppellectilus", "supplanto", "supplex", "supplicium", "suppono",
	"supra", "supremus", "surculus", "surgo", "surrectum", "surrexi",
	"sursumsuscipio", "suscito", "suspendi", "suspendo", "suspensum",
	"sustineo", "sustuli", "suum", "suus", "synagoga", "tabella", "tabellae",
	"tabernus", "tabesco", "tabgo", "tabula", "taceo", "tactum", "tactus",
	"tacuitacitum", "taedium", "talio", "talionis", "talis", "talus", "tam",
	"tamdiu", "tamen", "tametsi", "tamisium", "tamquam", "tandem", "tanquam",
	"tantillus", "tantum", "tantummodo", "tantus", "tardus", "te", "tectum",
	"tedium", "tego", "temeritas", "temperantia", "tempero", "tempestas",
	"templum", "temporis", "temptatio", "tempus", "tenax", "tendo", "teneo",
	"tener", "tenera", "tenerum", "tenuis", "tenus", "tepesco", "tepidus",
	"ter", "terebro", "teres", "terga", "tergeo", "tergiversatio", "tergo",
	"tergum", "tergus", "termes", "terminatio", "termino", "terminus", "tero",
	"terra", "terrarum", "terreo", "territo", "terror", "tersi", "tersum",
	"tersus", "tertius", "testimonium", "testis", "tetigi", "texi", "texo",
	"textilis", "textor", "textrix", "textus", "thalassinus", "theatrum",
	"theca", "thema", "thematis", "theologus", "thermae", "thesaurus",
	"thesis", "thorax", "thymbra", "thymum", "tibi", "timeo", "timidus",
	"timor", "titulus", "tolero", "tollo", "tondeo", "tonsor", "tonsum",
	"torqueo", "torrens", "tot", "totidem", "totiens", "toties", "totondi",
	"totus", "tracto", "tractum", "tradidi", "traditum", "trado", "traho",
	"trans", "transeo", "transfero", "translatum", "transmitto", "transtuli",
	"traxi", "tredecim", "tremo", "trepide", "tres", "tria", "tribuo",
	"tricesimus", "triduana", "triduanus", "triduum", "triginta", "tripudio",
	"tristis", "tritum", "triumphus", "trivi", "trucido", "truculenter", "tu",
	"tubineus", "tui", "tuli", "tum", "tumultus", "tumulus", "tunc", "turba",
	"turbatio", "turbatus", "turbo", "turpe", "turpis", "tutamen", "tutaminis",
	"tutis", "tyrannus", "uberrime", "ubi", "ulciscor", "ullus", "ulterius",
	"ultio", "ultionis", "ultraultra", "um", "umbra", "umerus", "umquam", "una",
	"unde", "undique", "universe", "universi", "universitas", "universum",
	"universus", "unus", "urbanus", "urbis", "urbs", "uredo", "us", "usitas",
	"usque", "ustilo", "ustulo", "usus", "ut", "uter", "uterque", "uti",
	"utilis", "utilitas", "utique", "utor", "utpote", "utproinde", "utrimque",
	"utrius", "utroque", "utrum", "uxor", "vaco", "vacuus", "vado", "vae",
	"valde", "valens", "valeo", "valetudo", "validus", "valiturus", "vallum",
	"valui", "vapulus", "varietas", "varius", "vehemens", "vehementer", "vel",
	"velle", "velociter", "velox", "velum", "velut", "veni", "venia", "venio",
	"ventito", "ventosus", "ventum", "ventus", "venustas", "ver", "vera",
	"verbera", "verbum", "vere", "verecundia", "vereor", "vergo", "veris",
	"veritas", "vero", "versus", "verto", "verumptamen", "verumtamen", "verus",
	"vesco", "vescor", "vesica", "vesper", "vespera", "vespillo", "vester",
	"vestigium", "vestio", "vestis", "vestitum", "vestivi", "vestra", "vestri",
	"vestrum", "vetus", "via", "vici", "vicinus", "vicissitudo", "victor",
	"victoria", "victum", "victus", "videlicet", "video", "videor", "vidi",
	"viduata", "viduo", "vigilo", "vigor", "vilicus", "vilis", "vilitas",
	"villa", "villicus", "vinco", "vinculum", "vindicatum", "vindico",
	"vinitor", "vinum", "vir", "vires", "virga", "virgo", "viridis",
	"viriliter", "virtus", "vis", "viscus", "visum", "vita", "vitiosus",
	"vitium", "vito", "vivo", "vix", "vixi", "vobis", "vociferor", "vocis",
	"voco", "volaticus", "volatilis", "volens", "volo", "volpes", "voltur",
	"volturius", "volubilis", "volubiliter", "voluntarius", "voluntas",
	"volup", "voluptarius", "voluptas", "voluptuosus", "volutabrum", "volva",
	"vomer", "vomica", "vomito", "vorago", "vorax", "voro", "vos", "votum",
	"voveo", "vovi", "vox", "vulariter", "vulgaris", "vulgivagus", "vulgo",
	"vulgus", "vulnero", "vulnus", "vulpes", "vulticulus", "vultuosus",
	"vultur", "vulturius", "vultus", "vulva", "xiphias", "ymber"}
