/**
 * Created by jbisa on 8/20/17.
 */

'use strict';


let _slots = {};

/**
 * Encapsulates the adunit, size and placement parameters of an ad slot.
 */
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

    /**
     * Creates a slot and adds it to the internal slot list.
     * @param adUnit {string}
     * @param width {string}
     * @param height {string}
     * @param divId {string}
     */
    static define(adUnit, width, height, divId) {
        _slots[divId] = new Slot(adUnit, width, height, divId);
    }

    /**
     * Find and return the Slot by its divId.
     * @param divId {string}
     * @returns {Slot}
     */
    static find(divId) {
        return _slots[divId];
    }
}
