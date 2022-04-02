## Project : Tradelist (Sprint 2)
* Contributors : Siju Sakaria, Yashasvi Mutteneni, Mansi Singh, Sharwari Marathe

## Backend 
The programming language used is Golang. ORM library for Golang used is GORM and the database used is SQLite.



## Frontend
All of these components are developed using Angular and Material.
* Login Form component : The login form is used for authenticating returning users, this component is developed using Material and Angular. It uses material components such as MatCard Module, MatFormFieldModule, MatSelectModule, Form Builder. 
* Create User Component : The create user form is used for registering new users, this component was also developed using the same components mentioned above. Users can toggle between the two forms depending on their requirement.
* Filter structure for Grid Component : The filter structure contains the fields through which the posts will be filtered. This component uses the Ng2SearchPipeModule which creates a filter and allows filtering based on the search keywords. 
* Index Page : The index page contains all the components that we have developed tp this date. This is a rough sketch of how out website would look like in the end. End to end user testing and unit testing was performed on each component. 
* Contact Form : This form is used by users visiting the page for viewing posts. The users can use this form to contact the owner who posted the ad. This form uses the ReactiveFormsModule to ensure form validation. 
* Report/Flag Form : Like the name suggests the Report/Flag form can be used by the users to report inappropriate posts. This ensures that all the posts on the website are age and public appropriare. The form uses the Reactive Forms Module for form validation.

* Cypress Test cases for Create Post, Login, Signup Forms
![Alt text](cypresstest2.png)

