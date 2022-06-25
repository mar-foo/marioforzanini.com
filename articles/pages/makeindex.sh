#!/bin/sh
for art in $(find . -name "*.md"); do
    cfg="$(basename $art .md).cfg"
    cat <<EOF
<span style="float:left;"><a href="$(basename $art .md).html">$(grep title $cfg | cut -d\= -f2)</a></span><span style="float:right; margin-right: 5px;">$(stat -c "%y" $art | cut -d\  -f1)</span><br>
EOF
done
