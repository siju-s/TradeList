import {Component, ElementRef, ViewChild} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {Post, PostService} from "../post.service";

export class Location {
  constructor(public Id: number, public locationName: string) {
  }
}

export class Type {
  constructor(public Id: number, public typeName: string) {
  }
}
export class SubCategory {
  constructor(public Id: number, public subcategoryName: string) {
  }
}

@Component({
  selector: 'app-formpost',
  templateUrl: './formpost.component.html',
  styleUrls: ['./formpost.component.css']
})

export class FormpostComponent {
  constructor(private formBuilder: FormBuilder, private postService: PostService) {
  }
  
  // allSubCategories = []
  // locations = [];
  allLocations = [
    new Location(100, 'Gainesville'),
    new Location(101, 'SanJose'),
    new Location(102, 'NewYork'),
    new Location(103, 'Seattle')
  ];
  allTypes = [
    new Type(1, 'Job'),
    new Type(2, 'Property'),
    new Type(3, 'For Sale'),
  ];
  // allTypes = [
  //   new Type(1, 'Job offered'),
  //   new Type(1, 'Gig offered'),
  //   new Type(1, 'Resume/ Job wanted'),
  //   new Type(1, 'House offered'),
  //   new Type(1, 'House wanted'),
  //   new Type(, 'For sale by owner'),
  //   new Type(107, 'For sale by dealer'),
  //   new Type(108, 'Wanted by owner'),
  //   new Type(109, 'Wanted by dealer'),
  //   new Type(110, 'Service offered'),
  //   new Type(111, 'Community'),
  //   new Type(112, 'Event/ class'),

  // ];
  allSubCategories = [
     new SubCategory(1,  "Accounting"),
		 new SubCategory(1, "HR"),
		 new SubCategory(1,  "Legal"),
		 new SubCategory(1,  "Customer Service"),
		 new SubCategory(1,  "Healthcare"),
		 new SubCategory(1,  "Hospitality"),
		 new SubCategory(1,  "Housekeeping"),
		 new SubCategory(1,  "Software"),
		 new SubCategory(1,  "Accounting"),
     new SubCategory(2, "For Sale"),
     new SubCategory(2, "To Rent"),
     new SubCategory(2,  "To Share"),
     new SubCategory(2,  "Sublet"),
     new SubCategory(2,  "Storage"),
     new SubCategory( 3,  "Appliances"),
     new SubCategory( 3,  "Audio equipment"),
     new SubCategory( 3,  "Books"),
     new SubCategory( 3,  "Clothes"),
     new SubCategory( 3,  "Computers"),
     new SubCategory( 3,  "Furniture"),
		 new SubCategory( 3,  "Gym equipment"),
     new SubCategory( 3,  "Sports equipment")
  ];
  
  
  profileForm = this.formBuilder.group({
    Description: [''],
    Compensation: [''],
    FirstName: [''],
    LastName: [''],
    address: [''],
    email: (''),
    phone: ['']
  });

  createPost() {
    console.log('Form data is ', this.profileForm.value);
    const post: Post = {
      Sellerid: 1,  //Mock
      Categoryid: 1,
      Subcategoryid: 1,
      Title: 'Test',
      Description: 'Test desc'
    };
    this.postService.createPost(post)
  }

}
