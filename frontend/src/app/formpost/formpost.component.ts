import {Component} from '@angular/core';
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

@Component({
  selector: 'app-formpost',
  templateUrl: './formpost.component.html',
  styleUrls: ['./formpost.component.css']
})

export class FormpostComponent {
  constructor(private formBuilder: FormBuilder, private postService: PostService) {
  }

  // locations = [];
  allLocations = [
    new Location(100, 'Gainesville'),
    new Location(101, 'SanJose'),
    new Location(102, 'NewYork'),
    new Location(103, 'Seattle')
  ];
  allTypes = [
    new Type(100, 'Job offered'),
    new Type(101, 'Gig offered'),
    new Type(102, 'Resume/ Job wanted'),
    new Type(104, 'House offered'),
    new Type(105, 'House wanted'),
    new Type(106, 'For sale by owner'),
    new Type(107, 'For sale by dealer'),
    new Type(108, 'Wanted by owner'),
    new Type(109, 'Wanted by dealer'),
    new Type(110, 'Service offered'),
    new Type(111, 'Community'),
    new Type(112, 'Event/ class'),

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
      sellerid: 1,  //Mock
      categoryid: 1,
      subcategoryid: 1,
      Title: 'Test',
      description: 'Test desc'
    };
    this.postService.createPost(post)
  }

}