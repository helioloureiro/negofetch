package main

import (
	"fmt"

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
	switch system {
	case "AIX":
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
            -/osssossossso/-	
	`
	case "Hash":
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
      ~   ######   ~
	`

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
	case "macOS":
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
	default:
		return golorama.Reset()
	}
}

func reset() string {
	return golorama.Reset()
}
