import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';

export class Location {
  constructor(public Id:number, public locationName:string) {
  }
}

export class Type {
  constructor(public Id:number, public typeName:string) {
  }
}
// export class MyErrorStateMatcher implements ErrorStateMatcher {
//   isErrorState(control: FormControl | null, form: FormGroupDirective | NgForm | null): boolean {
//     const isSubmitted = form && form.submitted;
//     return !!(control && control.invalid && (control.dirty || control.touched || isSubmitted));
//   }
// }

@Component({
  selector: 'app-formpost',
  templateUrl: './formpost.component.html',
  styleUrls: ['./formpost.component.css']
})
// export class FormpostComponent implements OnInit {

//   constructor() { }

//   ngOnInit(): void {
//   }

// }
export class FormpostComponent  {
  constructor(private formBuilder:FormBuilder){}
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
    allLocations : [
      new Location(100, 'Gainesville'),
      new Location(101, 'SanJose'),
      new Location(102, 'NewYork'),
      new Location(103, 'Seattle')
  ]
      // {value: 'steak-0', viewValue: 'Steak'},
      // {value: 'pizza-1', viewValue: 'Pizza'},
      // {value: 'tacos-2', viewValue: 'Tacos'},
    ,
    allTypes : [
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

    ],
    Description: [''],
    Compensation: [''],
    FirstName:[''],
    LastName: [''],
    address:[''],
    email : (''),
    phone:['']
  });


  saveForm(){
    console.log('Form data is ', this.profileForm.value);
  }

}


// import {Component} from '@angular/core';

// interface Food {
//   value: string;
//   viewValue: string;
// }

// /**
//  * @title Basic select
//  */
// @Component({
//   selector: 'select-overview-example',
//   templateUrl: 'select-overview-example.html',
// })
// export class SelectOverviewExample {
//   foods: Food[] = [
//     {value: 'steak-0', viewValue: 'Steak'},
//     {value: 'pizza-1', viewValue: 'Pizza'},
//     {value: 'tacos-2', viewValue: 'Tacos'},
//   ];
// }
