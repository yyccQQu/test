"use strict";


angular.module('app.layout', ['ui.router'])
    .config(function ($stateProvider, $urlRouterProvider) {
        $stateProvider
            .state('app', {
                abstract: true,
                views: {
                    root: {
                        controller: 'layoutCtrl',
                        templateUrl: 'views/layout/layout.tpl.html'
                    }
                }
            })

        $urlRouterProvider.otherwise('/real-login');

    });

