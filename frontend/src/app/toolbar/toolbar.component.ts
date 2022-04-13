import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.css']
})
export class ToolbarComponent implements OnInit {
  isLoggedIn = false;

  constructor() { }

  ngOnInit(): void {
    this.isLoggedIn = localStorage.getItem('user') != null
    console.log("Logged in:" + this.isLoggedIn)
  }

  display = false;
  onPress() {
    //this.display = true;

    //To toggle the component
    this.display = !this.display;
  }

}
