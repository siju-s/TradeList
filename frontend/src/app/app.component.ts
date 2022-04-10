import {Component} from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Tradelist';
  isLoggedIn = false;

  constructor() {
    this.isLoggedIn = localStorage.getItem('user') != null
    console.log(this.isLoggedIn)
  }
}
