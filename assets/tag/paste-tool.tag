<paste-tool>
    <div class="btn-group pull-right">
        <button class="btn btn-default" role="button" onclick={ openRaw }>{ i18n('raw') }</button>
        <yield/>
    </div>

    <style scoped>
    </style>

    <script>
        this.mixin('i18n');
        this.mixin('appExport');

        this.openRaw = function(e) {
            this.openExport(this.i18n('raw'), 'raw');
        }.bind(this);
    </script>
</paste-tool>