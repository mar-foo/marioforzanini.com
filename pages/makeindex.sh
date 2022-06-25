#!/bin/sh
for art in $(find ../articles/pages/ -name "*.md"); do
    cfg="$(basename "$art" .md).cfg"
    html="$(basename "$art" .md).html"
    cat <<EOF
<span style="float:left;"><a href="/articles/$html">$(grep title ../articles/pages/"$cfg" | cut -d\= -f2)</a></span><span style="float:right;margin-right: 5px;">$(stat -c "%y" "$art" | cut -d\  -f1)</span><br>
EOF
done
