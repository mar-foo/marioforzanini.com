#!/bin/sh
for art in $(find . -name "*.md"); do
    cfg="$(basename $art .md).cfg"
    cat <<EOF
<span style="float:left;"><a href="$(basename $art .md).html">$(grep title $cfg | cut -d\= -f2)</a></span><span style="float:right; margin-right: 5px;">$(grep "^date" "$cfg" | cut -d= -f2)</span><br>
EOF
done
