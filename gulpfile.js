var gulp = require('gulp');
var rename = require('gulp-rename');
var concat = require('gulp-concat');
var jshint = require('gulp-jshint');
var uglify = require('gulp-uglify');
var clean = require('gulp-clean');
var copy = require('gulp-copy');
var recess = require('gulp-recess');
var html2js = require('gulp-html2js');
var shell = require('gulp-shell');
// var gulpif = require('gulp-if');
var cssmin = require('gulp-cssmin');
var usemin = require('gulp-usemin');
var replace = require('gulp-replace');

var del = require('del');
var gulpIgnore = require('gulp-ignore');

gulp.task("clean", function () {
    return del([
        'dist'
    ], {
        force: true
    });
});

gulp.task("html2js", function () {
    gulp.src('app/**/*.html')
        .pipe(html2js('templates.js', {
            adapter: 'angular',
            name: 'templates'
        }))
        .pipe(gulp.dest('dist/'));
});


gulp.task("assets", ['clean'], function () {
    var fonts = 'dist/fonts/';
    gulp.src([
        'bower_components/bootstrap/fonts/*.{ttf,woff,woff2,eof,svg}',
        'bower_components/font-awesome/fonts/*.{ttf,woff,woff2,eof,svg}',
        'bower_components/rdash-ui/dist/fonts/*.{ttf,woff,woff2,eof,svg}'
    ])
        .pipe(gulp.dest(fonts));

    var images = 'dist/images/';
    gulp.src([
        'bower_components/jquery.gritter/images/**',
        'assets/images/**'
    ])
        .pipe(gulpIgnore.exclude('trees.jpg'))
        .pipe(gulp.dest(images));

    var ico = 'dist/ico/';
    gulp.src([
        'assets/ico/**'
    ])
        .pipe(gulp.dest(ico));

});