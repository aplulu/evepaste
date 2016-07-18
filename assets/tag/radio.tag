<radio>
    <label class="btn btn-primary{ selected ? ' active' : '' }">
        <input type="radio" name="{ parent.name }" value={ opts.value } autocomplete="off" checked={ selected ? 'checked' : '' } onclick={ click }> { opts.label }
    </label>

    <style scoped>
        label {
            position: relative;
            float: left;
        }

        input[type=radio] {
            position: absolute;
            clip: rect(0,0,0,0);
            pointer-events: none;
        }
    </style>

    <script>
        this.on('update', function() {
            this.selected = this.parent.value == opts.value
        });

        this.click = function(e) {
            this.trigger('valueUpdated', opts.value);
        }.bind(this);
    </script>
</radio>