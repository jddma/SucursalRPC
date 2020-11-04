$(document).ready(function() {

    $("#btn").click(function () {
        formSubmit();
    });

    $(".login-form").keypress(function(e) {
        if(e.which == 13){
            formSubmit();
        }
    });

    function formSubmit() {

        if ($("#email")[0].checkValidity())
        {
            if ($("#password")[0].checkValidity())
            {

                let requestData = {};
                requestData["email"] = $("#email").val();
                requestData["password"] = $("#password").val();

                let requestSettings = {
                    "url": "/controllers/login",
                    "method": "POST",
                    "timeout": 0,
                    "headers": {
                        "Content-Type": "application/json"
                    },
                    "data": JSON.stringify(requestData),
                };

                $.ajax(requestSettings).done(function (response) {

                    switch (response) {
                        case "error":
                            alert("Credencales incorrectas")
                            break;

                        case "success":
                            location.href = "/panel";
                            break
                    }

                });
            }
            else
            {
                alert("Rellene todos los datos")
            }
        }
        else
        {
            alert("Ingrese un correo válido.")
        }

    }

});