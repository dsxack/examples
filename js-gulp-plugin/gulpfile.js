var gulp = require("gulp");
var through = require("through2");
var del = require("del");
var concat = require("gulp-concat");

/**
 * Задача на очистку билда
 */
gulp.task('clear', function (callback) {
    del(['build/**/*']).then(function () {
        callback();
    });
});

/**
 * Задача по сборке билда
 */
gulp.task('build', ['clear'], function () {
    gulp.src('src/**/*.js')
    /** наш код здесь **/
    .pipe(gulp.dest('build'));
});

/**
 * Дефолтная задача – сборка билда
 */
gulp.task('default', ['build']);
