import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { ActivatedRoute } from '@angular/router';
import { NgbCarouselConfig } from '@ng-bootstrap/ng-bootstrap';
import { Categories, JobPost, Post, PostService, Response, Subcategories, Image } from '../formpost/post.service';
import DateUtils from '../helpers/date-helper';
import {LoginService, User} from "../loginform/login.service";
import { DataService } from '../shared/DataService';

@Component({
  selector: 'app-userprofile',
  templateUrl: './userprofile.component.html',
  styleUrls: ['./userprofile.component.css'],
  providers: [NgbCarouselConfig]
})
export class UserprofileComponent implements OnInit {
  // isLoggedIn = false;
  user: User;
  isLoggedIn = false;
  categories: Categories[]
  subcategories: Subcategories[]
  post: Array<Post> = []
  jobPost: Array<JobPost> = []
  postImageMap = new Map<number, Array<Image>>()
  gridColumns = 3;
  titles : Title[]
  selectedid = -1
  selectList =[]

  toggleGridColumns() {
    this.gridColumns = this.gridColumns === 3 ? 4 : 3;
  }

  
  constructor(private postService: PostService, private dataService: DataService, private route: ActivatedRoute, private datePipe: DatePipe) {
    
      }
      // ID
  ngOnInit(): void {
    this.user = JSON.parse(localStorage.getItem('user')!) as User
    this.isLoggedIn = localStorage.getItem('user') != null
    console.log(this.user)
    console.log("Logged in:" + this.isLoggedIn)
    
    this.initObservers()
        this.loadPosts(this.route)
        this.route.params.subscribe(_ => {
          console.log("loadposts")
          this.loadPosts(this.route);
        });
  }

  private loadPosts(route: ActivatedRoute) {
    // let subcategoryId = Number(route.snapshot.paramMap.get('id'))
    this.user = JSON.parse(localStorage.getItem('user')!) as User
    let subcategoryId = Number(this.user.ID)
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
      this.postService.getPostsForUserId(subcategoryId).subscribe(data => {
        if (data.data) {
          this.handlePostData(data);
        }
        console.log('Hello',data)
      })
    }
  }
  getDeleteID(postid : any){
    // this.postService.deleteForUser(postid, userid).subscribe(data => {
    //   if (data.data) {
    //     this.handlePostData(data);
    //   }
    // });
    // **********
      // console.log(postid)
      // this.post = this.post.filter(item => item.ID !== postid);
      // console.log("The id", this.post)
// ************
      // this.postService.deleteForUser(postid, userid).subscribe(data => {
      //   // this.post = this.post.filter(item => item.ID !== postid);
      // });


      // this.postService.deleteForUser(postid, userid).subscribe(data => {
      //   if (data.data) {
      //     this.handlePostData(data);
      //   }
      //   console.log('Hello Delete',data)
      // })

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
        console.log("Individual item", postItem.ID)
        this.post.push(postItem)
        postItem.CreatedAt = DateUtils.getPostDate(this.datePipe, postItem.CreatedAt!);
        this.handleImage(postItem);
      }
    } else {
      this.post = data.data
      for (let item of this.post) {
        this.handleImage(item);
        console.log("Individual item", item.ID)
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

