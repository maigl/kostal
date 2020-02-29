#!/bin/bash

color_picker_url=https://coolors.co/c41b5c-08415c-6b818c-f1bf98-eee5e9

color1=$(echo $color_picker_url | sed 's@https://coolors.co/@@' | awk -F'-' '{print $1}')
color2=$(echo $color_picker_url | sed 's@https://coolors.co/@@' | awk -F'-' '{print $2}')
color3=$(echo $color_picker_url | sed 's@https://coolors.co/@@' | awk -F'-' '{print $3}')
color4=$(echo $color_picker_url | sed 's@https://coolors.co/@@' | awk -F'-' '{print $4}')

html=$(cat web/frame.html | sed "s/COLOR1/$color1/" | sed "s/COLOR2/$color2/" | sed "s/COLOR3/$color3/" | sed "s/COLOR4/$color4/" )

cat << EOF
package frame

var html = \`
$html
\`
EOF

