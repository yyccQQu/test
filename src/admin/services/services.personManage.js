angular.module("app.Person", [])
  .service("personManageService", personManageService);

personManageService.$inject = ['APP_CONFIG', 'httpSvc'];

function personManageService(APP_CONFIG, httpSvc, PlatformService) {

  var getRole = PlatformService.getRole,
      getRoleStatus = PlatformService.getRoleStatus,
      getRoleDel = PlatformService.getRoleDel,
      getRoleAdd = PlatformService.getRoleAdd,
      getRolePermissionGet = PlatformService.getRolePermissionGet,
      getRolePermissionPost = PlatformService.getRolePermissionPost,
      getRoleMenuGet = PlatformService.getRoleMenuGet,
      getRoleMenuPost = PlatformService.getRoleMenuPost

  return {
      getRole: getRole, // 角色管理	GET/role
      getRoleStatus: getRoleStatus, // 角色管理--修改状态	PUT /role/status
      getRoleDel: getRoleDel, // 角色管理--删除	DELETE/role
      getRoleAdd: getRoleAdd, // 角色管理--添加	POST/role
      getRolePermissionGet: getRolePermissionGet, // 角色管理--权限配置--修改	GET/role/permission
      getRolePermissionPost: getRolePermissionPost, // 角色管理--权限配置--修改	POST/role/permission
      getRoleMenuGet: getRoleMenuGet, // 角色管理--菜单	GET/role/menu
      getRoleMenuPost: getRoleMenuPost // 角色管理--菜单--修改	POST/role/menu
  }

  //
  // // 角色管理--菜单	GET/role/menu
  // function getRoleMenuGet(postData) {
  //     return httpSvc.get(APP_CONFIG.apiUrls.ROLE_MENU_GET, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理--菜单--修改	POST/role/menu
  // function getRoleMenuPost(postData) {
  //     return httpSvc.post(APP_CONFIG.apiUrls.ROLE_MENU_POST, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理--权限配置--修改	GET/role/permission
  // function getRolePermissionGet(postData) {
  //     return httpSvc.get(APP_CONFIG.apiUrls.ROLE_PERMISSION_GET, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理--权限配置--修改	POST/role/permission
  // function getRolePermissionPost(postData) {
  //     return httpSvc.post(APP_CONFIG.apiUrls.ROLE_PERMISSION_POST, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理	GET/role
  // function getRole(postData) {
  //     return httpSvc.get(APP_CONFIG.apiUrls.GET_ROLE, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理--添加	POST/role
  // function getRoleAdd(postData) {
  //     return httpSvc.post(APP_CONFIG.apiUrls.POST_ROLE, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理--修改状态	PUT /role/status
  // function getRoleStatus(postData) {
  //     return httpSvc.put(APP_CONFIG.apiUrls.ROLE_STATUS, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }
  // // 角色管理--删除	DELETE/role
  // function getRoleDel(postData) {
  //     return httpSvc.del(APP_CONFIG.apiUrls.DEL_ROLE, postData).then(getDataComplete)
  //         .catch(getDataFailed);
  //
  //     function getDataComplete(response) {
  //         return response;
  //     }
  //
  //     function getDataFailed(error) {
  //         console.log('XHR Failed for getAvengers.' + error);
  //
  //     }
  // }

}
