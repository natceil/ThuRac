const passwordOld = document.getElementById('password-old');
const passwordNew = document.getElementById('password-new');
const passwordConfirm = document.getElementById('password-confirm');
const eyeOn = document.getElementById('eye-on');
const eyeOff = document.getElementById('eye-off');

function changeBorder(id) {
    const object = document.getElementById(id);
    object.classList.remove('login-input');
    object.classList.add('login-input-change');
};

eyeOn.addEventListener('click', function () {
    eyeOn.classList.add('no-show');
    eyeOff.classList.remove('no-show');
    passwordOld.type = "text";
});


eyeOff.addEventListener('click', function () {
    eyeOff.classList.add('no-show');
    eyeOn.classList.remove('no-show');
    passwordOld.type = "password";
});

const element = document.getElementById('form-change');
element.addEventListener('submit', event => {
    event.preventDefault();
    const noti = document.getElementById('noti');
    noti.innerText = "";
    if (passwordConfirm.value.trim()==="") {
        noti.innerText="Mật khẩu xác nhận không được để trống";
    }
    if (passwordNew.value.trim()==="") {
        noti.innerText="Mật khẩu mới không được để trống";
    }
    if (passwordOld.value.trim()==="") {
        noti.innerText="Bạn phải nhậm mật khẩu cũ ";
    }
    
    
    if (passwordOld.value.trim()!==""&&passwordNew.value.trim()!==""&&passwordConfirm.value.trim()!=="") {
        fetch(window.location.origin+'/auth/changepass', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                passwordOld: passwordOld.value,
                passwordNew: passwordNew.value,
                passwordConfirm: passwordConfirm.value
            }),
            credentials: 'include',
          })
            .then((res) => {
              return res.json();
            })
            .then((data) => {
              noti.innerText = "";
              if (data.success) {
                alert("Mật khẩu đã được đổi thành công");
                window.location = window.location.origin;

              } else {
                noti.innerText = data.message;
              }
            })
            .catch((error) => {
              console.log(error);
            });
    }
});

