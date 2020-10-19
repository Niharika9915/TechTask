
//building a webserver//
package main
 import(
     "fmt"
     "net/http"
 )

 func ScheduleMeeting (w http.ResponseWriter, r*http.Request){
     fmt.Fprintf(w, "Schedule Meeting")
 }

 func main(){
     http.HandleFunc("/", ScheduleMeeting)
     http.ListenAndServe(":8080", nil)
 }

 //HTML template//
 <!DOCTYPE html>
<html>
<head>
<title>Start a Meeting
/meetings </title>
</head>
<body>
<p>Enter Id {{<element id="id">}}</p>
<p>title {
    display: none;
}
</p>
<label for="Number of participants"> Quantity (between 1 and 100):</label>
<input type="number" id="participants"
name="participants" min="1" max="100">
<p>Start Time {{.Time}}</p>
<p>End Time {{.Time}}</p>
<p> var n = event.timeStamp;
event.timeStamp </p>
</body>
</html>


//template.go uses homepage.html to link to the final server//


package main

import (
  "html/template"
  "log template"
  "net/http"
  "attributes"
)

type PageVariables struct {
	Start Time        string
	End Time         string
    Id                string
    title             string
    participants      int 
    Timestamp         char
}

func main() {
	http.HandleFunc("/", WebPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func WebPage(w http.ResponseWriter, r *http.Request){

    now := time.Now() 
    HomePageVars := PageVariables{ 
      Start time: now.Format("15:04:05"),
      End Time: now.Format("15:04:05"),
       Id:  (element id="id"),             
      title:   now.Format("Enter your name")          
      participants: type="number" id="participants"
     name="participants" min="1" max="100"     
      Timestamp:  event.timeStamp
    }

    
    err = t.Execute(w, HomePageVars) 
    if err != nil { 
  	  log.Print("template executed error: ", err) 

        //RESPONDING TO USER's REQUEST//

    <!DOCTYPE html>
<html>

<head>
<script type="text/javascript' </script>
<title>{{.PageTitle}}</title>
</head>

<script type="text/javascript">
 $(document).ready(function() {
   $('input[name="Enter your name"]').change(function()
   $('input[Email="Enter your Email Id"]').change(function()
   
   {
     $('form').submit();
   });
});
</script>
<body>
<script type='text/javascript'>
 $(document).ready(function() {
   $('input[RSVP="Enter your choice"]').change(function()
  {{with $1:=.PageRadioButtons}}
  <p> Which do you prefer</p>

    <form action="/selected" method="post">
         {{range $1}}
           <input type="radio" choice={{.yes}} choice={{.no}} {{if .yes}}
            {{end}} {{if .no}}checked{{end}} {{end}} {{if .MayBe}}checked{{end}} {{end}} {{if .Not Answered}}checked{{end}}> {{.Text}}
         {{end}}
    </form>
  {{end}}

</body>
</html>

//Final Execution//

package main 

import (
  "fmt"
	"net/http"
	"os"
)




func test(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Join the Schedule meeting")
}