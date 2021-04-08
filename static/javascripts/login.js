const username = document.getElementById('username');
const password = document.getElementById('password');
const eyeOn = document.getElementById('eye-on');
const eyeOff = document.getElementById('eye-off');
// Change color border for input login form 
function changeBorder(id) {
    const object = document.getElementById(id);
    object.classList.remove('login-input');
    object.classList.add('login-input-change');
};

// When click out of input element change to old style
document.onclick = function(a) {
    if (a.target.id!=='username'&&username.value==="") {
        username.classList.remove('login-input-change');
        username.classList.add('login-input');
    }
    if (a.target.id!=='password'&&password.value==="") {
        password.classList.remove('login-input-change');
        password.classList.add('login-input');
    }
} 

eyeOn.addEventListener('click', function () {
    eyeOn.classList.add('no-show');
    eyeOff.classList.remove('no-show');
    password.type = "text";
});


eyeOff.addEventListener('click', function () {
    eyeOff.classList.add('no-show');
    eyeOn.classList.remove('no-show');
    password.type = "password";
});


const element = document.getElementById('form-login');
element.addEventListener('submit', event => {
  event.preventDefault();
  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;
  const noti = document.getElementById('noti');
  noti.innerText = "";
  if (username.trim()==="") {
      noti.innerText = "Tên đăng nhập không được bỏ trống";
  }
  if (password.trim()==="") {
      noti.innerText = "Mật khẩu không được bỏ trống";
  }
  if (username.trim()!==""&&password.trim()!=="") {
    fetch('http://192.168.25.199:8000/Auth/Login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          account: username,
          password: password,
        }),
      })
        .then((res) => {
          return res.json();
        })
        .then((data) => {
            console.log(data)
          noti.innerText = "";
          if (data.success) {
            console.log('success');
          } else {
            noti.innerText = data.message;
          }
        })
        .catch((error) => {
          console.log(error);
        });
  }
});