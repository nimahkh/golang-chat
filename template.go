package main

import "html/template"

var Html = template.Must(template.New("chat_room").Parse(`
<html> 
<head> 
    <title>{{.roomid}}</title>
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Itim">
    <link rel="stylesheet" type="text/css" href="http://meyerweb.com/eric/tools/css/reset/reset.css">
    <link rel="stylesheet" type="text/css" href="/assets/style.css">
    <script src="http://ajax.googleapis.com/ajax/libs/jquery/1.7/jquery.js"></script> 
    <script src="http://malsup.github.com/jquery.form.js"></script> 
    <script> 
        $('#message_form').focus();
        $(document).ready(function() { 
			const alarm = document.getElementById('play');
            // bind 'myForm' and provide a simple callback function 
            $('#myForm').ajaxForm(function() {
                $('#message_form').val('');
                $('#message_form').focus();
            });
            if (!!window.EventSource) {
                var source = new EventSource('/stream/{{.roomid}}/{{.userid}}');
                source.addEventListener('message', function(e) {
					const userName= e.data.split(":")[0];
					const darker = userName==={{.isme}} && "darker" 
					const data= "<div class='container "+darker+"'>"+e.data+"</div>"
                    $('#messages').append(data + "</br>");
                    $('html, body').animate({scrollTop:$(document).height()}, 'slow');
					if({{.isme}} !== e.data.split(":")[0]){
                      alarm.volume= 0.3
					  alarm.play()
					}
                }, false);
            } else {
                alert("NOT SUPPORTED");
            }
        });	
    </script> 
    </head>
    <body>
	<audio id="play" src="/assets/alarm.mp3" preload="auto"></audio>
	<div class="main-container">
    	<div id="messages"></div>
		<form id="myForm" action="/room/{{.roomid}}/{{.userid}}" method="post"> 
		User: <input id="user_form" name="user" value="{{.userid}}">
		Message: <input id="message_form" name="message">
		<input type="submit" value="Submit"> 
		</form>
	</div>

</body>
</html>
`))
