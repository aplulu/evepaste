'use strict';

var gulp = require('gulp'),
    $ = require('gulp-load-plugins')(),
    bower = require('main-bower-files'),
    streamqueue = require('streamqueue'),
    browserSync = require('browser-sync').create();

var conf = {
    dest: {
        css: 'main/static/css/',
        js: 'main/static/js/',
        fonts: 'main/static/fonts/'
    },
    src: {
        js: 'assets/js/',
        tag: 'assets/tag/',
        css: 'assets/css/',
        less: 'assets/less/'
    },
    bower: {
        paths: {
            bowerJson: 'bower.json'
        },
        overrides: {
            'bootswatch-dist': {
                main: [
                    'js/bootstrap.js',
                    'css/*.min.css',
                    'fonts/*.*'
                ]
            },
            jquery: {
                ignore: true
            }
        }
    }
};

gulp.task('fonts', function () {
    return gulp.src('./bower_components/**/*.{eot,svg,ttf,woff,woff2}')
        .pipe($.flatten())
        .pipe(gulp.dest(conf.dest.fonts));
});

gulp.task('css', function() {
    var bowerLessFilter = $.filter('**/*.less', {restore: true});
    var bowerStream = gulp.src(bower(conf.bower))
        .pipe(bowerLessFilter)
        .pipe($.less())
        .pipe(bowerLessFilter.restore);

    var assetLessFilter = $.filter('**/*.less', {restore: true});
    var assetStream = gulp.src([conf.src.css + '*.css', conf.src.less + '*.less'])
        .pipe(assetLessFilter)
        .pipe($.less())
        .pipe(assetLessFilter.restore);

    return streamqueue({ objectMode: true }, bowerStream, assetStream)
        .pipe($.plumber({
            errorHandler: $.notify.onError('<%= error.message %>')
        }))
        .pipe($.filter('**/*.css'))
        .pipe($.concat('application.css'))
        .pipe($.cleanCss())
        .pipe(gulp.dest(conf.dest.css))
        .pipe(browserSync.stream());
});

gulp.task('js', function() {
    var jsFilter = $.filter('**/*.js');

    var bowerStream = gulp.src(bower(conf.bower))
        .pipe(jsFilter);

    var tagStream = gulp.src(conf.src.tag + '*.tag')
        .pipe($.plumber({
            errorHandler: $.notify.onError('<%= error.message %>')
        }))
        .pipe($.riot({
            compact: true
        }));

    var jsStream = gulp.src(conf.src.js + '*.js');

    return streamqueue({ objectMode: true }, bowerStream, tagStream, jsStream)
        .pipe($.plumber({
            errorHandler: $.notify.onError('<%= error.message %>')
        }))
        .pipe($.concat('application.js'))
        .pipe($.uglify({
            preserveComments: 'license'
        }))
        .pipe(gulp.dest(conf.dest.js));
});


gulp.task('watch', function() {
    gulp.watch([conf.src.css + '*.css', conf.src.less + '*.less']).on('change', function(event) {
        console.log('File ' + event.path + ' was ' + event.type);
        gulp.start('css');
    });
    gulp.watch([conf.src.js + '*.js', conf.src.tag + '*.tag']).on('change', function(event) {
        console.log('File ' + event.path + ' was ' + event.type);
        gulp.start('js');
        browserSync.reload();
    });
});

gulp.task('browsersync', ['watch'], function() {
    return browserSync.init({
        proxy: 'localhost:8080'
    });
});

gulp.task('build', ['css', 'js', 'fonts']);