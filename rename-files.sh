#/bin/bash

IFS=$'\n'; set -f

for f in $(find ./web/static/gallery -name '*.png'); do 
    fileName=$(echo "$f" | tr '_' '-' | sed -e 's/-th//')
    if [ $f != $fileName ]; then
        mv $f $fileName
    fi
done

unset IFS; set +f