#!/bin/bash

#
# Created by jbisa on 8/20/17.
#

/usr/local/bin/closure-compiler \
    --language_in ECMASCRIPT6 \
    --compilation_level ADVANCED_OPTIMIZATIONS \
    --js src/plugins.js src/adservice.js src/slot.js src/main.js \
    --js_output_file dist/adtag_min.js

