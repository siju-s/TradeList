import {Component, OnInit} from '@angular/core';
import {Categories, PostService, Subcategories} from "../formpost/post.service";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  componentName = "NavbarComponent"
  categories: Categories[]
  subcategories: Subcategories[]
  categorySubMap = new Map<number, Array<Subcategories>>()

  constructor(private postService: PostService) {

  }

  ngOnInit(): void {
    this.postService.fetchCategories().subscribe(data => {
      this.categories = data.data
      this.subcategories = []
      for (let i = 0; i < this.categories.length; i++) {
        this.fetchSubcategories(this.categories[i].CategoryId)
      }
      console.log(this.subcategories)
    })
  }

  fetchSubcategories(categoryId: number) {
    this.postService.fetchSubcategories(categoryId).subscribe(data => {
      for (let i = 0; i < data.data.length; i++) {
        this.subcategories.push(data.data[i])
      }
      this.categorySubMap.set(categoryId, data.data)
    })
    console.log(this.categorySubMap)
  }

}
