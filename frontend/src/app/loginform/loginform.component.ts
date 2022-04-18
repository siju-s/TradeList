import {Component, OnInit} from '@angular/core';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {Login, LoginService, User} from "./login.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-loginform',
  templateUrl: './loginform.component.html',
  styleUrls: ['./loginform.component.css']
})
export class LoginformComponent implements OnInit {
  isLoginFailed = false
  errorMessage = ''
  isFormSubmitted = false
  form: any = {
    email: null,
    password: null
  };
  isLoggedIn = false;
  roles: string[] = [];
  isSuccessful = false
  isSignupFailed = false

  constructor(private modalService: NgbModal, private loginService: LoginService, private router: Router) {

  }

  ngOnInit(): void {
  }

  get f() {
    return this.form.controls;
  }

  submit() {
    this.isFormSubmitted = true
    const {email, password} = this.form;
    console.log(this.form);
    const login: Login = {
      Email: email,
      Password: password
    }
    this.loginService.login(login).subscribe(data => {
      console.log(data)
      const status = data["status"]
      this.isLoginFailed = status != 200;
      this.isLoggedIn = !this.isLoginFailed
      this.errorMessage = data["message"]
      if (this.isLoggedIn) {
        this.onLoggedIn(data["data"]);
      }
    })
  }

  private onLoggedIn(user: User) {
    console.log(user)
    localStorage.setItem('user', JSON.stringify(user))
    this.loginService.setUser(user)
    this.router.navigate([''])
      .then(() => {
        window.location.reload();
      });
  }

  signup() {
    console.log(this.form);
    const user: User = {
      Contact: {
        FirstName: this.form.firstname,
        LastName: this.form.lastname,
        Email: this.form.email,
        Password: this.form.password
      }
    }
    this.loginService.signup(user).subscribe(data => {
      console.log(data.data)
      const status = data["status"]
      this.isSuccessful = status == 201;
      this.errorMessage = data["message"]
      if (this.isSuccessful) {
        this.onLoggedIn(data.data)
      }
      else {
        this.isSignupFailed = true
      }
    })
  }

  reloadPage(): void {
    window.location.reload();
  }
}

