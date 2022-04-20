import {AfterViewInit, Component} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {Categories, Job, JobPost, Location, Post, PostService, Subcategories} from "../formpost/post.service";
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {LoginService} from "../loginform/login.service";
import {Router} from '@angular/router';
import {FormlyFieldConfig} from "@ngx-formly/core";

@Component({
  selector: 'app-editpost',
  templateUrl: './editpost.component.html',
  styleUrls: ['./editpost.component.css']
})

export class EditpostComponent implements AfterViewInit {

  categories: Categories[]
  subcategories: Subcategories[]
  locations: Location[] = [];
  payTypes: Array<any> = [];
  locationTypes: Array<any> = [];
  jobTypes: Array<any> = [];

  selectedCategory: Categories;
  selectedSubcategory: Subcategories;
  selectedLocation: Location;
  selectedCategoryId: number = 0;
  selectedTitle: string = '';
  isPostCreationFailed = false;
  isSuccessful = false;
  errorMessage: string;
  title = 'appBootstrap';
  isLoggedIn: boolean;

  form: FormGroup;
  model: any = {};
  options = {
    formState: {
      mainModel: this.model,
    },
  };

  fields: FormlyFieldConfig[] = [
    {
      key: 'Title',
      type: 'input',
      templateOptions: {
        label: 'Title',
        required: true
      }

    },
    {
      key: 'Description',
      type: 'textarea',
      templateOptions: {
        label: 'Description',
        required: true
      }

    },
    {
      key: 'category',
      type: 'select',
      templateOptions: {
        label: 'Category',
        required: true,
        valueProp: 'CategoryId',
        labelProp: 'Name',
        change: (field, _) => {
          this.selectedCategory = this.categories.find(item => item.CategoryId == field.formControl!.value)!
          console.log(this.selectedCategory)
          this.fetchSubcategories(this.getSubcategoryField(), field.formControl!.value)
        }
      },
      hooks: {
        onInit: (field) => this.loadOptions(field)
      }
    },
    {
      key: 'subcategory',
      type: 'select',
      templateOptions: {
        label: 'Subcategory',
        required: true,
        valueProp: 'SubcategoryId',
        labelProp: 'Name',
        change: (field, _) => {
          this.selectedSubcategory = this.subcategories.find(item => item.SubcategoryId == field.formControl!.value)!
          console.log(this.selectedSubcategory)
        }
      },
      hideExpression: '!model.category',
    },
    {
      key: 'location',
      type: 'select',
      templateOptions: {
        label: 'Location',
        required: true,
        valueProp: 'Name',
        labelProp: 'Name',
        change: (field, _) => {
          this.selectedLocation = this.locations.find(item => item.Name == field.formControl!.value)!
          console.log(this.selectedLocation)
        }
      },
      hooks: {
        onInit: (field) => this.loadOptions(field)
      }
    },
    {
      className: 'col-6 float-left',
      key: 'salary',
      type: 'input',
      templateOptions: {
        label: 'Salary',
        required: true,
      },
      hideExpression: (model: any, _: any) => {
        if (model.category) {
          return model.category != 1
        } else {
          return true;
        }
      },
    },
    {
      className: 'col-6 float-left',
      key: 'paytype',
      type: 'select',
      templateOptions: {
        label: 'PayType',
        valueProp: 'value',
        labelProp: 'label',
        required: true
      },

      hideExpression: (model: any, _: any) => {
        if (model.category) {
          return model.category != 1
        } else {
          return true;
        }
      },
      hooks: {
        onInit: (field) => this.loadOptions(field)
      }
    },

    {
      key: 'locationtype',
      type: 'select',
      templateOptions: {
        label: 'Location Type',
        valueProp: 'value',
        labelProp: 'label',
        required: true
      },

      hideExpression: (model: any, _: any) => {
        if (model.category) {
          return model.category != 1
        } else {
          return true;
        }
      },
      hooks: {
        onInit: (field) => this.loadOptions(field)
      }
    },

    {
      key: 'jobtype',
      type: 'select',
      templateOptions: {
        label: 'Job Type',
        valueProp: 'value',
        labelProp: 'label',
        required: true
      },

      hideExpression: (model: any, _: any) => {
        if (model.category) {
          return model.category != 1
        } else {
          return true;
        }
      },
      hooks: {
        onInit: (field) => this.loadOptions(field)
      }
    },
    {
      key: "file",
      type: "file",
      templateOptions: {
        multiple: true
      }
    }
  ]

  constructor(private formBuilder: FormBuilder, private postService: PostService, private loginService: LoginService,
              private modalService: NgbModal, private router: Router) {
    this.form = formBuilder.group({})
  }

  ngOnInit() {
    this.isLoggedIn = localStorage.getItem('user') != null
    if (!this.isLoggedIn) {
      this.router.navigate(['/login'])
      return
    }
    var user = localStorage.getItem('user')
    if (user != null) {
      console.log("User id:"+ JSON.parse(user)["ID"])
    }
  }

  loadOptions(field?: FormlyFieldConfig) {
    if (!field || !field.templateOptions) {
      return;
    }
    switch (field.key) {
      case 'paytype' : {
        this.setFieldData(field, this.postService.getPayTypes(), this.payTypes)
        field.formControl!.setValue(this.payTypes[0].value)
        break;
      }
      case 'locationtype' : {
        this.setFieldData(field, this.postService.getLocationTypes(), this.locationTypes)
        field.formControl!.setValue(this.locationTypes[0].value)
        break;
      }
      case 'jobtype' : {
        this.setFieldData(field, this.postService.getJobTypes(), this.jobTypes)
        field.formControl!.setValue(this.jobTypes[0].value)
        break;
      }
      case 'category': {
        this.initCategories(field)
        break;
      }
      case 'location': {
        this.initLocations(field)
        break;
      }
    }
    console.log(this.payTypes)
  }

  initCategories(field: FormlyFieldConfig) {
    this.postService.fetchCategories().subscribe(data => {
        this.categories = []
        this.setFieldData(field, data.data, this.categories)
        console.log(this.categories)
        this.selectedCategory = this.categories[0]
        field.formControl!.setValue(this.selectedCategory.CategoryId)
        this.fetchSubcategories(this.getSubcategoryField(), this.selectedCategory.CategoryId)
      }
    )
  }

  initLocations(field: FormlyFieldConfig) {
    this.postService.fetchLocations().subscribe(data => {
        this.locations = []
        this.setFieldData(field, data.data, this.locations)
        console.log(this.locations)
        this.selectedLocation = this.locations[0]
        field.formControl!.setValue(this.selectedLocation.Name)
      }
    )
  }

  private getSubcategoryField() {
    return this.fields.find(field => field.key == 'subcategory')!;
  }

//   filterSubById(id:any) {
//     return this.subCategories.filter(item => item.parentId === id);
// }

  setFieldData(field: FormlyFieldConfig, data: any, resultArr: any[]) {
    data.forEach((item: any) => {
      resultArr.push(item);
    })
    field.templateOptions!.options = resultArr
  }

  private fetchSubcategories(field: FormlyFieldConfig, categoryId: number) {
    this.postService.fetchSubcategories(categoryId).subscribe(data => {
        this.subcategories = []

        this.setFieldData(field, data.data, this.subcategories)
        console.log(this.subcategories)
        this.selectedSubcategory = this.subcategories[0]
        field.formControl!.setValue(this.selectedSubcategory.SubcategoryId)
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
    const data = this.form.value;
    console.log('Form data is ', data);
    console.log(localStorage.getItem('user'))

    console.log(data.file);
    let sellerid = 0;

    const user = localStorage.getItem('user')

    console.log(user)

    if (user != null || user != undefined) {
      sellerid = JSON.parse(user)["ID"]
    }

    if (sellerid == 0) {
      return
    }

    const post: Post = {
      SellerId: sellerid,  //Mock
      CategoryId: this.selectedCategory.CategoryId,
      SubcategoryId: this.selectedSubcategory.SubcategoryId,
      Title: data.Title,
      Description: data.Description,
    };

    const job: Job = {
      Salary: data.salary,
      Pay: data.paytype,
      Type: data.jobtype,
      Location: data.locationtype,
      Place: this.selectedLocation.Name
    }

    const jobPost: JobPost = {
      Post: post,
      Job: job
    }

    console.log(jobPost)

    this.postService.createPost(jobPost, data.file).subscribe(data => {
      console.log(data)
      this.isSuccessful = data["status"] == 200
      this.isPostCreationFailed = !this.isSuccessful
      this.errorMessage = data["message"]
      if (this.isSuccessful) {
        this.form.reset()
      }
    })
  }

  ngAfterViewInit(): void {
  }

}

