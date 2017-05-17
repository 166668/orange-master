'use strict';

myApp.factory('notifyService',
    function () {
        return {
            showInfo: function(msg) {
                $.notify(msg,"success");
            },
            showError: function(msg) {
                $.notify(msg,"error");
            }
        }
    }
);
