var gulp = require("gulp");
var eslint = require("gulp-eslint");

var applyLintPaths = {
  allSrcJs: "src/**/*.js",
  gulpFile: "gulpfile.js"
};

/**
 * lint
 */
gulp.task("lint", function () {
  return (
    gulp.src([
      applyLintPaths.allSrcJs,
      applyLintPaths.gulpFile
    ])
      .pipe(eslint())
      .pipe(eslint.format())
      .pipe(eslint.failAfterError())
  );
});

gulp.task("lint-watch", function () {
  return (
    gulp.watch(applyLintPaths.allSrcJs, gulp.task("lint"))
  );
});