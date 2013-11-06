## go-strftime

golang date time format like ruby layout


## install
go get github.com/weidewang/go-strftime


## test

<pre>

package main

import (
      "fmt"
      "github.com/weidewang/go-strftime"
      "time"
      )

func main() {

t := time.Now()
     s := strftime.Strftime(&t, "%Y-%m-%d %H:%M:%S")
     fmt.Println(s)

}



</pre>
