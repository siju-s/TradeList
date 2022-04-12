import {Component, OnInit, Renderer2} from '@angular/core';
import {Categories, PostService, Subcategories} from "../formpost/post.service";
import {DataService} from "../shared/DataService";
import {MatMenuTrigger} from "@angular/material/menu";
import {MatButton} from "@angular/material/button";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  componentName = "NavbarComponent"
  enteredButton = false;
  isMatMenuOpen = false;
  prevButtonTrigger: MatMenuTrigger;
  categories: Categories[]
  subcategories: Subcategories[]
  categorySubMap = new Map<number, Array<Subcategories>>()

  constructor(private postService: PostService, private dataService: DataService, private renderer2: Renderer2) {

  }

  ngOnInit(): void {
    this.postService.fetchCategories().subscribe(data => {
      this.categories = data.data
      this.dataService.categories.next(this.categories)
      this.subcategories = []
      for (let i = 0; i < this.categories.length; i++) {
        this.fetchSubcategories(this.categories[i].CategoryId)
      }
      console.log(this.categories)
    })
  }

  fetchSubcategories(categoryId: number) {
    this.postService.fetchSubcategories(categoryId).subscribe(data => {
      for (let i = 0; i < data.data.length; i++) {
        this.subcategories.push(data.data[i])
      }
      this.categorySubMap.set(categoryId, data.data)
    })
    this.dataService.subcategories.next(this.subcategories)

    // console.log(this.categorySubMap)
  }

  menuenter() {
    this.isMatMenuOpen = true;
  }

  menuLeave(trigger: MatMenuTrigger, button: MatButton) {
    setTimeout(() => {
      if (!this.enteredButton) {
        this.isMatMenuOpen = false;
        trigger.closeMenu();
        this.renderer2.removeClass(button['_elementRef'].nativeElement, 'cdk-focused');
        this.renderer2.removeClass(button['_elementRef'].nativeElement, 'cdk-program-focused');
      } else {
        this.isMatMenuOpen = false;
      }
    }, 80)
  }

  buttonEnter(trigger: MatMenuTrigger) {
    setTimeout(() => {
      if (this.prevButtonTrigger && this.prevButtonTrigger != trigger) {
        this.prevButtonTrigger.closeMenu();
        this.prevButtonTrigger = trigger;
        this.isMatMenuOpen = false;
        trigger.openMenu();
      } else if (!this.isMatMenuOpen) {
        this.enteredButton = true;
        this.prevButtonTrigger = trigger
        trigger.openMenu();
      } else {
        this.enteredButton = true;
        this.prevButtonTrigger = trigger
      }
    })
  }

  buttonLeave(trigger: MatMenuTrigger, button: MatButton) {
    setTimeout(() => {
      if (this.enteredButton && !this.isMatMenuOpen) {
        trigger.closeMenu();
        this.renderer2.removeClass(button['_elementRef'].nativeElement, 'cdk-focused');
        this.renderer2.removeClass(button['_elementRef'].nativeElement, 'cdk-program-focused');
      }
      if (!this.isMatMenuOpen) {
        trigger.closeMenu();
        this.renderer2.removeClass(button['_elementRef'].nativeElement, 'cdk-focused');
        this.renderer2.removeClass(button['_elementRef'].nativeElement, 'cdk-program-focused');
      } else {
        this.enteredButton = false;
      }
    }, 100)
  }

}
