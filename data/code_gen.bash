#/bin/bash

datafile="./kostal.data.csv"

cat << EOF
package data

// GENERATED FILE: see code_gen.bash
var Registers = map[string]*Register{
EOF

while IFS= read -r line
do
    IFS=' ' read -r -a tmp <<< "$line"
    len="${#tmp[@]}"
    description="${tmp[2]}"
    for i in $(seq 3 $((len-6))); do
        description+=" ${tmp[$i]}"
    done

cat << EOF
    "${tmp[1]}": &Register{
        Addr: ${tmp[1]},
        Unit: "${tmp[(($len-5))]}",
        Format: "${tmp[(($len-4))]}",
        Length: ${tmp[(($len-3))]},
        Description: "$description",
    },
EOF
done < $datafile

echo "}"
