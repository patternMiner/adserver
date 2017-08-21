/**
 * Created by jbisa on 8/20/17.
 */

'use strict';

const AD_URL = 'http://localhost:8080/ad';

export class AdService {
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
                console.log(ad);
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
                xhr.send();
            });
        }

        doc.getElementById(slot.getDivId()).appendChild(i);

        let adUrl = `${AD_URL}?adunit=${slot.getAdUnit()}&width=${slot.getWidth()}&height=${slot.getHeight()}`;

        fetchContent(adUrl).then(processContent, processContent);
    }
}
