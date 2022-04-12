import {Component, OnInit} from '@angular/core';
import {NgbCarouselConfig} from '@ng-bootstrap/ng-bootstrap';
import {Categories, Image, JobPost, Post, PostService, Response, Subcategories} from "../formpost/post.service";
import {ActivatedRoute} from "@angular/router";
import {DataService} from "../shared/DataService";

@Component({
  selector: 'app-grid',
  templateUrl: './grid.component.html',
  styleUrls: ['./grid.component.css'],
  providers: [NgbCarouselConfig]
})
export class GridComponent implements OnInit {
  filterTerm!: string;
  componentName = "user"
  title = 'Card View Demo';
  gridColumns = 3;
  postImageMap = new Map<number, Array<Image>>()
  categories: Categories[]
  subcategories: Subcategories[]

  toggleGridColumns() {
    this.gridColumns = this.gridColumns === 3 ? 4 : 3;
  }

  // constructor(config: NgbCarouselConfig) {
  //   // config.interval = 2000;
  //   // config.keyboard = true;
  //   // config.pauseOnHover = true;
  // }
  // constructor() { }

  post: Array<Post> = []
  jobPost: Array<JobPost> = []


  constructor(private postService: PostService, private dataService: DataService, private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    console.log("ngOnInit")
    this.initObservers()
    let subcategoryId = Number(this.route.snapshot.paramMap.get('id'))
    console.log(subcategoryId)
    if (subcategoryId === 0) {
      this.postService.getPosts().subscribe(data => {
        if (data.data) {
          this.handlePostData(data);
        }
        console.log(data)
      })
      return
    }
    this.postService.getPostsForSubcategory(subcategoryId).subscribe(data => {
      if (data.data) {
        this.handlePostData(data);
      }
      console.log(data)
    })
  }

  initObservers() {
    this.categories = []
    this.subcategories = []
    this.dataService.categories.subscribe((data) => {
      this.categories = data
      console.log(data)
    })
    this.dataService.subcategories.subscribe((data) => {
      this.subcategories = data
      console.log(data)
    })
  }

  private handlePostData(data: Response) {
    let response = data.data
    if (response == null || response.length === 0) {
      return
    }
    if (this.postService.instanceOfJobPost(response[0])) {
      this.jobPost = data.data
      for (let i = 0; i < this.jobPost.length; i++) {
        let postItem = this.jobPost[i].Post;
        this.post.push(postItem)
        this.handleImage(postItem);
      }
    } else {
      this.post = data.data
      for (let item of this.post) {
        this.handleImage(item);
      }
      console.log(this.post)
    }
    console.log(this.postImageMap)
  }

  private handleImage(postItem: Post) {
    if (postItem.Image && postItem.Image.length > 0) {
      const imageList = [];
      for (let index = 0; index < postItem.Image.length; index++) {
        imageList.push(postItem.Image[index])
      }
      this.postImageMap.set(postItem.ID!, imageList)
    }
  }

  getCategoryName(categoryId: number): string {
    // console.log("getCategoryName id:" + categoryId)
    // console.log(this.categories)
    if (this.categories.length > 0) {
      return this.categories.find(item => item.CategoryId == categoryId)!.Name
    }
    return ""
  }

  getSubcategoryName(subcategoryId: number): string {
    // console.log("getCategoryName id:" + subcategoryId)
    if (this.subcategories.length > 0) {
      return this.subcategories.find(item => item.SubcategoryId == subcategoryId)!.Name
    }
    return ""
  }
}
