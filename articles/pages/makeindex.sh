#!/bin/sh
echo '<ul>'
for art in $(find . -name "*.md"); do
    cfg="$(basename $art .md).cfg"
    cat <<EOF
<span style="float:left;"><a href="$(basename $art .md).html">$(grep title $cfg | cut -d\= -f2)</a></span><span style="float:right;">$(stat -c "%y" $art | cut -d\  -f1)</span><br>
EOF
done

echo '</ul>'
