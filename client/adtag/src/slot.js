/**
 * Created by jbisa on 8/20/17.
 */

'use strict';

import { AdService } from './adservice.js';

let _slots = {}

export class Slot {
    constructor(adUnit, width, height, divId) {
        this.adUnit = adUnit;
        this.height = height;
        this.width = width;
        this.divId = divId;
    }

    getAdUnit() {
        return this.adUnit;
    }

    getHeight() {
        return this.height;
    }

    getWidth() {
        return this.width;
    }

    getDivId() {
        return this.divId;
    }

    // create a slot and add it to the internal slot list.
    static define(adUnit, width, height, divId) {
        _slots[divId] = new Slot(adUnit, width, height, divId);
    }

    // find the given slot by divId, then fetch and render the ad in an iframe.
    static display(divId, doc) {
        let slot = _slots[divId];
        if (slot) {
            AdService.showAd(doc, slot);
        }
    }
}
