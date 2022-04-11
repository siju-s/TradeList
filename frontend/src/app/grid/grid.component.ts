import {Component, OnInit} from '@angular/core';
import {NgbCarouselConfig} from '@ng-bootstrap/ng-bootstrap';
import {JobPost, Post, PostService, Response} from "../formpost/post.service";
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
  images = [

    {title: 'Second Slide', short: 'Second Slide Short', src: "assets/images/image.jpeg"},
    {title: 'First Slide', short: 'First Slide Short', src: "assets/images/floorplans copy.png"},
    {title: 'Third Slide', short: 'Third Slide Short', src: "assets/images/amenities.jpg"}
  ];
  title = 'Card View Demo';
  gridColumns = 3;

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
        this.post.push(this.jobPost[i].Post)
      }
    }
  }
}
