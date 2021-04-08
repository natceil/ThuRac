package auth

import (
	"golangSql/m/models"
	"regexp"
	"strconv"
	"time"
)

const(
	layoutDate = "02-01-2006" //YYYY-MM-DD
)

func CheckDate(date string) bool{
	reDate := regexp.MustCompile("([0][1-9]|[12][0-9]|3[01])-([0][1-9]|1[012])-((19|20)\\d\\d)")
	//check not future
	today := time.Now()

	checkTimeExist, _ := time.Parse(layoutDate,date)

	//check future success
	days := checkTimeExist.Sub(today).Hours() / 24

	countRow := models.CheckDateExist(checkTimeExist.Format("2006-01-02"))

	if  reDate.MatchString(date) == false{
		return false
	}
	//success
	if days < 0{
		return false
	}
	//success
	if countRow > 0{
		return false
	}
	return true
}

func CheckNum(num string) bool{
	reNum := regexp.MustCompile("^[0-9]{1,6}$")
	number, _ := strconv.Atoi(num)

	if reNum.MatchString(num) == false{
		return false
	}

	if number < 0 {
		return false
	}

	return true
}






