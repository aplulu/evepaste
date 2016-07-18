<dscan-result>
    <div class="row row-info">
        <div class="col-md-7">
            <dl>
                <div if="{ scanResult['solarsystem_name'] }">
                    <dt>{ i18n('solarsystem') }:</dt>
                    <dd>{ scanResult['solarsystem_name'] }</dd>
                    <dt>{ i18n('constellation') }:</dt>
                    <dd>{ scanResult['constellation_name'] }</dd>
                    <dt>{ i18n('region') }:</dt>
                    <dd>{ scanResult['region_name'] }</dd>
                    <dt>{ i18n('security') }:</dt>
                    <dd>{ security() }</dd>
                </div>
                <dt>{ i18n('ships') }:</dt>
                <dd>
                    <dl>
                        <dt>{ i18n('ongrid') }:</dt>
                        <dd>{ shipsOnGrid }</dd>
                        <dt>{ i18n('offgrid') }:</dt>
                        <dd>{ shipsOffGrid }</dd>
                        <dt>{ i18n('total') }:</dt>
                        <dd>{ shipsTotal }</dd>
                    </dl>
                </dd>
                <div if={ hasControlTower }>
                    <dt>{ i18n('controltower') }:</dt>
                    <dd>
                        <dl>
                            <dt>{ i18n('online') }:</dt>
                            <dd>{ forceFieldCount }</dd>
                            <dt>{ i18n('offline') }:</dt>
                            <dd>{ controlTowerCount - forceFieldCount }</dd>
                            <dt>{ i18n('total') }:</dt>
                            <dd>{ controlTowerCount }</dd>
                        </dl>
                    </dd>
                </div>
            </dl>
        </div>
        <div class="col-md-5">
            <h5>{ i18n('filters') }</h5>
            <radio-group name="filter_category" value="{ filterCategory }" onchange={ changeFilterCategory }>
                <radio value="all" label={ parent.i18n('all') }/>
                <radio value="ships" label={ parent.i18n('ships') }/>
                <radio value="structures" label={ parent.i18n('structures') }/>
            </radio-group>

            <radio-group name="filter_grid" value="{ filterGrid }" onchange={ changeFilterGrid }>
                <radio value="all" label={ parent.i18n('all') }/>
                <radio value="ongrid" label={ parent.i18n('ongrid') }/>
                <radio value="offgrid" label={ parent.i18n('offgrid') }/>
            </radio-group>
        </div>
    </div>
    <div class="row">
        <div class="col-md-6">
            <h5>{ i18n('everything') }</h5>

            <ul class="list-scan list-scan-everything">
                <dscan-result-item each={ items['everything'] } name={ name } count={ count } color={ color } />
            </ul>
        </div>
        <div class="col-md-6">
            <h5>{ i18n('capitals') }</h5>

            <ul class="list-scan list-scan-capital">
                <dscan-result-item each={ items['capitals'] } name={ name } count={ count } color={ color } />
            </ul>

            <h5>{ i18n('ships') }</h5>

            <ul class="list-scan list-scan-ship">
                <dscan-result-item each={ items['ships'] } name={ name } count={ count } color={ color } />
            </ul>

            <h5>{ i18n('structures') }</h5>

            <ul class="list-scan list-scan-structure">
                <dscan-result-item each={ items['structures'] } name={ name } count={ count } color={ color } />
            </ul>

        </div>
    </div>

    <style scoped>
        .list-scan {
            padding-left: 0;
            margin-bottom: 20px;
        }

        dscan-result-item > li {
            line-height: 20px;
        }

        dscan-result-item:nth-of-type(odd) > li {
            background-color: #353a41;
        }

        h5 {
            font-weight: bold;
        }

        radio-group {
            display: block;
            margin-bottom: 15px;
        }

        dt {
            float: left;
            width: 120px;
            clear: left;
            text-align: right;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }

        dd {
            margin-left: 130px;
        }

        dd > dl {
            clear: both;
        }

        dd dt {
            width: 80px;
        }

        dd dd {
            margin-left: 90px;
        }

        .row-info {
            margin-bottom: 25px;
        }
    </style>

    <script>
        var self = this;
        var lists = ['everything', 'capitals', 'ships', 'structures'];
        var structureCategoryIds = [
            22, // Deployable
            23, // Starbase
            40, // Sovereignty Structures
            46, // Orbitals
            65 // Structure
        ];
        var capitalGroupIds = [
            30, // Titan
            485, // Dreadnought
            547, // Carrier
            659, // Supercarrier
            883, // Capital Industrial Ship
            902, // Jump Freighter
            941, // Industrial Command Ship
            1538 // Force Auxiliary
        ];

        this.items = {};
        this.shipsOnGrid = this.shipsOffGrid = this.shipsTotal = 0;
        this.filterGrid = 'all';
        this.filterCategory = 'all';
        this.hasControlTower = false;
        this.controlTowerCount = 0;
        this.forceFieldCount = 0;

        this.mixin('i18n');

        this.on('mount', function() {
            this.scanResult = JSON.parse(document.getElementById('scanResult').innerHTML);
            countShips();
            checkControlTower();
            renderAll();
        });

        this.changeFilterCategory = function(e) {
            this.filterCategory = e.target.value;
            renderAll();
        }.bind(this);

        this.changeFilterGrid = function(e) {
            this.filterGrid = e.target.value;
            renderAll();
        }.bind(this);

        this.security = function() {
            if (this.scanResult['solarsystem_security']) {
                return Math.round(this.scanResult['solarsystem_security'] * 10) / 10;
            }
        }.bind(this);

        function checkControlTower() {
            for (var i = 0; i < self.scanResult['types'].length; i++) {
                if (self.scanResult['types'][i]['group_id'] == 365) {
                    self.controlTowerCount += self.scanResult['types'][i]['total'];
                    self.hasControlTower = true;
                } else if (self.scanResult['types'][i]['type_id'] == 16103) {
                    self.forceFieldCount += self.scanResult['types'][i]['total'];
                }
            }
        }

        function countShips() {
            for (var i = 0; i < self.scanResult['groups'].length; i++) {
                if (self.scanResult['groups'][i]['category_id'] == 6) {
                    var g = self.scanResult['groups'][i];
                    self.shipsOnGrid += g['ongrid_count'];
                    self.shipsOffGrid += g['offgrid_count'];
                    self.shipsTotal += g['total'];
                }
            }
        }

        function renderList(type) {
            var data = [];

            if (type == 'everything') {
                data = self.scanResult['types'];
            } else if (type == 'capitals') {
                data = self.scanResult['groups'].filter(function(i) {
                    return i['category_id'] == 6 && capitalGroupIds.indexOf(i['group_id']) != -1;
                });
            } else if (type == 'ships') {
                data = self.scanResult['groups'].filter(function(i) {
                    return i['category_id'] == 6 && capitalGroupIds.indexOf(i['group_id']) == -1;
                });
            } else if (type == 'structures') {
                data = self.scanResult['groups'].filter(function(i) {
                    return structureCategoryIds.indexOf(i['category_id']) != -1;
                });
            }

            if (self.filterCategory == 'ships') {
                data = data.filter(function(i) {
                    return i['category_id'] == 6;
                });
            } else if (self.filterCategory == 'structures') {
                data = data.filter(function(i) {
                    return structureCategoryIds.indexOf(i['category_id']) != -1;
                });
            }

            var countField = 'total';
            if (self.filterGrid == 'ongrid') {
                data = data.filter(function(i) {
                    return i['ongrid_count'] > 0;
                });
                countField = 'ongrid_count';
            } else if (self.filterGrid == 'offgrid') {
                data = data.filter(function(i) {
                    return i['offgrid_count'] > 0;
                });
                countField = 'offgrid_count';
            }

            var max = 0;
            for (var i = 0; i < data.length; i++) {
                max = Math.max(max, data[i][countField])
            }

            self.items[type] = [];
            for (i = 0; i < data.length; i++) {
                var item = data[i];
                var name = item['type_name'] || item['group_name'];

                self.items[type].push({
                    name: name,
                    count: item[countField],
                    color: getColor([255,0,0], [0,128,0], item[countField] / max)
                });
            }

            // sort by count
            self.items[type].sort(function(a, b) {
                return b['count'] - a['count'];
            });
        }

        function renderAll() {
            for (var i = 0; i < lists.length; i++) {
                renderList(lists[i]);
            }
            self.update();
        }

        function getColor(color1, color2, weight) {
            var w = weight * 2 - 1;
            var w1 = (w + 1) / 2;
            var w2 = 1 - w1;
            var rgb = [Math.round(color1[0] * w1 + color2[0] * w2),
                Math.round(color1[1] * w1 + color2[1] * w2),
                Math.round(color1[2] * w1 + color2[2] * w2)];
            return 'rgb(' + rgb.join() + ')';
        }
    </script>
</dscan-result>