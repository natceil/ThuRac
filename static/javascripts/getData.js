getDataForHome(1);
async function getDataForHome(pageNumber) {
  let keyword = document.getElementById('cve-keyword').value;
  let show = document.getElementById('show-form').value;
  let loader = document.getElementById('loader');
  loader.style.visibility = "visible";
  fetch(window.location.origin + '/cve/find?pageSize=' + show + '&keyword=' + keyword + '&pageNumber=' + pageNumber, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  })
    .then((res) => {
      return res.json();
    })
    .then((data) => {
      if (data.success) {

        let tableBody = document.getElementById('result-display');
        tableBody.innerHTML = '';
        if (data.total == 0) {
          let newRow = `
                <tr>
                Không có kết quả phù hợp với tìm kiếm
                </tr>`;
          tableBody.insertAdjacentHTML('beforeend', newRow);

          loader.style.visibility = "hidden";
        } else {
          data.data.forEach(element => {
            let newRow = `
                  <tr>
                  <td><a href="/cve/detail/${element.nameCve}">${element.nameCve}</a></td>
                  <td>${element.description}</td>
                  </tr>
                  `;
            tableBody.insertAdjacentHTML('beforeend', newRow);
            //load pagging
            
          });
          
          let paggingDiv = document.getElementById('pagination');
          //reset to none
          paggingDiv.innerHTML='';
          if (data.totalPage <= 20) {
            for (let i = 1; i <= data.totalPage; i++) {
              let newPagging = `<li class="page-item page-items" id="${i}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${i})">${i}</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', newPagging);
            }
          } else {
            if (pageNumber<5) {
              for (let i = 1; i <= 5; i++) {
                let newPagging = `<li class="page-item page-items" id="${i}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${i})">${i}</a></li>`;
                paggingDiv.insertAdjacentHTML('beforeend', newPagging);
              }
              let blankPage = `<li class="page-item page-items"><a class="page-link">...</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', blankPage);
              let lastPage = `<li class="page-item page-items" id="${data.totalPage}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${data.totalPage})">${data.totalPage}</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', lastPage);
            } else if (pageNumber>=5&&pageNumber<=data.totalPage-3) {
              let firstPage = `<li class="page-item page-items" id="1-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(1)">1</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', firstPage);
              let blankPage = `<li class="page-item page-items"><a class="page-link">...</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', blankPage);
              let prePage =  `<li class="page-item page-items" id="${pageNumber-1}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${pageNumber-1})">${pageNumber-1}</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', prePage);
              let thisPage =  `<li class="page-item page-items" id="${pageNumber}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${pageNumber})">${pageNumber}</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', thisPage);
              let nextPage =  `<li class="page-item page-items" id="${pageNumber+1}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${pageNumber+1})">${pageNumber+1}</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', nextPage);
              paggingDiv.insertAdjacentHTML('beforeend', blankPage);
              let lastPage = `<li class="page-item page-items" id="${data.totalPage}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${data.totalPage})">${data.totalPage}</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', lastPage);
            } else if (pageNumber>data.totalPage-3) {
              let firstPage = `<li class="page-item page-items" id="1-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(1)">1</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', firstPage);
              let blankPage = `<li class="page-item page-items"><a class="page-link">...</a></li>`;
              paggingDiv.insertAdjacentHTML('beforeend', blankPage);
              for (let i=pageNumber-3;i<pageNumber;i++){
                let prePage =  `<li class="page-item page-items" id="${i}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${i})">${i}</a></li>`;
                paggingDiv.insertAdjacentHTML('beforeend', prePage);
              }
              for (let i = pageNumber;i<=data.totalPage;i++){
                let prePage =  `<li class="page-item page-items" id="${i}-page"><a class="page-link" onclick="event.preventDefault();getDataForHome(${i})">${i}</a></li>`;
                paggingDiv.insertAdjacentHTML('beforeend', prePage);
              }
            }
          }
          // active for pagging
          document.getElementById(`${pageNumber}-page`).classList.add("active");
          loader.style.visibility = "hidden";
        }
      } else {
        alert('Có lỗi xảy ra trong quá trình tải dữ liệu. Thử lại sau');
        loader.style.visibility = "hidden";
      }
    })
    .catch((error) => {
      console.log(error);
    });
}