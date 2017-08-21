/**
 * Created by jbisa on 8/20/17.
 */

'use strict';

import { Slot } from './slot.js';

// bootstrap the system.
(function(doc) {
    let display = (divId) => Slot.display(divId, doc),
        execute = (f) => f.call();
    if (window['adtag']) { // page loaded before the adtag library
        let adtag = window['adtag'];
        adtag['defineSlot'] = Slot.define;
        adtag['display'] = display;
        adtag['queue'].map(execute);
    } else {
        let adtag = { // adtag library loaded before the page body
            queue: {push: execute},
            defineSlot: Slot.define,
            display: display
        };
        window['adtag'] = adtag;
    }
})(document);
