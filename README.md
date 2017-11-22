# DataQuery Project 2017 

Lecturer        - Ian McCloughlin

Student         - Michael Kidd

Subject         - Data Query and Representation

Project Title   - Eliza Chatbot


Objective       - 

In this project you will create a chatbot web application in Go. 
Users will be able to visit the web application through their browser, type in sentences 
such as “Hi, my name is John.” and the web application will reply with sentences such as 
“Hello John, how are you?” You are free to use your artistic license in creating the chatbot, 
such as by giving it a certain type of personality, but you should be guided by the ELIZA program. 
If you wish to use any technique to enhance your chatbot, such as machine learning methods, 
that is okay as long as the chatbot works. However, reasonable care must be taken to ensure 
that the chatbot you create does not act in an offensive manner, and students are reminded of 
the GMIT Code of Student Conduct.

We need to make a web application that can seem to have a conversation with a real person,
We were advised of the Chatbot Eliza that would mimic the style of a psychiatrist. The chatbot needs to
take input from the user and then generate a response using regular expressions, then output that response
to the webpage. To do this we must use ajax and jquery so that the webpage does not reload when the process
starts. 

Implementation - I Started the project by serving a blank index page to the localhost using port 8080,
then I created template page that was served to in a seperate folder called chat. I created a style on this page
similar to whatapps and iMessage, by floating the users text right and Elizas response left. I then placed a box around
the text and coloured it similar to iMessage (Grey and Green). We were advised of using a bootstrap to create the UI, But I decided to design my own, it may not be as sleek as the boostrap version, but I decide it would be a better option, if anything was to be changed or edited at least I could do it much more easily as I designed it.

I then Used Ajax to take the users input from an input box, Preventing its default action which would cause the page to reload. The input is retrieved by the Golang server program and using a regular expression it is changed to give an answer for the chatbot. 

At this point I do not have the reflections working but hope to have them working before the deadline.
The Regular expressions strings I had found in a python version of Eliza and adapted them to a 2D slice and changed the regex info to be relevent.

So far I am happy with the result. In the project brief it mentions using machine learning with the chatbot, which if the project was of a higher weight or had alot more time, then I would possibly look into accomplishing this, But we must use the time we are given wisely.

Download Process 
- Open up a command line and use the command
- change directory to the location you wish to use
- eg. "cd Documents"
- type the following command to clone the repo
- git clone https://github.com/G00236920/DataQuery
- cd to the Eliza directory once downloaded, then
- run command or termnal "go build Eliza.go".
- then run command "./eliza"

You will now be able to go to 
- localhost:8080
and see the Webserver hosted, Click on start chat and Start talking to Eliza.



Author - Michael Kidd
