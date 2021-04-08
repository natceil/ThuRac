var flag = false;
function add() {
    if (flag) {
        return
    }
    const table = document.getElementById("values");
    let row2 = table.insertRow(0);
    row2.innerHTML = `<button type="button" class="btn btn-outline-dark mt-3 mb-3" id="add-new-date" onclick="addToDB()">Thêm mới</button>`
    let row = table.insertRow(0);
    let today = new Date();
    let dd = String(today.getDate()).padStart(2, '0');
    let mm = String(today.getMonth() + 1).padStart(2, '0');
    let yyyy = today.getFullYear();

    today = dd + '/' + mm + '/' + yyyy;
    let tableElement = `<td contenteditable><b>${today}</b></td>`;
    for (i = 1; i <= 16; i++) {
        tableElement = tableElement + `<td contenteditable id="insert-${i}"></td>`;
    }
    row.innerHTML = tableElement;
    flag = true;

}   
function addToDB() {
    alert('Dữ liệu gửi lên sẽ không thể sửa, bạn có chắc chắn ?');
    alert(document.getElementById("insert-1").textContent);
}