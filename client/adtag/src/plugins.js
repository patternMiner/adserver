/**
 * Created by jbisa on 8/20/17.
 */

'use strict';

// Avoid `console` errors in browsers that lack a console.
(function() {
    let methods = [
        'assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error',
        'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log',
        'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd',
        'timeline', 'timelineEnd', 'timeStamp', 'trace', 'warn'
    ];
    let noop = () => {};
    let console = (window.console = window.console || {});

    methods
        .filter((method) => !console[method])
        .map((method) => console[method] = noop);
}());
