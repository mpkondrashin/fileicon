package fileicon

import (
	"fmt"
	"strings"
)

var svgData = `<svg
version="1.1"
id="svg1"
width="96"
height="128"
viewBox="0 0 96 128"
sodipodi:docname="vivid.ai"
inkscape:version="1.3.2 (091e20e, 2023-11-25)"
xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
xmlns:xlink="http://www.w3.org/1999/xlink"
xmlns="http://www.w3.org/2000/svg"
xmlns:svg="http://www.w3.org/2000/svg">
  <defs
     id="defs1">
    <clipPath
       clipPathUnits="userSpaceOnUse"
       id="clipPath2">
      <path
         d="M 0,96 H 72 V 0 H 0 Z"
         transform="translate(0,-94.000002)"
         id="path2" />
    </clipPath>
    <clipPath
       clipPathUnits="userSpaceOnUse"
       id="clipPath4">
      <path
         d="M 0,96 H 72 V 0 H 0 Z"
         transform="translate(-71.940402,-72.000002)"
         id="path4" />
    </clipPath>
    <clipPath
       clipPathUnits="userSpaceOnUse"
       id="clipPath6">
      <path
         d="M 0,96 H 72 V 0 H 0 Z"
         id="path6" />
    </clipPath>
    <clipPath
       clipPathUnits="userSpaceOnUse"
       id="clipPath9">
      <path
         d="M 0,96 H 72 V 0 H 0 Z"
         transform="translate(-51.750001,-47.500001)"
         id="path9" />
    </clipPath>
    <clipPath
       clipPathUnits="userSpaceOnUse"
       id="clipPath11">
      <path
         d="M 0,96 H 72 V 0 H 0 Z"
         transform="translate(-43.000001,-30.700201)"
         id="path11" />
    </clipPath>
  </defs>
  <sodipodi:namedview
     id="namedview1"
     pagecolor="#ffffff"
     bordercolor="#000000"
     borderopacity="0.25"
     inkscape:showpageshadow="2"
     inkscape:pageopacity="0.0"
     inkscape:pagecheckerboard="0"
     inkscape:deskcolor="#d1d1d1"
     inkscape:zoom="1.84375"
     inkscape:cx="47.728814"
     inkscape:cy="64"
     inkscape:window-width="1408"
     inkscape:window-height="449"
     inkscape:window-x="0"
     inkscape:window-y="38"
     inkscape:window-maximized="0"
     inkscape:current-layer="layer-MC0">
    <inkscape:page
       x="0"
       y="0"
       inkscape:label="1"
       id="page1"
       width="96"
       height="128"
       margin="0"
       bleed="0" />
  </sodipodi:namedview>
  <g
     id="layer-MC0"
     inkscape:groupmode="layer"
     inkscape:label="Layer 1">
    <path
       id="path1"
       d="m 0,0 v -92 c 0,-1.1 0.9,-2 2,-2 h 68 c 1.104,0 2,0.9 2,2 v 66 H 50 c -1.104,0 -2,0.896 -2,2 V 2 H 2 C 0.9,2 0,1.104 0,0"
       style="fill:#666666 icc-color(sRGB-IEC61966-2, 0.1, 0.3999939, 0.3999939, 0.3999939);fill-opacity:1;fill-rule:nonzero;stroke:none"
       transform="matrix(1.3333333,0,0,-1.3333333,0,2.6666667)"
       clip-path="url(#clipPath2)" />
    <path
       id="path3"
       d="M 0,0 C -0.061,0.368 -0.196,0.724 -0.448,1.004 L -19.94,22.764 V 0 Z"
       style="fill:#666666 icc-color(sRGB-IEC61966-2, 0.1, 0.3999939, 0.3999939, 0.3999939);fill-opacity:1;fill-rule:nonzero;stroke:none"
       transform="matrix(1.3333333,0,0,-1.3333333,95.920533,32)"
       clip-path="url(#clipPath4)" />
    <path
       id="path5"
       d="M 66,6 H 6 v 49 h 60 z"
       style="fill:#ffffff icc-color(sRGB-IEC61966-2, 0.1, 1, 1, 1);fill-opacity:1;fill-rule:nonzero;stroke:none"
       transform="matrix(1.3333333,0,0,-1.3333333,0,128)"
       clip-path="url(#clipPath6)" />
  </g>
    <text x="16" y="95"  font-size="%d" style="fill:white;font-weight:bold;font-family:Arial">%s</text>

</svg>`

// 3 - 30
// 4 - 22
// 5 - 17
func VanilaIcon(ext string) string {
	size := 90 / len(ext)
	if size > 45 {
		size = 45
	}
	return fmt.Sprintf(svgData, size, strings.ToUpper(ext))
}
