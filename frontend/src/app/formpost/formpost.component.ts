import {AfterViewInit, Component, ElementRef, ViewChild} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {Job, JobPost, Post, PostService} from "../post.service";
import {UploadFilesComponent} from "../upload.component";
import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';
export class Location {
  constructor(public Id: number, public locationName: string) {
  }
}

// export class Category {
//   constructor(public Id: number, public title: string) {
//   }
// }
// export class SubCategory {
//   constructor(public Id: number, public title: string) {
//   }
// }

@Component({
  selector: 'app-formpost',
  templateUrl: './formpost.component.html',
  styleUrls: ['./formpost.component.css']
})

export class FormpostComponent implements AfterViewInit {

  @ViewChild(UploadFilesComponent) child?: UploadFilesComponent;

  constructor(private formBuilder: FormBuilder, private postService: PostService, private modalService: NgbModal) {
  }
  title = 'appBootstrap';
  closeResult: string = '';
  open(content:any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
  } 
 
  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return  `with: ${reason}`;
    }
  }
  // allSubCategories = []
  // locations = [];
  allLocations = [
    new Location(100, 'Gainesville'),
    new Location(101, 'SanJose'),
    new Location(102, 'NewYork'),
    new Location(103, 'Seattle')
  ];
  // mainCategory = new Category(1, 'Job');
  // mainSubcategory = new SubCategory(1,  "Accounting");
  // mainGroups = [
  //   new Category(1, 'Job'),
  //   new Category(2, 'Property'),
  //   new Category(3, 'For Sale'),
  // ];

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
  // subCategories = [
  //    new SubCategory(1,  "Accounting"),
  //   new SubCategory(1, "HR"),
  //   new SubCategory(1,  "Legal"),
  //   new SubCategory(1,  "Customer Service"),
  //   new SubCategory(1,  "Healthcare"),
  //   new SubCategory(1,  "Hospitality"),
  //   new SubCategory(1,  "Housekeeping"),
  //   new SubCategory(1,  "Software"),
  //    new SubCategory(2, "For Sale"),
  //    new SubCategory(2, "To Rent"),
  //    new SubCategory(2,  "To Share"),
  //    new SubCategory(2,  "Sublet"),
  //    new SubCategory(2,  "Storage"),
  //    new SubCategory( 3,  "Appliances"),
  //    new SubCategory( 3,  "Audio equipment"),
  //    new SubCategory( 3,  "Books"),
  //    new SubCategory( 3,  "Clothes"),
  //    new SubCategory( 3,  "Computers"),
  //    new SubCategory( 3,  "Furniture"),
  //   new SubCategory( 3,  "Gym equipment"),
  //    new SubCategory( 3,  "Sports equipment")
  // ];
  // selected(){
  //   alert("hey there")
  // }
  // filterSubById(id: number) {
  //   return this.subCategories.filter(item => item.Id === id);
  // }
//   mainCategory = {
//     title: 'abc',
//     id: 1
//   };
//   subCategory = {
//     title: 'xxx',
//     parentId: 1
//   };

//   mainGroups = [
//     {
//       title: 'abc',
//       id: 1
//     },
//     {
//       title: 'def',
//       id: 2
//     }
//   ]

//   subCategories = [
//     {
//       title: 'xxx',
//       parentId: 1
//     },
//     {
//       title: 'yyy',
//       parentId: 1
//     },
//     {
//       title: 'zzz',
//       parentId: 2
//     }
//   ]

//   filterSubById(id:any) {
//     return this.subCategories.filter(item => item.parentId === id);
// }
  selectedDay: string ='';
  selectedId:number=0;
  mainGroups = [
        {
          title: 'Job',
          id: 1
        },
        {
          title: 'Property',
          id: 2
        },
        {
          title: 'For Sale',
          id: 3
        }
      ]

  selectChangeHandler(event : any){
    this.selectedDay = event.target.value;
    for(let i=0; i<this.mainGroups.length;i++){
      if(this.selectedDay == this.mainGroups[i].title){
        this.selectedId = this.mainGroups[i].id
      }
    }
  }
  // subCategories = [
  //    new SubCategory(1,  "Accounting"),
  //   new SubCategory(1, "HR"),
  //   new SubCategory(1,  "Legal"),
  //   new SubCategory(1,  "Customer Service"),
  //   new SubCategory(1,  "Healthcare"),
  //   new SubCategory(1,  "Hospitality"),
  //   new SubCategory(1,  "Housekeeping"),
  //   new SubCategory(1,  "Software"),
  //    new SubCategory(2, "For Sale"),
  //    new SubCategory(2, "To Rent"),
  //    new SubCategory(2,  "To Share"),
  //    new SubCategory(2,  "Sublet"),
  //    new SubCategory(2,  "Storage"),
  //    new SubCategory( 3,  "Appliances"),
  //    new SubCategory( 3,  "Audio equipment"),
  //    new SubCategory( 3,  "Books"),
  //    new SubCategory( 3,  "Clothes"),
  //    new SubCategory( 3,  "Computers"),
  //    new SubCategory( 3,  "Furniture"),
  //   new SubCategory( 3,  "Gym equipment"),
  //    new SubCategory( 3,  "Sports equipment")
  // ];
  subCategories = [
        {
          title: 'Accounting',
          parentId: 1
        },
        {
          title: 'HR',
          parentId: 1
        },
        {
          title: 'Legal',
          parentId: 1
        },
        {
          title: 'Customer Service',
          parentId: 1
        },
        {
          title: 'Healthcare',
          parentId: 1
        },
        {
          title: 'Hospitality',
          parentId: 1
        },
        {
          title: 'Software',
          parentId: 1
        },
        {
          title: 'For Sale',
          parentId: 2
        },
        {
          title: 'To Rent',
          parentId: 2
        },
        {
          title: 'To Share',
          parentId: 2
        },
        {
          title: 'Sublet',
          parentId: 2
        },

        {
          title: 'Storage',
          parentId: 2
        },
        {
          title: 'Appliances',
          parentId: 3
        },
        {
          title: 'Audio equipment',
          parentId: 3
        },
        {
          title: 'Books',
          parentId: 3
        },
        {
          title: 'Clothes',
          parentId: 3
        },
        {
          title: 'Computers',
          parentId: 3
        },
        {
          title: 'Furniture',
          parentId: 3
        },
        {
          title: 'Gym equipment',
          parentId: 3
        },
        {
          title: 'Sports equipment',
          parentId: 3
        }
      ]
  filterSubById() {
        return this.subCategories.filter(item => item.parentId === this.selectedId);
    }



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

    // console.log(this.uploadFilesComponent.getSelectedFiles());

    const files = this.child?.getSelectedFiles();

    console.log(files);

   // TODO Make sure dummy data is replaced with data from form
    const post: Post = {
      Sellerid: 1,  //Mock
      Categoryid: 1,
      Subcategoryid: 1,
      Title: 'Test',
      Description: 'Test desc',
    };

    const job : Job = {
      Salary: 500,
      Pay: "monthly",
      Type: "fulltime",
      Location: "remote",
      Place: "Gainesville"
    }

    const jobPost : JobPost = {
      Post : post,
      Job : job
    }

    this.postService.createPost(jobPost, files)
  }

  ngAfterViewInit(): void {
  }

}

