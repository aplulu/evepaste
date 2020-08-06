<radio-group>
    <div class="btn-group">
        <yield/>
    </div>

    <style scoped>
        .btn-group {
            position: relative;
            display: inline-block;
            vertical-align: middle;
        }
        
        radio:not(:first-child):not(:last-child):not(.dropdown-toggle) .btn {
            border-radius: 0;
        }

        radio:first-child:not(:last-child):not(.dropdown-toggle) .btn {
            border-bottom-right-radius: 0;
            border-top-right-radius: 0;
        }

        radio:last-child:not(:first-child) .btn {
            border-bottom-left-radius: 0;
            border-top-left-radius: 0;
        }

        radio .btn {
            margin-left: -1px;
        }

        radio:first-child .btn {
            margin-left: 0;
        }


    </style>

    <script>
        var self = this;
        this.value = this.opts.value;
        this.name = this.opts.name;

        this.on('mount', function() {
            var radios = this.tags['radio'];
            if (radios) {
                radios.forEach(function(r) {
                    r.on('valueUpdated', self.valueUpdated);
                });
            }
        });

        this.valueUpdated = function(v) {
            this.value = v;
            this.update();
            //this.triggerDomEvent('change')
        }.bind(this);
    </script>
</radio-group>