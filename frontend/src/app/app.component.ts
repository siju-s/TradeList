import { Component } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Tradelist';
  displaylogin = false;
  displaysignup = false;
  onPresslogin() {
  
    this.displaylogin = !this.displaylogin;
  }
  onPresssignup() {
    this.displaysignup = !this.displaysignup;
  }
}
