import {AfterViewInit, Component, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {Categories, Job, JobPost, Location, Post, PostService, Subcategories} from "./post.service";
import {UploadFilesComponent} from "../upload.component";
import {ModalDismissReasons, NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {LoginService} from "../loginform/login.service";
import {Router} from '@angular/router';

@Component({
  selector: 'app-formpost',
  templateUrl: './formpost.component.html',
  styleUrls: ['./formpost.component.css']
})

export class FormpostComponent implements AfterViewInit {

  @ViewChild(UploadFilesComponent) child?: UploadFilesComponent;
  categories: Categories[]
  subcategories: Subcategories[]
  locations: Location[];
  selectedCategory: Categories;
  selectedSubcategory: Subcategories;
  selectedLocation: Location;
  selectedCategoryId: number = 0;
  selectedTitle: string = '';
  isPostCreationFailed = false;
  isSuccessful = false;
  errorMessage: string;
  title = 'appBootstrap';
  closeResult: string = '';
  profileForm: FormGroup;
  isLoggedIn : boolean;

  constructor(private formBuilder: FormBuilder, private postService: PostService, private loginService: LoginService,
              private modalService: NgbModal, private router: Router) {
    this.profileForm = formBuilder.group({
      Title: [''],
      Location: [''],
      Category: [''],
      SubCategory: [''],
      Price: [''],
      Email: [''],
      PhoneNo: [''],
      Description: ['']
    })
  }

  ngOnInit() {
    this.isLoggedIn = localStorage.getItem('user') != null
    if (!this.isLoggedIn) {
      this.router.navigate(['/login'])
      return
    }
    this.postService.fetchCategories().subscribe(data => {
        this.categories = data.data
        this.selectedCategory = this.categories[0]
        console.log(this.categories)
        this.fetchSubcategories()
      }
    )
    this.postService.fetchLocations().subscribe(data => {
      this.locations = data.data
      this.selectedLocation = this.locations[0]
      console.log(this.locations);
    })
  }

//   filterSubById(id:any) {
//     return this.subCategories.filter(item => item.parentId === id);
// }


  selectChangeHandler(event: any) {
    console.log(this.selectedCategory)
    this.fetchSubcategories();
  }

  private fetchSubcategories() {
    this.postService.fetchSubcategories(this.selectedCategory.CategoryId).subscribe(data => {
        this.subcategories = data.data
        this.selectedSubcategory = this.subcategories[0]
        console.log(this.subcategories)
      }
    )
  }

  getTitle(event: any) {
    this.selectedTitle = event.target.value;
  }


  // filterSubById() {
  //   return this.subCategories.filter(item => item.parentId === this.selectedCategoryId);
  // }

  createPost() {
    const data = this.profileForm.value;
    console.log('Form data is ', data);
    console.log(localStorage.getItem('user'))

    const files = this.child?.getSelectedFiles();

    console.log(files);
    let sellerid = 0;

    const user = localStorage.getItem('user')

    console.log(this.loginService.getUser())

    if (user != null || user != undefined) {
      sellerid = JSON.parse(user)["ID"]
    }

    if (sellerid == 0) {
      return
    }

    const post: Post = {
      Sellerid: sellerid,  //Mock
      Categoryid: this.selectedCategory.CategoryId,
      Subcategoryid: this.selectedSubcategory.SubcategoryId,
      Title: data.Title,
      Description: data.Description,
    };

    const job: Job = {
      Salary: 500,
      Pay: "monthly",
      Type: "fulltime",
      Location: "remote",
      Place: this.selectedLocation.Name
    }

    const jobPost: JobPost = {
      Post: post,
      Job: job
    }

    console.log(jobPost)

    this.postService.createPost(jobPost, files).subscribe(data => {
      console.log(data)
      this.isSuccessful = data["status"] == 200
      this.isPostCreationFailed = !this.isSuccessful
      this.errorMessage = data["message"]
    })
  }

  ngAfterViewInit(): void {
  }

}

