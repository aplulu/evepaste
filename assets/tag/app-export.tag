var AppExportMixin = {
    exportObservable: riot.observable(),
    openExport: function(title, type) {
        this.exportObservable.trigger('openExport', title, type);
    }
};
riot.mixin('appExport', AppExportMixin);
<app-export>
    <div id="myModal" class="modal {modal-show: dialogShowing}" tabindex="-1">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close" onclick={ closeDialog }><span aria-hidden="true">×</span></button>
                    <h4 class="modal-title" id="myModalLabel">{title}</h4>
                </div>
                <div class="modal-body">
                    <div if={ isMessage }>
                        <p>{message}</p>
                    </div>
                    <div if={ !isMessage }>
                        <textarea rows="12" id="modalTextarea"></textarea>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal" onclick={ closeDialog }>{ i18n('close') }</button>
                </div>
            </div>
        </div>
    </div>
    <!-- <div class="modal-backdrop fade" onclick={ closeDialog }></div> -->

    <style scoped>
        .modal {
            display: block;
            position: fixed;
            top: 50%;
            left: 50%;
            width: 50%;
            max-width: 630px;
            min-width: 320px;
            height: auto;
            visibility: hidden;
            opacity: 0;
            z-index: 1050;
            -webkit-transition: all 0.3s;
            -moz-transition: all 0.3s;
            transition: all 0.3s;
            -webkit-transform: translateX(-50%) translateY(-50%);
            transform: translateX(-50%) translateY(-50%);
            overflow: visible;
        }

        .modal.modal-show {
            visibility: visible;
            opacity: 1;
        }

        .modal-backdrop {
            display: none;
        }

        .modal-show ~ .modal-backdrop {
            display: block;
            opacity: 0.5;
        }


        .modal-dialog {
            -webkit-transform: scale(0.7);
            -moz-transform: scale(0.7);
            -ms-transform: scale(0.7);
            transform: scale(0.7);
            opacity: 0;
            -webkit-transition: all 0.3s;
            -moz-transition: all 0.3s;
            transition: all 0.3s;
        }

        .modal-show .modal-dialog {
            -webkit-transform: scale(1);
            -moz-transform: scale(1);
            -ms-transform: scale(1);
            transform: scale(1);
            opacity: 1;
        }

        .modal-content textarea {
            display: block;
            width: 100%;
            min-height: 50px;
        }
    </style>

    <script>
        var self = this;

        this.isMessage = false;

        this.mixin(AppExportMixin);
        this.mixin('appLoading');
        this.mixin('i18n');

        this.exportObservable.on('openExport', function(title, type) {
            url = '/p/' + self.opts['pasteId'] + '/' + type;

            var method = 'get';
            var dataType = 'text';

            if (type == 'crest') {
                method = 'post';
                dataType = 'json';
            }

            self.startLoading();

            jQuery.ajax({
                method: method,
                url: url,
                dataType: dataType,
                cache: true
            }).done(function(data) {
                self.openDialog(title, type, data);
            }).fail(function(xhr) {
                if (xhr['responseJSON'] && xhr['responseJSON']['error_description']) {
                    self.openErrorDialog(title, xhr['responseJSON']['error_description']);
                }
            }).always(function() {
                self.stopLoading();
            });
        });

        this.on('update', function() {
            // IE Bug workaround...
            var textArea = document.getElementById('modalTextarea');
            if (textArea) {
                textArea.value = self.text;
            }
        });

        this.openDialog = function(title, type, data) {
            this.dialogShowing = true;
            this.title = title;
            if (type == 'crest') {
                this.isMessage = true;
                this.message = data['message'];
            } else {
                this.isMessage = false;
                this.text = data;
            }
            this.update();
        }.bind(this);

        this.openErrorDialog = function(title, message) {
            this.dialogShowing = true;
            this.title = title;
            this.isMessage = true;
            this.message = message;
            this.update();
        }.bind(this);

        this.closeDialog = function() {
            this.dialogShowing = false;
        }.bind(this);
    </script>
</app-export>