import {Component, OnInit} from '@angular/core';
import {LoginService, User} from "../loginform/login.service";


@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.css']
})
export class ToolbarComponent implements OnInit {
  isLoggedIn = false;
  user: User;

  constructor(private loginService: LoginService) {
  }

  ngOnInit(): void {
    this.user = JSON.parse(localStorage.getItem('user')!) as User
    this.isLoggedIn = localStorage.getItem('user') != null
    console.log(this.user)
    console.log("Logged in:" + this.isLoggedIn)
  }

  logout(): void {
    localStorage.clear();
    this.loginService.logout();
    window.location.reload();
  }

}
