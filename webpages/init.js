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

document.addEventListener("DOMContentLoaded", () => {
    window.accessToken = "";
    window.Profile = null;
    const liffId = "2004358435-rDXj330L";
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
        liff.init({
            liffId: liffId,
        }).then(() => {
            if (liff.isLoggedIn()) {
                runApp()
            } else {
                liff.login();
            }
        }).catch((err) => {
            liff.login();
        });

    } catch (err) {

    } finally {
        closeLoading();
    }



});