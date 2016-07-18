(function() {
    'use strict';


    var i18nMixin = {
        langs: null,
        i18n: function(key) {
            return this['langs'][key] || key;
        }
    };

    var langs = document.getElementById('langs');
    if (langs) {
        i18nMixin.langs = JSON.parse(langs.innerHTML);
    } else {
        i18nMixin.langs = {};
    }

    riot.mixin('i18n', i18nMixin);


    riot.mount('*');
})();