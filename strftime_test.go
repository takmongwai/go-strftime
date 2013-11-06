package strftime

import (
	"testing"
  "time"
  "log"
  
)

var t time.Time


func init(){
  t, _ = time.Parse("2006-01-02 15:04:05.999999999 MST -0700", "2013-02-06 09:37:42.753734088 CST +0800")
  log.Printf("method: %s %v","init()",t)
}

func test_format(test *testing.T,format string,assert string){
  if rs := Strftime(&t,format);rs != assert{
    test.Errorf("format: %s ,want: %s ,got: %s ",format,assert,rs)
  }
}

func TestFormat(t *testing.T){
  cs := map[string]string {
    //year
   "%Y":"2013",
   "%G":"2013",
   "%C":"20",
   "%y":"13",
   "%g":"13",
   //month
   "%m":"02",
   "%_m":" 2",
   "%-m":"2",
   "%B":"February",
   "%^B":"FEBRUARY",
   "%b":"Feb",
   "%^b":"FEB",
   "%h":"Feb",
   
   //day
   "%d":"06",
   "%-d":"6",
   "%e":" 6",
   //day of the year
   "%j":"37",
   
   
   //Time
   //Hour
   "%H":"09",
   "%k":" 9",
   "%I":"09",
   "%l":" 9",
   "%P":"am",
   "%p":"AM",
   
   //Minute
   "%M":"37",
   //Second
   "%S":"42",
   //Millisecond of the second
   "%L":".753",
   //Fractional seconds digits
   "%1N":".7",
   "%2N":".75",
   "%3N":".753",
   "%4N":".7537",
   "%5N":".75373",
   "%6N":".753734",
   "%7N":".753734",
   "%8N":".75373408",
   "%9N":".753734088",
   //Time zone
   "%z":"+0800",
   "%Z":"CST",
   "%A":"Wednesday",
   "%^A":"WEDNESDAY",
   "%a":"Wed",
   "%^a":"WED",
   "%w":"3",
   "%u":"3",
   "%V":"06",
   "%s":"1360114662",
   "%Q":"1360114662753",
   "%t":"\t",
   "%n":"\n",
   "%c":"Wed Feb  6 09:37:42 2013",
   "%D":"02/06/13",
   "%v":" 6-Feb-13",
   "%X":"09:37:42",
   "%T":"09:37:42",
   "%r":"09:37:42 AM",
   "%R":"09:37",
   
   
   "%Y-%m":"2013-02",
   "%Y-%_m":"2013- 2",
   "%Y-%-m":"2013-2",
   
   "%y-%m":"13-02",
   "%y-%_m":"13- 2",
   "%y-%-m":"13-2",
   
   
   "%Y-%m-%d":"2013-02-06",
   "%Y-%_m-%d":"2013- 2-06",
   "%Y-%-m-%d":"2013-2-06",
   "%Y-%-m-%-d":"2013-2-6",
  }
  for f,a := range cs{
    test_format(t,f,a)
  }
}

