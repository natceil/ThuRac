// This is client side for home

// Get data for chart
getTotalYears();
getTotalProduct();

function getTotalYears() {
    fetch(window.location.origin+'/cve/static', {
        method: 'GET',
        headers: {
            'Content-Type': 'Application/json',
        },
    })
        .then(response => response.json())
        .then(data => {
            let dataStatic = [];
                for (let i = 0; i < 5; i++) {
                    let objectWanna = {
                        label: Object.keys(data.data)[i],
                        y: Object.values(data.data)[i],
                        indexLabel: ""+Object.values(data.data)[i]
                    }
                    dataStatic.push(objectWanna);
                }
            let chart1 = new CanvasJS.Chart("result-years", {
                animationEnabled: true,
                theme: "light2",
                title: {
                    text: "Tổng số CVE qua các năm "
                },
                axisY: {
                    includeZero: true
                },
                data: [{
                    type: "column",
                    yValueFormatString: "#,##0#\"\"",
                    dataPoints: dataStatic
                }]
            });
            chart1.render();
            
        })
        .catch((error) => {
            console.log('Error:', error);
        });

}


function getTotalProduct() {
    fetch(window.location.origin+'/cve/products', {
        method: 'GET',
        headers: {
            'Content-Type': 'Application/json',
        },
    })
        .then(response => response.json())
        .then(data => {

            let chart = new CanvasJS.Chart("result-product", {
                animationEnabled: true,
                theme: "light1",
                title: {
                    text: "Số lượng CVE ảnh hưởng đến các  sản phẩm phổ biến trong "+data.data.currYear
                },
                axisY: {
                    includeZero: true
                },
                data: [{
                    type: "column",
                    yValueFormatString: "#,##0#\"\"",
                    dataPoints: [
                        {
                            label: "Windows",
                            y: data.data.windows,
                            indexLabel: ""+data.data.windows
                        },
                        {
                            label: "Android",
                            y: data.data.android,
                            indexLabel: ""+data.data.android
                        },
                        {
                            label: "Google",
                            y: data.data.google,
                            indexLabel: ""+data.data.google
                        },
                        {
                            label: "Linux",
                            y: data.data.linux,
                            indexLabel: ""+data.data.linux
                        },
                        {
                            label: "Cisco",
                            y: data.data.cisco,
                            indexLabel: ""+data.data.cisco
                        }
                    ]
                }]
            });
            chart.render();
            
        })
        .catch((error) => {
            console.log('Error:', error);
        });

}