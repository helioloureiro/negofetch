package main

import (
	"fmt"
	"strings"

	"github.com/helioloureiro/golorama"
)

func macOSLogo() string {
	// set_colors 2 3 1 1 5 4
	c := setColors(2, 3, 1, 1, 5, 4)
	return `
 ` + c[1] + `                    'c.
                  ,xNMM.
                .OMMMMo
                OMMM0,
      .;loddo:' loolloddol;.
    cKMMMMMMMMMMNWMMMMMMMMMM0:
 ` + c[2] + ` .KMMMMMMMMMMMMMMMMMMMMMMMWd.
  XMMMMMMMMMMMMMMMMMMMMMMMX.
 ` + c[3] + `;MMMMMMMMMMMMMMMMMMMMMMMM:
 :MMMMMMMMMMMMMMMMMMMMMMMM:
 ` + c[4] + `.MMMMMMMMMMMMMMMMMMMMMMMMX.
  kMMMMMMMMMMMMMMMMMMMMMMMMWd.
  ` + c[5] + `.XMMMMMMMMMMMMMMMMMMMMMMMMMMk
   .XMMMMMMMMMMMMMMMMMMMMMMMMK.
     ` + c[6] + `kMMMMMMMMMMMMMMMMMMMMMMd
      ;KMMMMMMMWXXWMMMMMMMk.
        .cooc,.    .,coo:.
`
}

func getLogo(system string) string {
	// system name will be matched in lower case
	system = strings.ToLower(system)
	switch system {
	case "aix":
		c := setColors(2, 7)
		return c[1] + `
	       ':+ssssossossss+-'
        .oys///oyhddddhyo///sy+.
      /yo:+hNNNNNNNNNNNNNNNNh+:oy/
    :h/:yNNNNNNNNNNNNNNNNNNNNNNy-+h:
  'ys.yNNNNNNNNNNNNNNNNNNNNNNNNNNy.ys
 'h+-mNNNNNNNNNNNNNNNNNNNNNNNNNNNNm-oh
 h+-NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN.oy
/d'mNNNNNNN/::mNNNd::m+:/dNNNo::dNNNd'm:
h//NNNNNNN: . .NNNh  mNo  od. -dNNNNN:+y
N.sNNNNNN+ -N/ -NNh  mNNd.   sNNNNNNNo-m
N.sNNNNNs  +oo  /Nh  mNNs' ' /mNNNNNNo-m
h//NNNNh  ossss' +h  md- .hm/ 'sNNNNN:+y
:d'mNNN+/yNNNNNd//y//h//oNNNNy//sNNNd'm-
 yo-NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNm.ss
 'h+-mNNNNNNNNNNNNNNNNNNNNNNNNNNNNm-oy
   sy.yNNNNNNNNNNNNNNNNNNNNNNNNNNs.yo
    :h+-yNNNNNNNNNNNNNNNNNNNNNNs-oh-
      :ys:/yNNNNNNNNNNNNNNNmy/:sy:
        .+ys///osyhhhhys+///sy+.
            -/osssossossso/-	`
	case "hash":
		// it is set_color 123 on original...
		c := setColors(1, 2, 3)
		return c[1] + `
		+   ######   +
	  ###   ######   ###
	#####   ######   #####
   ######   ######   ######
  
  ####### '"###### '"########
  #######   ######   ########
  #######   ######   ########
  
   ###### '"###### '"######
	#####   ######   #####
	  ###   ######   ###
		~   ######   ~	`

	case "alpine_small":
		c := setColors(4, 7)
		return `
		` + c[1] + `   /\\ /\\
		/` + c[2] + `/ ` + c[1] + `\\  \\
	   /` + c[2] + `/   ` + c[1] + `\\  \\
	  /` + c[2] + `//    ` + c[1] + `\\  \\
	  ` + c[2] + `//      ` + c[1] + `\\  \\
			   \\
		`
	case "alpine":
		c := setColors(4, 5, 7, 6)
		return c[1] + `
		.hddddddddddddddddddddddh.
		:dddddddddddddddddddddddddd:
	   /dddddddddddddddddddddddddddd/
	  +dddddddddddddddddddddddddddddd+
	'sdddddddddddddddddddddddddddddddds'
   'ydddddddddddd++hdddddddddddddddddddy'
  .hddddddddddd+'  '+ddddh:-sdddddddddddh.
  hdddddddddd+'      '+y:    .sddddddddddh
  ddddddddh+'   '//'   '.'     -sddddddddd
  ddddddh+'   '/hddh/'   ':s-    -sddddddd
  ddddh+'   '/+/dddddh/'   '+s-    -sddddd
  ddd+'   '/o' :dddddddh/'   'oy-    .yddd
  hdddyo+ohddyosdddddddddho+oydddy++ohdddh
  .hddddddddddddddddddddddddddddddddddddh.
   'yddddddddddddddddddddddddddddddddddy'
	'sdddddddddddddddddddddddddddddddds'
	  +dddddddddddddddddddddddddddddd+
	   /dddddddddddddddddddddddddddd/
		:dddddddddddddddddddddddddd:
		 .hddddddddddddddddddddddh.		
`
	case "alter":
		c := setColors(6, 6)
		return `
` + c[1] + `                      %,
                    ^WWWw
                   'wwwwww
                  !wwwwwwww
                 #'wwwwwwwww
                @wwwwwwwwwwww
               wwwwwwwwwwwwwww
              wwwwwwwwwwwwwwwww
             wwwwwwwwwwwwwwwwwww
            wwwwwwwwwwwwwwwwwwww,
           w~1i.wwwwwwwwwwwwwwwww,
         3~:~1lli.wwwwwwwwwwwwwwww.
        :~~:~?ttttzwwwwwwwwwwwwwwww
       #<~:~~~~?llllltO-.wwwwwwwwwww
      #~:~~:~:~~?ltlltlttO-.wwwwwwwww
     @~:~~:~:~:~~(zttlltltlOda.wwwwwww
    @~:~~: ~:~~:~:(zltlltlO    a,wwwwww
   8~~:~~:~~~~:~~~~_1ltltu          ,www
  5~~:~~:~~:~~:~~:~~~_1ltq             N,,
 g~:~~:~~~:~~:~~:~:~~~~1q                N,		
`
	case "amazon":
		c := setColors(3, 7)
		return c[1] + `
		'-/oydNNdyo:.'
		'.:+shmMMMMMMMMMMMMMMmhs+:.'
	  -+hNNMMMMMMMMMMMMMMMMMMMMMMNNho-
  .''      -/+shmNNMMMMMMNNmhs+/-      ''.
  dNmhs+:.       '.:/oo/:.'       .:+shmNd
  dMMMMMMMNdhs+:..        ..:+shdNMMMMMMMd
  dMMMMMMMMMMMMMMNds    odNMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  dMMMMMMMMMMMMMMMMh    yMMMMMMMMMMMMMMMMd
  .:+ydNMMMMMMMMMMMh    yMMMMMMMMMMMNdy+:.
	   '.:+shNMMMMMh    yMMMMMNhs+:''
			  '-+shy    shs+:'
`
	case "anarchy":
		c := setColors(7, 4)
		return `
                         ` + c[2] + `..` + c[1] + `
                        ` + c[2] + `..` + c[1] + `
                      ` + c[2] + `:..` + c[1] + `
                    ` + c[2] + `:+++.` + c[1] + `
              .:::++` + c[2] + `++++` + c[1] + `+::.
          .:+######` + c[2] + `++++` + c[1] + `######+:.
       .+#########` + c[2] + `+++++` + c[1] + `##########:.
     .+##########` + c[2] + `+++++++` + c[1] + `##` + c[2] + `+` + c[1] + `#########+.
    +###########` + c[2] + `+++++++++` + c[1] + `############:
   +##########` + c[2] + `++++++` + c[1] + `#` + c[2] + `++++` + c[1] + `#` + c[2] + `+` + c[1] + `###########+
  +###########` + c[2] + `+++++` + c[1] + `###` + c[2] + `++++` + c[1] + `#` + c[2] + `+` + c[1] + `###########+
 :##########` + c[2] + `+` + c[1] + `#` + c[2] + `++++` + c[1] + `####` + c[2] + `++++` + c[1] + `#` + c[2] + `+` + c[1] + `############:
 ###########` + c[2] + `+++++` + c[1] + `#####` + c[2] + `+++++` + c[1] + `#` + c[2] + `+` + c[1] + `###` + c[2] + `++` + c[1] + `######+
.##########` + c[2] + `++++++` + c[1] + `#####` + c[2] + `++++++++++++` + c[1] + `#######.
.##########` + c[2] + `+++++++++++++++++++` + c[1] + `###########.
 #####` + c[2] + `++++++++++++++` + c[1] + `###` + c[2] + `++++++++` + c[1] + `#########+
 :###` + c[2] + `++++++++++` + c[1] + `#########` + c[2] + `+++++++` + c[1] + `#########:
  +######` + c[2] + `+++++` + c[1] + `##########` + c[2] + `++++++++` + c[1] + `#######+
   +####` + c[2] + `+++++` + c[1] + `###########` + c[2] + `+++++++++` + c[1] + `#####+
    :##` + c[2] + `++++++` + c[1] + `############` + c[2] + `++++++++++` + c[1] + `##:
     .` + c[2] + `++++++` + c[1] + `#############` + c[2] + `++++++++++` + c[1] + `+.
      :` + c[2] + `++++` + c[1] + `###############` + c[2] + `+++++++` + c[1] + `::
     .` + c[2] + `++. .:+` + c[1] + `##############` + c[2] + `+++++++` + c[1] + `..
     ` + c[2] + `.:.` + c[1] + `      ..::++++++::..:` + c[2] + `++++` + c[1] + `+.
     ` + c[2] + `.` + c[1] + `                       ` + c[2] + `.:+++` + c[1] + `.
                                ` + c[2] + `.:` + c[1] + `:
                                   ` + c[2] + `..` + c[1] + `
                                    ` + c[2] + `..` + c[1] + `

`
	case "android":
		c := setColors(2, 7)
		return c[1] + `         -o          o-
		+hydNNNNdyh+
	  +mMMMMMMMMMMMMm+
	'dMM` + c[2] + `m:` + c[1] + `NMMMMMMN` + c[2] + `:m` + c[1] + `MMd'
	hMMMMMMMMMMMMMMMMMMh
..  yyyyyyyyyyyyyyyyyyyy  ..
.mMMm'MMMMMMMMMMMMMMMMMMMM'mMMm.
:MMMM-MMMMMMMMMMMMMMMMMMMM-MMMM:
:MMMM-MMMMMMMMMMMMMMMMMMMM-MMMM:
:MMMM-MMMMMMMMMMMMMMMMMMMM-MMMM:
:MMMM-MMMMMMMMMMMMMMMMMMMM-MMMM:
-MMMM-MMMMMMMMMMMMMMMMMMMM-MMMM-
+yy+ MMMMMMMMMMMMMMMMMMMM +yy+
	mMMMMMMMMMMMMMMMMMMm
	'/++MMMMh++hMMMM++/'
		MMMMo  oMMMM
		MMMMo  oMMMM
		oNMm-  -mMNs
		`
	case "antergos":
		c := setColors(4, 6)
		return c[2] + `              '.-/::/-''
		.-/osssssssso/.
	   :osyysssssssyyys+-
	'.+yyyysssssssssyyyyy+.
   '/syyyyyssssssssssyyyyys-'
  '/yhyyyyysss` + c[1] + `++` + c[2] + `ssosyyyyhhy/'
 .ohhhyyyys` + c[1] + `o++/+o` + c[2] + `so` + c[1] + `+` + c[2] + `syy` + c[1] + `+` + c[2] + `shhhho.
.shhhhys` + c[1] + `oo++//+` + c[2] + `sss` + c[1] + `+++` + c[2] + `yyy` + c[1] + `+s` + c[2] + `hhhhs.
-yhhhhs` + c[1] + `+++++++o` + c[2] + `ssso` + c[1] + `+++` + c[2] + `yyy` + c[1] + `s+o` + c[2] + `hhddy:
-yddhhy` + c[1] + `o+++++o` + c[2] + `syyss` + c[1] + `++++` + c[2] + `yyy` + c[1] + `yooy` + c[2] + `hdddy-
.yddddhs` + c[1] + `o++o` + c[2] + `syyyyys` + c[1] + `+++++` + c[2] + `yyhh` + c[1] + `sos` + c[2] + `hddddy'
'odddddhyosyhyyyyyy` + c[1] + `++++++` + c[2] + `yhhhyosddddddo
.dmdddddhhhhhhhyyyo` + c[1] + `+++++` + c[2] + `shhhhhohddddmmh.
ddmmdddddhhhhhhhso` + c[1] + `++++++` + c[2] + `yhhhhhhdddddmmdy
dmmmdddddddhhhyso` + c[1] + `++++++` + c[2] + `shhhhhddddddmmmmh
-dmmmdddddddhhys` + c[1] + `o++++o` + c[2] + `shhhhdddddddmmmmd-
.smmmmddddddddhhhhhhhhhdddddddddmmmms.
'+ydmmmdddddddddddddddddddmmmmdy/.
  '.:+ooyyddddddddddddyyso+:.'
  `

	case "antix":
		c := setColors(1, 7, 6)
		return c[1] + `
                    \
         , - ~ ^ ~ - \        /
     , '              \ ' ,  /
   ,                   \   '/
  ,                     \  / ,
 ,___,                   \/   ,
 /   |   _  _  _|_ o     /\   ,
|,   |  / |/ |  |  |    /  \  ,
 \,_/\_/  |  |_/|_/|_/_/    \,
   ,                  /     ,\
     ,               /  , '   \
      ' - , _ _ _ ,  '

`
	case "arcolinux":
		c := setColors(7, 4)
		return c[2] + `                    /-
                   ooo:
                  yoooo/
                 yooooooo
                yooooooooo
               yooooooooooo
             .yooooooooooooo
            .oooooooooooooooo
           .oooooooarcoooooooo
          .ooooooooo-oooooooooo
         .ooooooooo-  oooooooooo
        :ooooooooo.    :ooooooooo
       :ooooooooo.      :ooooooooo
      :oooarcooo         .oooarcooo
     :ooooooooy           .ooooooooo
    :ooooooooo   ` + c[1] + `/ooooooooooooooooooo` + c[2] + `
   :ooooooooo      ` + c[1] + `.-ooooooooooooooooo.` + c[2] + `
  ooooooooo-             ` + c[1] + `-ooooooooooooo.` + c[2] + `
 ooooooooo-                 ` + c[1] + `.-oooooooooo.` + c[2] + `
ooooooooo.                     ` + c[1] + `-ooooooooo` + c[2] + `
`
	case "archbox":
		c := setColors(2, 7, 1)
		return c[1] + `              ...:+oh/:::..
         ..-/oshhhhhh'   '::::-.
     .:/ohhhhhhhhhhhh'        '-::::.
 .+shhhhhhhhhhhhhhhhh'             '.::-.
 /'-:+shhhhhhhhhhhhhh'            .-/+shh
 /      .:/ohhhhhhhhh'       .:/ohhhhhhhh
 /           '-:+shhh'  ..:+shhhhhhhhhhhh
 /                 .:ohhhhhhhhhhhhhhhhhhh
 /                  'hhhhhhhhhhhhhhhhhhhh
 /                  'hhhhhhhhhhhhhhhhhhhh
 /                  'hhhhhhhhhhhhhhhhhhhh
 /                  'hhhhhhhhhhhhhhhhhhhh
 /      .+o+        'hhhhhhhhhhhhhhhhhhhh
 /     -hhhhh       'hhhhhhhhhhhhhhhhhhhh
 /     ohhhhho      'hhhhhhhhhhhhhhhhhhhh
 /:::+'hhhhoos'     'hhhhhhhhhhhhhhhhhs+'
    '--/:'   /:     'hhhhhhhhhhhho/-
             -/:.   'hhhhhhs+:-'
                ::::/ho/-'
			   `
	case "archlabs":
		c := setColors(6, 6, 7, 1)
		return c[1] + `                     'c'
                    'kKk,
                   .dKKKx.
                  .oKXKXKd.
                 .l0XXXXKKo.
                 c0KXXXXKX0l.
                :0XKKOxxOKX0l.
               :OXKOc. .c0XX0l.
              :OK0o. ` + c[4] + `...` + c[1] + `'dKKX0l.
             :OX0c  ` + c[4] + `;xOx'` + c[1] + `'dKXX0l.
            :0KKo.` + c[4] + `.o0XXKd'.` + c[1] + `lKXX0l.
           c0XKd.` + c[4] + `.oKXXXXKd..` + c[1] + `oKKX0l.
         .c0XKk;` + c[4] + `.l0K0OO0XKd..` + c[1] + `oKXXKo.
        .l0XXXk:` + c[4] + `,dKx,.'l0XKo.` + c[1] + `.kXXXKo.
       .o0XXXX0d,` + c[4] + `:x;   .oKKx'` + c[1] + `.dXKXXKd.
      .oKXXXXKK0c.` + c[4] + `;.    :00c'` + c[1] + `cOXXXXXKd.
     .dKXXXXXXXXk,` + c[4] + `.     cKx'` + c[1] + `'xKXXXXXXKx'
    'xKXXXXK0kdl:.     ` + c[4] + `.ok; ` + c[1] + `.cdk0KKXXXKx'
   'xKK0koc,..         ` + c[4] + `'c, ` + c[1] + `    ..,cok0KKk,
  ,xko:'.             ` + c[4] + `.. ` + c[1] + `           .':okx;
 .,'.                                   .',.

`
	case "archstrike":
		c := setColors(8, 6)
		return c[1] +
			`                   *   
                  **.
                 ****
                ******
                *******
              ** *******
             **** *******
            ` + c[1] + `****` + c[2] + `_____` + c[1] + `***` + c[2] + `/` + c[1] + `*
           ***` + c[2] + `/` + c[1] + `*******` + c[2] + `//` + c[1] + `***
          **` + c[2] + `/` + c[1] + `********` + c[2] + `///` + c[1] + `*` + c[2] + `/` + c[1] + `**
         **` + c[2] + `/` + c[1] + `*******` + c[2] + `////` + c[1] + `***` + c[2] + `/` + c[1] + `**
        **` + c[2] + `/` + c[1] + `****` + c[2] + `//////.,` + c[1] + `****` + c[2] + `/` + c[1] + `**
       ***` + c[2] + `/` + c[1] + `*****` + c[2] + `/////////` + c[1] + `**` + c[2] + `/` + c[1] + `***
      ****` + c[2] + `/` + c[1] + `****    ` + c[2] + `/////` + c[1] + `***` + c[2] + `/` + c[1] + `****
     ******` + c[2] + `/` + c[1] + `***  ` + c[2] + `////   ` + c[1] + `**` + c[2] + `/` + c[1] + `******
    ********` + c[2] + `/` + c[1] + `* ` + c[2] + `///      ` + c[1] + `*` + c[2] + `/` + c[1] + `********
  ,******     ` + c[2] + `// ______ /    ` + c[1] + `******,

		`
	case "arch":
		c := setColors(6, 6, 7, 1)
		return c[1] +
			`                   -'
                  .o+'
                 'ooo/
                '+oooo:
               '+oooooo:
               -+oooooo+:
             '/:-:++oooo+:
            '/++++/+++++++:
           '/++++++++++++++:
          '/+++o` + c[2] + `oooooooo` + c[1] + `oooo/'
` + c[2] + `         ` + c[1] + `./` + c[2] + `ooosssso++osssssso` + c[1] + `+'
` + c[2] + `        .oossssso-''''/ossssss+'
       -osssssso.      :ssssssso.
      :osssssss/        osssso+++.
     /ossssssss/        +ssssooo/-
   '/ossssso+/:-        -:/+osssso+-
  '+sso+:-'                 '.-/+oso:
 '++:.                           '-/+/
 .'                                 '/
`
	case "macos":
		return macOSLogo()
	default:
		return fmt.Sprintf("System not found: %s", system)
	}
}

func printLogo(logo string) {
	fmt.Println(logo)
}

// setColors: it uses almost same schema colors as
// from original neofetch, but as array
func setColors(colors ...int) []string {
	result := make([]string, len(colors)+1)
	result[0] = ""
	for id, colorID := range colors {
		result[id+1] = colorConverter(colorID)
	}
	return result
}

// colorConverter: colors schema isn't matching with golorama
// so a conversion is needed
func colorConverter(color int) string {
	switch color {
	case LIGHTRED:
		return golorama.GetCSI(golorama.LIGHTRED)
	case LIGHTGREEN:
		return golorama.GetCSI(golorama.LIGHTGREEN)
	case LIGHTYELLOW:
		return golorama.GetCSI(golorama.LIGHTYELLOW)
	case LIGHTBLUE:
		return golorama.GetCSI(golorama.LIGHTBLUE)
	case LIGHTMAGENTA:
		return golorama.GetCSI(golorama.LIGHTMAGENTA)
	case LIGHTCYAN:
		return golorama.GetCSI(golorama.LIGHTCYAN)
	case RED:
		return golorama.GetCSI(golorama.RED)
	case GREEN:
		return golorama.GetCSI(golorama.GREEN)
	case YELLOW:
		return golorama.GetCSI(golorama.YELLOW)
	case BLUE:
		return golorama.GetCSI(golorama.BLUE)
	case MAGENTA:
		return golorama.GetCSI(golorama.MAGENTA)
	case CYAN:
		return golorama.GetCSI(golorama.CYAN)
	default:
		return golorama.Reset()
	}
}

func reset() string {
	return golorama.Reset()
}
