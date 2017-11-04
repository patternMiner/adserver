/**
 * Created by jbisa on 8/20/17.
 */

'use strict';

/** @define {string} */
var PROTOCOL = 'http';

/** @define {string} */
var DOMAIN_NAME = 'localhost';

/** @define {string} */
var PORT = '8080';

/** @define {string} */
var SERVICE_NAME = 'ad';

/**
 * Service to asynchronously fetch and render ad creatives based on adunit, size, and targeting parameters.
 */
export class AdService {
    /**
     * Fetches and renders the ad creative in an iframe on the given doc.
     * The ad creative is fetched based on the adunit, size, and targeting data specific to the given slot.
     */
    static showAd(doc, slot) {
        let i = doc.createElement('iframe');
        i.scrolling = 'auto';
        i.style = 'border: none;';
        i.src = 'about:self';
        i.width = slot.getWidth();
        i.height = slot.getHeight();

        let processContent = (responseText) => {
            let jsonResponse = JSON.parse(responseText);
            if(jsonResponse['Items']) {
                let ad = jsonResponse['Items'][0];
                i.src = ad['Url'];
            }
        }

        let fetchContent = (url) => {
            let xhr = new XMLHttpRequest();
            xhr.overrideMimeType("application/json");
            return new Promise((resolve, reject) => {
                xhr.onreadystatechange = function () {
                    if (xhr.readyState === 4) {
                        if (xhr.status === 200) {
                            resolve(xhr.responseText);
                        } else {
                            reject(xhr.responseText);
                        }
                    }
                };
                xhr.open('GET', url, true);
                try {
                    xhr.send();
                } catch (e) {
                    console.log(e);
                }
            });
        }

        doc.getElementById(slot.getDivId()).appendChild(i);

        // Format the ad fetch url, and kick off the asynchronous fetch of the creative.
        let adUrl = `${PROTOCOL}://${DOMAIN_NAME}:${PORT}/${SERVICE_NAME}`,
            parameters = `?adunit=${slot.getAdUnit()}&width=${slot.getWidth()}&height=${slot.getHeight()}`;
        fetchContent(adUrl + parameters).then(processContent, processContent);
    }
}
