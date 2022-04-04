import {AfterViewInit, Component, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {Categories, Job, JobPost, Location, Post, PostService, Subcategories} from "./post.service";
import {UploadFilesComponent} from "../upload.component";
import {ModalDismissReasons, NgbModal} from '@ng-bootstrap/ng-bootstrap';

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

  title = 'appBootstrap';
  closeResult: string = '';
  profileForm: FormGroup;

  constructor(private formBuilder: FormBuilder, private postService: PostService, private modalService: NgbModal) {
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
    this.postService.fetchCategories().subscribe(data => {
        this.categories = data.data
        console.log(this.categories)
      }
    )
    this.postService.fetchLocations().subscribe(data => {
      this.locations = data.data
      console.log(this.locations);
    })
  }

  open(content: any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${FormpostComponent.getDismissReason(reason)}`;
    });
  }

  private static getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return `with: ${reason}`;
    }
  }

//   filterSubById(id:any) {
//     return this.subCategories.filter(item => item.parentId === id);
// }


  selectChangeHandler(event: any) {
    console.log(this.selectedCategory)

    this.postService.fetchSubcategories(this.selectedCategory.CategoryId).subscribe(data => {
        this.subcategories = data.data
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

    const files = this.child?.getSelectedFiles();

    console.log(files);

    const post: Post = {
      Sellerid: 1,  //Mock
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

    this.postService.createPost(jobPost, files)
  }

  ngAfterViewInit(): void {
  }

}

