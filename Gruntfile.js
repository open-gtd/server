module.exports = function(grunt) {

    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),

        copy: {
            main: {
                files: [
                    // includes files within path
                    {expand: true, src: ['path/**/dist'], dest: 'web/assets/vendor', filter: 'isFile'}
                ]
            }
        },

        bowerInstall: {

            target: {

                // Point to the files that should be updated when
                // you run `grunt bower-install`
                src: [
                    'app/Resources/**/*.twig',
                    'app/styles/main.scss',  // .scss & .sass support...
                    'app/config.yml'         // and .yml & .yaml support out of the box!
                ]

                // Optional:
                // ---------
                // cwd: '',
                // dependencies: true,
                // devDependencies: false,
                // exclude: [],
                // fileTypes: {},
                // ignorePath: '',
                // overrides: {}
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks('grunt-bower-install');

    // Default task(s).
    grunt.registerTask('default', ['copy', 'bowerInstall']);

};