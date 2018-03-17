angular.module('app.Person',['ui.router']).config(function($stateProvider){
  $stateProvider
    .state('app.Person', {
      abstract: true,
      data: {
        title: 'Person'
      },
      resolve: {
          srcipts: function(lazyScript) {
              return lazyScript.register([
                  'vendor.ui.js'
              ]);
          }
      }
    })

    .state('app.Person.personal', {
      url: '/Person/personal',
      data: {
        //title: '角色管理'
        title: 'P_1'
      },
      views: {
        "content@app": {
          controller: 'PersonalCtrl',
          templateUrl: 'views/Person/views/personal.html'
        }
      }
    })

})
