$(document).ready(function () {

    let action;
    
    $("#create").click(function (e) { 

        create();
        action = "addaccount";
        
    });

    $("#add").click(function (e) { 

        addOrWithdrawals();
        action = "addmoney";
        
    });

    $("#withdrawals").click(function (e) { 

        addOrWithdrawals();
        action = "withdrawals";

    });

    $("#delete").click(function (e) {

        deleteAccount();
        action = "deleteaccount";

    });

    $("#edit").click(function (e) {

        edit();
        action = "modifyaccount"

    });

    $("#btn").click(function (e) {

        formSubmit();

    });

    function create() {

        $("#document").removeClass("deactivate");
        $("#mount").removeClass("deactivate");
        $("#name").removeClass("deactivate");

        $("#new-document").addClass("deactivate");
        
    }

    function addOrWithdrawals() {

        $("#document").removeClass("deactivate");
        $("#mount").removeClass("deactivate");

        $("#name").addClass("deactivate");
        $("#new-document").addClass("deactivate");
        
    }

    function deleteAccount() {

        $("#document").removeClass("deactivate");

        $("#mount").addClass("deactivate");
        $("#name").addClass("deactivate");
        $("#new-document").addClass("deactivate");

    }

    function edit() {

        $("#document").removeClass("deactivate");
        $("#new-document").removeClass("deactivate");

        $("#mount").addClass("deactivate");
        $("#name").addClass("deactivate");

    }

    function formSubmit() {

        let formData = {};
        let method;
        formData["document"] = $("#document").val();

        if (action == "addaccount") {
            formData["mount"] = $("#mount").val();
            formData["name"] = $("#name").val();
            method = "POST";
        }

        if (action == "addmoney") {
            formData["mountToAdd"] = $("#mount").val();
            method = "PUT";
        }

        if (action == "withdrawals") {
            formData["mountToRemove"] = $("#mount").val();
            method = "PUT";
        }

        if (action == "modifyaccount") {
            formData["newDocument"] = $("#new-document").val();
            method = "PUT";
        }

        if (action == "deleteaccount") {
            method = "DELETE";
        }

        let requestSettings = {
            "url": "/controllers/" + action ,
            "method": method,
            "timeout": 0,
            "headers": {
                "Content-Type": "application/json"
            },
            "data": JSON.stringify(formData),
        };

        $.ajax(requestSettings).done(function (response) {


        });

    }

});