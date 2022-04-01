// import { Component, OnInit } from '@angular/core';

// @Component({
//   selector: 'app-grid',
//   templateUrl: './grid.component.html',
//   styleUrls: ['./grid.component.css']
// })
// export class GridComponent implements OnInit {
//   title = 'Card View Demo';
//   gridColumns = 3;
//   toggleGridColumns() {
//     this.gridColumns = this.gridColumns === 3 ? 4 : 3;
//   }
//   constructor() { }

//   ngOnInit(): void {
//   }

// }
// ###################
import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import { NgbCarouselConfig } from '@ng-bootstrap/ng-bootstrap';
import { ContactformComponent } from '../contactform/contactform.component';
import {Post, PostService} from "../post.service";
@Component({
  selector: 'app-grid',
  templateUrl: './grid.component.html',
  styleUrls: ['./grid.component.css'],
  providers: [NgbCarouselConfig]
})
export class GridComponent implements OnInit {
  componentName = "user"
  images = [

    {title: 'Second Slide', short: 'Second Slide Short', src: "assets/images/image.jpeg"},
    {title: 'First Slide', short: 'First Slide Short', src: "assets/images/floorplans copy.png"},
    {title: 'Third Slide', short: 'Third Slide Short', src: "assets/images/amenities.jpg"}
  ];
  title = 'Card View Demo';
  gridColumns = 3;
  toggleGridColumns() {
    this.gridColumns = this.gridColumns === 3 ? 4 : 3;
  }
  // constructor(config: NgbCarouselConfig) {
  //   // config.interval = 2000;
  //   // config.keyboard = true;
  //   // config.pauseOnHover = true;
  // }
  // constructor() { }

  post:Array<Post> = []

  constructor(private postService: PostService, private changeDetection: ChangeDetectorRef) { }

  ngOnInit(): void {
    this.postService.getPosts().subscribe(data => {
      this.post = data.data
      console.log(data.data[0])
      this.changeDetection.detectChanges()
    })
  }


  // ngOnInit(): void {
  // }

}
// ***********
// import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
// import {Post, PostService} from "../post.service";
// @Component({
//   selector: 'app-grid',
//   templateUrl: './grid.component.html',
//   styleUrls: ['./grid.component.css']
// })
// export class GridComponent implements OnInit {

//   post:Array<Post> = []

//   constructor(private postService: PostService, private changeDetection: ChangeDetectorRef) { }

//   ngOnInit(): void {
//     this.postService.getPosts().subscribe(data => {
//       this.post = data.data
//       console.log(data.data[0])
//       this.changeDetection.detectChanges()
//     })

// }
// }
// ********
// import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
// import {Post, PostService} from "../post.service";

// @Component({
//   selector: 'app-viewpost',
//   templateUrl: './viewpost.component.html',
//   styleUrls: ['./viewpost.component.css']
// })
// export class ViewpostComponent implements OnInit {

//   post:Array<Post> = []

//   constructor(private postService: PostService, private changeDetection: ChangeDetectorRef) { }

//   ngOnInit(): void {
//     this.postService.getPosts().subscribe(data => {
//       this.post = data.data
//       console.log(data.data[0])
//       this.changeDetection.detectChanges()
//     })
//   }