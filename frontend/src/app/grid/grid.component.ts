import {Component, OnInit} from '@angular/core';
import {NgbCarouselConfig} from '@ng-bootstrap/ng-bootstrap';
import {Categories, Image, JobPost, Post, PostService, Response, Subcategories} from "../formpost/post.service";
import {ActivatedRoute} from "@angular/router";
import {DataService} from "../shared/DataService";
import {DatePipe} from '@angular/common';
import DateUtils from "../helpers/date-helper";

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


  constructor(private postService: PostService, private dataService: DataService, private route: ActivatedRoute, private datePipe: DatePipe) {
  }

  ngOnInit(): void {
    this.initObservers()
    this.loadPosts(this.route)
    this.route.params.subscribe(_ => {
      console.log("loadposts")
      this.loadPosts(this.route);
    });
  }

  private loadPosts(route: ActivatedRoute) {
    let subcategoryId = Number(route.snapshot.paramMap.get('id'))
    this.post = []
    this.jobPost = []
    console.log("loadPosts subcategory:" + subcategoryId)
    if (subcategoryId === 0) {
      this.postService.getPosts().subscribe(data => {
        if (data.data) {
          this.handlePostData(data);
        }
        console.log(data)

      })
    } else {
      this.postService.getPostsForSubcategory(subcategoryId).subscribe(data => {
        if (data.data) {
          this.handlePostData(data);
        }
        console.log(data)
      })
    }
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
        postItem.CreatedAt = DateUtils.getPostDate(this.datePipe, postItem.CreatedAt!);
        this.handleImage(postItem);
      }
    } else {
      this.post = data.data
      for (let item of this.post) {
        this.handleImage(item);
        item.CreatedAt = DateUtils.getPostDate(this.datePipe, item.CreatedAt!);
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

  getFormattedDate(date: Date) {
    var year = date.getFullYear();

    var month = (1 + date.getMonth()).toString();
    month = month.length > 1 ? month : '0' + month;

    var day = date.getDate().toString();
    day = day.length > 1 ? day : '0' + day;

    return month + '-' + day + '-' + year;
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
