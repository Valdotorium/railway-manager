package time


type Time struct {
	Hour   int
	Minute int
	Second int
}
type Date struct {
	Year  int
	Month int
	Day   int
	Daytime Time
}

//now make functions that update time and Date
func (t *Time) AddSeconds(s int) {
    t.Second = t.Second + s
    if t.Second >= 60 {
        t.Second -= 60
        t.Minute++
        if t.Minute >= 60 {
            t.Minute -= 60
            t.Hour++
            if t.Hour >= 24 {
                t.Hour -= 24
            }
        }
    }
}

func (t *Time) AddMinutes(m int) {
	t.AddSeconds(m * 60)
}

func (t *Time) AddHours(h int) {
    t.AddMinutes(h * 60)
}
func GetDaysInMonth(month int, year int) int{
    switch month {
    case 1, 3, 5, 7, 8, 10, 12:
        return 31
    case 4, 6, 9, 11:
        return 30
    case 2:
        if (year % 4 == 0 && year % 100!= 0) || year % 400 == 0 {
            return 29
        } else {
            return 28
        }
    default:
        return 0
    }
}
func (d *Date) AddDays(days int) {
	//the same as addSeconds(), but with days, months etc..
	d.Day += days
	if d.Day > GetDaysInMonth(d.Month, d.Year) {
		d.Day -= GetDaysInMonth(d.Month, d.Year)
        d.Month++
		if d.Month > 12 {
            d.Month -= 12
            d.Year++
        }
    }
}

func (d *Date) AddMonths(months int) {
    d.AddDays(months * GetDaysInMonth(d.Month, d.Year))
}

func (d *Date) AddYears(years int) {
    d.AddMonths(years * 12)
}