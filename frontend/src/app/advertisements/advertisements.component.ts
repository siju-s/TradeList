import { Component, OnInit } from '@angular/core';
import { NgbCarouselConfig } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-advertisements',
  templateUrl: './advertisements.component.html',
  styleUrls: ['./advertisements.component.css'],
  providers: [NgbCarouselConfig]
})
export class AdvertisementsComponent implements OnInit {



  ngOnInit(): void {
  }

  title = 'ng-carousel-demo';
   
  images = [
    {title: 'First Slide', short: 'First Slide Short', src: "assets/images/carouselone.jpeg"},
    {title: 'Second Slide', short: 'Second Slide Short', src: "assets/images/carouseltwo.jpeg"},
    {title: 'Third Slide', short: 'Third Slide Short', src: "assets/images/carouselthree.jpeg"}
  ];
   
  constructor(config: NgbCarouselConfig) {
    config.interval = 2000;
    config.keyboard = true;
    config.pauseOnHover = true;
  }
}
