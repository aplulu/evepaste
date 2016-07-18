<fit-export>
    <div class="form-group">
        <label>{ i18n('export_format') }</label>
        <select name="export_format" class="form-control" onclick={ changeExportFormat }>
            <option value="eft">{ i18n('eft_ingame') }</option>
            <option value="crest" selected={ opts['logged'] == 'true' } disabled={ opts['logged'] != 'true' }>{ i18n('crest_login_required') }</option>
            <option value="clf">{ i18n('clf') }</option>
            <option value="dna">{ i18n('dna') }</option>
        </select>
    </div>

    <div class="form-group" if={ exportFormat == 'eft' }>
        <label>{ i18n('language') }</label>
        <select name="eft_lang" class="form-control" onclick={ changeEftLang }>
            <option each={lang in eveLangs} value="{lang}" selected={ 'en' == lang }>{ i18n(lang) }</option>
        </select>
    </div>

    <div class="form-group">
        <button class="btn btn-default" onclick={ export }>{ i18n('export') }</button>
    </div>

    <style scoped>
    </style>

    <script>
        var self = this;
        this.exportFormat = 'eft';
        this.eftLang = 'en';
        this.eveLangs = [];

        this.on('mount', function() {
            self.eveLangs = self.opts['eveLangs'].split('|');
            self.eftLang = self.opts['curLang'];
            if (self.opts['logged'] == 'true') {
                self.exportFormat = 'crest';
            }
            self.update();
        });

        this.changeExportFormat = function(e) {
            this.exportFormat = e.target.value;
        }.bind(this);

        this.changeEftLang = function(e) {
            this.eftLang = e.target.value;
        }.bind(this);

        this.export = function(e) {
            var url = this.exportFormat;
            if (this.exportFormat == 'eft') {
                url += '/' + this.eftLang;
            }

            this.openExport(this.i18n(this.exportFormat), url);
        }.bind(this);

        this.mixin('i18n');
        this.mixin('appExport');
    </script>
</fit-export>