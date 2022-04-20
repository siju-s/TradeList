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
  forceClearStorage = true

  constructor(private loginService: LoginService) {
  }

  ngOnInit(): void {
    const clearStorage = localStorage.getItem('forceClearStorage')
    // Flag to clear storage ONCE ONLY to avoid user data not saved due to login issue
    if (clearStorage == null) {
      localStorage.clear()
      localStorage.setItem('forceClearStorage', true.toString())
    }
    this.user = JSON.parse(localStorage.getItem('user')!) as User
    this.isLoggedIn = localStorage.getItem('user') != null
    console.log(this.user)
    console.log("Logged in:" + this.isLoggedIn)
  }

  logout(): void {
    localStorage.removeItem('user');
    this.loginService.logout();
    window.location.reload();
  }

}
