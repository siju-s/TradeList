import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import {Post, PostService} from "../post.service";

@Component({
  selector: 'app-viewpost',
  templateUrl: './viewpost.component.html',
  styleUrls: ['./viewpost.component.css']
})
export class ViewpostComponent implements OnInit {

  post:Array<Post> = []

  constructor(private postService: PostService, private changeDetection: ChangeDetectorRef) { }

  ngOnInit(): void {
    this.postService.getPosts().subscribe(data => {
      this.post = data.data
      console.log(data.data[0])
      this.changeDetection.detectChanges()
    })
  }
  //
  // public postItem (index: number, item: Post) {
  //   return item.title;
  // }


}
