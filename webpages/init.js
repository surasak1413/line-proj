function showLoading() {
    loadingSwal = Swal.fire({
        title: 'Loading...',
        allowOutsideClick: false,
        allowEscapeKey: false,
        showConfirmButton: false,
        willOpen: () => {
            Swal.showLoading();
        }
    });
}

function closeLoading() {
    if (loadingSwal) {
        loadingSwal.close();
    }
}


function showMessage(type = "success", message) {

    Swal.fire({
        title: message,
        icon: type || "success"
    });
}

document.addEventListener("DOMContentLoaded", async () => {
    window.accessToken = "";
    window.Profile = null;

    function runApp() {
        window.accessToken = liff.getAccessToken();
        liff.getProfile().then(profile => {
            window.Profile = {
                user_id: profile.userId || '',
                picture_url: profile.pictureUrl || '',
                display_name: profile.displayName || '',
                status_message: profile.statusMessage || '',
            }
        });
    }

    try {
        showLoading();
        await $.ajax({
            url: "/v1/get-liff-id",
            type: "POST",
            contentType: "application/json",
            success: function (response) {
                window.liffId = response.id || "";

                liff.init({
                    liffId: window.liffId,
                }).then(() => {
                    if (liff.isLoggedIn()) {
                        runApp()
                    } else {
                        liff.login();
                    }
                }).catch((err) => {
                    liff.login();
                });

            },
            error: function (jqXHR, textStatus, errorThrown) {
            }
        });


    } catch (err) {
        console.warn(err)
    } finally {
        closeLoading()
    }



});