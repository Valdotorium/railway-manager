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
func (t *time) addSeconds(s int) {
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

func (t *time) addMinutes(m int) {
	t.addSeconds(m * 60)
}

func (t *time) addHours(h int) {
    t.addMinutes(h * 60)
}
func getDaysInMonth(month int, year int) int{
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
func (d *Date) addDays(days int) {
	//the same as addSeconds(), but with days, months etc..
	d.Day += days
	if d.Day > getDaysInMonth(d.Month, d.Year) {
		d.Day -= getDaysInMonth(d.Month, d.Year)
        d.Month++
		if d.Month > 12 {
            d.Month -= 12
            d.Year++
        }
    }
}

func (d *Date) addMonths(months int) {
    d.addDays(months * getDaysInMonth(d.Month, d.Year))
}

func (d *Date) addYears(years int) {
    d.addMonths(years * 12)
}