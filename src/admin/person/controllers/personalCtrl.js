
angular.module('app.Person').controller('PersonalCtrl', function(DTOptionsBuilder, DTColumnBuilder, $http){
    $http.get('src/api/personDatabase/personBase.json').success(function(data){console.log(data)})

    this.standardOptions = DTOptionsBuilder
        .fromSource('api/personDatabase/personBase.json')
         //Add Bootstrap compatibility
        .withDOM("<'dt-toolbar'<'col-xs-12 col-sm-6'f><'col-sm-6 col-xs-12 hidden-xs'l>r>" +
            "t" +
            "<'dt-toolbar-footer'<'col-sm-6 col-xs-12 hidden-xs'i><'col-xs-12 col-sm-6'p>>")
        .withBootstrap();
    this.standardColumns = [
        DTColumnBuilder.newColumn('id').withClass('text-danger'),
        DTColumnBuilder.newColumn('name'),
        DTColumnBuilder.newColumn('state'),
        DTColumnBuilder.newColumn('descr'),
        DTColumnBuilder.newColumn('doto')
        // DTColumnBuilder.newColumn('city'),
        // DTColumnBuilder.newColumn('date')
    ];


});
