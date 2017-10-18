# Go Problem Sheet 2

This project is based off a problem sheet we have been given by our lecturer for the purposes of learning the go language, while using curl and Web Applications 




**Getting Started**





**List of Progressions**


—Part 1—


	Objective: 

	Create a web application in Go that responds with the text “Guessing game”. 		
	This should be the response body irrespective of what request is received. 		
	Explain in your README how to examine the response, including the headers, 
	using curl.

	Implementation: 
	
	To find a response in Curl :

	- Open Terminal or Command Prompt
	- Type :  curl -I http://localhost:8080/test

	- RESULT : 

		HTTP/1.1 200 OK						—Response
		Date: Tue, 17 Oct 2017 17:09:00 GMT			-Time
		Content-Length: 23					-Char length	
		Content-Type: text/plain; charset=utf-8			—Charset and type



		Guessing Game test 					-Body

	

—Part 2—


	Objective: 
	
	Change your web application to make “Guessing game” a level 1 		
	heading in HTML. Test your application, making sure that the HTML is rendered by 		
	your browser. If it isn’t, fix it.


	Implementation: 
	
	To find a response in Curl :

	- Open Terminal or Command Prompt
	- Type :  curl http://localhost:8080/test

	- RESULT : 

		Guessing Game test 					-Body


	This time the Text will be a level 1 Header.

—Part 3—

	Objective: 

	Change the web application to serve a web page rather than hard-coding the text 	
	into the web application executable. Use the Bootstrap starter template, changing 	
	the header to say “Guessing game”. Add a link on the page to the relative URL /		
	guess with the text “New game”. Have this page served as the root resource in the 	
	web application.


	Implementation:
	
	To find a response in Curl :

	- Open Terminal or Command Prompt
	- Type :  curl http://localhost:8080/

	- RESULT : 

	A webpage is shown with the heading "Guessing Game", the page contains a link to another page 
	with the heading "New Game". The "Guess" page has a link to homepage.
	For some reason it would not let me use "guess" with lower case. I will update if I find out why.



—Part 4—

—Part 5—

—Part 6—

—Part 7—

—Part 8—

—Part 9— 




**Versioning**

- Version 1.0 - I do not intend to make any changes to the code here



**Authors**

- Michael Kidd - G00236920@gmit.ie




**License**

This project is licensed under Nothing License - As I have no reason to license this and the code is not very difficult to learn or original in its implementation.




**Acknowledgments**

