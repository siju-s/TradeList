import {Component, OnInit} from '@angular/core';
import {NgbCarouselConfig} from '@ng-bootstrap/ng-bootstrap';
import {Image, JobPost, Post, PostService, Response} from "../formpost/post.service";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-grid',
  templateUrl: './grid.component.html',
  styleUrls: ['./grid.component.css'],
  providers: [NgbCarouselConfig]
})
export class GridComponent implements OnInit {
  filterTerm!: string;
  componentName = "user"
  // images = [
  //
  //   {title: 'Second Slide', short: 'Second Slide Short', Url: "assets/images/image.jpeg"},
  //   {title: 'First Slide', short: 'First Slide Short', Url: "assets/images/floorplans copy.png"},
  //   {title: 'Third Slide', short: 'Third Slide Short', Url: "assets/images/amenities.jpg"}
  // ];
  // images = []
  title = 'Card View Demo';
  gridColumns = 3;
  postImageMap = new Map<number, Array<Image>>()

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


  constructor(private postService: PostService, private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    let subcategoryId = Number(this.route.snapshot.paramMap.get('id'))
    console.log(subcategoryId)
    this.postService.getPostsForSubcategory(subcategoryId).subscribe(data => {
      if (data.data) {
        this.handlePostData(data);
      }
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
        if (postItem.Image && postItem.Image.length > 0) {
            var imageList = []
            for (let index = 0; index < postItem.Image.length; index++) {
                 imageList.push(postItem.Image[index])
            }
            this.postImageMap.set(postItem.ID!, imageList)
        }
      }
    }
    console.log(this.postImageMap)
  }
}
