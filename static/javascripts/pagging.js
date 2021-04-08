const item = document.querySelectorAll(".page-items");
const max = item.length;
const searchParams = new URLSearchParams(window.location.search);
if (max > 20) {
    hiden();
    const first = document.getElementById('blank-pagging-first');
    const last = document.getElementById('blank-pagging-last');
    function hiden() {
        for (let j = 9; j < item.length - 1; j++) {
            item[j].style.display = "none";
        }
    }

    function hidenLast() {
        for (let j = 1; j < item.length - 10; j++) {
            item[j].style.display = "none";
        }
    }

    function hidenAll() {
        for (let j = 1; j < item.length - 1; j++) {
            item[j].style.display = "none";
        }
    }

    var number;
    const pageNumbers = searchParams.get('pageNumber');
    if (!pageNumbers) {
        number = 1;
    }
    else if (!pageNumbers.match(/^\d+$/)) {
        number = 1;
    }
    else {
        number = parseInt(pageNumbers)
    }

    if (number >= 9 && number <= max - 11) {
        hidenAll();
        document.getElementById((number - 1) + "-page").style.display = "block";
        document.getElementById((number - 2) + "-page").style.display = "block";
        document.getElementById((number - 3) + "-page").style.display = "block";
        document.getElementById((number - 4) + "-page").style.display = "block";
        document.getElementById(number + "-page").style.display = "block";
        document.getElementById((number + 1) + "-page").style.display = "block";
        document.getElementById((number + 2) + "-page").style.display = "block";
        document.getElementById((number + 3) + "-page").style.display = "block";
        document.getElementById((number + 4) + "-page").style.display = "block";
        document.getElementById((number + 5) + "-page").style.display = "block";
        document.getElementById((number + 6) + "-page").style.display = "block";
        document.getElementById((number + 7) + "-page").style.display = "block";
        document.getElementById((number + 8) + "-page").style.display = "block";
        document.getElementById((number + 9) + "-page").style.display = "block";
    }
    if (number < 9) {
        first.style.display = "none";
    }

    if (number > max - 11) {
        hidenLast();
        document.getElementById(number + "-page").style.display = "block";
        document.getElementById((number - 1) + "-page").style.display = "block";
        document.getElementById((number - 2) + "-page").style.display = "block";
        document.getElementById((number - 3) + "-page").style.display = "block";
        document.getElementById((number - 4) + "-page").style.display = "block";
        let maxLoop = max - number;
        for (let i = 1; i <= maxLoop; i++) {
            document.getElementById((number + i) + "-page").style.display = "block";
        }

        last.style.display = "none";
    }

}


    // Get value of pageNumber in URL
    const pageNumber = searchParams.get('pageNumber');
    // If pageNumber does not exist, server will set pageNumber equal 1. Therefore, active page 1
    if (!pageNumber) {
        document.getElementById("1-page").classList.add("active");
    }
    // If pageNumber does not a number, server will set pageNumber equal 1. Therefore, active page 1
    else if (!pageNumber.match(/^\d+$/)) {
        document.getElementById("1-page").classList.add("active");
    }
    // In other cases, active page equals pageNumber params in URL 
    else {
        document.getElementById(pageNumber + "-page").classList.add("active");
    }