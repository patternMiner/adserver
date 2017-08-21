/**
 * Created by jbisa on 8/20/17.
 */

'use strict';

import { AdService } from './adservice.js';
import { Slot } from './slot.js';


/**
 * Bootstraps the adtag api, and synchronizes with the loading of adtags library.
 *
 * adtag api:
 *
 *   adtag.defineSlot(adUnit, width, height, divId)
 *
 *   adtag.queue(queueItem)
 *
 *   adtag.display(divId)
 *
 * @param doc {HTMLDocument}
 */
(function(doc) {
    /**
     * Finds the ad slot by divId, then arranges for the retrieval and rendition of a suitable ad creative.
     * @param divId {string}
     * @private
     */
    let display = (divId) => {
        let slot = Slot.find(divId);
        if (slot) {
            AdService.showAd(doc, slot);
        }
    }

    /**
     * Executes a given queue item function.
     * @param f {function}
     */
    let execute = (f) => f.call();

    if (window['adtag']) { // page loaded before the adtag library finished loading.
        let adtag = window['adtag'];
        adtag['defineSlot'] = Slot.define;
        adtag['display'] = display;
        adtag['queue'].map(execute);
    } else {
        let adtag = { // adtag library loaded before the page body finished loading.
            queue: {push: execute},
            defineSlot: Slot.define,
            display: display
        };
        window['adtag'] = adtag;
    }
})(document);
