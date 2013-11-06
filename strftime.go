package strftime

import (
  "fmt"
  "strconv"
  "strings"
  "time"
)

/**
把 ruby 的日期格式化样式转换成go的样式

ruby 的样式见 http://ruby-doc.org/stdlib-2.0.0/libdoc/date/rdoc/DateTime.html

*/

//用ruby样式格式化时间

func fslen(len int) string {
  var lenrs []string
  if len < 0 {
    len = 1
  }
  if len > 9 {
    len = 9
  }
  lenrs = append(lenrs, ".")
  for lf := 0; lf < len; lf++ {
    lenrs = append(lenrs, "9")
  }
  return strings.Join(lenrs, "")
}


// func Strftime2(t *time.Time,format string)(r string) {
//   return
// }

func Strftime(t *time.Time, format string) (r string) {
  var t_format []string
  app := func(f string, idx *int, step int) {
    t_format = append(t_format, f)
    *idx += step
  }

  b := []rune(format)
  for i := 0; i < len(b); i++ {
    if b[i] == '%' && b[i+1] != '%' {
      switch b[i+1] {
      //year
      case 'Y', 'G':
        app(t.Format("2006"), &i, 1)
        continue
      case 'C': //TODO
        s, _ := strconv.Atoi(t.Format("2006"))
        app(strconv.Itoa(s/100), &i, 1)
        continue
      case 'y', 'g':
        app(t.Format("06"), &i, 1)
        continue
        //month
      case 'm':
        app(t.Format("01"), &i, 1)
        continue
      case '_': //处理下划线的,有_m
        if b[i+2] == 'm' {
          app(fmt.Sprintf("%2s", t.Format("1")), &i, 2)
          continue
        }
      case '-': //处理 -,有 -m,-d
        if b[i+2] == 'm' {
          app(t.Format("1"), &i, 2)
          continue
        } else if b[i+2] == 'd' {
          app(t.Format("2"), &i, 2)
          continue
        }
      case 'B':
        app(t.Format("January"), &i, 1)
        continue
      case 'b', 'h':
        app(t.Format("Jan"), &i, 1)
        continue
      case '^':
        if b[i+2] == 'B' {
          app(strings.ToUpper(t.Format("January")), &i, 2)
          continue
        } else if b[i+2] == 'b' {
          app(strings.ToUpper(t.Format("Jan")), &i, 2)
          continue
        } else if b[i+2] == 'A' {
          app(strings.ToUpper(t.Weekday().String()), &i, 2)
          continue
        } else if b[i+2] == 'a' {
          app(strings.ToUpper(t.Weekday().String()[:3]), &i, 2)
          continue
        }
      case 'd':
        app(t.Format("02"), &i, 1)
        continue
      case 'e':
        app(fmt.Sprintf("%2s", t.Format("2")), &i, 1)
        continue
      case 'j':
        app(strconv.Itoa(t.YearDay()), &i, 1)
        continue
      case 'H':
        app(t.Format("15"), &i, 1)
        continue
      case 'k':
        hk, _ := strconv.Atoi(t.Format("15"))
        app(fmt.Sprintf("%2s", strconv.Itoa(hk)), &i, 1)
        continue
      case 'I':
        hki, _ := strconv.Atoi(t.Format("3"))
        app(fmt.Sprintf("%02s", strconv.Itoa(hki)), &i, 1)
        continue
      case 'l':
        hki, _ := strconv.Atoi(t.Format("3"))
        app(fmt.Sprintf("%2s", strconv.Itoa(hki)), &i, 1)
        continue
      case 'P':
        app(t.Format("am"), &i, 1)
        continue
      case 'p':
        app(t.Format("AM"), &i, 1)
        continue
      case 'M':
        app(t.Format("04"), &i, 1)
        continue
      case 'S':
        app(t.Format("05"), &i, 1)
        continue
      case 'L':
        app(t.Format(".999"), &i, 1)
        continue
      case '1', '2', '3', '4', '5', '6', '7', '8', '9':
        if b[i+2] == 'N' {
          if l, err := strconv.Atoi(string(b[i+1])); err == nil {
            app(t.Format(fslen(l)), &i, 2)
          }
        }
        continue
      case 'z':
        app(t.Format("-0700"), &i, 1)
        continue
      case 'Z':
        app(t.Format("MST"), &i, 1)
        continue
      case 'A':
        app(t.Weekday().String(), &i, 1)
        continue
      case 'a':
        app(t.Weekday().String()[:3], &i, 1)
        continue
      case 'w', 'u': //w sunday is 0,0..6 u monday is 1,1..7
        wd := int(t.Weekday() - time.Sunday)
        if wd == 0 && b[i] == 'u' {
          app(strconv.Itoa(7), &i, 1)
        } else {
          app(strconv.Itoa(wd), &i, 1)
        }
        continue
      case 'V':
        _, w := t.ISOWeek()
        app(fmt.Sprintf("%02s", strconv.Itoa(w)), &i, 1)
        continue
      case 's':
        app(strconv.FormatInt(t.Unix(), 10), &i, 1)
        continue
      case 'Q':
        app(strconv.FormatInt(t.UnixNano()/1000000, 10), &i, 1)
        continue
      case 'n':
        app("\n", &i, 1)
        continue
      case 't':
        app("\t", &i, 1)
        continue
        //Combination
      case 'c': //%c - date and time (%a %b %e %T %Y)
        app(fmt.Sprintf(
          "%s %s %s %s:%s:%s %s",
          t.Weekday().String()[:3],
          t.Format("Jan"),
          fmt.Sprintf("%2s", t.Format("2")),
          t.Format("15"),
          t.Format("04"),
          t.Format("05"),
          t.Format("2006")),
          &i, 1)
        continue
      case 'D', 'x':
        app(fmt.Sprintf("%s/%s/%s", t.Format("01"), t.Format("02"), t.Format("06")), &i, 1)
        continue
      case 'F': //%Y-%m-%d
        app(fmt.Sprintf("%s-%s-%s", t.Format("2006"), t.Format("01"), t.Format("02")), &i, 1)
        continue
      case 'v': //%e-%b-%y
        app(fmt.Sprintf("%s-%s-%s", fmt.Sprintf("%2s", t.Format("2")), t.Format("Jan"), t.Format("06")), &i, 1)
        continue
      case 'X', 'T':
        app(fmt.Sprintf("%s:%s:%s", t.Format("15"), t.Format("04"), t.Format("05")), &i, 1)
        continue
      case 'r': //%r - 12-hour time (%I:%M:%S %p)
        hki, _ := strconv.Atoi(t.Format("3"))
        app(fmt.Sprintf("%02s:%s:%s %s",
          strconv.Itoa(hki),
          t.Format("04"),
          t.Format("05"),
          t.Format("AM")),
          &i, 1)
        continue
      case 'R':
        app(fmt.Sprintf("%s:%s", t.Format("15"), t.Format("04")), &i, 1)
        continue
      }
    }
    t_format = append(t_format, string(b[i]))
  }
  r = strings.Join(t_format, "")
  return
}
