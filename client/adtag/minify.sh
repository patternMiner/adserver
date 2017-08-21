#!/bin/bash

#
# Created by jbisa on 8/20/17.
#

/Applications/WebStorm.app/Contents/jdk/Contents/Home/jre/bin/java -jar \
    /Users/jbisa/closure-compiler/closure-compiler-v20170806.jar \
    --language_in ECMASCRIPT6 \
    --compilation_level ADVANCED_OPTIMIZATIONS \
    --js src/plugins.js src/adservice.js src/slot.js src/main.js \
    --js_output_file dist/adtag_min.js
