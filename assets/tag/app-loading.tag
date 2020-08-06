var AppLoadingMixin = {
    loadingObservable: riot.observable(),
    startLoading: function() {
        this.loadingObservable.trigger('start');
    },
    stopLoading: function() {
        this.loadingObservable.trigger('stop');
    }
};
riot.mixin('appLoading', AppLoadingMixin);
<app-loading>
    <div class="bar-container" id="barContainer">
        <div class="bar" id="bar">
            <div class="progress"></div>
        </div>
        <div class="spinner"></div>
    </div>

    <style scoped>
        .bar-container {
            pointer-events: none;
            display: none;
        }

        .barContainer.active {
            display: block;
        }

        .bar {
            background: #239ddd;
            position: fixed;
            z-index: 1031;
            top: 0;
            left: 0;
            width: 100%;
            height: 2px;
        }


        .progress {
            display: block;
            position: absolute;
            right: 0px;
            width: 100px;
            height: 100%;
            box-shadow: 0 0 10px #239ddd, 0 0 5px #239ddd;
            opacity: 1.0;
            -webkit-transform: rotate(3deg) translate(0px, -4px);
            -ms-transform: rotate(3deg) translate(0px, -4px);
            transform: rotate(3deg) translate(0px, -4px);
        }

        .spinner {
            display: block;
            position: fixed;
            z-index: 10010;
            top: 15px;
            right: 15px;
            width: 14px;
            height: 14px;
            border: solid 2px transparent;
            border-top-color: #239ddd;
            border-left-color: #239ddd;
            border-radius: 50% !important;
            -webkit-animation: page-progress-spinner 400ms linear infinite;
            -moz-animation: page-progress-spinner 400ms linear infinite;
            -ms-animation: page-progress-spinner 400ms linear infinite;
            -o-animation: page-progress-spinner 400ms linear infinite;
            animation: page-progress-spinner 400ms linear infinite;
        }

        @-webkit-keyframes page-progress-spinner {
            0% { -webkit-transform: rotate(0deg); transform: rotate(0deg); }
            100% { -webkit-transform: rotate(360deg); transform: rotate(360deg); }
        }
        @-moz-keyframes page-progress-spinner {
            0% { -moz-transform: rotate(0deg); transform: rotate(0deg); }
            100% { -moz-transform: rotate(360deg); transform: rotate(360deg); }
        }
        @-o-keyframes page-progress-spinner {
            0% { -o-transform: rotate(0deg); transform: rotate(0deg); }
            100% { -o-transform: rotate(360deg); transform: rotate(360deg); }
        }
        @-ms-keyframes page-progress-spinner {
            0% { -ms-transform: rotate(0deg); transform: rotate(0deg); }
            100% { -ms-transform: rotate(360deg); transform: rotate(360deg); }
        }
        @keyframes page-progress-spinner {
            0% { transform: rotate(0deg); transform: rotate(0deg); }
            100% { transform: rotate(360deg); transform: rotate(360deg); }
        }
    </style>

    <script>
        var self = this;
        this.active = false;
        this.options = {
            minimum: 0.07,
            speed: 200,
            trickle: true,
            trickleRate: 0.02,
            trickleSpeed: 800
        };

        this.mixin(AppLoadingMixin);

        this.on('mount', function() {
            self.bar = jQuery('#bar');
            self.barContainer = jQuery('#barContainer');
        });

        this.loadingObservable.on('start', function() {
            self.start();
        });

        this.loadingObservable.on('stop', function() {
            self.end();
        });

        this.start = function() {
            if (!this.progress) {
                this.set(0);
            }

            self.bar.css('margin-left', '0%');
            self.bar.css('transition', 'none');

            self.barContainer.css('transition', 'none');
            self.barContainer.css('opacity', 1);
            this.barContainer.show();

            var timer = function() {
                setTimeout(function() {
                    if (!self.progress) {
                        return;
                    }

                    self.trickle();
                    timer();
                }, self.options.trickleSpeed);
            };

            if (this.options.trickle) {
                timer();
            }
        }.bind(this);

        this.end = function() {
            if (!this.progress) {
                return;
            }
            this.inc(0.3 + 0.5 * Math.random());
            return this.set(1);
        }.bind(this);

        this.set = function(n) {
            n = this.clamp(n, this.options.minimum, 1);
            this.progress = (n === 1 ? null : n);

            var speed = this.options.speed;

            this.queue(function(next) {
                self.bar.css('margin-left', ((-1 + n) * 100) + '%');
                self.bar.css('transition', 'all ' + speed + 'ms ease');

                if (n === 1) {
                    self.barContainer.css('transition', 'none');
                    self.barContainer.css('opacity', 1);

                    setTimeout(function() {
                        self.barContainer.css('transition', 'all ' + speed + 'ms linear');
                        self.barContainer.css('opacity', 0);

                        setTimeout(function() {
                            self.remove();
                            next();
                        }, speed);
                    }, speed);
                } else {
                    setTimeout(next, speed);
                }
            });
        }.bind(this);

        this.inc = function(count) {
            if (this.progress) {
                var n = this.progress;
                if (count == undefined) {
                    count = (1 - n) * this._clamp(Math.random() * n, 0.1, 0.95);
                }
                n = this.clamp(n + count, 0, 0.994);
                return this.set(n);
            } else {
                return this.start();
            }
        }.bind(this);

        this.trickle = function() {
            return this.inc(Math.random() * this.options.trickleRate);
        }.bind(this);

        this.remove = function() {
            this.active = false;
            self.barContainer.hide();
        }.bind(this);

        this.clamp = function(n, min, max) {
            if (n < min) return min;
            if (n > max) return max;
            return n;
        }.bind(this);

        this.queue = (function() {
            var pending = [];

            function next() {
                var fn = pending.shift();
                if (fn) {
                    fn(next);
                }
            }

            return function(fn) {
                pending.push(fn);
                if (pending.length == 1) next();
            };
        })();
    </script>
</app-loading>